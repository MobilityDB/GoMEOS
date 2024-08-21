package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"

*/
import "C"
import "unsafe"

// ------------------------- Geom ---------------------------
type Geom struct {
	_inner *C.GSERIALIZED
}

// ------------------------- Input ---------------------------
func NewGeom(geom_str string, typemod int) Geom {
	c_geom_str := C.CString(geom_str)
	defer C.free(unsafe.Pointer(c_geom_str))
	c_geom := C.pgis_geometry_in(c_geom_str, C.int(typemod))
	g := Geom{_inner: c_geom}
	return g
}

func PgisGeometryIn(input string, typemod int) *Geom {
	c_geom_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_geom_str))
	c_geom := C.pgis_geometry_in(c_geom_str, C.int(typemod))
	g := &Geom{_inner: c_geom}
	return g
}

func PgisGeographyIn(input string, typemod int) *Geom {
	c_geom_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_geom_str))
	c_geom := C.pgis_geography_in(c_geom_str, C.int(typemod))
	g := &Geom{_inner: c_geom}
	return g
}

func GeographyFromHexEwkb(input string) *Geom {
	c_geom_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_geom_str))
	c_geom := C.geography_from_hexewkb(c_geom_str)
	g := &Geom{_inner: c_geom}
	return g
}

func GeometryFromHexEwkb(input string) *Geom {
	c_geom_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_geom_str))
	c_geom := C.geometry_from_hexewkb(c_geom_str)
	g := &Geom{_inner: c_geom}
	return g
}

func GeographyFromText(input string, srid int) *Geom {
	c_geom_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_geom_str))
	c_geom := C.geography_from_text(c_geom_str, C.int(srid))
	g := &Geom{_inner: c_geom}
	return g
}

func GeometryFromText(input string, srid int) *Geom {
	c_geom_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_geom_str))
	c_geom := C.geometry_from_text(c_geom_str, C.int(srid))
	g := &Geom{_inner: c_geom}
	return g
}

func GeoFromGeojson(input string) *Geom {
	c_geom_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_geom_str))
	c_geom := C.geo_from_geojson(c_geom_str)
	g := &Geom{_inner: c_geom}
	return g
}

// ------------------------- Output ---------------------------
func (geom *Geom) GeoOut() string {
	c_tgmpi_out := C.geo_out(geom._inner)
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

func GeoAsText(g *Geom, precision int) string {
	return C.GoString(C.geo_as_text(g._inner, C.int(precision)))
}

func GeoAsGeojson(g *Geom, option int, precision int, srs string) string {
	return C.GoString(C.geo_as_geojson(g._inner, C.int(option), C.int(precision), C.CString(srs)))
}

func GeoAsEWKT(g *Geom, precision int) string {
	return C.GoString(C.geo_as_ewkt(g._inner, C.int(precision)))
}

func GeoAsHexEwkb(g *Geom, endian string) string {
	return C.GoString(C.geo_as_hexewkb(g._inner, C.CString(endian)))
}

// ------------------------- Operation ---------------------------
func GeoSame(g1 *Geom, g2 *Geom) bool {
	return bool(C.geo_same(g1._inner, g2._inner))
}
