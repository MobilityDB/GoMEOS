"""GoMEOS code generator.

Drives idiomatic Go wrapper generation from meos-idl.json (produced by the
MEOS-API parser).  Output is one Go source file per MEOS public header,
emitted to ``generated/`` and consumed by the hand-written ergonomic surface
in the package root.

The generator is intentionally minimal in its first iteration: it covers the
scalar-in / opaque-pointer-out signatures (the bulk of the API) and leaves
edge-case shapes (multiple return values, output parameters, array
arguments, generic-method dispatch) for follow-up commits.  Functions whose
shape is not yet handled emit a ``// TODO`` placeholder so the diff against
the hand-written surface stays auditable.

Run from the repo root:

    python3 tools/codegen.py
"""

from __future__ import annotations

import json
import re
import sys
from dataclasses import dataclass
from pathlib import Path

HEADER_FILES = [
    "meos.h",
    "meos_catalog.h",
    "meos_geo.h",
    "meos_internal.h",
    "meos_internal_geo.h",
    "meos_npoint.h",
    "meos_cbuffer.h",
    "meos_pose.h",
    "meos_rgeo.h",
    "meos_h3.h",
    "meos_quadbin.h",
    "meos_json.h",
    "meos_pointcloud.h",
]

# Forward-declared opaque types we never wrap (mirrors the cdef skip list
# the PyMEOS-CFFI builder applies).
OPAQUE_TYPES = ("json_object", "GEOSContextHandle_t")

# C-name skip list: functions handled by hand or intentionally hidden.
# CODEGEN-REGULAR-EXCEPTION: this is the documented audit bucket for the
# error-plumbing functions the binding never exposes (it drives errors through
# its own handler). meos_error joins it because it is the sole VARIADIC public
# MEOS function (``void meos_error(int, int, const char *, ...)``) and cgo
# cannot call a C variadic function at all -- it is unwrappable by construction,
# not skipped to pad coverage. Every other catalog function is emitted.
SKIPPED_FUNCTIONS = {
    "py_error_handler",
    "meos_initialize_error_handler",
    "meos_error",
    # The errno accessors are the plumbing this generator's own error
    # projection is built on: every emitted wrapper clears the errno before
    # its call and reads it after.  Wrapping them as ordinary functions would
    # both expose that plumbing as public API and be self-defeating -- a
    # generated MeosErrno() would reset the errno immediately before reading
    # it, and so always report success.  Callers get errors as Go errors.
    "meos_errno",
    "meos_errno_set",
    "meos_errno_restore",
    "meos_errno_reset",
}

# The error projection ------------------------------------------------------
#
# MEOS reports a failure out of band: the failing call sets a thread-local
# errno and returns a sentinel that is the maximum of its own return type.
# That sentinel can never answer the question by itself -- INT_MAX is a
# legitimate count, DBL_MAX a legitimate distance, and a bool has no spare
# value at all, so its failure is indistinguishable from a definite false.
# The out-of-band state is the only discriminator.
#
# Go's idiom for that channel is the trailing error result, so every emitted
# wrapper returns one.  This is applied to 100% of the emitted surface with no
# per-function exception: a rule that held for only some functions would need
# a hand-maintained list, and the property would stop being checkable.
_ERRNO_RESET = "\tC.meos_errno_reset()"
_ERRNO_GUARD = "\tif _err = meosError(); _err != nil {\n\t\treturn\n\t}"


def _free_returned_string(return_c: str) -> str:
    """The free for a C string MEOS hands to the caller, or nothing.

    A ``char *`` result belongs to the caller: MEOS's own example frees it
    (``meos/examples/03_berlinmod_assemble.c`` does ``free(trip_str)`` on the
    result of ``temporal_as_hexwkb``).  ``C.GoString`` copies into Go memory, so
    without this the C allocation is unreachable the moment the wrapper returns.

    ``const char *`` is the opposite case and must NOT be freed: MEOS hands back
    a name it owns (``temporal_interp``, ``temporal_subtype``,
    ``temporal_basetype_name``).  The C declaration is the whole discriminator,
    so nothing here is inferred from a function's name or group.

    The free is deferred, so it runs after ``C.GoString`` has copied, and it sits
    after the errno guard so a failed call -- which returns no buffer -- never
    reaches it.
    """
    return "\n\tdefer C.free(unsafe.Pointer(_cret))" if return_c == "char *" else ""


def _result_signature(all_returns: list[str]) -> str:
    """Render the Go result list, with the projected error appended.

    The results are named so the error guard can return the zero value of
    every result with a bare ``return``, which keeps the emitted guard one
    shape for every function rather than a per-type zero-value expression.
    """
    named = [f"_r{i} {go_type}" for i, go_type in enumerate(all_returns)]
    named.append("_err error")
    return f" ({', '.join(named)})"


@dataclass
class Signature:
    """The Go signature this generator emits for one MEOS function.

    A caller of a generated wrapper needs the shape the generator CHOSE, which is not
    the C parameter list: an out-parameter is folded into an extra result, an array and
    its length collapse into one slice, and a ``char *`` becomes a ``string``.  Reading
    the C signature instead reconstructs a shape the wrapper does not have.
    """

    c_name: str
    go_name: str
    params: list[tuple[str, str]]   # (Go parameter name, Go type), in order
    returns: list[str]              # Go result types, WITHOUT the trailing error


# Every wrapper this generator emits, by MEOS C name.  ``objectgen.py`` reads it so the
# object layer and the flat layer cannot disagree about a folded out-parameter: the
# object layer never parses a C parameter list, and a function absent from here has no
# wrapper to call.  Filled as each function is emitted, so a caller that emits one
# function gets that one entry and a full ``generate()`` run gets them all.
SIGNATURES: dict[str, Signature] = {}


def _record(c_name: str, go_name: str, go_args: list[str],
            returns: list[str]) -> None:
    """Record the signature just assembled, splitting each ``"name type"`` apart."""
    params = []
    for arg in go_args:
        name, _, go_type = arg.partition(" ")
        params.append((name, go_type))
    SIGNATURES[c_name] = Signature(c_name, go_name, params, list(returns))


# Type mapping ----------------------------------------------------------

@dataclass
class TypeMapping:
    go_type: str           # Go-side parameter / return type
    c_cast: str | None     # ``C.<cast>(x)`` template, ``{}`` is the Go expr
    from_c: str | None     # Convert a C result to Go (``{}`` is the C expr)


