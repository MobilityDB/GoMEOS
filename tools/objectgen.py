#!/usr/bin/env python3
# BINDING-HEADER-PARSE-OK: this generator reads only the meos-idl.json catalog,
# never a C header; the catalog is the binding's single source of truth.
"""Generate the GoMEOS object layer from MEOS-API's meos-idl.json.

The catalog's ``objectModel`` is the ecosystem-wide statement of the class hierarchy
implicit in MEOS: ``lattice`` carries the temporal tree, ``companions`` the Box,
Collection and Value hierarchies, and ``classes.<Class>.methods`` assigns every public
MEOS function to the class it is a method of, under its canonical camelCase ``ooName``.
This projects that model onto Go -- one type per model class, the model's parent edge
spelled as struct embedding, and one method per assigned function delegating to the
wrapper ``codegen.py`` emits for the same function.

THE WRAPPER SIGNATURES COME FROM ``codegen.SIGNATURES``, never from the raw C parameter
list, so the two generators cannot disagree about a folded out-parameter: MEOS declares
``bool tbool_value_at_timestamptz(..., bool *value)`` and the emitted wrapper takes three
arguments and answers two values, which no C signature states.  That table is filled by
RUNNING the flat generator, so which functions have a wrapper is decided in exactly one
place; this file holds no second copy of what that generator skips.

Go spells inheritance as embedding, so ``TBool`` embeds ``TAlpha`` embeds ``Temporal``
and a parent's methods are promoted. Every instance is a handle to the MEOS value, and
the two packages exchange one across ``unsafe.Pointer`` -- cgo types are package-scoped,
so ``*C.Temporal`` in ``functions`` and here are different Go types and the pointer is
the only shared currency.

Usage:
    python3 tools/objectgen.py path/to/meos-idl.json [--report]

Regenerates functions/ (through codegen) and writes types/*.go, replacing that directory.
"""

from __future__ import annotations

import argparse
import json
import shutil
import sys
from collections import defaultdict
from pathlib import Path

import codegen

PACKAGE = "types"
FUNCTIONS_IMPORT = "github.com/MobilityDB/GoMEOS/functions"

# The C structs that carry the Temporal header and are discriminated by its ``subtype``
# field. A pointer to any of them is a Temporal at the surface, exactly as MEOS.NET's
# object layer resolves them, because the concrete class is the product leaf x subtype
# and the leaf family carries by far the larger surface.
TEMPORAL_STRUCTS = ("Temporal", "TInstant", "TSequence", "TSequenceSet")

GO_KEYWORDS = {
    "break", "case", "chan", "const", "continue", "default", "defer", "else",
    "fallthrough", "for", "func", "go", "goto", "if", "import", "interface", "map",
    "package", "range", "return", "select", "struct", "switch", "type", "var",
}

# Acronym runs kept upper case so a Go name reads the way the catalog's own camelCase
# name does: ``asMFJSON`` becomes ``AsMFJSON`` rather than ``AsMfjson``.
ACRONYMS = {"mfjson", "geojson", "wkb", "wkt", "ewkb", "ewkt", "hexwkb", "hexewkb",
            "srid", "json", "jsonb", "mvt", "gml", "kml", "de9im", "id", "url"}


def pascal(name: str) -> str:
    """The Go exported spelling of a catalog camelCase ``ooName``."""
    out, word = [], []

    def flush():
        if not word:
            return
        w = "".join(word)
        out.append(w.upper() if w.lower() in ACRONYMS else w[:1].upper() + w[1:])
        word.clear()

    for ch in name:
        if ch in "_-":
            flush()
            continue
        if ch.isupper() and word:
            flush()
        word.append(ch)
    flush()
    return "".join(out) or name


def ident(name: str) -> str:
    """A Go parameter name, kept clear of the keywords."""
    return name + "_" if name in GO_KEYWORDS else name


def zero(go_type: str) -> str:
    """The zero value of a Go type, for the error and absent returns."""
    if go_type.startswith("*") or go_type.startswith("[]") or go_type == "unsafe.Pointer":
        return "nil"
    if go_type == "string":
        return '""'
    if go_type == "bool":
        return "false"
    return "0"


