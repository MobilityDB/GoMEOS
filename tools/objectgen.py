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
from dataclasses import dataclass
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
        self.meostype: dict[str, str] = {}      # class -> the MeosType it stands for
        self.wrap_root: dict[str, str] = {}
        self._build()

    def _build(self) -> None:
        # A leaf naming one type stands for it, as #_build of MEOS.NET's
        # tools/objectgen.py reads the lattice and the companions.
        for name, node in self.om["lattice"].items():
            self.parent[name] = node["parent"]
            temptypes = node.get("temptypes") or []
            if node.get("kind") == "leaf" and len(temptypes) == 1:
                self.meostype[name] = temptypes[0]

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
                if node.get("temptype"):
                    self.meostype[name] = node["temptype"]

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


@dataclass
class Emitted:
    """A method #method_for emitted: the C function behind it, its Go name, its
    parameter declarations and names, its result clause, and whether it is a
    method on the value or a function of the class."""
    function: str
    name: str
    decl: list
    args: list
    result: str
    instance: bool


class Generator:
    def __init__(self, model: Model):
        self.m = model
        self.deferred: dict[str, list[str]] = defaultdict(list)
        self.emitted = 0
        self.last: Emitted | None = None
        # class -> {C function: instance method}, its ancestors' included, since
        # an embedded parent passes its methods down.
        self.class_functions: dict[str | None, dict[str, Emitted]] = {}
        # C function -> function of a class, for every class emitted before.
        self.static_functions: dict[str, Emitted] = {}
        # class -> the text writer its String calls.
        self.class_text_writer: dict[str | None, str | None] = {}
        # Every function of the package, since Go keeps one name per package.
        self.package_names: set[str] = set()
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
        # A string the flat wrapper takes as *string is one MEOS accepts as NULL, and
        # its nil reaches the wrapper unchanged; #qualify spells it as it stands
        return go_type in self.scalars or go_type in ("string", "*string", "bool",
                                                       "unsafe.Pointer")

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

        # A byte buffer travels as a []byte, as MEOS.NET's method_for reads the
        # catalog's shape: a `size_t *` the shape lists in `outParams` states the
        # length of the buffer the call answers, so it leaves the signature and
        # the buffer comes back as a []byte (functions.TakeBytes); a `uint8_t`
        # array the shape lists in `inputArrays` is taken as a []byte, the
        # parameter its length comes from filled with len().
        shape = f.get("shape") or {}
        length_out = next(
            (codegen._go_param_name(p["name"]) for p in f.get("params", [])
             if p["name"] in (shape.get("outParams") or [])
             and p["cType"].replace(" ", "") == "size_t*"), None)
        byte_inputs = {
            codegen._go_param_name(a["param"]): codegen._go_param_name(a["lengthFrom"]["name"])
            for a in shape.get("inputArrays") or []
            if (a.get("element") or {}).get("c") == "uint8_t"
            and (a.get("lengthFrom") or {}).get("kind") == "param"}
        length_of = {length: arr for arr, length in byte_inputs.items()}
        prelude = []

        args, decl = [], []
        for pname, ptype in params:
            if pname == length_out:
                prelude.append("\tvar _size uint64")
                args.append("unsafe.Pointer(&_size)")
                continue
            if pname in byte_inputs:
                local = f"_p_{pname}"
                prelude.append(f"\tvar {local} unsafe.Pointer\n"
                               f"\tif len({ident(pname)}) > 0 {{\n"
                               f"\t\t{local} = unsafe.Pointer(&{ident(pname)}[0])\n\t}}")
                decl.append(f"{ident(pname)} []byte")
                args.append(local)
                continue
            if pname in length_of:
                args.append(f"{self.qualify(ptype)}(len({ident(length_of[pname])}))")
                continue
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
        # A method may not share the name of the struct its class embeds, the
        # field Go gives that parent (JsonbSet embeds Set, and jsonbset_set reads
        # `set`), so such a method takes the name of its SQL function, which
        # carries the type as the SQL dialect writes it (jsonbsetSet), spelled as
        # #pascal spells every method.
        if recv is not None and name == self.m.parent.get(cls):
            name = pascal(f.get("sqlfn") or cls[:1].lower() + cls[1:] + name)

        # THE OUT-PARAMETER FOLD, and the shape is the repository's own: MEOS answers
        # a bool saying whether a value is there and writes the value out, so the
        # wrapper returns (found, value, error) and the method answers (value, ok,
        # error) -- absence is an optional result, never an error.
        folded = (len(sig.returns) == 2 and sig.returns[0] == "bool"
                  and f["returnType"]["c"].replace("const ", "").strip() == "bool")

        rets, body = [], list(prelude)
        call = f"functions.{sig.go_name}({', '.join(args)})"
        answers_bytes = (length_out is not None and sig.returns == ["unsafe.Pointer"]
                         and f["returnType"]["c"].replace("const ", "").replace(" ", "")
                         == "uint8_t*")

        if not sig.returns:
            body.append(f"\treturn {call}")
        elif answers_bytes:
            rets = ["[]byte"]
            body.append(f"\t_r0, _err := {call}")
            body.append("\tif _err != nil {\n\t\treturn nil, _err\n\t}")
            body.append("\treturn functions.TakeBytes(_r0, _size), nil")
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
        # What #portable_methods and #string_method call it by.
        self.last = Emitted(fname, name, decl, [d.split(" ", 1)[0] for d in decl],
                            result, recv is not None)
        return f"// {name} is MEOS {fname}.\n{head}\n" + "\n".join(body) + "\n}\n"

    # The portable spelling of each encoding the catalog's typeEncodings names,
    # the SQL dialect's (asText, asEWKT, asWKB, asHexWKB, asEWKB, asHexEWKB,
    # asMFJSON and their From readers), as #FORMAT_SPELLING of MEOS.NET's
    # tools/objectgen.py spells them.
    FORMAT_SPELLING = {"text": "Text", "ewkt": "EWKT", "wkb": "WKB",
                       "hexwkb": "HexWKB", "ewkb": "EWKB", "hexewkb": "HexEWKB",
                       "mfjson": "MFJSON"}

    def encoding(self, cls: str) -> tuple[dict | None, str | None]:
        """The typeEncodings entry of the class's C type and its SQL type name,
        the MeosType without its T_, lower-cased, as meos_catalog.c names it."""
        enc = self.m.idl.get("typeEncodings", {}).get(self.m.ctype.get(cls))
        mt = self.m.meostype.get(cls)
        return enc, (mt[2:].lower() if mt else None)

    def portable_methods(self, cls: str, own: list, functions: dict) -> list[str]:
        """The portable reader and writer of every encoding typeEncodings gives the
        class's type, as #portable_methods of MEOS.NET's tools/objectgen.py emits
        them: the function of the type's SQL name in readers/writers, else the C
        type's decoders/encoders. Where the class or an ancestor has no method of
        that name, As<Format> forwards to the method behind the writer and
        <Class>From<Format> to the function behind the reader, so
        TGeomPointFromEWKB is TemporalFromWKB, one function under two names as in
        SQL."""
        enc, sqltype = self.encoding(cls)
        if not enc or not sqltype:
            return []
        taken = {e.name for e in own} | {e.name for e in functions.values()}
        out = []
        for e in enc.get("encodings", []):
            spelling = self.FORMAT_SPELLING.get(e)
            if not spelling:
                continue
            for verb, by_type, by_class in (("As", "writers", "encoders"),
                                            ("From", "readers", "decoders")):
                fn = (enc.get(by_type, {}).get(e, {}).get(sqltype)
                      or enc.get(by_class, {}).get(e))
                if not fn:
                    continue
                if verb == "As":
                    target, name = functions.get(fn), verb + spelling
                    call = f"x.{target.name}" if target else None
                else:
                    target, name = self.static_functions.get(fn), cls + verb + spelling
                    call = target.name if target else None
                if target is None or name in taken or name in self.package_names:
                    continue
                result, body = target.result, f"\treturn {call}({', '.join(target.args)})"
                # A reader answers the class it reads: a function returning an
                # ancestor (TemporalFromWKB answers *Temporal) is wrapped in this
                # class, so TGeomPointFromEWKB answers a *TGeomPoint, nil kept nil.
                answered = result[2:].split(",", 1)[0] if result.startswith("(*") else None
                if (verb == "From" and answered and answered != cls
                        and answered in self.m.ancestors(cls)):
                    result = f"(*{cls}, error)"
                    body = (f"\tr, err := {call}({', '.join(target.args)})\n"
                            f"\tif err != nil || r == nil {{\n\t\treturn nil, err\n\t}}\n"
                            f"\treturn {cls}FromPointer(r.Pointer()), nil")
                head = (f"func (x *{cls}) {name}({', '.join(target.decl)}) {result}"
                        if verb == "As" else
                        f"func {name}({', '.join(target.decl)}) {result}")
                out.append(f"// {name} is the portable {verb.lower()}{spelling} of "
                           f"MEOS {fn}.\n{head} {{\n{body}\n}}\n")
                taken.add(name)
                if verb == "From":
                    self.package_names.add(name)
                self.emitted += 1
        return out

    def text_writer(self, cls: str) -> str | None:
        """The text writer typeEncodings names for the class, as #text_writer of
        MEOS.NET's tools/objectgen.py reads it: the writer of its SQL type, else
        the text encoder of its C type, else, for a C type the catalog keeps out
        of its codecs (text, which MEOS-API counts with the strings), the class's
        own *_out."""
        enc, sqltype = self.encoding(cls)
        if enc is None:
            return next((f for f in self.class_functions.get(cls, {})
                         if f.endswith("_out")), None)
        if sqltype:
            writer = enc.get("writers", {}).get("text", {}).get(sqltype)
            if writer:
                return writer
        return enc.get("encoders", {}).get("text")

    def string_method(self, cls: str, functions: dict) -> list[str]:
        """String over the text writer #text_writer names, the function behind the
        SQL asText of the type (rule 2 of the bindings in the I/O plan, MobilityDB
        #3353): AsText for a spatial value, which writes its points as WKT, and
        Out for the others; a class writing what it inherits emits none. A Go
        String answers no error, so a MEOS error is written as fmt writes a
        failure, %!s(MEOS error: ...), never as an empty string."""
        inherited = self.class_text_writer.get(self.m.parent.get(cls))
        writer = self.text_writer(cls) or inherited
        self.class_text_writer[cls] = writer
        if writer is None or writer == inherited:
            return []
        target = functions.get(writer)
        if target is None or target.result != "(string, error)":
            return []
        args = []
        if target.decl:
            if [d.split(" ", 1)[1] for d in target.decl] != ["int"]:
                return []
            digits = self.default_digits(cls)
            if digits is None:
                return []
            args = [str(digits)]
        return [f"// String is the text MEOS writes this value as, MEOS {writer}.\n"
                f"func (x *{cls}) String() string {{\n"
                f"\ts, err := x.{target.name}({', '.join(args)})\n"
                f"\tif err != nil {{\n\t\treturn \"%!s(MEOS error: \" + err.Error() + \")\"\n\t}}\n"
                f"\treturn s\n}}\n"]

    def default_digits(self, cls: str) -> int | None:
        """The digits a text writer takes by default: the text encoding's own,
        else MEOS's OUT_DEFAULT_DECIMAL_DIGITS, as MEOS.NET's ToString reads them."""
        enc, _ = self.encoding(cls)
        aux = {a["name"]: a["default"]
               for a in (enc or {}).get("encoderAux", {}).get("text", [])
               or (enc or {}).get("out_aux", [])}
        if "maxdd" in aux:
            return aux["maxdd"]
        return next((m.get("value") for m in self.m.idl.get("macros", [])
                     if m.get("name") == "OUT_DEFAULT_DECIMAL_DIGITS"), None)

    def emit_class(self, cls: str) -> str:
        spec = self.m.om["classes"].get(cls) or {"methods": []}
        parent = self.m.parent.get(cls)

        methods, seen, own = [], set(), []
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
            own.append(self.last)

        functions = dict(self.class_functions.get(parent, {}))
        functions.update({e.function: e for e in own if e.instance})
        self.class_functions[cls] = functions
        self.static_functions.update({e.function: e for e in own if not e.instance})
        self.package_names |= {e.name for e in own if not e.instance}
        methods += self.portable_methods(cls, own, functions)
        methods += self.string_method(cls, functions)

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
    # A parent before its children, so a class reaches the methods its ancestors
    # emit (#emit_class reads them for String and the portable forms).
    for cls in sorted(model.classes(), key=lambda c: (model.depth(c), c)):
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