# Scalars and well-known opaque pointers.  Pointer types map to the package's
# wrapper structs (``*STBox``, ``*Geom``, etc.) where the wrapper exposes an
# ``Inner()`` method returning the C pointer.
TYPE_MAP: dict[str, TypeMapping] = {
    "void":              TypeMapping("",         None,                                None),
    "bool":              TypeMapping("bool",     "C.bool({})",                        "bool({})"),
    "int":               TypeMapping("int",      "C.int({})",                         "int({})"),
    "int8":              TypeMapping("int8",     "C.int8({})",                        "int8({})"),
    "int16":             TypeMapping("int16",    "C.int16({})",                       "int16({})"),
    "int32":             TypeMapping("int32",    "C.int32({})",                       "int32({})"),
    "int32_t":           TypeMapping("int32",    "C.int32_t({})",                     "int32({})"),
    "int64":             TypeMapping("int64",    "C.int64({})",                       "int64({})"),
    "int64_t":           TypeMapping("int64",    "C.int64_t({})",                     "int64({})"),
    "uint8":             TypeMapping("uint8",    "C.uint8({})",                       "uint8({})"),
    "uint8_t":           TypeMapping("uint8",    "C.uint8_t({})",                     "uint8({})"),
    "uint16":            TypeMapping("uint16",   "C.uint16({})",                      "uint16({})"),
    "uint32":            TypeMapping("uint32",   "C.uint32({})",                      "uint32({})"),
    "uint32_t":          TypeMapping("uint32",   "C.uint32_t({})",                    "uint32({})"),
    "uint64":            TypeMapping("uint64",   "C.uint64({})",                      "uint64({})"),
    "uint64_t":          TypeMapping("uint64",   "C.uint64_t({})",                    "uint64({})"),
    "Quadbin":           TypeMapping("uint64",   "C.Quadbin({})",                     "uint64({})"),
    "double":            TypeMapping("float64",  "C.double({})",                      "float64({})"),
    "size_t":            TypeMapping("uint",     "C.size_t({})",                      "uint({})"),
    "DateADT":           TypeMapping("int32",    "C.DateADT({})",                     "int32({})"),
    "Timestamp":         TypeMapping("int64",    "C.Timestamp({})",                   "int64({})"),
    "TimestampTz":       TypeMapping("int64",    "C.TimestampTz({})",                 "int64({})"),
    "TimeADT":           TypeMapping("int64",    "C.TimeADT({})",                     "int64({})"),
    "TimeOffset":        TypeMapping("int64",    "C.TimeOffset({})",                  "int64({})"),
    # Opaque pointer family.  Every binding that consumes meos-idl.json
    # (PyMEOS-CFFI as `_ffi.CData`, MEOS.NET as `IntPtr`, meos-rs as
    # `*mut c_void`) surfaces these as raw pointers and lets the caller
    # supply a value obtained elsewhere (typically another wrapped
    # function's return).  GoMEOS follows the same convention.
    "void *":                  TypeMapping("unsafe.Pointer", "unsafe.Pointer({})",  "unsafe.Pointer({})"),
    "const void *":            TypeMapping("unsafe.Pointer", "unsafe.Pointer({})",  "unsafe.Pointer({})"),
    # cgo types ``void **`` as ``*unsafe.Pointer``; the Go side is an opaque
    # unsafe.Pointer cast to that at the call.
    "void **":                 TypeMapping("unsafe.Pointer", "(*unsafe.Pointer)(unsafe.Pointer({}))",  "unsafe.Pointer({})"),
    "const void **":           TypeMapping("unsafe.Pointer", "(*unsafe.Pointer)(unsafe.Pointer({}))",  "unsafe.Pointer({})"),
    # Function-pointer typedefs.  MEOS exposes these as opaque handles
    # the caller obtains from another MEOS call (e.g. taking the address
    # of a typed comparator).  Skiplist primitives and the error handler
    # are the only signatures that take one today.
    # Function-pointer typedefs: JMEOS surfaces them as an opaque Pointer.  cgo
    # represents every C function pointer uniformly as ``*[0]byte``, so the Go
    # side is unsafe.Pointer cast to ``*[0]byte`` at the call.
    "datum_func2":             TypeMapping("unsafe.Pointer", "(*[0]byte)({})",       "{}"),
    "error_handler_fn":        TypeMapping("unsafe.Pointer", "(*[0]byte)({})",       "{}"),
    "int (*)(void *, void *)": TypeMapping("unsafe.Pointer", "(*[0]byte)({})",       "{}"),
    "void *(*)(void *, void *)": TypeMapping("unsafe.Pointer", "(*[0]byte)({})",     "{}"),
    # NO ENUM ROW IS WRITTEN HERE: ``configure()`` derives one per catalog enum, so an
    # enum MEOS adds is mapped the day it lands and one it retires stops being mapped.
    "char *":            TypeMapping("string",   "C.CString({})",                     "C.GoString({})"),
    "const char *":      TypeMapping("string",   "C.CString({})",                     "C.GoString({})"),
    # PostgreSQL ``text`` is a varlena envelope around a Go string.  The
    # hand-written ``cstring2text`` / ``text2cstring`` helpers handle the
    # palloc/pfree, mirroring the same idiom used by the hand-written
    # surface.
    # PostgreSQL ``text`` is a varlena envelope around a Go string. MEOS
    # exposes the canonical converters ``cstring_to_text`` / ``text_to_cstring``
    # (public, exported, in the catalog -- the older ``cstring2text`` /
    # ``text2cstring`` spellings are retired): a Go string becomes
    # ``C.cstring_to_text(C.CString(s))`` in and ``C.GoString(C.text_to_cstring(t))``
    # out.
    "text *":            TypeMapping("string",   "C.cstring_to_text(C.CString({}))",  "C.GoString(C.text_to_cstring({}))"),
    "const text *":      TypeMapping("string",   "C.cstring_to_text(C.CString({}))",  "C.GoString(C.text_to_cstring({}))"),
}

# Opaque MEOS struct pointers mapped to GoMEOS wrapper types.  When the
# pointer is an input, the caller passes the wrapper and the codegen emits
# ``.Inner()``.  When it is a return, the codegen wraps the pointer in a new
# instance.  ``Temporal`` is the interface type and uses ``CreateTemporal``.
WRAPPER_TYPES: dict[str, tuple[str, str]] = {
    # C struct name (without trailing space/star) -> (Go wrapper type, ctor expr).
    # The ctor expr uses a literal ``$res`` placeholder for the C result so we
    # do not collide with Python ``str.format`` parsing the embedded braces of
    # Go composite literals.
    # Every MEOS struct is a uniform opaque pointer handle (the meos-rs /
    # JMEOS pattern: the generated FFI never builds an OO interface or a
    # runtime subtype dispatch -- that idiomatic layer is retired). A
    # Temporal* return is just a *Temporal handle; the caller passes it back
    # to another MEOS call via Inner().
    "Temporal":          ("*Temporal",         "&Temporal{_inner: $res}"),
    "TInstant":          ("*TInstant",         "&TInstant{_inner: $res}"),
    "TSequence":         ("*TSequence",        "&TSequence{_inner: $res}"),
    "TSequenceSet":      ("*TSequenceSet",     "&TSequenceSet{_inner: $res}"),
    "STBox":             ("*STBox",            "&STBox{_inner: $res}"),
    "TBox":              ("*TBox",             "&TBox{_inner: $res}"),
    "Span":              ("*Span",             "&Span{_inner: $res}"),
    "SpanSet":           ("*SpanSet",          "&SpanSet{_inner: $res}"),
    "Set":               ("*Set",              "&Set{_inner: $res}"),
    "GSERIALIZED":       ("*Geom",             "&Geom{_inner: $res}"),
    "Interval":          ("*Interval",         "&Interval{_inner: $res}"),
    "Npoint":            ("*Npoint",           "&Npoint{_inner: $res}"),
    "Nsegment":          ("*Nsegment",         "&Nsegment{_inner: $res}"),
    "Cbuffer":           ("*Cbuffer",          "&Cbuffer{_inner: $res}"),
    "Pose":              ("*Pose",             "&Pose{_inner: $res}"),
    "Rgeo":              ("*Rgeo",             "&Rgeo{_inner: $res}"),
    "Jsonb":             ("*Jsonb",            "&Jsonb{_inner: $res}"),
    "JsonPath":          ("*JsonPath",         "&JsonPath{_inner: $res}"),
    "Pcpoint":           ("*Pcpoint",          "&Pcpoint{_inner: $res}"),
    "Pcpatch":           ("*Pcpatch",          "&Pcpatch{_inner: $res}"),
    "TPCBox":            ("*TPCBox",           "&TPCBox{_inner: $res}"),
    "MeosArray":         ("*MeosArray",        "&MeosArray{_inner: $res}"),
    "PCSCHEMA":          ("*PCSchema",         "&PCSchema{_inner: $res}"),
    "SkipList":          ("*SkipList",         "&SkipList{_inner: $res}"),
    "RTree":             ("*RTree",            "&RTree{_inner: $res}"),
    "Match":             ("*Match",            "&Match{_inner: $res}"),
    "BOX3D":             ("*Box3D",            "&Box3D{_inner: $res}"),
    "GBOX":              ("*GBox",             "&GBox{_inner: $res}"),
    "AFFINE":            ("*AFFINE",           "&AFFINE{_inner: $res}"),
    "PJ_CONTEXT":        ("*PJContext",        "&PJContext{_inner: $res}"),
    "gsl_rng":           ("*GslRng",           "&GslRng{_inner: $res}"),
}


