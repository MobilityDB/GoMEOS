package gomeos

/*
#include "meos.h"
#include "meos_catalog.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"

// ------------------------- Accessors ----------------------------------
func TSequenceLowerInclude[TS TSequence](temp TS) bool {
	inner := C.cast_temporal_to_tsequence(temp.Inner())
	return bool(inner.period.lower_inc)
}

func TSequenceUpperInclude[TS TSequence](temp TS) bool {
	inner := C.cast_temporal_to_tsequence(temp.Inner())
	return bool(inner.period.upper_inc)
}
