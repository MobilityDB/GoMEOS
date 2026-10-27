# GoMEOS code generator

`codegen.py` reads `tools/meos-idl.json` (produced by the
[MEOS-API](https://github.com/MobilityDB/MEOS-API) parser) and emits one
idiomatic Go wrapper per MEOS public C function into `tools/_preview/`.
The directory is prefixed with `_` so `go build ./...` ignores it; the
output is a Draft artifact that lets reviewers see the shape of the
generated surface without disturbing the hand-written package.

## Running

```
python3 tools/codegen.py
```

## Coverage today

The generator binds every non-vendored catalog function of a bound family
(the catalog's `family`, all on, `<FAMILY>=OFF` in the environment to drop
one) whose header the libmeos prefix installs, one Go file per header, the
headers taken from the catalog's `file` field (`header_plan`).  Against
MobilityDB master it emits 4612 wrappers; 144 shapes stay unresolved, 191
are excluded as `Datum`-bearing internal helpers (the hand-written surface
exposes those through typed overloads which the codegen cannot synthesise
from the IDL), and the 227 internal functions of `pgtypes.h` and
`pgsql_compat.h`, headers the prefix does not install, are named in its
report and not bound.  The skiplist primitives that take
function-pointer arguments (`comp_fn`, `merge_fn`, `datum_func2`) or
operate on `void **` element arrays surface as `unsafe.Pointer`,
matching how PyMEOS-CFFI exposes them as `_ffi.CData` and MEOS.NET as
`IntPtr` — callers provide the raw pointer obtained from another
wrapped MEOS call (for example a typed comparator's address).

The covered shapes are: scalar inputs, wrapped opaque pointers
(`Temporal`, `STBox`, `TBox`, `Span`, `SpanSet`, `Set`, `GSERIALIZED`,
`TInstant`, `TSequence`, `TSequenceSet`, `Npoint`, `Nsegment`,
`SkipList`, `RTree`, `Match`, `BOX3D`, `GBOX`, `AFFINE`, `PJ_CONTEXT`,
`Interval`), C and PostgreSQL text strings, output
parameters named `result` / `value` (scalar and wrapped-pointer
out-params surfaced as additional Go return values), counted-array
inputs (`T **` / `const T *` paired with `int count` or `size_t size`)
lowered to Go slices, counted-array returns (`T **` paired with
`int *count` or matching the input slice length) lowered to Go slices
including byte buffers and `char **` / `text **` string arrays,
`unsafe.Pointer` for `void *` arguments, and metadata-driven shapes
described below.

## meta/meos-meta.json shape annotations

The remaining ecosystem-wide editorial decisions live in MEOS-API's
`meta/meos-meta.json` under each function's `shape` key, merged into
`meos-idl.json` at IDL-generation time.  All bindings consume the same
catalog.

Annotation kinds:

* `shape.arrayReturn.lengthFrom = { kind: "accessor", func, arg, castTo? }` —
  the function returns an array whose length is obtained by calling a
  sibling accessor on one of its inputs.  Used by the `*_values` family
  (`bigintset_values` -> `set_num_values(s)`), the `*_insts_p` /
  `*_sequences_p` accessors (cast to `const Temporal *`), and the
  `*_spanarr` / `*_sps` spanset accessors.
* `shape.arrayReturn.lengthFrom = { kind: "param", name }` — the
  function returns an array whose length is written to an output
  parameter of the same call.
* `shape.outputArrays = [{ param, ... }]` — additional parallel output
  arrays sharing the primary length.  Used by the `*_split` family
  (`time_bins`, `value_bins`) including the triple-pointer
  `GSERIALIZED ***` output of `tgeo_space_split` / `tgeo_space_time_split`.
* `shape.namedOutputs = ["subtype", "radius"]` — scalar out-parameters
  whose name is neither `result` nor `value`.
* `shape.arrayInputGroup = { params, count, nullable }` — N parallel
  input arrays sharing one count (`tpointseq_make_coords`).  Nullable
  members accept Go `nil`.
* `shape.skip = "<reason>"` — bindings omit the function entirely.
  Retained as an escape hatch but no entries use it today; the
  skiplist family now surfaces through `unsafe.Pointer` instead.

## Refreshing the IDL

`tools/meos-idl.json` is vendored.  When MEOS bumps, regenerate with
MEOS-API against the new headers and copy the output back:

```
cp ../MEOS-API/output/meos-idl.json tools/meos-idl.json
python3 tools/codegen.py
```

The vendored IDL relies on MEOS-API PR #1 (the stdbool stub) being in
place so `bool` returns and `bool *` outputs do not get demoted to
`int`.

The vendored IDL relies on MEOS-API PR #1 (the stdbool stub) being in
place so `bool` returns and `bool *` outputs do not get demoted to `int`.