# CODEGEN-REGULAR-EXCEPTION: ENUM_GO_NAME is a NAME MAP, not an exclusion.
# emit_types() iterates EVERY catalog enum and emits a Go type for each, and
# configure() maps every one of them for parameters and results; this dict only
# fixes the Go SPELLING where PascalCase would not produce the name the package
# already publishes -- `interpType` reaches Go as `Interpolation`, which no
# mechanical rule yields and which is a published name
# ([[never-change-api-surface-autonomously]]). Every other enum falls back to
# `_go_name(c_name)`. Zero enums are dropped or left unclassified.
ENUM_GO_NAME: dict[str, str] = {
    "interpType": "Interpolation",
    "MeosType": "MeosType",
    "MeosOper": "MeosOper",
    "nullHandleType": "NullHandleType",
    "tempSubtype": "TempSubtype",
    "SkipListType": "SkipListType",
}

# The catalog's enums, filled in by ``configure``. An enum crosses cgo as its own
# Go type, and taking the SET from the catalog is what keeps an enum MEOS adds
# from arriving unmapped the way a hand list leaves it -- the shape MEOS.NET's
# tools/codegen.py states at its own ENUM_TYPES.
ENUM_TYPES: set[str] = set()


def configure(idl: dict) -> None:
    """Take from the catalog the type facts the mapping above reads.

    A HAND LIST GOES SHORT AND ALSO GOES STALE, in both directions at once: the
    rows this replaces named 6 of the 14 enums the catalog declares -- leaving
    `IndexSearchOp` (9 functions), `SPTreeKind` (8) and `MeosPixType` (4)
    unmapped, so 21 wrappers emitted a TODO stub reading `unsupported param` --
    and carried 2 more, `RTreeSearchOp` and `errorLevel`, for enums the catalog
    no longer declares at all and no function names.

    The Go spelling comes from ENUM_GO_NAME where the published name is not what
    PascalCase yields, so this and ``emit_types`` cannot disagree about a name:
    both read the same helper.
    """
    ENUM_TYPES.clear()
    ENUM_TYPES.update(e["name"] for e in idl.get("enums", []) if e.get("name"))
    for c_name in ENUM_TYPES:
        go_name = ENUM_GO_NAME.get(c_name, _go_name(c_name))
        TYPE_MAP[c_name] = TypeMapping(
            go_name, f"C.{c_name}({{}})", f"{go_name}({{}})")


def emit_types(idl: dict, corpus: str) -> str:
    """Emit the type layer the flat wrappers depend on: one opaque handle
    struct per MEOS wrapper type and one Go type + const block per catalog
    enum.  This is the meos-rs / JMEOS opaque-handle pattern (no OO interface,
    no runtime subtype dispatch) expressed in Go: every MEOS struct is a
    ``*C.<T>`` behind an ``Inner()`` accessor, and the caller only ever passes
    the handle back to another MEOS call.

    ``corpus`` is the full body of the emitted wrappers.  A wrapper type or
    enum is emitted ONLY if it is referenced there -- so a WRAPPER_TYPES entry
    no live function uses (no matching ``*C.<T>`` exists) and an enum whose
    family header is out of scope (e.g. the raster ``MeosPixType``) are dropped
    rather than declaring a Go alias to a C type cgo cannot see."""
    def referenced(go_name: str) -> bool:
        return re.search(rf"\b{re.escape(go_name)}\b", corpus) is not None

    out: list[str] = []

    # -- opaque wrapper structs (one per distinct referenced wrapper type) ---
    out.append("// -------------------- opaque handle types --------------------")
    seen: set[str] = set()
    for c_name, (go_type, _ctor) in WRAPPER_TYPES.items():
        go_name = go_type.lstrip("*")
        if go_name in seen or not referenced(go_name):
            continue
        seen.add(go_name)
        out.append(f"type {go_name} struct {{ _inner *C.{c_name} }}")
        out.append(
            f"func (x *{go_name}) Inner() *C.{c_name} {{ "
            f"if x == nil {{ return nil }}; return x._inner }}"
        )
        # A cgo type belongs to the package that declares it, so this package's
        # `*C.{c_name}` and a caller's own `*C.{c_name}` are DIFFERENT Go types
        # and Inner() cannot cross a package boundary. The handle is an opaque
        # MEOS pointer either way, so the pointer itself is the boundary -- the
        # same contract MEOS.NET's generated handles carry as IntPtr. A nil
        # handle yields a nil pointer and a nil pointer yields a nil handle, so
        # a MEOS entry answering NULL cannot become a dereference here.
        out.append(
            f"func (x *{go_name}) Pointer() unsafe.Pointer {{ "
            f"if x == nil {{ return nil }}; return unsafe.Pointer(x._inner) }}"
        )
        out.append(
            f"func {go_name}FromPointer(p unsafe.Pointer) *{go_name} {{ "
            f"if p == nil {{ return nil }}; "
            f"return &{go_name}{{_inner: (*C.{c_name})(p)}} }}"
        )

    # -- enums: Go type aliased to the C enum + a numeric const block --------
    out.append("\n// -------------------- enums --------------------")
    for e in idl.get("enums", []):
        c_name = e["name"]
        go_name = ENUM_GO_NAME.get(c_name, _go_name(c_name))
        if not referenced(go_name):
            continue
        out.append(f"type {go_name} C.{c_name}")
        values = e.get("values", [])
        if values:
            out.append("const (")
            for v in values:
                out.append(f"\t{go_name}_{v['name']} {go_name} = {int(v['value'])}")
            out.append(")")

    header = _PER_HEADER_PREAMBLE + "\n".join(out) + "\n"
    return header


def _references_opaque(entry: dict) -> bool:
    if any(t in entry["returnType"]["c"] for t in OPAQUE_TYPES):
        return True
    return any(any(t in p["cType"] for t in OPAQUE_TYPES) for p in entry["params"])


def _is_datum_internal(entry: dict) -> bool:
    """Functions that take or return a raw Datum are MEOS-internal helpers
    that the hand-written surface re-exposes through typed overloads.  The
    codegen cannot produce those overloads automatically and reporting them
    as TODO would conflate metadata work with real shape gaps."""
    if "Datum" in entry["returnType"]["c"].split():
        return True
    for p in entry["params"]:
        if "Datum" in p["cType"].split():
            return True
    return False


def _strip_qualifiers(c_type: str) -> tuple[str, int]:
    """Return ``(base_type, pointer_level)`` stripped of ``const`` and ``*``."""
    s = c_type.replace("const ", "").strip()
    stars = s.count("*")
    return s.replace("*", "").strip(), stars