class Model:
    """The class hierarchy the catalog defines, indexed for emission."""

    def __init__(self, idl: dict):
        self.idl = idl
        self.om = idl["objectModel"]
        self.functions = {f["name"]: f for f in idl["functions"]}
        self.parent: dict[str, str | None] = {}
        self.ctype: dict[str, str] = {}
        self.wrap_root: dict[str, str] = {}
        self._build()

    def _build(self) -> None:
        for name, node in self.om["lattice"].items():
            self.parent[name] = node["parent"]

        # The template subtypes are classes of their own, under the root.
        for value in self.om["axes"]["subtype"]["values"]:
            cls = value.get("class")
            if cls:
                self.parent[cls] = "Temporal"

        for family in [k for k in self.om["companions"] if not k.startswith("_")]:
            for name, node in self.om["companions"][family]["nodes"].items():
                if name.startswith("_"):
                    continue
                self.parent[name] = node["parent"]

        # What a class's instances point to is the catalog's to say, so every class
        # takes the cType the model derives from MEOS's own signatures. A binding
        # holding that map itself cannot reach a class the model gains.
        for name, spec in self.om["classes"].items():
            if spec.get("cType"):
                self.ctype[name] = spec["cType"]
            self.parent.setdefault(name, None)

        # A pointer to a C struct is wrapped in the SHALLOWEST class carrying it, so a
        # GSERIALIZED * is a Geo and its Geometry and Geography leaves stay distinct
        # classes under it, exactly as a Set * is a Set.
        for name, ct in self.ctype.items():
            best = self.wrap_root.get(ct)
            if best is None or self.depth(name) < self.depth(best):
                self.wrap_root[ct] = name

    def depth(self, cls: str) -> int:
        d, seen = 0, {cls}
        p = self.parent.get(cls)
        while p and p not in seen:
            seen.add(p)
            d += 1
            p = self.parent.get(p)
        return d

    def ancestors(self, cls: str) -> list[str]:
        out, seen = [], {cls}
        p = self.parent.get(cls)
        while p and p not in seen:
            out.append(p)
            seen.add(p)
            p = self.parent.get(p)
        return out

    def classes(self) -> list[str]:
        """Every type the layer declares, superclass included.

        A companion ROOT -- Box, Collection, Value -- is a parent edge in the model
        and carries no ``classes`` entry of its own, so emitting only what ``classes``
        names leaves each child embedding a type nothing declares. The set to emit is
        therefore the classes UNION everything named as a parent.
        """
        named = set(self.om["classes"])
        parents = {p for p in self.parent.values() if p}
        return sorted(named | parents)

    def class_for_ctype(self, struct: str) -> str | None:
        """The class a single pointer to this C struct is wrapped in."""
        if struct in TEMPORAL_STRUCTS:
            return "Temporal"
        return self.wrap_root.get(struct)


