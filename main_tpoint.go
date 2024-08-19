package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"
import "unsafe"

// ------------------------- Output ----------------------------------

func TPointOut[TP TPoint](tp TP, maxdd int) string {
	c_point := C.tpoint_out(tp.Inner(), C.int(maxdd))
	defer C.free(unsafe.Pointer(c_point))
	point_out := C.GoString(c_point)
	return point_out
}

func TPointAsText[TP TPoint](tp TP, maxdd int) string {
	c_text := C.tpoint_as_text(tp.Inner(), C.int(maxdd))
	defer C.free(unsafe.Pointer(c_text))
	text_out := C.GoString(c_text)
	return text_out
}

func TPointAsEWKT[T TPoint](temp T, maxdd int) string {
	return C.GoString(C.tpoint_as_ewkt(temp.Inner(), C.int(maxdd)))
}

func TPointGeoAsEWKT[T TPoint](temp T, maxdd int) string {
	c_inst := C.cast_temporal_to_tinstant(temp.Inner())
	c_value := c_inst.value
	c_geo := (*C.GSERIALIZED)(unsafe.Pointer(&c_value))
	return C.GoString(C.geo_as_ewkt(c_geo, C.int(maxdd)))
}

// ------------------------- Accessors -------------------------------------

func TPointToSTBox[TP TPoint](tp TP) *STBox {
	st_box := C.tpoint_to_stbox(tp.Inner())
	return &STBox{_inner: st_box}
}

func TPointTransform[T TPoint](temp T, output T, srid_to int) T {
	c_temp := C.tpoint_transform(temp.Inner(), C.int(srid_to))
	output.Init(c_temp)
	return output
}

func TPointTrajectory[TP TPoint](tp TP) *Geom {
	trajectory := C.tpoint_trajectory(tp.Inner())
	return &Geom{_inner: trajectory}
}