def _go_type_for(c_type: str) -> tuple[str | None, str | None, str | None]:
    """Look up a C type in the mapping tables.

    Returns ``(go_type, c_cast, from_c)``; any field may be ``None`` when the
    type is not yet handled.
    """
    if c_type in TYPE_MAP:
        m = TYPE_MAP[c_type]
        # Normalise scalar/string templates to the same ``$x`` placeholder
        # the wrapper-type entries use, so emit_function only deals with one
        # substitution dialect.
        c_cast = m.c_cast.replace("{}", "$x") if m.c_cast else None
        from_c = m.from_c.replace("{}", "$x") if m.from_c else None
        return m.go_type, c_cast, from_c
    base, stars = _strip_qualifiers(c_type)
    if stars == 1 and base in WRAPPER_TYPES:
        go_type, ctor = WRAPPER_TYPES[base]
        # Pass through the wrapper as an input; convert back as a return.
        c_cast = "$x.Inner()" if not go_type.startswith("*") else "$x._inner"
        return go_type, c_cast, ctor.replace("$res", "$x")
    # JMEOS-flat fallback: any remaining pointer (an array ``T **``, a scalar
    # out-pointer, a ``void *``) is an opaque handle -- JMEOS maps every such
    # pointer to a raw ``Pointer`` and never unpacks it (the array->collection
    # conversion is a separate layer).  cgo needs the concrete pointer type at
    # the boundary, so the Go side is ``unsafe.Pointer`` and the call casts to
    # the exact C pointer type.
    if stars >= 1:
        if base == "void":
            # cgo maps ``void *`` -> unsafe.Pointer, ``void **`` ->
            # *unsafe.Pointer, etc.  A bare ``void *`` needs no cast.
            if stars == 1:
                return "unsafe.Pointer", "unsafe.Pointer($x)", "unsafe.Pointer($x)"
            cgotype = "*" * (stars - 1) + "unsafe.Pointer"
            return "unsafe.Pointer", f"({cgotype})(unsafe.Pointer($x))", "unsafe.Pointer($x)"
        cgotype = "*" * stars + f"C.{base}"
        return "unsafe.Pointer", f"({cgotype})(unsafe.Pointer($x))", "unsafe.Pointer($x)"
    return None, None, None


# Name conversion -------------------------------------------------------

def _go_name(c_name: str) -> str:
    """Convert ``snake_case`` to ``PascalCase`` while keeping initialisms."""
    parts = c_name.split("_")
    out = []
    for p in parts:
        if not p:
            continue
        # Preserve common acronyms in upper case for readability.
        if p.upper() in {"WKB", "WKT", "MFJSON", "JSON", "SRID", "EWKT", "EWKB",
                         "STBOX", "TBOX", "ID", "X", "Y", "Z", "T", "MFJ"}:
            out.append(p.upper())
        else:
            out.append(p[:1].upper() + p[1:])
    return "".join(out) or c_name


_GO_RESERVED = {"type", "func", "interface", "select", "case", "chan", "goto",
                "package", "import", "go", "defer", "return", "range", "var",
                "const", "for", "if", "else", "switch", "break", "continue",
                "default", "fallthrough", "map", "struct", "string"}


def _go_param_name(c_name: str) -> str:
    if c_name in _GO_RESERVED:
        return c_name + "_"
    return c_name


def _c_callee(c_name: str) -> str:
    """The identifier used at the ``C.<name>`` call site.

    cgo parses ``C.union_x`` as the C ``union x`` TYPE, not the function
    ``union_x`` -- so every MEOS ``union_*`` set/span operator is reached
    through a ``#define g<name> <name>`` alias emitted in the cgo preamble
    (see ``_union_alias_block``), mirroring the hand-written cast.h ``gunion_``
    macros."""
    if c_name.startswith("union_"):
        return "g" + c_name
    return c_name


def _union_alias_block(idl: dict) -> str:
    names = sorted(f["name"] for f in idl["functions"]
                   if f["name"].startswith("union_"))
    if not names:
        return ""
    defines = "\n".join(f"#define g{n} {n}" for n in names)
    # A ``//`` line comment (not ``/* */``) -- this block is spliced INSIDE the
    # cgo preamble's C comment, and C comments do not nest.
    return ("\n// cgo reads C.union_* as a union TYPE; alias the union "
            "operators so C.g<name> resolves them as functions.\n" + defines)


# Parameter classification ---------------------------------------------

# Param-name heuristics for the MEOS C convention: these names, paired with
# the right pointer level, signal an out-parameter or a length-of-array
# companion to the preceding pointer-of-pointers input.
_OUT_NAMES = {"result", "value"}
_LENGTH_NAMES = {"count", "size", "ngeoms", "wkb_size"}
_SIZE_OUT_NAMES = {"size_out", "size", "len"}


def _classify_param(p: dict, params: list[dict], i: int) -> str:
    """Return one of ``INPUT`` / ``OUTPUT_SCALAR`` / ``OUTPUT_COUNT`` /
    ``OUTPUT_SIZE`` / ``ARRAY_INPUT`` / ``ARRAY_LENGTH``.

    The classification follows the MEOS C convention: parameters named
    ``count`` / ``size`` paired with a preceding double-pointer or array
    pointer are array lengths; pointer parameters named ``result`` /
    ``value`` are out-parameters; ``int *count`` / ``size_t *size_out`` at
    the tail are count outputs."""
    name = p["name"]
    ctype = p["cType"]
    base, stars = _strip_qualifiers(ctype)

    # ``T *result`` / ``T *value`` with a concrete scalar/wrapper base -> a
    # hidden out-parameter the wrapper allocates and returns (JMEOS hides
    # ``result`` the same way).  A ``void *`` is never a typed output we can
    # allocate -- it is an opaque INPUT (e.g. meos_array_add's ``void *value``).
    if stars >= 1 and name in _OUT_NAMES and base != "void":
        return "OUTPUT_SCALAR"

    # Everything else -- including array pointers ``T **`` and their ``count``
    # companions -- is a flat opaque INPUT.  Mirroring JMEOS, the generated FFI
    # does NOT unpack arrays into Go slices or return the count as an extra
    # value; the caller passes the array/count pointers through as opaque
    # handles (a higher layer does the collection conversion).
    return "INPUT"


# Emission --------------------------------------------------------------

@dataclass
class EmittedFunc:
    name: str          # Go-side name
    code: str          # full Go source for the function
    skipped: bool      # true if emitted as a TODO stub


def _array_return_info(return_c: str):
    """Return ``(base_type, ptr_level)`` for ``T **`` / ``T *`` returns that
    pair with an OUTPUT_COUNT param, or ``None`` if the return is not an
    array shape."""
    base, stars = _strip_qualifiers(return_c)
    if stars in (1, 2):
        return base, stars
    return None


def _go_slice_elem(base: str, stars: int) -> tuple[str, str] | None:
    """For an array return ``base`` + ``stars``, return ``(elem_go_type,
    elem_ctor_template)`` where the template uses ``$x`` for the C element.

    The element type for a ``T **`` return is the wrapped pointer (``*Foo``);
    for a ``T *`` return of scalar elements it is the Go primitive."""
    if stars == 2 and base in WRAPPER_TYPES:
        go_type, ctor = WRAPPER_TYPES[base]
        return go_type, ctor.replace("$res", "$x")
    if stars == 2 and base == "char":
        # ``char **`` -> []string; each element is a C string.
        return "string", "C.GoString($x)"
    if stars == 2 and base == "text":
        # ``text **`` -> []string via text2cstring per element.
        return "string", "C.GoString(C.text_to_cstring($x))"
    if stars == 1:
        scalar = TYPE_MAP.get(base)
        if scalar is not None:
            # Normalise the ``{}`` placeholder used inside TYPE_MAP to the
            # ``$x`` placeholder the emitter substitutes per element.
            ctor = (scalar.from_c or "$x").replace("{}", "$x")
            return scalar.go_type, ctor
        if base == "uint8_t":
            # ``uint8_t *`` byte buffer return -> []byte handled separately.
            return None
    return None