class Generator:
    def __init__(self, model: Model):
        self.m = model
        self.deferred: dict[str, list[str]] = defaultdict(list)
        self.emitted = 0
        # The Go handle each wrapper struct reaches the flat layer as, read from the
        # flat generator rather than restated: WRAPPER_TYPES is what decides it. Its
        # value carries the star and need not echo the struct -- GSERIALIZED reaches
        # Go as `*Geom` -- so the handle NAME is that value without the star, and it
        # is what `functions.<Handle>FromPointer` is called.
        self.handle_of = {c: go.lstrip("*") for c, (go, _) in codegen.WRAPPER_TYPES.items()}
        self.struct_of = {go: c for c, (go, _) in codegen.WRAPPER_TYPES.items()}
        self.scalars = set(t.go_type for t in codegen.TYPE_MAP.values())

    # -- marshalling ------------------------------------------------------

    def class_of_handle(self, go_type: str) -> str | None:
        """The model class behind a wrapper's ``*Handle`` Go type, if there is one."""
        if not go_type.startswith("*"):
            return None
        struct = self.struct_of.get(go_type)
        return self.m.class_for_ctype(struct) if struct else None

    def passthrough(self, go_type: str) -> bool:
        """Whether this layer hands the wrapper's own type straight through.

        A MEOS struct the model gives no class is one of these. `Interval` is
        PostgreSQL's and absent from the MeosType enum, `SkipList`, `PJ_CONTEXT` and
        `gsl_rng` are runtime handles MEOS registers in no enum, and `AFFINE`, `BOX3D`,
        `GBOX` and `Match` state their whole layout in scalars -- so no class stands for
        any of them and none should. The flat handle IS the type, and its semantics are
        already the ones MEOS.NET states for a by-value struct: the call hands MEOS an
        address and the memory stays MEOS's. Spelling a value struct here instead is not
        available anyway -- this package imports no cgo, so it cannot materialise a C
        struct, and only the `functions` package can.
        """
        return (go_type in self.scalars
                or go_type in ("string", "bool", "unsafe.Pointer")
                or go_type in self.struct_of)

    def qualify(self, go_type: str) -> str:
        """A passthrough type as this package must spell it.

        The flat generator declares the catalog's enums and named scalars as Go types
        of its OWN package -- `Interpolation`, `NullHandleType`, `MeosType` -- so a
        method carrying one names it `functions.X` here. Go's predeclared types are all
        lower case, which is the discriminator, so a new enum needs no list entry.
        """
        prefix, base = "", go_type
        while base.startswith("[]"):
            prefix, base = prefix + "[]", base[2:]
        if base.startswith("*"):
            prefix, base = prefix + "*", base[1:]
        if base[:1].isupper():
            return f"{prefix}functions.{base}"
        return go_type

    def map_param(self, go_type: str, name: str) -> tuple[str, str] | None:
        """``(go type in this layer, the expression handed to the wrapper)``."""
        cls = self.class_of_handle(go_type)
        if cls:
            # THE DECLARED TYPE AND THE HANDLE ARE NOT THE SAME QUESTION. The surface
            # takes the CLASS a pointer is wrapped in -- TInstant, TSequence and
            # TSequenceSet all surface as Temporal -- while the wrapper wants the
            # handle IT declares, so the conversion reads the wrapper's own struct.
            # Only the pointer crosses, so the two need not agree.
            handle = self.handle_of.get(self.struct_of.get(go_type, ""))
            if handle is None:
                return None
            return (f"*{cls}", f"functions.{handle}FromPointer({name}.Pointer())")
        if self.passthrough(go_type):
            return (self.qualify(go_type), name)
        return None

    def map_return(self, go_type: str, expr: str) -> tuple[str, str] | None:
        """``(go type in this layer, the expression producing it)``."""
        cls = self.class_of_handle(go_type)
        if cls:
            return (f"*{cls}", f"{cls}FromPointer({expr}.Pointer())")
        if self.passthrough(go_type):
            return (self.qualify(go_type), expr)
        return None

    # -- emission ---------------------------------------------------------

    def method_for(self, cls: str, entry: dict) -> str | None:
        oo = entry["ooName"]
        fname = entry["function"]
        sig = codegen.SIGNATURES.get(fname)
        if sig is None:
            self.deferred[cls].append(f"{oo}: {fname} has no emitted wrapper")
            return None
        f = self.m.functions.get(fname)
        if f is None:
            self.deferred[cls].append(f"{oo}: no catalog function {fname}")
            return None

        # A first parameter that is this class (or an ancestor of it) makes the
        # function a METHOD ON the value; anything else is a plain function of the
        # class, which is how Go spells a static.
        family = {cls, *self.m.ancestors(cls)}
        params = list(sig.params)
        recv = None
        if params and self.class_of_handle(params[0][1]) in family:
            recv = params.pop(0)

        args, decl = [], []
        for pname, ptype in params:
            mapped = self.map_param(ptype, ident(pname))
            if mapped is None:
                self.deferred[cls].append(f"{oo}: parameter {pname} is {ptype}")
                return None
            gotype, expr = mapped
            decl.append(f"{ident(pname)} {gotype}")
            args.append(expr)

        if recv is not None:
            handle = self.handle_of.get(self.struct_of.get(recv[1], ""))
            if handle is None:
                self.deferred[cls].append(f"{oo}: receiver {recv[1]} has no handle")
                return None
            args.insert(0, f"functions.{handle}FromPointer(x.Pointer())")

        name = pascal(oo) if recv is not None else cls + pascal(oo)

        # THE OUT-PARAMETER FOLD, and the shape is the repository's own: MEOS answers
        # a bool saying whether a value is there and writes the value out, so the
        # wrapper returns (found, value, error) and the method answers (value, ok,
        # error) -- absence is an optional result, never an error.
        folded = (len(sig.returns) == 2 and sig.returns[0] == "bool"
                  and f["returnType"]["c"].replace("const ", "").strip() == "bool")

        rets, body = [], []
        call = f"functions.{sig.go_name}({', '.join(args)})"

        if not sig.returns:
            body.append(f"\treturn {call}")
        elif folded:
            mapped = self.map_return(sig.returns[1], "_value")
            if mapped is None:
                self.deferred[cls].append(f"{oo}: result {sig.returns[1]}")
                return None
            gotype, expr = mapped
            rets = [gotype, "bool"]
            body.append(f"\t_found, _value, _err := {call}")
            body.append(f"\tif _err != nil {{\n\t\treturn {zero(gotype)}, false, _err\n\t}}")
            body.append(f"\tif !_found {{\n\t\treturn {zero(gotype)}, false, nil\n\t}}")
            body.append(f"\treturn {expr}, true, nil")
        else:
            names = [f"_r{i}" for i in range(len(sig.returns))]
            outs = []
            for i, rt in enumerate(sig.returns):
                mapped = self.map_return(rt, names[i])
                if mapped is None:
                    self.deferred[cls].append(f"{oo}: result {rt}")
                    return None
                rets.append(mapped[0])
                outs.append(mapped[1])
            body.append(f"\t{', '.join(names)}, _err := {call}")
            body.append("\tif _err != nil {\n\t\treturn "
                        + ", ".join(zero(r) for r in rets) + ", _err\n\t}")
            body.append(f"\treturn {', '.join(outs)}, nil")

        result = f"({', '.join(rets + ['error'])})" if rets else "error"
        head = (f"func (x *{cls}) {name}({', '.join(decl)}) {result} {{"
                if recv is not None else
                f"func {name}({', '.join(decl)}) {result} {{")
        self.emitted += 1
        return f"// {name} is MEOS {fname}.\n{head}\n" + "\n".join(body) + "\n}\n"

    def emit_class(self, cls: str) -> str:
        spec = self.m.om["classes"].get(cls) or {"methods": []}
        parent = self.m.parent.get(cls)

        methods, seen = [], set()
        for entry in spec["methods"]:
            if entry.get("ooExclude"):
                continue
            code = self.method_for(cls, entry)
            if code is None:
                continue
            head = code.split("\n")[1]
            if head in seen:
                continue
            seen.add(head)
            methods.append(code)

        uses_functions = any("functions." in m for m in methods)
        imports = ['\t"unsafe"']
        if uses_functions:
            imports.append(f'\n\t"{FUNCTIONS_IMPORT}"')

        out = [f"package {PACKAGE}\n",
               "// Code generated by tools/objectgen.py from meos-idl.json. DO NOT EDIT.\n",
               "import (\n" + "\n".join(imports) + "\n)\n"]

        doc = (self.m.om["lattice"].get(cls) or {}).get("doc") or ""
        if doc:
            out.append(f"// {cls} is {doc[0].lower() + doc[1:]}\n")

        out.append(f"type {cls} struct {{\n\t{parent if parent else 'handle'}\n}}\n")
        out.append(
            f"// {cls}FromPointer wraps a MEOS pointer, answering nil for nil so an\n"
            f"// absent MEOS result stays absent rather than becoming a live handle.\n"
            f"func {cls}FromPointer(p unsafe.Pointer) *{cls} {{\n"
            f"\tif p == nil {{\n\t\treturn nil\n\t}}\n"
            f"\tv := &{cls}{{}}\n\tv.ptr = p\n\treturn v\n}}\n")
        out.extend(methods)
        return "\n".join(out)

    def base_file(self) -> str:
        return (
            f"package {PACKAGE}\n\n"
            "// Code generated by tools/objectgen.py from meos-idl.json. DO NOT EDIT.\n\n"
            'import "unsafe"\n\n'
            "// handle is the MEOS value every class in this package stands for.\n"
            "//\n"
            "// cgo types are package-scoped, so the *C.Temporal the functions package\n"
            "// holds and one declared here would be DIFFERENT Go types. The untyped\n"
            "// pointer is the only currency the two packages share, which is why each\n"
            "// class carries one and hands it across at every call.\n"
            "type handle struct {\n\tptr unsafe.Pointer\n}\n\n"
            "// Pointer answers the MEOS value this handle stands for, or nil.\n"
            "func (h *handle) Pointer() unsafe.Pointer {\n"
            "\tif h == nil {\n\t\treturn nil\n\t}\n\treturn h.ptr\n}\n")


