package gomeos

/*
#include "meos.h"
#include "meos_catalog.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"

func TSequenceMake[TI TInstant, TS TSequence](instants []TI, count int, lower_inc bool, upper_inc bool, interp Interpolation, normalize bool, output TS) TS {
	var c_instants []*C.TInstant
	// Fill the C array
	for _, inst := range instants {
		tinst := C.cast_temporal_to_tinstant(inst.Inner())
		c_instants = append(c_instants, tinst)
	}
	var input **C.TInstant = &(c_instants[0])
	c_geogseq := C.tsequence_make(input, C.int(count), C.bool(lower_inc), C.bool(upper_inc), C.interpType(interp), C.bool(normalize))
	c_geotemp := C.cast_tsequence_to_temporal(c_geogseq)
	output.Init(c_geotemp)
	return output
}

func TSequenceLowerInclude[TS TSequence](temp TS) bool {
	inner := C.cast_temporal_to_tsequence(temp.Inner())
	return bool(inner.period.lower_inc)
}

func TSequenceUpperInclude[TS TSequence](temp TS) bool {
	inner := C.cast_temporal_to_tsequence(temp.Inner())
	return bool(inner.period.upper_inc)
}
