// cast.c
#include "cast.h"

TInstant *cast_temporal_to_tinstant(Temporal *temp) {
    TInstant *tinst = (TInstant *) temp;
    return tinst;
}

TSequence *cast_temporal_to_tsequence(Temporal *temp) {
    TSequence *tseq = (TSequence *) temp;
    return tseq;
}

Temporal *cast_tinstant_to_temporal(TInstant *tinst) {
    Temporal *temp = (Temporal *) tinst;
    return temp;
}


Temporal *cast_tsequence_to_temporal(TSequence *tseq) {
    Temporal *temp = (Temporal *) tseq;
    return temp;
}

GSERIALIZED *cast_pointer_to_geo(Datum *p) {
    GSERIALIZED *gs = (GSERIALIZED *) p;
    return gs;
}