def _emit_array_input_group(entry: dict, group: dict) -> EmittedFunc:
    """Emit a wrapper for functions that take N parallel input arrays sharing
    one count parameter (e.g. ``tpointseq_make_coords``).  Each array param
    becomes a Go slice; nullable members accept ``nil`` for the empty case."""
    c_name = entry["name"]
    go_name = _go_name(c_name)
    return_c = entry["returnType"]["c"]
    params = entry["params"]
    array_params = set(group["params"])
    count_param = group["count"]
    nullable = set(group.get("nullable", []))

    ret_go, _, ret_from_c = _go_type_for(return_c)
    if ret_go is None and return_c != "void":
        return EmittedFunc(go_name, _todo_stub(c_name, "arrayInputGroup return shape " + return_c), True)

    go_args, inner_args, deferred = [], [], []
    primary = None
    for p in params:
        pname = _go_param_name(p["name"])
        ptype = p["cType"]
        base, stars = _strip_qualifiers(ptype)
        if p["name"] in array_params:
            primary = primary or pname
            go_t = TYPE_MAP[base].go_type if base in TYPE_MAP else None
            if go_t is None:
                return EmittedFunc(go_name, _todo_stub(c_name, "arrayInputGroup element " + base), True)
            go_args.append(f"{pname} []{go_t}")
            local = f"_c_{pname}"
            if p["name"] in nullable:
                deferred.append(f"var {local} []C.{base}")
                deferred.append(f"if {pname} != nil {{")
                deferred.append(f"\t{local} = make([]C.{base}, len({pname}))")
                deferred.append(f"\tfor _i, _v := range {pname} {{ {local}[_i] = C.{base}(_v) }}")
                deferred.append(f"}}")
                inner_args.append(f"_ptr_or_nil_{base}({local})")
            else:
                deferred.append(f"{local} := make([]C.{base}, len({pname}))")
                deferred.append(f"for _i, _v := range {pname} {{ {local}[_i] = C.{base}(_v) }}")
                inner_args.append(f"&{local}[0]")
        elif p["name"] == count_param:
            base, _ = _strip_qualifiers(ptype)
            inner_args.append(f"C.{base}(len({primary}))")
        else:
            go_t, c_cast, _ = _go_type_for(ptype)
            if go_t is None:
                return EmittedFunc(go_name, _todo_stub(c_name, "arrayInputGroup other param " + ptype), True)
            go_args.append(f"{pname} {go_t}")
            inner_args.append(c_cast.replace("$x", pname) if c_cast else pname)

    sig_args = ", ".join(go_args)
    call = f"C.{_c_callee(c_name)}({', '.join(inner_args)})"
    body = "\n".join("\t" + d for d in deferred)
    ret_sig = _result_signature([] if return_c == "void" else [ret_go])
    if return_c == "void":
        tail = f"{_ERRNO_RESET}\n\t{call}\n{_ERRNO_GUARD}\n\treturn"
    else:
        tail = (f"{_ERRNO_RESET}\n\t_cret := {call}\n{_ERRNO_GUARD}"
                f"{_free_returned_string(return_c)}\n"
                f"\treturn {ret_from_c.replace('$x', '_cret')}, nil")
    code = (
        f"// {go_name} wraps MEOS C function {c_name}.\n"
        f"func {go_name}({sig_args}){ret_sig} {{\n"
        f"{body}\n{tail}\n}}\n"
    )
    _record(c_name, go_name, go_args, [] if return_c == "void" else [ret_go])
    return EmittedFunc(go_name, code, False)


