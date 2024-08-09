package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

func TnumberTwavg[T Temporal](temp T) float64 {
	res := C.tnumber_twavg(temp.Inner())
	return float64(res)
}

func TemporalFromHexWKB[T Temporal](empty_temp T, s string) T {
	c_temp := C.temporal_from_hexwkb(C.CString(s))
	empty_temp.Init(c_temp)
	return empty_temp
}

func TemporalInstantN[T Temporal, I TInstant](temp T, inst I, n int) I {
	c_inst := C.temporal_instant_n(temp.Inner(), C.int(n))
	c_temp := inst.InstToTemporal(c_inst)
	inst.Init(c_temp)
	return inst
}

func TPointLength[T Temporal](temp T) float64 {
	return float64(C.tpoint_length(temp.Inner()))
}

func TpointAtGeomTime[T Temporal](temp T, new_temp T, geom *Geom) T {
	c_temp := C.tpoint_at_geom_time(temp.Inner(), geom._inner, nil, nil)
	new_temp.Init(c_temp)
	return new_temp
}

func TpointMinusGeomTime[T Temporal](temp T, new_temp T, geom *Geom) T {
	c_temp := C.tpoint_minus_geom_time(temp.Inner(), geom._inner, nil, nil)
	new_temp.Init(c_temp)
	return new_temp
}
