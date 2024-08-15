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


#include "meos.h"
#include <stdio.h>
#include <stdlib.h>

TInstant *cast_temporal_to_tinstant(Temporal *temp);
TSequence *cast_temporal_to_tsequence(Temporal *temp);
Temporal *cast_tinstant_to_temporal(TInstant *tinst);
Temporal *cast_tsequence_to_temporal(TSequence *tseq);
GSERIALIZED *cast_pointer_to_geo(Datum *p);

#endif // CAST_H