def emit_function(entry: dict) -> EmittedFunc:
    c_name = entry["name"]
    go_name = _go_name(c_name)
    return_c = entry["returnType"]["c"]
    params = entry["params"]
    shape = entry.get("shape", {})

    # ``shape.skip`` (from meta/meos-meta.json) marks functions the codegen
    # is told to omit entirely (function-pointer args, void ** internals).
    if "skip" in shape:
        return EmittedFunc(go_name, _todo_stub(c_name, "skip: " + shape["skip"]), True)

    # Metadata can mark scalar parameters as outputs even when their name is
    # not the canonical ``result`` / ``value``.  Carry the set down to
    # ``_classify_param`` via a closure-friendly local override.
    named_outputs = set(shape.get("namedOutputs", []))

    # ``shape.arrayInputGroup`` describes N parallel input arrays that share
    # one count.  We hand-roll a separate emission path for these because
    # the heuristic loop expects a single (array, count) pair.
    if "arrayInputGroup" in shape:
        return _emit_array_input_group(entry, shape["arrayInputGroup"])

    # ``shape.outputArrays`` declares parallel out-parameters sharing the
    # primary length.  Hold the metadata for the param walk to recognise
    # them and for the body-assembly pass to unpack them after the call.
    # JMEOS-flat: output arrays are passed through as opaque pointer params,
    # not unpacked into Go slices, so no declared-output-array set is used.
    declared_output_arrays: dict = {}

    # Pre-classify every parameter so the input loop can skip ones that
    # become outputs or array-length companions.
    def classify_one(p, i):
        if p["name"] in named_outputs:
            return "OUTPUT_SCALAR"
        return _classify_param(p, params, i)
    classes = [classify_one(p, i) for i, p in enumerate(params)]

    # Identify a paired array return.  Three shapes, in priority order:
    #  (m) shape.arrayReturn metadata -> length from sibling accessor
    #  (a) explicit OUTPUT_COUNT param -> slice length comes from C
    #  (b) an ARRAY_LENGTH input -> slice length matches the input slice
    has_out_count = "OUTPUT_COUNT" in classes
    has_in_length = "ARRAY_LENGTH" in classes
    is_array_return = False
    array_length_source: str | None = None  # ``c_count`` or ``len(input_slice)``
    array_elem_go = array_elem_ctor = None

    # JMEOS-flat: an array return is a raw opaque pointer, never a length-driven
    # Go slice, so the accessor-length metadata path stays disabled.
    array_meta = None
    if array_meta and array_meta.get("kind") == "accessor" and return_c != "void":
        info = _array_return_info(return_c)
        if info is not None:
            base, stars = info
            elem = _go_slice_elem(base, stars)
            if elem is None and stars == 2 and base == "text":
                # ``text **`` array returns lower to []string.
                elem = ("string", "C.GoString(C.text_to_cstring($x))")
            if elem is not None:
                is_array_return = True
                array_elem_go, array_elem_ctor = elem
                accessor_func = array_meta["func"]
                accessor_arg = _go_param_name(array_meta["arg"])
                cast_to = array_meta.get("castTo")
                arg_expr = f"{accessor_arg}.Inner()" if cast_to else f"{accessor_arg}.Inner()"
                # When castTo is set, force the accessor's first param to that
                # type via cgo's unsafe.Pointer ladder.  Otherwise pass the
                # wrapper's Inner() directly.
                if cast_to:
                    cast_base = _strip_qualifiers(cast_to)[0]
                    arg_expr = f"(*C.{cast_base})(unsafe.Pointer({accessor_arg}.Inner()))"
                array_length_source = f"int(C.{accessor_func}({arg_expr}))"
    if (has_out_count or has_in_length) and return_c not in ("void",):
        info = _array_return_info(return_c)
        if info is not None:
            base, stars = info
            elem = _go_slice_elem(base, stars)
            if elem is not None:
                is_array_return = True
                array_elem_go, array_elem_ctor = elem
            elif stars == 1 and base == "uint8_t":
                # WKB-style ``uint8_t *`` paired with ``size_t *size_out``.
                is_array_return = True
                array_elem_go = "byte"
                array_elem_ctor = "byte($x)"
        if is_array_return and not has_out_count:
            # Length is taken from the input slice ``len(...)``; find the
            # first ARRAY_LENGTH index and read the preceding param name.
            length_idx = next(i for i, c in enumerate(classes) if c == "ARRAY_LENGTH")
            array_length_source = f"len({_go_param_name(params[length_idx - 1]['name'])})"

    # Resolve the return type when not an array shape.
    if is_array_return:
        ret_go = f"[]{array_elem_go}"
        ret_from_c = None
    else:
        ret_go, _, ret_from_c = _go_type_for(return_c)
        if ret_go is None and return_c != "void":
            return EmittedFunc(go_name, _todo_stub(c_name, "unsupported return type " + return_c), True)

    # Walk parameters, building Go-side signature args and C-side call args.
    go_args = []
    inner_args = []
    deferred = []
    extra_returns: list[tuple[str, str, str]] = []  # (go_type, var_name, c_to_go expr)
    array_length_for: list[int] = []  # indices of ARRAY_LENGTH params (skipped from sig)
    array_count_local: str | None = None  # the OUTPUT_COUNT local for array returns

    for i, p in enumerate(params):
        pname = _go_param_name(p["name"])
        ptype = p["cType"]
        cls = classes[i]

        if cls == "ARRAY_LENGTH":
            # Length is derived from the preceding slice's len(); locate it.
            prev_name = _go_param_name(params[i - 1]["name"])
            base, _ = _strip_qualifiers(ptype)
            inner_args.append(f"C.{base}(len({prev_name}))")
            array_length_for.append(i - 1)
            continue

        if cls == "OUTPUT_COUNT":
            base, _ = _strip_qualifiers(ptype)
            local = f"_out_{pname}"
            deferred.append(f"var {local} C.{base}")
            inner_args.append(f"&{local}")
            # If this count is paired with an array return, it is not
            # surfaced as a separate Go return -- it controls the slice
            # length.
            if is_array_return:
                array_count_local = local
            else:
                extra_returns.append(
                    ("int" if base == "int" else "uint", local,
                     f"int({local})" if base == "int" else f"uint({local})"))
            continue

        if cls == "OUTPUT_SCALAR":
            base, stars = _strip_qualifiers(ptype)
            # Wrapped opaque output (e.g. ``GSERIALIZED ** result``).
            if stars == 2 and base in WRAPPER_TYPES:
                go_type, ctor = WRAPPER_TYPES[base]
                local = f"_out_{pname}"
                deferred.append(f"var {local} *C.{base}")
                inner_args.append(f"&{local}")
                extra_returns.append((go_type, local, ctor.replace("$res", local)))
                continue
            # text** -> string
            if stars == 2 and base == "text":
                local = f"_out_{pname}"
                deferred.append(f"var {local} *C.text")
                inner_args.append(f"&{local}")
                extra_returns.append(("string", local, f"C.GoString(C.text_to_cstring({local}))"))
                continue
            # Scalar slice output (``TimestampTz **time_bins`` and friends)
            # is left as TODO: the parallel output array needs post-call
            # unsafe.Slice + element conversion alongside the primary array
            # return, and that wiring is not yet plumbed through the
            # body-assembly pass.  Affects the ``*_split`` and
            # ``tpoint_as_mvtgeom`` families.
            # Wrapped opaque single-pointer output (``Span *result`` etc.):
            # MEOS writes into a caller-allocated struct; allocate one and
            # surface the pointer back as the typed wrapper.
            if stars == 1 and base in WRAPPER_TYPES:
                go_type, ctor = WRAPPER_TYPES[base]
                local = f"_out_{pname}"
                deferred.append(f"var {local} C.{base}")
                inner_args.append(f"&{local}")
                extra_returns.append((go_type, local, ctor.replace("$res", f"&{local}")))
                continue
            # Scalar output (`double * result`, `bool * value`, etc.)
            if stars == 1 and base in TYPE_MAP:
                go_type = TYPE_MAP[base].go_type
                local = f"_out_{pname}"
                deferred.append(f"var {local} C.{base}")
                inner_args.append(f"&{local}")
                # TYPE_MAP entries use the legacy ``{}`` placeholder; the
                # emitter's substitution dialect is ``$x``.
                from_c = (TYPE_MAP[base].from_c or "$x").replace("{}", "$x")
                extra_returns.append((go_type, local, from_c.replace("$x", local)))
                continue
            return EmittedFunc(go_name, _todo_stub(c_name, "unhandled OUTPUT_SCALAR shape " + ptype), True)

        # Metadata-declared output array (e.g. ``time_bins TimestampTz **``,
        # ``space_bins GSERIALIZED ***``).  Allocate the appropriately-typed
        # local; the body-assembly pass below converts to a Go slice using
        # the primary length.
        if p["name"] in declared_output_arrays:
            base, stars = _strip_qualifiers(ptype)
            local = f"_out_{pname}"
            if stars == 2 and base in TYPE_MAP and base != "text":
                deferred.append(f"var {local} *C.{base}")
                inner_args.append(f"&{local}")
                extra_returns.append((
                    f"[]{TYPE_MAP[base].go_type}",
                    local,
                    f"@SLICE_SCALAR:{local}:{base}",
                ))
                continue
            if stars == 3 and base in WRAPPER_TYPES:
                go_elem = WRAPPER_TYPES[base][0]
                deferred.append(f"var {local} **C.{base}")
                inner_args.append(f"&{local}")
                extra_returns.append((
                    f"[]{go_elem}",
                    local,
                    f"@SLICE_WRAPPED:{local}:{base}",
                ))
                continue
            if stars == 2 and base in WRAPPER_TYPES:
                go_elem = WRAPPER_TYPES[base][0]
                deferred.append(f"var {local} *C.{base}")
                inner_args.append(f"&{local}")
                extra_returns.append((
                    f"[]{go_elem}",
                    local,
                    f"@SLICE_WRAPPED_2:{local}:{base}",
                ))
                continue

        # INPUT path -------------------------------------------------------
        base, stars = _strip_qualifiers(ptype)

        # Counted-array input: ``T ** xs`` paired with the next ``int
        # count`` (already marked ARRAY_LENGTH).  Lowered to a Go slice of
        # the wrapped element.
        if stars == 2 and base in WRAPPER_TYPES and i + 1 < len(params) and classes[i + 1] == "ARRAY_LENGTH":
            elem_go, _ = WRAPPER_TYPES[base]
            slice_go = f"[]{elem_go}"
            local = f"_c_{pname}"
            deferred.append(f"{local} := make([]*C.{base}, len({pname}))")
            deferred.append(f"for _i, _v := range {pname} {{ {local}[_i] = _v._inner }}")
            ptr_expr = f"(**C.{base})(unsafe.Pointer(&{local}[0]))" if base != "uint8_t" else f"&{local}[0]"
            go_args.append(f"{pname} {slice_go}")
            inner_args.append(ptr_expr)
            continue

        # text** array input + count -> []string (textset_make and family).
        if stars == 2 and base == "text" and i + 1 < len(params) and classes[i + 1] == "ARRAY_LENGTH":
            local = f"_c_{pname}"
            deferred.append(f"{local} := make([]*C.text, len({pname}))")
            deferred.append(f"for _i, _v := range {pname} {{ {local}[_i] = C.cstring_to_text(C.CString(_v)) }}")
            go_args.append(f"{pname} []string")
            inner_args.append(f"(**C.text)(unsafe.Pointer(&{local}[0]))")
            continue

        # Scalar slice input: ``const int *values`` etc. paired with count.
        if stars == 1 and "const" in ptype and base in {"int", "int64", "double", "DateADT", "TimestampTz"} and i + 1 < len(params) and classes[i + 1] == "ARRAY_LENGTH":
            slice_go = f"[]{TYPE_MAP[base].go_type}"
            local = f"_c_{pname}"
            deferred.append(f"{local} := make([]C.{base}, len({pname}))")
            deferred.append(f"for _i, _v := range {pname} {{ {local}[_i] = C.{base}(_v) }}")
            go_args.append(f"{pname} {slice_go}")
            inner_args.append(f"&{local}[0]")
            continue

        # Byte-buffer input: ``const uint8_t *wkb`` + ``size_t size``.
        if stars == 1 and "const" in ptype and base == "uint8_t" and i + 1 < len(params) and classes[i + 1] == "ARRAY_LENGTH":
            go_args.append(f"{pname} []byte")
            inner_args.append(f"(*C.uint8_t)(unsafe.Pointer(&{pname}[0]))")
            continue

        go_t, c_cast, _ = _go_type_for(ptype)
        if go_t is None:
            return EmittedFunc(go_name, _todo_stub(c_name, "unsupported param " + ptype), True)
        go_args.append(f"{pname} {go_t}")
        if c_cast is None:
            inner_args.append(pname)
        else:
            cast_expr = c_cast.replace("$x", pname)
            if go_t == "string":
                local = f"_c_{pname}"
                deferred.append(f"{local} := {cast_expr}")
                deferred.append(f"defer C.free(unsafe.Pointer({local}))")
                inner_args.append(local)
            else:
                inner_args.append(cast_expr)

    # Assemble Go signature ----------------------------------------------
    sig_args = ", ".join(go_args)
    all_returns: list[str] = []
    if is_array_return:
        all_returns.append(ret_go)
    elif return_c != "void":
        all_returns.append(ret_go)
    for go_type, _, _ in extra_returns:
        all_returns.append(go_type)

    ret_sig = _result_signature(all_returns)

    # Assemble Go body ----------------------------------------------------
    body_lines = []
    body_lines.extend("\t" + d for d in deferred)
    body_lines.append(_ERRNO_RESET)
    call = f"C.{_c_callee(c_name)}({', '.join(inner_args)})"

    if is_array_return:
        # Slice length resolution order:
        #  1. shape.arrayReturn metadata (array_length_source pre-set)
        #  2. OUTPUT_COUNT C local captured during param walk
        #  3. len() of an ARRAY_LENGTH input slice
        if array_count_local:
            count_expr = f"int({array_count_local})"
        else:
            count_expr = array_length_source
        body_lines.append(f"\tres := {call}")
        body_lines.append(_ERRNO_GUARD)
        body_lines.append(f"\t_n := {count_expr}")
        if array_elem_go == "byte":
            # WKB byte buffer; cast (*C.uint8_t -> []byte via unsafe.Slice).
            body_lines.append("\t_slice := unsafe.Slice((*byte)(unsafe.Pointer(res)), _n)")
            body_lines.append("\t_out := make([]byte, _n)")
            body_lines.append("\tcopy(_out, _slice)")
            body_lines.append("\treturn _out, nil")
        else:
            base, stars = _array_return_info(return_c)
            elem_c = f"*C.{base}" if stars == 2 else f"C.{base}"
            body_lines.append(f"\t_slice := unsafe.Slice((*{elem_c})(unsafe.Pointer(res)), _n)")
            body_lines.append(f"\t_out := make({ret_go}, _n)")
            body_lines.append("\tfor _i, _e := range _slice {")
            body_lines.append(f"\t\t_out[_i] = {array_elem_ctor.replace('$x', '_e')}")
            body_lines.append("\t}")
            body_lines.append("\treturn _out, nil")
    elif return_c == "void":
        if extra_returns:
            body_lines.append(f"\t{call}")
            body_lines.append(_ERRNO_GUARD)
            tail = ", ".join(c2g for _, _, c2g in extra_returns)
            tail = _expand_slice_markers(tail, body_lines, array_count_local, array_length_source)
            body_lines.append(f"\treturn {tail}, nil")
        else:
            body_lines.append(f"\t{call}")
            body_lines.append(_ERRNO_GUARD)
            body_lines.append("\treturn nil")
    else:
        # ``_cret`` (not ``res``) so the result local never collides with a C
        # parameter named ``res`` (e.g. h3_uncompact_cells' resolution arg).
        body_lines.append(f"\t_cret := {call}")
        body_lines.append(_ERRNO_GUARD)
        free_cret = _free_returned_string(return_c)
        if free_cret:
            body_lines.append(free_cret.lstrip("\n"))
        return_expr = ret_from_c.replace("$x", "_cret")
        if extra_returns:
            tail = ", ".join(c2g for _, _, c2g in extra_returns)
            tail = _expand_slice_markers(tail, body_lines, array_count_local, array_length_source)
            body_lines.append(f"\treturn {return_expr}, {tail}, nil")
        else:
            body_lines.append(f"\treturn {return_expr}, nil")

    code = (
        f"// {go_name} wraps MEOS C function {c_name}.\n"
        f"func {go_name}({sig_args}){ret_sig} {{\n"
        + "\n".join(body_lines)
        + "\n}\n"
    )
    _record(c_name, go_name, go_args, all_returns)
    return EmittedFunc(go_name, code, False)


