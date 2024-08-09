package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>


*/
import "C"
import "unsafe"

// ------------------------- TFloatInst ---------------------------
type Geom struct {
	_inner *C.GSERIALIZED
}

func NewGeom(geom_str string, typemod int) Geom {
	c_geom_str := C.CString(geom_str)
	defer C.free(unsafe.Pointer(c_geom_str))
	c_geom := C.pgis_geometry_in(c_geom_str, C.int(typemod))
	g := Geom{_inner: c_geom}
	return g
}

func (geom *Geom) GeomOut() string {
	c_tgmpi_out := C.geo_out(geom._inner)
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}