def main() -> int:
    repo = Path(__file__).resolve().parent.parent

    ap = argparse.ArgumentParser()
    ap.add_argument("catalog", type=Path, nargs="?",
                    default=repo / "tools" / "meos-idl.json",
                    help="the catalog to project (default: tools/meos-idl.json, "
                         "where the refresh stages it)")
    ap.add_argument("--report", action="store_true",
                    help="print what each class deferred and why")
    args = ap.parse_args()

    # Fill codegen.SIGNATURES by RUNNING the flat generator, so which functions have a
    # wrapper is decided in one place. Re-deriving that here would be a second copy of
    # its exclusions, and the two would drift.
    stats = codegen.generate(args.catalog, repo / "functions")
    print(f"Flat layer: {stats['emitted']} wrappers, "
          f"{len(codegen.SIGNATURES)} signatures recorded", file=sys.stderr)

    idl = json.loads(args.catalog.read_text())
    model = Model(idl)
    gen = Generator(model)

    out_dir = repo / PACKAGE
    if out_dir.exists():
        shutil.rmtree(out_dir)
    out_dir.mkdir(parents=True)
    (out_dir / "meos.go").write_text(gen.base_file())
    for cls in model.classes():
        (out_dir / f"{cls.lower()}.go").write_text(gen.emit_class(cls))

    total_deferred = sum(len(v) for v in gen.deferred.values())
    print(f"Object layer: {len(model.classes())} classes, {gen.emitted} methods, "
          f"{total_deferred} deferred", file=sys.stderr)
    if args.report:
        for cls in sorted(gen.deferred, key=lambda c: -len(gen.deferred[c])):
            print(f"--- {cls} ({len(gen.deferred[cls])} deferred)", file=sys.stderr)
            for reason in gen.deferred[cls]:
                print(f"      {reason}", file=sys.stderr)
    return 0


if __name__ == "__main__":
    sys.exit(main())