def _expand_slice_markers(tail: str, body_lines: list[str], count_local: str | None,
                          array_length_source: str | None) -> str:
    """Resolve ``@SLICE_*`` placeholders left by output-array emission.

    The marker carries everything we need (kind, C local, base type) to
    synthesise both the unpacking statements and the final Go expression
    used in the return tuple.  Statements are appended to ``body_lines``;
    the substituted expression replaces the marker in the tail string."""
    if "@SLICE_" not in tail:
        return tail
    length_expr = f"int({count_local})" if count_local else array_length_source

    def _resolve(marker: str) -> str:
        kind, local, base = marker.split(":", 2)
        out_local = f"{local}_go"
        if kind == "@SLICE_SCALAR":
            elem_c = f"C.{base}"
            elem_go = TYPE_MAP[base].go_type
            from_c = (TYPE_MAP[base].from_c or "$x").replace("{}", "$x").replace("$x", "_e")
            body_lines.append(f"\t_slice_{local} := unsafe.Slice({local}, {length_expr})")
            body_lines.append(f"\t{out_local} := make([]{elem_go}, {length_expr})")
            body_lines.append(f"\tfor _i, _e := range _slice_{local} {{ {out_local}[_i] = {from_c} }}")
            return out_local
        if kind == "@SLICE_WRAPPED_2":
            go_elem, ctor = WRAPPER_TYPES[base]
            ctor_expr = ctor.replace("$res", "_e")
            elem_c = f"*C.{base}"
            body_lines.append(f"\t_slice_{local} := unsafe.Slice({local}, {length_expr})")
            body_lines.append(f"\t{out_local} := make([]{go_elem}, {length_expr})")
            body_lines.append(f"\tfor _i, _e := range _slice_{local} {{ {out_local}[_i] = {ctor_expr} }}")
            return out_local
        if kind == "@SLICE_WRAPPED":
            # Triple pointer: ``T ***`` means the C function writes a
            # ``T **`` (an array of N pointers) to ``*local``.
            go_elem, ctor = WRAPPER_TYPES[base]
            ctor_expr = ctor.replace("$res", "_e")
            body_lines.append(f"\t_slice_{local} := unsafe.Slice({local}, {length_expr})")
            body_lines.append(f"\t{out_local} := make([]{go_elem}, {length_expr})")
            body_lines.append(f"\tfor _i, _e := range _slice_{local} {{ {out_local}[_i] = {ctor_expr} }}")
            return out_local
        return marker

    # Substitute every marker; ordering is preserved because we walk the
    # comma-separated expression left to right.
    out_pieces = []
    for piece in tail.split(", "):
        if piece.startswith("@SLICE_"):
            out_pieces.append(_resolve(piece))
        else:
            out_pieces.append(piece)
    return ", ".join(out_pieces)


