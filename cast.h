// cast.h
#ifndef CAST_H
#define CAST_H
#define gunion_set_int union_set_int
#define gunion_set_set union_set_set
#define gunion_set_float union_set_float
#define gunion_set_date union_set_date
#define gunion_span_int union_span_int
#define gunion_span_float union_span_float
#define gunion_span_date union_span_date
#define gunion_span_span union_span_span
#define gunion_span_spanset union_span_spanset
#define gunion_spanset_int union_spanset_int
#define gunion_spanset_float union_spanset_float
#define gunion_spanset_date union_spanset_date
#define gunion_spanset_span union_spanset_span
#define gunion_spanset_spanset union_spanset_spanset


#include <stddef.h>
#include "meos.h"
#include "meos_geo.h"
#include "meos_catalog.h"
#include <stdio.h>
#include <stdlib.h>

TInstant *cast_temporal_to_tinstant(Temporal *temp);
TSequence *cast_temporal_to_tsequence(Temporal *temp);
Temporal *cast_tinstant_to_temporal(TInstant *tinst);
Temporal *cast_tsequence_to_temporal(TSequence *tseq);
Temporal *cast_tsequenceset_to_temporal(TSequenceSet *tseqset);
GSERIALIZED *cast_pointer_to_geo(Datum *p);

/* Thin wrapper around the MEOS "temptype supports linear interp" query:
 * cgo cannot expose the underlying function directly because its second
 * parameter is named `type`, a Go reserved word (the codegen flags it
 * as an unsupported-MeosType-param case in tools/_preview). Taking a
 * Temporal* and reading its temptype keeps the MeosType-typed parameter
 * inside C.
 *
 * MEOS renamed temptype_continuous -> temptype_supports_linear in
 * MobilityDB#1005; forward-declare it here so the wrapper compiles
 * regardless of which name the vendored meos_catalog.h carries. */
extern bool temptype_supports_linear(MeosType type);
static inline bool gomeos_temporal_continuous(const Temporal *t) {
  return temptype_supports_linear(t->temptype);
}

#endif // CAST_H