def _todo_stub(c_name: str, reason: str) -> str:
    return (
        f"// TODO {c_name}: {reason}\n"
        f"// func {_go_name(c_name)}(...) {{ /* not yet handled by codegen */ }}\n"
    )


# Driver ----------------------------------------------------------------

# The package-global #cgo directives (CFLAGS/LDFLAGS) live once here in cgo.go.
# cgo resolves each Go file's C references against THAT file's own preamble, so
# every generated per-header file repeats the MEOS #includes (see _PER_HEADER_PREAMBLE).
_CGO_FILE = """package functions

/*
// The installed libmeos describes itself: `pkg-config --cflags meos` reports
// its include dir together with the family macros it is compiled with, and
// `--libs` reports where to link it from. Both matter, because the public
// headers gate declarations on those macros (meos.h holds #if MEOS, #if
// POINTCLOUD and #if JSON blocks), so a hand-kept -D list that disagrees with
// the library either declares a symbol that cannot link or hides a family the
// library provides.
//
// Asking pkg-config also makes the binding relocatable: PKG_CONFIG_PATH selects
// which libmeos to build against, so a build points at a private prefix rather
// than inheriting whatever occupies a machine-wide directory. A prefix outside
// the loader's default search path needs LD_LIBRARY_PATH (or ldconfig).
#cgo pkg-config: meos

// H3 ships no pkg-config file, so the one dependency whose header the public
// MEOS headers expose (meos_h3.h includes <h3api.h>) names its include dir
// here. GEOS and GSL, the other exposed dependencies, resolve from the default
// include path.
#cgo darwin CFLAGS: -I/opt/homebrew/include/h3
#cgo linux CFLAGS: -I/usr/include/h3

#include <stddef.h>
#include "meos.h"
#include "meos_catalog.h"
#include "meos_geo.h"
#include "meos_internal.h"
#include "meos_internal_geo.h"
#include "meos_npoint.h"
#include "meos_cbuffer.h"
#include "meos_pose.h"
#include "meos_rgeo.h"
#include "meos_h3.h"
#include "meos_quadbin.h"
#include "meos_json.h"
#include "meos_pointcloud.h"
*/
import "C"

import "fmt"

// MeosError reports that the MEOS call behind a wrapper failed, carrying the
// error code MEOS recorded for it.
type MeosError struct {
	Code int
}

func (e *MeosError) Error() string {
	return fmt.Sprintf("meos: error %d", e.Code)
}

// meosError reads the out-of-band state MEOS sets when a call fails, and
// clears it so a stale code cannot condemn the next, innocent call.  It
// returns nil when the call succeeded.
//
// The returned value can never answer this on its own: every MEOS sentinel is
// a legitimate value of its own return type -- INT_MAX is a count, DBL_MAX is
// a distance, and a bool has no spare value at all.  So no generated wrapper
// compares against a sentinel; each consults this instead.
//
// Install meos_initialize_noexit_error_handler (wrapped as
// MeosInitializeNoexitErrorHandler) once per process before calling anything
// else, on every thread that calls MEOS: the default handler ends the
// process, and meos_initialize resets the handler per thread.
func meosError() error {
	if code := int(C.meos_errno()); code != 0 {
		C.meos_errno_reset()
		return &MeosError{Code: code}
	}
	return nil
}
"""

_PER_HEADER_PREAMBLE = """package functions

/*
#include <stddef.h>
#include "meos.h"
#include "meos_catalog.h"
#include "meos_geo.h"
#include "meos_internal.h"
#include "meos_internal_geo.h"
#include "meos_npoint.h"
#include "meos_cbuffer.h"
#include "meos_pose.h"
#include "meos_rgeo.h"
#include "meos_h3.h"
#include "meos_quadbin.h"
#include "meos_json.h"
#include "meos_pointcloud.h"
*/
import "C"
import (
\t"unsafe"
)

var _ = unsafe.Pointer(nil)

"""


def generate(idl_path: Path, out_dir: Path) -> dict:
    idl = json.loads(idl_path.read_text())
    configure(idl)
    out_dir.mkdir(parents=True, exist_ok=True)

    entries_by_file: dict[str, list[dict]] = {h: [] for h in HEADER_FILES}
    for entry in idl["functions"]:
        if entry["file"] in entries_by_file:
            entries_by_file[entry["file"]].append(entry)

    stats = {"emitted": 0, "skipped": 0, "explicit_skip": 0, "datum": 0, "by_header": {}}
    corpus_parts: list[str] = []
    # The shared per-file cgo preamble, with the union-operator aliases spliced
    # into the C block so every file that calls a ``C.gunion_*`` resolves it.
    preamble = _PER_HEADER_PREAMBLE.replace("*/", _union_alias_block(idl) + "\n*/", 1)
    for header in HEADER_FILES:
        emitted_funcs = []
        local_emit = local_skip = local_explicit = local_datum = 0
        for entry in entries_by_file[header]:
            if entry["name"] in SKIPPED_FUNCTIONS:
                continue
            if _references_opaque(entry):
                continue
            if _is_datum_internal(entry):
                local_datum += 1
                continue
            ef = emit_function(entry)
            emitted_funcs.append(ef)
            if ef.skipped:
                # Differentiate explicit meta-driven skips (declared
                # non-wrappable) from honest TODOs (shape the generator
                # could not resolve).
                if "skip:" in ef.code:
                    local_explicit += 1
                else:
                    local_skip += 1
            else:
                local_emit += 1
        stats["by_header"][header] = (local_emit, local_skip, local_explicit, local_datum)
        stats["emitted"] += local_emit
        stats["skipped"] += local_skip
        stats["explicit_skip"] += local_explicit
        stats["datum"] += local_datum

        out_file = out_dir / f"meos_{header.replace('.h', '')}.go"
        funcs_code = "\n\n".join(e.code for e in emitted_funcs)
        corpus_parts.append(funcs_code)
        body = preamble + funcs_code + "\n"
        out_file.write_text(body)

    # The single cgo preamble file with all #includes lives next to the
    # generated wrappers; Go's cgo machinery merges directives across files.
    (out_dir / "cgo.go").write_text(_CGO_FILE)

    # The type layer (opaque handle structs + enums) the wrappers depend on,
    # restricted to the types the emitted wrappers actually reference.
    (out_dir / "types.go").write_text(emit_types(idl, "\n".join(corpus_parts)))
    return stats


if __name__ == "__main__":
    here = Path(__file__).parent
    # The generated flat FFI layer is the ``functions`` package at the repo
    # root (mirrors JMEOS ``functions`` / ``pymeos_cffi.functions``): the thin
    # I/O shell over MEOS that ``go build ./...`` compiles alongside the OO
    # ``gomeos`` package.
    stats = generate(here / "meos-idl.json", here.parent / "functions")
    print(f"Emitted {stats['emitted']} idiomatic wrappers")
    print(f"Skipped {stats['skipped']} as TODO (unresolved signature shape)")
    print(f"Excluded {stats['explicit_skip']} via meta/meos-meta.json shape.skip declarations")
    print(f"Excluded {stats['datum']} as Datum-bearing internal helpers")
    for h, (e, s, sk, d) in stats["by_header"].items():
        print(f"  {h}: {e} emitted, {s} TODO, {sk} explicit-skip, {d} Datum")
