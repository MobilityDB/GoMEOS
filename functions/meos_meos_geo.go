package functions

/*
#include <stddef.h>
#include "meos.h"
#include "meos_catalog.h"
#include "meos_cbuffer.h"
#include "meos_cellindex.h"
#include "meos_geo.h"
#include "meos_h3.h"
#include "meos_internal.h"
#include "meos_internal_geo.h"
#include "meos_json.h"
#include "meos_npoint.h"
#include "meos_pointcloud.h"
#include "meos_pose.h"
#include "meos_quadbin.h"
#include "meos_raster.h"
#include "meos_rgeo.h"
#include "meos_s2cell.h"
#include "pg_basetypes.h"
#include "pg_bool.h"
#include "pg_date.h"
#include "pg_float.h"
#include "pg_int.h"
#include "pg_interval.h"
#include "pg_text.h"
#include "pg_time.h"
#include "pg_timestamp.h"

// cgo reads C.union_* as a union TYPE; alias the union operators so C.g<name> resolves them as functions.
#define gunion_bigint_set union_bigint_set
#define gunion_bigint_span union_bigint_span
#define gunion_bigint_spanset union_bigint_spanset
#define gunion_cbuffer_set union_cbuffer_set
#define gunion_date_set union_date_set
#define gunion_date_span union_date_span
#define gunion_date_spanset union_date_spanset
#define gunion_float_set union_float_set
#define gunion_float_span union_float_span
#define gunion_float_spanset union_float_spanset
#define gunion_geo_set union_geo_set
#define gunion_h3index_set union_h3index_set
#define gunion_int_set union_int_set
#define gunion_int_span union_int_span
#define gunion_int_spanset union_int_spanset
#define gunion_jsonb_set union_jsonb_set
#define gunion_npoint_set union_npoint_set
#define gunion_pcpatch_set union_pcpatch_set
#define gunion_pcpoint_set union_pcpoint_set
#define gunion_pose_set union_pose_set
#define gunion_posechain_set union_posechain_set
#define gunion_quadbin_set union_quadbin_set
#define gunion_s2cell_set union_s2cell_set
#define gunion_set_bigint union_set_bigint
#define gunion_set_cbuffer union_set_cbuffer
#define gunion_set_date union_set_date
#define gunion_set_float union_set_float
#define gunion_set_geo union_set_geo
#define gunion_set_h3index union_set_h3index
#define gunion_set_int union_set_int
#define gunion_set_jsonb union_set_jsonb
#define gunion_set_npoint union_set_npoint
#define gunion_set_pcpatch union_set_pcpatch
#define gunion_set_pcpoint union_set_pcpoint
#define gunion_set_pose union_set_pose
#define gunion_set_posechain union_set_posechain
#define gunion_set_quadbin union_set_quadbin
#define gunion_set_s2cell union_set_s2cell
#define gunion_set_set union_set_set
#define gunion_set_text union_set_text
#define gunion_set_timestamptz union_set_timestamptz
#define gunion_set_value union_set_value
#define gunion_span_bigint union_span_bigint
#define gunion_span_date union_span_date
#define gunion_span_float union_span_float
#define gunion_span_int union_span_int
#define gunion_span_span union_span_span
#define gunion_span_spanset union_span_spanset
#define gunion_span_timestamptz union_span_timestamptz
#define gunion_span_value union_span_value
#define gunion_spanset_bigint union_spanset_bigint
#define gunion_spanset_date union_spanset_date
#define gunion_spanset_float union_spanset_float
#define gunion_spanset_int union_spanset_int
#define gunion_spanset_span union_spanset_span
#define gunion_spanset_spanset union_spanset_spanset
#define gunion_spanset_timestamptz union_spanset_timestamptz
#define gunion_spanset_value union_spanset_value
#define gunion_stbox_stbox union_stbox_stbox
#define gunion_tbox_tbox union_tbox_tbox
#define gunion_text_set union_text_set
#define gunion_timestamptz_set union_timestamptz_set
#define gunion_timestamptz_span union_timestamptz_span
#define gunion_timestamptz_spanset union_timestamptz_spanset
#define gunion_tpcbox_tpcbox union_tpcbox_tpcbox
#define gunion_value_set union_value_set
#define gunion_value_span union_value_span
#define gunion_value_spanset union_value_spanset
*/
import "C"
import (
	"unsafe"
)

var _ = unsafe.Pointer(nil)

// Box3dMake wraps MEOS C function box3d_make.
func Box3dMake(xmin float64, xmax float64, ymin float64, ymax float64, zmin float64, zmax float64, srid int32) (_r0 *Box3D, _err error) {
	C.meos_errno_reset()
	_cret := C.box3d_make(C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.double(zmin), C.double(zmax), C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Box3D{_inner: _cret}, nil
}


// Box3dIn wraps MEOS C function box3d_in.
func Box3dIn(str string) (_r0 *Box3D, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.box3d_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Box3D{_inner: _cret}, nil
}


// Box3dOut wraps MEOS C function box3d_out.
func Box3dOut(box *Box3D, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.box3d_out(box._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// GboxMake wraps MEOS C function gbox_make.
func GboxMake(hasz bool, hasm bool, geodetic bool, xmin float64, xmax float64, ymin float64, ymax float64, zmin float64, zmax float64, mmin float64, mmax float64) (_r0 *GBox, _err error) {
	C.meos_errno_reset()
	_cret := C.gbox_make(C.bool(hasz), C.bool(hasm), C.bool(geodetic), C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.double(zmin), C.double(zmax), C.double(mmin), C.double(mmax))
	if _err = meosError(); _err != nil {
		return
	}
	return &GBox{_inner: _cret}, nil
}


// GboxIn wraps MEOS C function gbox_in.
func GboxIn(str string) (_r0 *GBox, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.gbox_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &GBox{_inner: _cret}, nil
}


// GboxOut wraps MEOS C function gbox_out.
func GboxOut(box *GBox, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.gbox_out(box._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// GeoAsEWKB wraps MEOS C function geo_as_ewkb.
func GeoAsEWKB(gs *Geom, endian *string, size unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	var _c_endian *C.char
	if endian != nil {
		_c_endian = C.CString(*endian)
		defer C.free(unsafe.Pointer(_c_endian))
	}
	C.meos_errno_reset()
	_cret := C.geo_as_ewkb(gs._inner, _c_endian, (*C.size_t)(unsafe.Pointer(size)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// GeoAsEWKT wraps MEOS C function geo_as_ewkt.
func GeoAsEWKT(gs *Geom, precision int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_as_ewkt(gs._inner, C.int(precision))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// GeoAsGeojson wraps MEOS C function geo_as_geojson.
func GeoAsGeojson(gs *Geom, option int, precision int, srs *string) (_r0 string, _err error) {
	var _c_srs *C.char
	if srs != nil {
		_c_srs = C.CString(*srs)
		defer C.free(unsafe.Pointer(_c_srs))
	}
	C.meos_errno_reset()
	_cret := C.geo_as_geojson(gs._inner, C.int(option), C.int(precision), _c_srs)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// GeoAsHexewkb wraps MEOS C function geo_as_hexewkb.
func GeoAsHexewkb(gs *Geom, endian *string) (_r0 string, _err error) {
	var _c_endian *C.char
	if endian != nil {
		_c_endian = C.CString(*endian)
		defer C.free(unsafe.Pointer(_c_endian))
	}
	C.meos_errno_reset()
	_cret := C.geo_as_hexewkb(gs._inner, _c_endian)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// GeoAsText wraps MEOS C function geo_as_text.
func GeoAsText(gs *Geom, precision int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_as_text(gs._inner, C.int(precision))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// GeoFromEWKB wraps MEOS C function geo_from_ewkb.
func GeoFromEWKB(wkb unsafe.Pointer, wkb_size uint, srid int32) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_from_ewkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(wkb_size), C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoFromGeojson wraps MEOS C function geo_from_geojson.
func GeoFromGeojson(geojson string) (_r0 *Geom, _err error) {
	_c_geojson := C.CString(geojson)
	defer C.free(unsafe.Pointer(_c_geojson))
	C.meos_errno_reset()
	_cret := C.geo_from_geojson(_c_geojson)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoFromText wraps MEOS C function geo_from_text.
func GeoFromText(wkt string, srid int32) (_r0 *Geom, _err error) {
	_c_wkt := C.CString(wkt)
	defer C.free(unsafe.Pointer(_c_wkt))
	C.meos_errno_reset()
	_cret := C.geo_from_text(_c_wkt, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoOut wraps MEOS C function geo_out.
func GeoOut(gs *Geom) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_out(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// GeogFromHexewkb wraps MEOS C function geog_from_hexewkb.
func GeogFromHexewkb(hexwkb string) (_r0 *Geom, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.geog_from_hexewkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeogIn wraps MEOS C function geog_in.
func GeogIn(str string, typmod int32) (_r0 *Geom, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.geog_in(_c_str, C.int32(typmod))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomFromHexewkb wraps MEOS C function geom_from_hexewkb.
func GeomFromHexewkb(hexwkb string) (_r0 *Geom, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.geom_from_hexewkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomIn wraps MEOS C function geom_in.
func GeomIn(str string, typmod int32) (_r0 *Geom, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.geom_in(_c_str, C.int32(typmod))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoCopy wraps MEOS C function geo_copy.
func GeoCopy(gs *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_copy(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeogpointMake2d wraps MEOS C function geogpoint_make2d.
func GeogpointMake2d(srid int32, x float64, y float64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geogpoint_make2d(C.int32_t(srid), C.double(x), C.double(y))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeogpointMake3dz wraps MEOS C function geogpoint_make3dz.
func GeogpointMake3dz(srid int32, x float64, y float64, z float64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geogpoint_make3dz(C.int32_t(srid), C.double(x), C.double(y), C.double(z))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeompointMake2d wraps MEOS C function geompoint_make2d.
func GeompointMake2d(srid int32, x float64, y float64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geompoint_make2d(C.int32_t(srid), C.double(x), C.double(y))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeompointMake3dz wraps MEOS C function geompoint_make3dz.
func GeompointMake3dz(srid int32, x float64, y float64, z float64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geompoint_make3dz(C.int32_t(srid), C.double(x), C.double(y), C.double(z))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomToGeog wraps MEOS C function geom_to_geog.
func GeomToGeog(geom *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_to_geog(geom._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeogToGeom wraps MEOS C function geog_to_geom.
func GeogToGeom(geog *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geog_to_geom(geog._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoIsEmpty wraps MEOS C function geo_is_empty.
func GeoIsEmpty(gs *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_is_empty(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeoIsUnitary wraps MEOS C function geo_is_unitary.
func GeoIsUnitary(gs *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_is_unitary(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeoMemSize wraps MEOS C function geo_mem_size.
func GeoMemSize(gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_mem_size(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// GeoArea wraps MEOS C function geo_area.
func GeoArea(gs *Geom, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_area(gs._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// GeoCentroid wraps MEOS C function geo_centroid.
func GeoCentroid(gs *Geom, spheroid bool) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_centroid(gs._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoLength wraps MEOS C function geo_length.
func GeoLength(gs *Geom, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_length(gs._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// GeoPerimeter wraps MEOS C function geo_perimeter.
func GeoPerimeter(gs *Geom, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_perimeter(gs._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// GeoTypename wraps MEOS C function geo_typename.
func GeoTypename(type_ int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_typename(C.int(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// GeogArea wraps MEOS C function geog_area.
func GeogArea(gs *Geom, use_spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.geog_area(gs._inner, C.bool(use_spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// GeogCentroid wraps MEOS C function geog_centroid.
func GeogCentroid(gs *Geom, use_spheroid bool) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geog_centroid(gs._inner, C.bool(use_spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeogLength wraps MEOS C function geog_length.
func GeogLength(gs *Geom, use_spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.geog_length(gs._inner, C.bool(use_spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// GeogPerimeter wraps MEOS C function geog_perimeter.
func GeogPerimeter(gs *Geom, use_spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.geog_perimeter(gs._inner, C.bool(use_spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// GeomAzimuth wraps MEOS C function geom_azimuth.
func GeomAzimuth(gs1 *Geom, gs2 *Geom) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.geom_azimuth(gs1._inner, gs2._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// GeomArea wraps MEOS C function geom_area.
func GeomArea(gs *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_area(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// GeomLength wraps MEOS C function geom_length.
func GeomLength(gs *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_length(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// GeomPerimeter wraps MEOS C function geom_perimeter.
func GeomPerimeter(gs *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_perimeter(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// LineNumpoints wraps MEOS C function line_numpoints.
func LineNumpoints(gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.line_numpoints(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// LinePointN wraps MEOS C function line_point_n.
func LinePointN(geom *Geom, n int) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.line_point_n(geom._inner, C.int(n))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoAffine wraps MEOS C function geo_affine.
func GeoAffine(gs *Geom, a float64, b float64, c float64, d float64, e float64, f float64, g float64, h float64, i float64, xoff float64, yoff float64, zoff float64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_affine(gs._inner, C.double(a), C.double(b), C.double(c), C.double(d), C.double(e), C.double(f), C.double(g), C.double(h), C.double(i), C.double(xoff), C.double(yoff), C.double(zoff))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoReverse wraps MEOS C function geo_reverse.
func GeoReverse(gs *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_reverse(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoRotate wraps MEOS C function geo_rotate.
func GeoRotate(gs *Geom, angle float64, x0 float64, y0 float64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_rotate(gs._inner, C.double(angle), C.double(x0), C.double(y0))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoRotateGeo wraps MEOS C function geo_rotate_geo.
func GeoRotateGeo(gs *Geom, angle float64, origin *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_rotate_geo(gs._inner, C.double(angle), origin._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoRotateX wraps MEOS C function geo_rotate_x.
func GeoRotateX(gs *Geom, angle float64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_rotate_x(gs._inner, C.double(angle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoRotateY wraps MEOS C function geo_rotate_y.
func GeoRotateY(gs *Geom, angle float64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_rotate_y(gs._inner, C.double(angle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoRotateZ wraps MEOS C function geo_rotate_z.
func GeoRotateZ(gs *Geom, angle float64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_rotate_z(gs._inner, C.double(angle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoRound wraps MEOS C function geo_round.
func GeoRound(gs *Geom, maxdd int) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_round(gs._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoScale wraps MEOS C function geo_scale.
func GeoScale(gs *Geom, xfactor float64, yfactor float64, zfactor float64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_scale(gs._inner, C.double(xfactor), C.double(yfactor), C.double(zfactor))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoScaleGeo wraps MEOS C function geo_scale_geo.
func GeoScaleGeo(gs *Geom, scale *Geom, sorigin *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_scale_geo(gs._inner, scale._inner, sorigin._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoSegmentize wraps MEOS C function geo_segmentize.
func GeoSegmentize(gs *Geom, max_seg_length float64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_segmentize(gs._inner, C.double(max_seg_length))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoTranslate wraps MEOS C function geo_translate.
func GeoTranslate(gs *Geom, deltax float64, deltay float64, deltaz float64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_translate(gs._inner, C.double(deltax), C.double(deltay), C.double(deltaz))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoTransscale wraps MEOS C function geo_transscale.
func GeoTransscale(gs *Geom, deltax float64, deltay float64, xfactor float64, yfactor float64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_transscale(gs._inner, C.double(deltax), C.double(deltay), C.double(xfactor), C.double(yfactor))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoSetSRID wraps MEOS C function geo_set_srid.
func GeoSetSRID(gs *Geom, srid int32) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_set_srid(gs._inner, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoSRID wraps MEOS C function geo_srid.
func GeoSRID(gs *Geom) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_srid(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// GeoTransform wraps MEOS C function geo_transform.
func GeoTransform(geom *Geom, srid_to int32) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_transform(geom._inner, C.int32_t(srid_to))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoTransformPipeline wraps MEOS C function geo_transform_pipeline.
func GeoTransformPipeline(gs *Geom, pipelinestr string, srid int32, is_forward bool) (_r0 *Geom, _err error) {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	C.meos_errno_reset()
	_cret := C.geo_transform_pipeline(gs._inner, _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoCollectGarray wraps MEOS C function geo_collect_garray.
func GeoCollectGarray(gsarr unsafe.Pointer, count int) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_collect_garray((**C.GSERIALIZED)(unsafe.Pointer(gsarr)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoMakelineGarray wraps MEOS C function geo_makeline_garray.
func GeoMakelineGarray(gsarr unsafe.Pointer, count int) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_makeline_garray((**C.GSERIALIZED)(unsafe.Pointer(gsarr)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoNumPoints wraps MEOS C function geo_num_points.
func GeoNumPoints(gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_num_points(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// GeoNumGeos wraps MEOS C function geo_num_geos.
func GeoNumGeos(gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_num_geos(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// GeoGeoN wraps MEOS C function geo_geo_n.
func GeoGeoN(geom *Geom, n int) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_geo_n(geom._inner, C.int(n))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoPointarr wraps MEOS C function geo_pointarr.
func GeoPointarr(gs *Geom, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_pointarr(gs._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// GeoPoints wraps MEOS C function geo_points.
func GeoPoints(gs *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_points(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomArrayUnion wraps MEOS C function geom_array_union.
func GeomArrayUnion(gsarr unsafe.Pointer, count int) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_array_union((**C.GSERIALIZED)(unsafe.Pointer(gsarr)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeogArrayUnion wraps MEOS C function geog_array_union.
func GeogArrayUnion(gsarr unsafe.Pointer, count int) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geog_array_union((**C.GSERIALIZED)(unsafe.Pointer(gsarr)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomBoundary wraps MEOS C function geom_boundary.
func GeomBoundary(gs *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_boundary(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomBuffer wraps MEOS C function geom_buffer.
func GeomBuffer(gs *Geom, size float64, params string) (_r0 *Geom, _err error) {
	_c_params := C.CString(params)
	defer C.free(unsafe.Pointer(_c_params))
	C.meos_errno_reset()
	_cret := C.geom_buffer(gs._inner, C.double(size), _c_params)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomCentroid wraps MEOS C function geom_centroid.
func GeomCentroid(gs *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_centroid(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomConvexHull wraps MEOS C function geom_convex_hull.
func GeomConvexHull(gs *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_convex_hull(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomDifference2d wraps MEOS C function geom_difference2d.
func GeomDifference2d(gs1 *Geom, gs2 *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_difference2d(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomIsSimple wraps MEOS C function geom_is_simple.
func GeomIsSimple(gs *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_is_simple(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeomIntersection2d wraps MEOS C function geom_intersection2d.
func GeomIntersection2d(gs1 *Geom, gs2 *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_intersection2d(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomIntersection2dColl wraps MEOS C function geom_intersection2d_coll.
func GeomIntersection2dColl(gs1 *Geom, gs2 *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_intersection2d_coll(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomMinimumBoundingRadius wraps MEOS C function geom_minimum_bounding_radius.
func GeomMinimumBoundingRadius(geom *Geom, radius unsafe.Pointer) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_minimum_bounding_radius(geom._inner, (*C.double)(unsafe.Pointer(radius)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomOrientedEnvelope wraps MEOS C function geom_oriented_envelope.
func GeomOrientedEnvelope(gs *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_oriented_envelope(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomShortestline wraps MEOS C function geom_shortestline.
func GeomShortestline(gs1 *Geom, gs2 *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_shortestline(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomShortestline2d wraps MEOS C function geom_shortestline2d.
func GeomShortestline2d(gs1 *Geom, gs2 *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_shortestline2d(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomShortestline3d wraps MEOS C function geom_shortestline3d.
func GeomShortestline3d(gs1 *Geom, gs2 *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_shortestline3d(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomSymdifference2d wraps MEOS C function geom_symdifference2d.
func GeomSymdifference2d(gs1 *Geom, gs2 *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_symdifference2d(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomUnaryUnion wraps MEOS C function geom_unary_union.
func GeomUnaryUnion(gs *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_unary_union(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomUnion2d wraps MEOS C function geom_union2d.
func GeomUnion2d(gs1 *Geom, gs2 *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_union2d(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// LineInterpolatePoint wraps MEOS C function line_interpolate_point.
func LineInterpolatePoint(gs *Geom, distance_fraction float64, repeat bool) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.line_interpolate_point(gs._inner, C.double(distance_fraction), C.bool(repeat))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// LineLocatePoint wraps MEOS C function line_locate_point.
func LineLocatePoint(gs1 *Geom, gs2 *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.line_locate_point(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// LineSubstring wraps MEOS C function line_substring.
func LineSubstring(gs *Geom, from float64, to float64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.line_substring(gs._inner, C.double(from), C.double(to))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeogDwithin wraps MEOS C function geog_dwithin.
func GeogDwithin(g1 *Geom, g2 *Geom, tolerance float64, use_spheroid bool) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geog_dwithin(g1._inner, g2._inner, C.double(tolerance), C.bool(use_spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeogDisjoint wraps MEOS C function geog_disjoint.
func GeogDisjoint(gs1 *Geom, gs2 *Geom, use_spheroid bool) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geog_disjoint(gs1._inner, gs2._inner, C.bool(use_spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeogIntersects wraps MEOS C function geog_intersects.
func GeogIntersects(gs1 *Geom, gs2 *Geom, use_spheroid bool) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geog_intersects(gs1._inner, gs2._inner, C.bool(use_spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeomContains wraps MEOS C function geom_contains.
func GeomContains(gs1 *Geom, gs2 *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_contains(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeomCovers wraps MEOS C function geom_covers.
func GeomCovers(gs1 *Geom, gs2 *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_covers(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeomDisjoint wraps MEOS C function geom_disjoint.
func GeomDisjoint(gs1 *Geom, gs2 *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_disjoint(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeomDisjoint2d wraps MEOS C function geom_disjoint2d.
func GeomDisjoint2d(gs1 *Geom, gs2 *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_disjoint2d(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeomDwithin wraps MEOS C function geom_dwithin.
func GeomDwithin(gs1 *Geom, gs2 *Geom, tolerance float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_dwithin(gs1._inner, gs2._inner, C.double(tolerance))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeomDwithin2d wraps MEOS C function geom_dwithin2d.
func GeomDwithin2d(gs1 *Geom, gs2 *Geom, tolerance float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_dwithin2d(gs1._inner, gs2._inner, C.double(tolerance))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeomDwithin3d wraps MEOS C function geom_dwithin3d.
func GeomDwithin3d(gs1 *Geom, gs2 *Geom, tolerance float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_dwithin3d(gs1._inner, gs2._inner, C.double(tolerance))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeomIntersects wraps MEOS C function geom_intersects.
func GeomIntersects(gs1 *Geom, gs2 *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_intersects(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeomIntersects2d wraps MEOS C function geom_intersects2d.
func GeomIntersects2d(gs1 *Geom, gs2 *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_intersects2d(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeomIntersects3d wraps MEOS C function geom_intersects3d.
func GeomIntersects3d(gs1 *Geom, gs2 *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_intersects3d(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeomRelate wraps MEOS C function geom_relate.
func GeomRelate(gs1 *Geom, gs2 *Geom) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_relate(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// GeomRelatePattern wraps MEOS C function geom_relate_pattern.
func GeomRelatePattern(gs1 *Geom, gs2 *Geom, patt string) (_r0 bool, _err error) {
	_c_patt := C.CString(patt)
	defer C.free(unsafe.Pointer(_c_patt))
	C.meos_errno_reset()
	_cret := C.geom_relate_pattern(gs1._inner, gs2._inner, _c_patt)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeomTouches wraps MEOS C function geom_touches.
func GeomTouches(gs1 *Geom, gs2 *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_touches(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeoStboxes wraps MEOS C function geo_stboxes.
func GeoStboxes(gs *Geom, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_stboxes(gs._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// GeoSplitEachNStboxes wraps MEOS C function geo_split_each_n_stboxes.
func GeoSplitEachNStboxes(gs *Geom, elem_count int, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_split_each_n_stboxes(gs._inner, C.int(elem_count), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// GeoSplitNStboxes wraps MEOS C function geo_split_n_stboxes.
func GeoSplitNStboxes(gs *Geom, box_count int, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_split_n_stboxes(gs._inner, C.int(box_count), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// GeoDistance wraps MEOS C function geo_distance.
func GeoDistance(gs1 *Geom, gs2 *Geom, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_distance(gs1._inner, gs2._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// GeoShortestline wraps MEOS C function geo_shortestline.
func GeoShortestline(gs1 *Geom, gs2 *Geom, spheroid bool) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_shortestline(gs1._inner, gs2._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeogDistance wraps MEOS C function geog_distance.
func GeogDistance(g1 *Geom, g2 *Geom, use_spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.geog_distance(g1._inner, g2._inner, C.bool(use_spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// GeogShortestline wraps MEOS C function geog_shortestline.
func GeogShortestline(gs1 *Geom, gs2 *Geom, use_spheroid bool) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geog_shortestline(gs1._inner, gs2._inner, C.bool(use_spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomDistance wraps MEOS C function geom_distance.
func GeomDistance(gs1 *Geom, gs2 *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_distance(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// GeomDistance2d wraps MEOS C function geom_distance2d.
func GeomDistance2d(gs1 *Geom, gs2 *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_distance2d(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// GeomMaxDistance wraps MEOS C function geom_max_distance.
func GeomMaxDistance(gs1 *Geom, gs2 *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_max_distance(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// GeomMaxDistance2d wraps MEOS C function geom_max_distance2d.
func GeomMaxDistance2d(gs1 *Geom, gs2 *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_max_distance2d(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// GeomMaxDistance3d wraps MEOS C function geom_max_distance3d.
func GeomMaxDistance3d(gs1 *Geom, gs2 *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_max_distance3d(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// GeomDistance3d wraps MEOS C function geom_distance3d.
func GeomDistance3d(gs1 *Geom, gs2 *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_distance3d(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// GeoEquals wraps MEOS C function geo_equals.
func GeoEquals(gs1 *Geom, gs2 *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_equals(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// GeoSame wraps MEOS C function geo_same.
func GeoSame(gs1 *Geom, gs2 *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_same(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeogsetIn wraps MEOS C function geogset_in.
func GeogsetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.geogset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// GeomsetIn wraps MEOS C function geomset_in.
func GeomsetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.geomset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// SpatialsetOut wraps MEOS C function spatialset_out.
func SpatialsetOut(s *Set, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.spatialset_out(s._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// SpatialsetAsText wraps MEOS C function spatialset_as_text.
func SpatialsetAsText(set *Set, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.spatialset_as_text(set._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// SpatialsetAsEWKB wraps MEOS C function spatialset_as_ewkb.
func SpatialsetAsEWKB(s *Set, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.spatialset_as_ewkb(s._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SpatialsetAsEWKT wraps MEOS C function spatialset_as_ewkt.
func SpatialsetAsEWKT(set *Set, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.spatialset_as_ewkt(set._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// SpatialsetAsHexewkb wraps MEOS C function spatialset_as_hexewkb.
func SpatialsetAsHexewkb(s *Set, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.spatialset_as_hexewkb(s._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// GeosetMake wraps MEOS C function geoset_make.
func GeosetMake(values unsafe.Pointer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.geoset_make((**C.GSERIALIZED)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// GeoToSet wraps MEOS C function geo_to_set.
func GeoToSet(gs *Geom) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_to_set(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// GeosetToGeo wraps MEOS C function geoset_to_geo.
func GeosetToGeo(s *Set) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geoset_to_geo(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeosetEndValue wraps MEOS C function geoset_end_value.
func GeosetEndValue(s *Set) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geoset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeosetStartValue wraps MEOS C function geoset_start_value.
func GeosetStartValue(s *Set) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geoset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeosetValueN wraps MEOS C function geoset_value_n.
func GeosetValueN(s *Set, n int) (_r0 bool, _r1 *Geom, _err error) {
	var _out_result *C.GSERIALIZED
	C.meos_errno_reset()
	_cret := C.geoset_value_n(s._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Geom{_inner: _out_result}, nil
}


// GeosetValues wraps MEOS C function geoset_values.
func GeosetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.geoset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// ContainedGeoSet wraps MEOS C function contained_geo_set.
func ContainedGeoSet(gs *Geom, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_geo_set(gs._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSetGeo wraps MEOS C function contains_set_geo.
func ContainsSetGeo(s *Set, gs *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_geo(s._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeoUnionTransfn wraps MEOS C function geo_union_transfn.
func GeoUnionTransfn(state *Set, gs *Geom) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_union_transfn(state._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionGeoSet wraps MEOS C function intersection_geo_set.
func IntersectionGeoSet(gs *Geom, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_geo_set(gs._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetGeo wraps MEOS C function intersection_set_geo.
func IntersectionSetGeo(s *Set, gs *Geom) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_geo(s._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusGeoSet wraps MEOS C function minus_geo_set.
func MinusGeoSet(gs *Geom, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_geo_set(gs._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetGeo wraps MEOS C function minus_set_geo.
func MinusSetGeo(s *Set, gs *Geom) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_geo(s._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionGeoSet wraps MEOS C function union_geo_set.
func UnionGeoSet(gs *Geom, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_geo_set(gs._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetGeo wraps MEOS C function union_set_geo.
func UnionSetGeo(s *Set, gs *Geom) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_geo(s._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// DistanceGeosetGeoset wraps MEOS C function distance_geoset_geoset.
func DistanceGeosetGeoset(s1 *Set, s2 *Set, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_geoset_geoset(s1._inner, s2._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistanceSetGeo wraps MEOS C function distance_set_geo.
func DistanceSetGeo(s *Set, gs *Geom, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_set_geo(s._inner, gs._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// SpatialsetSetSRID wraps MEOS C function spatialset_set_srid.
func SpatialsetSetSRID(s *Set, srid int32) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.spatialset_set_srid(s._inner, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// SpatialsetSRID wraps MEOS C function spatialset_srid.
func SpatialsetSRID(s *Set) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.spatialset_srid(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// SpatialsetTransform wraps MEOS C function spatialset_transform.
func SpatialsetTransform(s *Set, srid int32) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.spatialset_transform(s._inner, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// SpatialsetTransformPipeline wraps MEOS C function spatialset_transform_pipeline.
func SpatialsetTransformPipeline(s *Set, pipelinestr string, srid int32, is_forward bool) (_r0 *Set, _err error) {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	C.meos_errno_reset()
	_cret := C.spatialset_transform_pipeline(s._inner, _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// STBOXAsEWKB wraps MEOS C function stbox_as_ewkb.
func STBOXAsEWKB(box *STBox, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_as_ewkb(box._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// STBOXAsHexewkb wraps MEOS C function stbox_as_hexewkb.
func STBOXAsHexewkb(box *STBox, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_as_hexewkb(box._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// STBOXAsHexwkb wraps MEOS C function stbox_as_hexwkb.
func STBOXAsHexwkb(box *STBox, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_as_hexwkb(box._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// STBOXAsWKB wraps MEOS C function stbox_as_wkb.
func STBOXAsWKB(box *STBox, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_as_wkb(box._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// STBOXFromHexwkb wraps MEOS C function stbox_from_hexwkb.
func STBOXFromHexwkb(hexwkb string) (_r0 *STBox, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.stbox_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXFromWKB wraps MEOS C function stbox_from_wkb.
func STBOXFromWKB(wkb unsafe.Pointer, size uint) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXIn wraps MEOS C function stbox_in.
func STBOXIn(str string) (_r0 *STBox, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.stbox_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXOut wraps MEOS C function stbox_out.
func STBOXOut(box *STBox, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_out(box._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// GeoTimestamptzToSTBOX wraps MEOS C function geo_timestamptz_to_stbox.
func GeoTimestamptzToSTBOX(gs *Geom, t int64) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_timestamptz_to_stbox(gs._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// GeoTstzspanToSTBOX wraps MEOS C function geo_tstzspan_to_stbox.
func GeoTstzspanToSTBOX(gs *Geom, s *Span) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_tstzspan_to_stbox(gs._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXCopy wraps MEOS C function stbox_copy.
func STBOXCopy(box *STBox) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_copy(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXMake wraps MEOS C function stbox_make.
func STBOXMake(hasx bool, hasz bool, geodetic bool, srid int32, xmin float64, xmax float64, ymin float64, ymax float64, zmin float64, zmax float64, s *Span) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_make(C.bool(hasx), C.bool(hasz), C.bool(geodetic), C.int32_t(srid), C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.double(zmin), C.double(zmax), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// GeoToSTBOX wraps MEOS C function geo_to_stbox.
func GeoToSTBOX(gs *Geom) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_to_stbox(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// SpatialsetToSTBOX wraps MEOS C function spatialset_to_stbox.
func SpatialsetToSTBOX(s *Set) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.spatialset_to_stbox(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXToBox3d wraps MEOS C function stbox_to_box3d.
func STBOXToBox3d(box *STBox) (_r0 *Box3D, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_to_box3d(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Box3D{_inner: _cret}, nil
}


// STBOXToGbox wraps MEOS C function stbox_to_gbox.
func STBOXToGbox(box *STBox) (_r0 *GBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_to_gbox(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &GBox{_inner: _cret}, nil
}


// STBOXToGeo wraps MEOS C function stbox_to_geo.
func STBOXToGeo(box *STBox) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_to_geo(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// STBOXToTstzspan wraps MEOS C function stbox_to_tstzspan.
func STBOXToTstzspan(box *STBox) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_to_tstzspan(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TimestamptzToSTBOX wraps MEOS C function timestamptz_to_stbox.
func TimestamptzToSTBOX(t int64) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_to_stbox(C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TstzsetToSTBOX wraps MEOS C function tstzset_to_stbox.
func TstzsetToSTBOX(s *Set) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzset_to_stbox(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TstzspanToSTBOX wraps MEOS C function tstzspan_to_stbox.
func TstzspanToSTBOX(s *Span) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspan_to_stbox(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TstzspansetToSTBOX wraps MEOS C function tstzspanset_to_stbox.
func TstzspansetToSTBOX(ss *SpanSet) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspanset_to_stbox(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXArea wraps MEOS C function stbox_area.
func STBOXArea(box *STBox, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_area(box._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// STBOXHash wraps MEOS C function stbox_hash.
func STBOXHash(box *STBox) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_hash(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// STBOXHashExtended wraps MEOS C function stbox_hash_extended.
func STBOXHashExtended(box *STBox, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_hash_extended(box._inner, C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// STBOXHast wraps MEOS C function stbox_hast.
func STBOXHast(box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_hast(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXHasx wraps MEOS C function stbox_hasx.
func STBOXHasx(box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_hasx(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXHasz wraps MEOS C function stbox_hasz.
func STBOXHasz(box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_hasz(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXIsgeodetic wraps MEOS C function stbox_isgeodetic.
func STBOXIsgeodetic(box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_isgeodetic(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXPerimeter wraps MEOS C function stbox_perimeter.
func STBOXPerimeter(box *STBox, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_perimeter(box._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// STBOXTmax wraps MEOS C function stbox_tmax.
func STBOXTmax(box *STBox) (_r0 bool, _r1 int64, _err error) {
	var _out_result C.TimestampTz
	C.meos_errno_reset()
	_cret := C.stbox_tmax(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int64(_out_result), nil
}


// STBOXTmaxInc wraps MEOS C function stbox_tmax_inc.
func STBOXTmaxInc(box *STBox) (_r0 bool, _r1 bool, _err error) {
	var _out_result C.bool
	C.meos_errno_reset()
	_cret := C.stbox_tmax_inc(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), bool(_out_result), nil
}


// STBOXTmin wraps MEOS C function stbox_tmin.
func STBOXTmin(box *STBox) (_r0 bool, _r1 int64, _err error) {
	var _out_result C.TimestampTz
	C.meos_errno_reset()
	_cret := C.stbox_tmin(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int64(_out_result), nil
}


// STBOXTminInc wraps MEOS C function stbox_tmin_inc.
func STBOXTminInc(box *STBox) (_r0 bool, _r1 bool, _err error) {
	var _out_result C.bool
	C.meos_errno_reset()
	_cret := C.stbox_tmin_inc(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), bool(_out_result), nil
}


// STBOXVolume wraps MEOS C function stbox_volume.
func STBOXVolume(box *STBox) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_volume(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// STBOXXmax wraps MEOS C function stbox_xmax.
func STBOXXmax(box *STBox) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.stbox_xmax(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// STBOXXmin wraps MEOS C function stbox_xmin.
func STBOXXmin(box *STBox) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.stbox_xmin(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// STBOXYmax wraps MEOS C function stbox_ymax.
func STBOXYmax(box *STBox) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.stbox_ymax(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// STBOXYmin wraps MEOS C function stbox_ymin.
func STBOXYmin(box *STBox) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.stbox_ymin(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// STBOXZmax wraps MEOS C function stbox_zmax.
func STBOXZmax(box *STBox) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.stbox_zmax(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// STBOXZmin wraps MEOS C function stbox_zmin.
func STBOXZmin(box *STBox) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.stbox_zmin(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// STBOXExpandSpace wraps MEOS C function stbox_expand_space.
func STBOXExpandSpace(box *STBox, d float64) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_expand_space(box._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXExpandTime wraps MEOS C function stbox_expand_time.
func STBOXExpandTime(box *STBox, interv *Interval) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_expand_time(box._inner, interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXGetSpace wraps MEOS C function stbox_get_space.
func STBOXGetSpace(box *STBox) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_get_space(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXQuadSplit wraps MEOS C function stbox_quad_split.
func STBOXQuadSplit(box *STBox, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_quad_split(box._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXRound wraps MEOS C function stbox_round.
func STBOXRound(box *STBox, maxdd int) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_round(box._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXShiftScaleTime wraps MEOS C function stbox_shift_scale_time.
func STBOXShiftScaleTime(box *STBox, shift *Interval, duration *Interval) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_shift_scale_time(box._inner, shift._inner, duration._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// StboxarrRound wraps MEOS C function stboxarr_round.
func StboxarrRound(boxarr *STBox, count int, maxdd int) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stboxarr_round(boxarr._inner, C.int(count), C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXSetSRID wraps MEOS C function stbox_set_srid.
func STBOXSetSRID(box *STBox, srid int32) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_set_srid(box._inner, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXSRID wraps MEOS C function stbox_srid.
func STBOXSRID(box *STBox) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_srid(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// STBOXTransform wraps MEOS C function stbox_transform.
func STBOXTransform(box *STBox, srid int32) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_transform(box._inner, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXTransformPipeline wraps MEOS C function stbox_transform_pipeline.
func STBOXTransformPipeline(box *STBox, pipelinestr string, srid int32, is_forward bool) (_r0 *STBox, _err error) {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	C.meos_errno_reset()
	_cret := C.stbox_transform_pipeline(box._inner, _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// AdjacentSTBOXSTBOX wraps MEOS C function adjacent_stbox_stbox.
func AdjacentSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedSTBOXSTBOX wraps MEOS C function contained_stbox_stbox.
func ContainedSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSTBOXSTBOX wraps MEOS C function contains_stbox_stbox.
func ContainsSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsSTBOXSTBOX wraps MEOS C function overlaps_stbox_stbox.
func OverlapsSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameSTBOXSTBOX wraps MEOS C function same_stbox_stbox.
func SameSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AboveSTBOXSTBOX wraps MEOS C function above_stbox_stbox.
func AboveSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.above_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterSTBOXSTBOX wraps MEOS C function after_stbox_stbox.
func AfterSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BackSTBOXSTBOX wraps MEOS C function back_stbox_stbox.
func BackSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.back_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeSTBOXSTBOX wraps MEOS C function before_stbox_stbox.
func BeforeSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BelowSTBOXSTBOX wraps MEOS C function below_stbox_stbox.
func BelowSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.below_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// FrontSTBOXSTBOX wraps MEOS C function front_stbox_stbox.
func FrontSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.front_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftSTBOXSTBOX wraps MEOS C function left_stbox_stbox.
func LeftSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OveraboveSTBOXSTBOX wraps MEOS C function overabove_stbox_stbox.
func OveraboveSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overabove_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterSTBOXSTBOX wraps MEOS C function overafter_stbox_stbox.
func OverafterSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbackSTBOXSTBOX wraps MEOS C function overback_stbox_stbox.
func OverbackSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overback_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeSTBOXSTBOX wraps MEOS C function overbefore_stbox_stbox.
func OverbeforeSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbelowSTBOXSTBOX wraps MEOS C function overbelow_stbox_stbox.
func OverbelowSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbelow_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverfrontSTBOXSTBOX wraps MEOS C function overfront_stbox_stbox.
func OverfrontSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overfront_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftSTBOXSTBOX wraps MEOS C function overleft_stbox_stbox.
func OverleftSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightSTBOXSTBOX wraps MEOS C function overright_stbox_stbox.
func OverrightSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightSTBOXSTBOX wraps MEOS C function right_stbox_stbox.
func RightSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// UnionSTBOXSTBOX wraps MEOS C function union_stbox_stbox.
func UnionSTBOXSTBOX(box1 *STBox, box2 *STBox, strict bool) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_stbox_stbox(box1._inner, box2._inner, C.bool(strict))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// IntersectionSTBOXSTBOX wraps MEOS C function intersection_stbox_stbox.
func IntersectionSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_stbox_stbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXCmp wraps MEOS C function stbox_cmp.
func STBOXCmp(box1 *STBox, box2 *STBox) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_cmp(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// STBOXEq wraps MEOS C function stbox_eq.
func STBOXEq(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_eq(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXGe wraps MEOS C function stbox_ge.
func STBOXGe(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_ge(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXGt wraps MEOS C function stbox_gt.
func STBOXGt(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_gt(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXLe wraps MEOS C function stbox_le.
func STBOXLe(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_le(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXLt wraps MEOS C function stbox_lt.
func STBOXLt(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_lt(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXNe wraps MEOS C function stbox_ne.
func STBOXNe(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_ne(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXSortkey wraps MEOS C function stbox_sortkey.
func STBOXSortkey(box *STBox, size unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_sortkey(box._inner, (*C.size_t)(unsafe.Pointer(size)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TspatialOut wraps MEOS C function tspatial_out.
func TspatialOut(temp *Temporal, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tspatial_out(temp._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TgeogpointFromMFJSON wraps MEOS C function tgeogpoint_from_mfjson.
func TgeogpointFromMFJSON(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeogpoint_from_mfjson(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeogpointIn wraps MEOS C function tgeogpoint_in.
func TgeogpointIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeogpoint_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeographyFromMFJSON wraps MEOS C function tgeography_from_mfjson.
func TgeographyFromMFJSON(mfjson string) (_r0 *Temporal, _err error) {
	_c_mfjson := C.CString(mfjson)
	defer C.free(unsafe.Pointer(_c_mfjson))
	C.meos_errno_reset()
	_cret := C.tgeography_from_mfjson(_c_mfjson)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeographyIn wraps MEOS C function tgeography_in.
func TgeographyIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeography_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeometryFromMFJSON wraps MEOS C function tgeometry_from_mfjson.
func TgeometryFromMFJSON(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeometry_from_mfjson(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeometryIn wraps MEOS C function tgeometry_in.
func TgeometryIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeometry_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeompointFromMFJSON wraps MEOS C function tgeompoint_from_mfjson.
func TgeompointFromMFJSON(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeompoint_from_mfjson(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeompointIn wraps MEOS C function tgeompoint_in.
func TgeompointIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeompoint_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TspatialAsEWKB wraps MEOS C function tspatial_as_ewkb.
func TspatialAsEWKB(temp *Temporal, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tspatial_as_ewkb(temp._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TspatialAsEWKT wraps MEOS C function tspatial_as_ewkt.
func TspatialAsEWKT(temp *Temporal, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tspatial_as_ewkt(temp._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TspatialAsHexewkb wraps MEOS C function tspatial_as_hexewkb.
func TspatialAsHexewkb(temp *Temporal, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tspatial_as_hexewkb(temp._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TspatialAsText wraps MEOS C function tspatial_as_text.
func TspatialAsText(temp *Temporal, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tspatial_as_text(temp._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TgeoFromBaseTemp wraps MEOS C function tgeo_from_base_temp.
func TgeoFromBaseTemp(gs *Geom, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_from_base_temp(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoinstMake wraps MEOS C function tgeoinst_make.
func TgeoinstMake(gs *Geom, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeoinst_make(gs._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TgeoseqFromBaseTstzset wraps MEOS C function tgeoseq_from_base_tstzset.
func TgeoseqFromBaseTstzset(gs *Geom, s *Set) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeoseq_from_base_tstzset(gs._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TgeoseqFromBaseTstzspan wraps MEOS C function tgeoseq_from_base_tstzspan.
func TgeoseqFromBaseTstzspan(gs *Geom, s *Span, interp Interpolation) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeoseq_from_base_tstzspan(gs._inner, s._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TgeoseqsetFromBaseTstzspanset wraps MEOS C function tgeoseqset_from_base_tstzspanset.
func TgeoseqsetFromBaseTstzspanset(gs *Geom, ss *SpanSet, interp Interpolation) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeoseqset_from_base_tstzspanset(gs._inner, ss._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TpointFromBaseTemp wraps MEOS C function tpoint_from_base_temp.
func TpointFromBaseTemp(gs *Geom, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_from_base_temp(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointinstMake wraps MEOS C function tpointinst_make.
func TpointinstMake(gs *Geom, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointinst_make(gs._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TpointseqFromBaseTstzset wraps MEOS C function tpointseq_from_base_tstzset.
func TpointseqFromBaseTstzset(gs *Geom, s *Set) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointseq_from_base_tstzset(gs._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TpointseqFromBaseTstzspan wraps MEOS C function tpointseq_from_base_tstzspan.
func TpointseqFromBaseTstzspan(gs *Geom, s *Span, interp Interpolation) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointseq_from_base_tstzspan(gs._inner, s._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TpointseqMakeCoords wraps MEOS C function tpointseq_make_coords.
func TpointseqMakeCoords(xcoords unsafe.Pointer, ycoords unsafe.Pointer, zcoords unsafe.Pointer, times unsafe.Pointer, count int, srid int32, geodetic bool, lower_inc bool, upper_inc bool, interp Interpolation, normalize bool) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointseq_make_coords((*C.double)(unsafe.Pointer(xcoords)), (*C.double)(unsafe.Pointer(ycoords)), (*C.double)(unsafe.Pointer(zcoords)), (*C.TimestampTz)(unsafe.Pointer(times)), C.int(count), C.int32_t(srid), C.bool(geodetic), C.bool(lower_inc), C.bool(upper_inc), C.interpType(interp), C.bool(normalize))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TpointseqsetFromBaseTstzspanset wraps MEOS C function tpointseqset_from_base_tstzspanset.
func TpointseqsetFromBaseTstzspanset(gs *Geom, ss *SpanSet, interp Interpolation) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointseqset_from_base_tstzspanset(gs._inner, ss._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// Box3dToSTBOX wraps MEOS C function box3d_to_stbox.
func Box3dToSTBOX(box *Box3D) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.box3d_to_stbox(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// GboxToSTBOX wraps MEOS C function gbox_to_stbox.
func GboxToSTBOX(box *GBox) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.gbox_to_stbox(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// GeomeasToTpoint wraps MEOS C function geomeas_to_tpoint.
func GeomeasToTpoint(gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.geomeas_to_tpoint(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeogpointToTgeography wraps MEOS C function tgeogpoint_to_tgeography.
func TgeogpointToTgeography(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeogpoint_to_tgeography(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeographyToTgeogpoint wraps MEOS C function tgeography_to_tgeogpoint.
func TgeographyToTgeogpoint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeography_to_tgeogpoint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeographyToTgeometry wraps MEOS C function tgeography_to_tgeometry.
func TgeographyToTgeometry(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeography_to_tgeometry(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeometryToTgeography wraps MEOS C function tgeometry_to_tgeography.
func TgeometryToTgeography(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeometry_to_tgeography(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeogpointToTgeompoint wraps MEOS C function tgeogpoint_to_tgeompoint.
func TgeogpointToTgeompoint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeogpoint_to_tgeompoint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeompointToTgeogpoint wraps MEOS C function tgeompoint_to_tgeogpoint.
func TgeompointToTgeogpoint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeompoint_to_tgeogpoint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeometryToTgeompoint wraps MEOS C function tgeometry_to_tgeompoint.
func TgeometryToTgeompoint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeometry_to_tgeompoint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeompointToTgeometry wraps MEOS C function tgeompoint_to_tgeometry.
func TgeompointToTgeometry(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeompoint_to_tgeometry(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointAsMvtgeom wraps MEOS C function tpoint_as_mvtgeom.
func TpointAsMvtgeom(temp *Temporal, bounds *STBox, extent int32, buffer int32, clip_geom bool, gsarr unsafe.Pointer, timesarr unsafe.Pointer, count unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_as_mvtgeom(temp._inner, bounds._inner, C.int32_t(extent), C.int32_t(buffer), C.bool(clip_geom), (**C.GSERIALIZED)(unsafe.Pointer(gsarr)), (**C.int64_t)(unsafe.Pointer(timesarr)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpointTfloatToGeomeas wraps MEOS C function tpoint_tfloat_to_geomeas.
func TpointTfloatToGeomeas(tpoint *Temporal, measure *Temporal, segmentize bool) (_r0 bool, _r1 *Geom, _err error) {
	var _out_result *C.GSERIALIZED
	C.meos_errno_reset()
	_cret := C.tpoint_tfloat_to_geomeas(tpoint._inner, measure._inner, C.bool(segmentize), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Geom{_inner: _out_result}, nil
}


// TspatialToSTBOX wraps MEOS C function tspatial_to_stbox.
func TspatialToSTBOX(temp *Temporal) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tspatial_to_stbox(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// BearingPointPoint wraps MEOS C function bearing_point_point.
func BearingPointPoint(gs1 *Geom, gs2 *Geom) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.bearing_point_point(gs1._inner, gs2._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// BearingTpointPoint wraps MEOS C function bearing_tpoint_point.
func BearingTpointPoint(temp *Temporal, gs *Geom, invert bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.bearing_tpoint_point(temp._inner, gs._inner, C.bool(invert))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// BearingTpointTpoint wraps MEOS C function bearing_tpoint_tpoint.
func BearingTpointTpoint(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.bearing_tpoint_tpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoCentroid wraps MEOS C function tgeo_centroid.
func TgeoCentroid(temp *Temporal, spheroid bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_centroid(temp._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoConvexHull wraps MEOS C function tgeo_convex_hull.
func TgeoConvexHull(temp *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_convex_hull(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TgeoEndValue wraps MEOS C function tgeo_end_value.
func TgeoEndValue(temp *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TgeoStartValue wraps MEOS C function tgeo_start_value.
func TgeoStartValue(temp *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TgeoTraversedArea wraps MEOS C function tgeo_traversed_area.
func TgeoTraversedArea(temp *Temporal, unary_union bool) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_traversed_area(temp._inner, C.bool(unary_union))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TgeoValue wraps MEOS C function tgeo_value.
func TgeoValue(temp *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TgeoValueAtTimestamptz wraps MEOS C function tgeo_value_at_timestamptz.
func TgeoValueAtTimestamptz(temp *Temporal, t int64, strict bool) (_r0 bool, _r1 *Geom, _err error) {
	var _out_result *C.GSERIALIZED
	C.meos_errno_reset()
	_cret := C.tgeo_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Geom{_inner: _out_result}, nil
}


// TgeoValueN wraps MEOS C function tgeo_value_n.
func TgeoValueN(temp *Temporal, n int) (_r0 bool, _r1 *Geom, _err error) {
	var _out_result *C.GSERIALIZED
	C.meos_errno_reset()
	_cret := C.tgeo_value_n(temp._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Geom{_inner: _out_result}, nil
}


// TgeoValues wraps MEOS C function tgeo_values.
func TgeoValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TgeoUnnest wraps MEOS C function tgeo_unnest.
func TgeoUnnest(temp *Temporal, values unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_unnest(temp._inner, (***C.GSERIALIZED)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TpointAngularDifference wraps MEOS C function tpoint_angular_difference.
func TpointAngularDifference(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_angular_difference(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointAzimuth wraps MEOS C function tpoint_azimuth.
func TpointAzimuth(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_azimuth(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointCumulativeLength wraps MEOS C function tpoint_cumulative_length.
func TpointCumulativeLength(temp *Temporal, spheroid bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_cumulative_length(temp._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointDirection wraps MEOS C function tpoint_direction.
func TpointDirection(temp *Temporal) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.tpoint_direction(temp._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// TpointGetX wraps MEOS C function tpoint_get_x.
func TpointGetX(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_get_x(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointGetY wraps MEOS C function tpoint_get_y.
func TpointGetY(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_get_y(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointGetZ wraps MEOS C function tpoint_get_z.
func TpointGetZ(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_get_z(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointIsSimple wraps MEOS C function tpoint_is_simple.
func TpointIsSimple(temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_is_simple(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpointLength wraps MEOS C function tpoint_length.
func TpointLength(temp *Temporal, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_length(temp._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TpointSpeed wraps MEOS C function tpoint_speed.
func TpointSpeed(temp *Temporal, spheroid bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_speed(temp._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointTrajectory wraps MEOS C function tpoint_trajectory.
func TpointTrajectory(temp *Temporal, unary_union bool) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_trajectory(temp._inner, C.bool(unary_union))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TpointTwcentroid wraps MEOS C function tpoint_twcentroid.
func TpointTwcentroid(temp *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_twcentroid(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TgeoAffine wraps MEOS C function tgeo_affine.
func TgeoAffine(temp *Temporal, a float64, b float64, c float64, d float64, e float64, f float64, g float64, h float64, i float64, xoff float64, yoff float64, zoff float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_affine(temp._inner, C.double(a), C.double(b), C.double(c), C.double(d), C.double(e), C.double(f), C.double(g), C.double(h), C.double(i), C.double(xoff), C.double(yoff), C.double(zoff))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoRotate wraps MEOS C function tgeo_rotate.
func TgeoRotate(temp *Temporal, angle float64, x0 float64, y0 float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_rotate(temp._inner, C.double(angle), C.double(x0), C.double(y0))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoRotateGeo wraps MEOS C function tgeo_rotate_geo.
func TgeoRotateGeo(temp *Temporal, angle float64, origin *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_rotate_geo(temp._inner, C.double(angle), origin._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoRotateX wraps MEOS C function tgeo_rotate_x.
func TgeoRotateX(temp *Temporal, angle float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_rotate_x(temp._inner, C.double(angle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoRotateY wraps MEOS C function tgeo_rotate_y.
func TgeoRotateY(temp *Temporal, angle float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_rotate_y(temp._inner, C.double(angle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoRotateZ wraps MEOS C function tgeo_rotate_z.
func TgeoRotateZ(temp *Temporal, angle float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_rotate_z(temp._inner, C.double(angle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoScale wraps MEOS C function tgeo_scale.
func TgeoScale(temp *Temporal, xfactor float64, yfactor float64, zfactor float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_scale(temp._inner, C.double(xfactor), C.double(yfactor), C.double(zfactor))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoScaleGeo wraps MEOS C function tgeo_scale_geo.
func TgeoScaleGeo(temp *Temporal, scale *Geom, sorigin *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_scale_geo(temp._inner, scale._inner, sorigin._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoTranslate wraps MEOS C function tgeo_translate.
func TgeoTranslate(temp *Temporal, deltax float64, deltay float64, deltaz float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_translate(temp._inner, C.double(deltax), C.double(deltay), C.double(deltaz))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoTransscale wraps MEOS C function tgeo_transscale.
func TgeoTransscale(temp *Temporal, deltax float64, deltay float64, xfactor float64, yfactor float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_transscale(temp._inner, C.double(deltax), C.double(deltay), C.double(xfactor), C.double(yfactor))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointMakeSimple wraps MEOS C function tpoint_make_simple.
func TpointMakeSimple(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_make_simple(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TspatialSRID wraps MEOS C function tspatial_srid.
func TspatialSRID(temp *Temporal) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.tspatial_srid(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// TspatialSetSRID wraps MEOS C function tspatial_set_srid.
func TspatialSetSRID(temp *Temporal, srid int32) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tspatial_set_srid(temp._inner, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TspatialTransform wraps MEOS C function tspatial_transform.
func TspatialTransform(temp *Temporal, srid int32) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tspatial_transform(temp._inner, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TspatialTransformPipeline wraps MEOS C function tspatial_transform_pipeline.
func TspatialTransformPipeline(temp *Temporal, pipelinestr string, srid int32, is_forward bool) (_r0 *Temporal, _err error) {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	C.meos_errno_reset()
	_cret := C.tspatial_transform_pipeline(temp._inner, _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoAtGeom wraps MEOS C function tgeo_at_geom.
func TgeoAtGeom(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_at_geom(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoAtSTBOX wraps MEOS C function tgeo_at_stbox.
func TgeoAtSTBOX(temp *Temporal, box *STBox, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_at_stbox(temp._inner, box._inner, C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoAtValue wraps MEOS C function tgeo_at_value.
func TgeoAtValue(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_at_value(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoMinusGeom wraps MEOS C function tgeo_minus_geom.
func TgeoMinusGeom(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_minus_geom(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoMinusSTBOX wraps MEOS C function tgeo_minus_stbox.
func TgeoMinusSTBOX(temp *Temporal, box *STBox, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_minus_stbox(temp._inner, box._inner, C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoMinusValue wraps MEOS C function tgeo_minus_value.
func TgeoMinusValue(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_minus_value(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointAtElevation wraps MEOS C function tpoint_at_elevation.
func TpointAtElevation(temp *Temporal, s *Span) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_at_elevation(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointAtGeom wraps MEOS C function tpoint_at_geom.
func TpointAtGeom(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_at_geom(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointAtValue wraps MEOS C function tpoint_at_value.
func TpointAtValue(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_at_value(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointMinusElevation wraps MEOS C function tpoint_minus_elevation.
func TpointMinusElevation(temp *Temporal, s *Span) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_minus_elevation(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointMinusGeom wraps MEOS C function tpoint_minus_geom.
func TpointMinusGeom(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_minus_geom(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointMinusValue wraps MEOS C function tpoint_minus_value.
func TpointMinusValue(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_minus_value(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// AlwaysEqGeoTgeo wraps MEOS C function always_eq_geo_tgeo.
func AlwaysEqGeoTgeo(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTgeoGeo wraps MEOS C function always_eq_tgeo_geo.
func AlwaysEqTgeoGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTgeoTgeo wraps MEOS C function always_eq_tgeo_tgeo.
func AlwaysEqTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeGeoTgeo wraps MEOS C function always_ne_geo_tgeo.
func AlwaysNeGeoTgeo(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTgeoGeo wraps MEOS C function always_ne_tgeo_geo.
func AlwaysNeTgeoGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTgeoTgeo wraps MEOS C function always_ne_tgeo_tgeo.
func AlwaysNeTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqGeoTgeo wraps MEOS C function ever_eq_geo_tgeo.
func EverEqGeoTgeo(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTgeoGeo wraps MEOS C function ever_eq_tgeo_geo.
func EverEqTgeoGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTgeoTgeo wraps MEOS C function ever_eq_tgeo_tgeo.
func EverEqTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeGeoTgeo wraps MEOS C function ever_ne_geo_tgeo.
func EverNeGeoTgeo(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTgeoGeo wraps MEOS C function ever_ne_tgeo_geo.
func EverNeTgeoGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTgeoTgeo wraps MEOS C function ever_ne_tgeo_tgeo.
func EverNeTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TeqGeoTgeo wraps MEOS C function teq_geo_tgeo.
func TeqGeoTgeo(gs *Geom, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTgeoGeo wraps MEOS C function teq_tgeo_geo.
func TeqTgeoGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneGeoTgeo wraps MEOS C function tne_geo_tgeo.
func TneGeoTgeo(gs *Geom, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTgeoGeo wraps MEOS C function tne_tgeo_geo.
func TneTgeoGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoStboxes wraps MEOS C function tgeo_stboxes.
func TgeoStboxes(temp *Temporal, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_stboxes(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TgeoSpaceBoxes wraps MEOS C function tgeo_space_boxes.
func TgeoSpaceBoxes(temp *Temporal, xsize float64, ysize float64, zsize float64, sorigin *Geom, bitmatrix bool, border_inc bool, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_space_boxes(temp._inner, C.double(xsize), C.double(ysize), C.double(zsize), sorigin._inner, C.bool(bitmatrix), C.bool(border_inc), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TgeoTimeBoxes wraps MEOS C function tgeo_time_boxes.
func TgeoTimeBoxes(temp *Temporal, duration *Interval, torigin int64, bitmatrix bool, border_inc bool, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_time_boxes(temp._inner, duration._inner, C.TimestampTz(torigin), C.bool(bitmatrix), C.bool(border_inc), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TgeoSpaceTimeBoxes wraps MEOS C function tgeo_space_time_boxes.
func TgeoSpaceTimeBoxes(temp *Temporal, xsize float64, ysize float64, zsize float64, duration *Interval, sorigin *Geom, torigin int64, bitmatrix bool, border_inc bool, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_space_time_boxes(temp._inner, C.double(xsize), C.double(ysize), C.double(zsize), duration._inner, sorigin._inner, C.TimestampTz(torigin), C.bool(bitmatrix), C.bool(border_inc), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TgeoSplitEachNStboxes wraps MEOS C function tgeo_split_each_n_stboxes.
func TgeoSplitEachNStboxes(temp *Temporal, elem_count int, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_split_each_n_stboxes(temp._inner, C.int(elem_count), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TgeoSplitNStboxes wraps MEOS C function tgeo_split_n_stboxes.
func TgeoSplitNStboxes(temp *Temporal, box_count int, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_split_n_stboxes(temp._inner, C.int(box_count), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// AdjacentSTBOXTspatial wraps MEOS C function adjacent_stbox_tspatial.
func AdjacentSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentTspatialSTBOX wraps MEOS C function adjacent_tspatial_stbox.
func AdjacentTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentTspatialTspatial wraps MEOS C function adjacent_tspatial_tspatial.
func AdjacentTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedSTBOXTspatial wraps MEOS C function contained_stbox_tspatial.
func ContainedSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTspatialSTBOX wraps MEOS C function contained_tspatial_stbox.
func ContainedTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTspatialTspatial wraps MEOS C function contained_tspatial_tspatial.
func ContainedTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSTBOXTspatial wraps MEOS C function contains_stbox_tspatial.
func ContainsSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsTspatialSTBOX wraps MEOS C function contains_tspatial_stbox.
func ContainsTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsTspatialTspatial wraps MEOS C function contains_tspatial_tspatial.
func ContainsTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsSTBOXTspatial wraps MEOS C function overlaps_stbox_tspatial.
func OverlapsSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsTspatialSTBOX wraps MEOS C function overlaps_tspatial_stbox.
func OverlapsTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsTspatialTspatial wraps MEOS C function overlaps_tspatial_tspatial.
func OverlapsTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameSTBOXTspatial wraps MEOS C function same_stbox_tspatial.
func SameSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameTspatialSTBOX wraps MEOS C function same_tspatial_stbox.
func SameTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameTspatialTspatial wraps MEOS C function same_tspatial_tspatial.
func SameTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AboveSTBOXTspatial wraps MEOS C function above_stbox_tspatial.
func AboveSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.above_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AboveTspatialSTBOX wraps MEOS C function above_tspatial_stbox.
func AboveTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.above_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AboveTspatialTspatial wraps MEOS C function above_tspatial_tspatial.
func AboveTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.above_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterSTBOXTspatial wraps MEOS C function after_stbox_tspatial.
func AfterSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterTspatialSTBOX wraps MEOS C function after_tspatial_stbox.
func AfterTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterTspatialTspatial wraps MEOS C function after_tspatial_tspatial.
func AfterTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BackSTBOXTspatial wraps MEOS C function back_stbox_tspatial.
func BackSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.back_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BackTspatialSTBOX wraps MEOS C function back_tspatial_stbox.
func BackTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.back_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BackTspatialTspatial wraps MEOS C function back_tspatial_tspatial.
func BackTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.back_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeSTBOXTspatial wraps MEOS C function before_stbox_tspatial.
func BeforeSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeTspatialSTBOX wraps MEOS C function before_tspatial_stbox.
func BeforeTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeTspatialTspatial wraps MEOS C function before_tspatial_tspatial.
func BeforeTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BelowSTBOXTspatial wraps MEOS C function below_stbox_tspatial.
func BelowSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.below_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BelowTspatialSTBOX wraps MEOS C function below_tspatial_stbox.
func BelowTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.below_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BelowTspatialTspatial wraps MEOS C function below_tspatial_tspatial.
func BelowTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.below_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// FrontSTBOXTspatial wraps MEOS C function front_stbox_tspatial.
func FrontSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.front_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// FrontTspatialSTBOX wraps MEOS C function front_tspatial_stbox.
func FrontTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.front_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// FrontTspatialTspatial wraps MEOS C function front_tspatial_tspatial.
func FrontTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.front_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftSTBOXTspatial wraps MEOS C function left_stbox_tspatial.
func LeftSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftTspatialSTBOX wraps MEOS C function left_tspatial_stbox.
func LeftTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftTspatialTspatial wraps MEOS C function left_tspatial_tspatial.
func LeftTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OveraboveSTBOXTspatial wraps MEOS C function overabove_stbox_tspatial.
func OveraboveSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overabove_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OveraboveTspatialSTBOX wraps MEOS C function overabove_tspatial_stbox.
func OveraboveTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overabove_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OveraboveTspatialTspatial wraps MEOS C function overabove_tspatial_tspatial.
func OveraboveTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overabove_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterSTBOXTspatial wraps MEOS C function overafter_stbox_tspatial.
func OverafterSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterTspatialSTBOX wraps MEOS C function overafter_tspatial_stbox.
func OverafterTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterTspatialTspatial wraps MEOS C function overafter_tspatial_tspatial.
func OverafterTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbackSTBOXTspatial wraps MEOS C function overback_stbox_tspatial.
func OverbackSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overback_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbackTspatialSTBOX wraps MEOS C function overback_tspatial_stbox.
func OverbackTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overback_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbackTspatialTspatial wraps MEOS C function overback_tspatial_tspatial.
func OverbackTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overback_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeSTBOXTspatial wraps MEOS C function overbefore_stbox_tspatial.
func OverbeforeSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeTspatialSTBOX wraps MEOS C function overbefore_tspatial_stbox.
func OverbeforeTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeTspatialTspatial wraps MEOS C function overbefore_tspatial_tspatial.
func OverbeforeTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbelowSTBOXTspatial wraps MEOS C function overbelow_stbox_tspatial.
func OverbelowSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbelow_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbelowTspatialSTBOX wraps MEOS C function overbelow_tspatial_stbox.
func OverbelowTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbelow_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbelowTspatialTspatial wraps MEOS C function overbelow_tspatial_tspatial.
func OverbelowTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbelow_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverfrontSTBOXTspatial wraps MEOS C function overfront_stbox_tspatial.
func OverfrontSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overfront_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverfrontTspatialSTBOX wraps MEOS C function overfront_tspatial_stbox.
func OverfrontTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overfront_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverfrontTspatialTspatial wraps MEOS C function overfront_tspatial_tspatial.
func OverfrontTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overfront_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftSTBOXTspatial wraps MEOS C function overleft_stbox_tspatial.
func OverleftSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftTspatialSTBOX wraps MEOS C function overleft_tspatial_stbox.
func OverleftTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftTspatialTspatial wraps MEOS C function overleft_tspatial_tspatial.
func OverleftTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightSTBOXTspatial wraps MEOS C function overright_stbox_tspatial.
func OverrightSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightTspatialSTBOX wraps MEOS C function overright_tspatial_stbox.
func OverrightTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightTspatialTspatial wraps MEOS C function overright_tspatial_tspatial.
func OverrightTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightSTBOXTspatial wraps MEOS C function right_stbox_tspatial.
func RightSTBOXTspatial(box *STBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_stbox_tspatial(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightTspatialSTBOX wraps MEOS C function right_tspatial_stbox.
func RightTspatialSTBOX(temp *Temporal, box *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_tspatial_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightTspatialTspatial wraps MEOS C function right_tspatial_tspatial.
func RightTspatialTspatial(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_tspatial_tspatial(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AcontainsGeoTgeo wraps MEOS C function acontains_geo_tgeo.
func AcontainsGeoTgeo(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acontains_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AcontainsTgeoGeo wraps MEOS C function acontains_tgeo_geo.
func AcontainsTgeoGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acontains_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AcontainsTgeoTgeo wraps MEOS C function acontains_tgeo_tgeo.
func AcontainsTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acontains_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AcoversGeoTgeo wraps MEOS C function acovers_geo_tgeo.
func AcoversGeoTgeo(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acovers_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AcoversTgeoGeo wraps MEOS C function acovers_tgeo_geo.
func AcoversTgeoGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acovers_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AcoversTgeoTgeo wraps MEOS C function acovers_tgeo_tgeo.
func AcoversTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acovers_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdisjointGeoTgeo wraps MEOS C function adisjoint_geo_tgeo.
func AdisjointGeoTgeo(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adisjoint_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdisjointTgeoGeo wraps MEOS C function adisjoint_tgeo_geo.
func AdisjointTgeoGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adisjoint_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdisjointTgeoTgeo wraps MEOS C function adisjoint_tgeo_tgeo.
func AdisjointTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adisjoint_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdwithinGeoTgeo wraps MEOS C function adwithin_geo_tgeo.
func AdwithinGeoTgeo(gs *Geom, temp *Temporal, dist float64, spheroid bool) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adwithin_geo_tgeo(gs._inner, temp._inner, C.double(dist), C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdwithinTgeoGeo wraps MEOS C function adwithin_tgeo_geo.
func AdwithinTgeoGeo(temp *Temporal, gs *Geom, dist float64, spheroid bool) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adwithin_tgeo_geo(temp._inner, gs._inner, C.double(dist), C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdwithinTgeoTgeo wraps MEOS C function adwithin_tgeo_tgeo.
func AdwithinTgeoTgeo(temp1 *Temporal, temp2 *Temporal, dist float64, spheroid bool) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adwithin_tgeo_tgeo(temp1._inner, temp2._inner, C.double(dist), C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AintersectsGeoTgeo wraps MEOS C function aintersects_geo_tgeo.
func AintersectsGeoTgeo(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.aintersects_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AintersectsTgeoGeo wraps MEOS C function aintersects_tgeo_geo.
func AintersectsTgeoGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.aintersects_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AintersectsTgeoTgeo wraps MEOS C function aintersects_tgeo_tgeo.
func AintersectsTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.aintersects_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AtouchesGeoTgeo wraps MEOS C function atouches_geo_tgeo.
func AtouchesGeoTgeo(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.atouches_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AtouchesTgeoGeo wraps MEOS C function atouches_tgeo_geo.
func AtouchesTgeoGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.atouches_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AtouchesTgeoTgeo wraps MEOS C function atouches_tgeo_tgeo.
func AtouchesTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.atouches_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AtouchesTpointGeo wraps MEOS C function atouches_tpoint_geo.
func AtouchesTpointGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.atouches_tpoint_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AtouchesGeoTpoint wraps MEOS C function atouches_geo_tpoint.
func AtouchesGeoTpoint(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.atouches_geo_tpoint(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcontainsGeoTgeo wraps MEOS C function econtains_geo_tgeo.
func EcontainsGeoTgeo(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.econtains_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcontainsTgeoGeo wraps MEOS C function econtains_tgeo_geo.
func EcontainsTgeoGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.econtains_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcontainsTgeoTgeo wraps MEOS C function econtains_tgeo_tgeo.
func EcontainsTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.econtains_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcoversGeoTgeo wraps MEOS C function ecovers_geo_tgeo.
func EcoversGeoTgeo(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ecovers_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcoversTgeoGeo wraps MEOS C function ecovers_tgeo_geo.
func EcoversTgeoGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ecovers_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcoversTgeoTgeo wraps MEOS C function ecovers_tgeo_tgeo.
func EcoversTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ecovers_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdisjointGeoTgeo wraps MEOS C function edisjoint_geo_tgeo.
func EdisjointGeoTgeo(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edisjoint_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdisjointTgeoGeo wraps MEOS C function edisjoint_tgeo_geo.
func EdisjointTgeoGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edisjoint_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdisjointTgeoTgeo wraps MEOS C function edisjoint_tgeo_tgeo.
func EdisjointTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edisjoint_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdwithinGeoTgeo wraps MEOS C function edwithin_geo_tgeo.
func EdwithinGeoTgeo(gs *Geom, temp *Temporal, dist float64, spheroid bool) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edwithin_geo_tgeo(gs._inner, temp._inner, C.double(dist), C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdwithinTgeoGeo wraps MEOS C function edwithin_tgeo_geo.
func EdwithinTgeoGeo(temp *Temporal, gs *Geom, dist float64, spheroid bool) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edwithin_tgeo_geo(temp._inner, gs._inner, C.double(dist), C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdwithinTgeoTgeo wraps MEOS C function edwithin_tgeo_tgeo.
func EdwithinTgeoTgeo(temp1 *Temporal, temp2 *Temporal, dist float64, spheroid bool) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edwithin_tgeo_tgeo(temp1._inner, temp2._inner, C.double(dist), C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EintersectsGeoTgeo wraps MEOS C function eintersects_geo_tgeo.
func EintersectsGeoTgeo(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.eintersects_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EintersectsTgeoGeo wraps MEOS C function eintersects_tgeo_geo.
func EintersectsTgeoGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.eintersects_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EintersectsTgeoTgeo wraps MEOS C function eintersects_tgeo_tgeo.
func EintersectsTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.eintersects_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EtouchesGeoTgeo wraps MEOS C function etouches_geo_tgeo.
func EtouchesGeoTgeo(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.etouches_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EtouchesTgeoGeo wraps MEOS C function etouches_tgeo_geo.
func EtouchesTgeoGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.etouches_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EtouchesTgeoTgeo wraps MEOS C function etouches_tgeo_tgeo.
func EtouchesTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.etouches_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EtouchesTpointGeo wraps MEOS C function etouches_tpoint_geo.
func EtouchesTpointGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.etouches_tpoint_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EtouchesGeoTpoint wraps MEOS C function etouches_geo_tpoint.
func EtouchesGeoTpoint(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.etouches_geo_tpoint(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TcontainsGeoTgeo wraps MEOS C function tcontains_geo_tgeo.
func TcontainsGeoTgeo(gs *Geom, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcontains_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcontainsTgeoGeo wraps MEOS C function tcontains_tgeo_geo.
func TcontainsTgeoGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcontains_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcontainsTgeoTgeo wraps MEOS C function tcontains_tgeo_tgeo.
func TcontainsTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcontains_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcoversGeoTgeo wraps MEOS C function tcovers_geo_tgeo.
func TcoversGeoTgeo(gs *Geom, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcovers_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcoversTgeoGeo wraps MEOS C function tcovers_tgeo_geo.
func TcoversTgeoGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcovers_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcoversTgeoTgeo wraps MEOS C function tcovers_tgeo_tgeo.
func TcoversTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcovers_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdisjointGeoTgeo wraps MEOS C function tdisjoint_geo_tgeo.
func TdisjointGeoTgeo(gs *Geom, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdisjoint_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdisjointTgeoGeo wraps MEOS C function tdisjoint_tgeo_geo.
func TdisjointTgeoGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdisjoint_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdisjointTgeoTgeo wraps MEOS C function tdisjoint_tgeo_tgeo.
func TdisjointTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdisjoint_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdwithinGeoTgeo wraps MEOS C function tdwithin_geo_tgeo.
func TdwithinGeoTgeo(gs *Geom, temp *Temporal, dist float64, spheroid bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdwithin_geo_tgeo(gs._inner, temp._inner, C.double(dist), C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdwithinTgeoGeo wraps MEOS C function tdwithin_tgeo_geo.
func TdwithinTgeoGeo(temp *Temporal, gs *Geom, dist float64, spheroid bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdwithin_tgeo_geo(temp._inner, gs._inner, C.double(dist), C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdwithinTgeoTgeo wraps MEOS C function tdwithin_tgeo_tgeo.
func TdwithinTgeoTgeo(temp1 *Temporal, temp2 *Temporal, dist float64, spheroid bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdwithin_tgeo_tgeo(temp1._inner, temp2._inner, C.double(dist), C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintersectsGeoTgeo wraps MEOS C function tintersects_geo_tgeo.
func TintersectsGeoTgeo(gs *Geom, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tintersects_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintersectsTgeoGeo wraps MEOS C function tintersects_tgeo_geo.
func TintersectsTgeoGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tintersects_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintersectsTgeoTgeo wraps MEOS C function tintersects_tgeo_tgeo.
func TintersectsTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tintersects_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TtouchesGeoTgeo wraps MEOS C function ttouches_geo_tgeo.
func TtouchesGeoTgeo(gs *Geom, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ttouches_geo_tgeo(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TtouchesTgeoGeo wraps MEOS C function ttouches_tgeo_geo.
func TtouchesTgeoGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ttouches_tgeo_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TtouchesTgeoTgeo wraps MEOS C function ttouches_tgeo_tgeo.
func TtouchesTgeoTgeo(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ttouches_tgeo_tgeo(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// EdwithinTgeoarrTgeoarr wraps MEOS C function edwithin_tgeoarr_tgeoarr.
func EdwithinTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, dist float64, spheroid bool, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.edwithin_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), C.double(dist), C.bool(spheroid), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// AdwithinTgeoarrTgeoarr wraps MEOS C function adwithin_tgeoarr_tgeoarr.
func AdwithinTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, dist float64, spheroid bool, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.adwithin_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), C.double(dist), C.bool(spheroid), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// EintersectsTgeoarrTgeoarr wraps MEOS C function eintersects_tgeoarr_tgeoarr.
func EintersectsTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.eintersects_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// AintersectsTgeoarrTgeoarr wraps MEOS C function aintersects_tgeoarr_tgeoarr.
func AintersectsTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.aintersects_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// EtouchesTgeoarrTgeoarr wraps MEOS C function etouches_tgeoarr_tgeoarr.
func EtouchesTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.etouches_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// AtouchesTgeoarrTgeoarr wraps MEOS C function atouches_tgeoarr_tgeoarr.
func AtouchesTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.atouches_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// EdisjointTgeoarrTgeoarr wraps MEOS C function edisjoint_tgeoarr_tgeoarr.
func EdisjointTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.edisjoint_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// AdisjointTgeoarrTgeoarr wraps MEOS C function adisjoint_tgeoarr_tgeoarr.
func AdisjointTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.adisjoint_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TdwithinTgeoarrTgeoarr wraps MEOS C function tdwithin_tgeoarr_tgeoarr.
func TdwithinTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, dist float64, spheroid bool, count unsafe.Pointer, periods unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tdwithin_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), C.double(dist), C.bool(spheroid), (*C.int)(unsafe.Pointer(count)), (***C.SpanSet)(unsafe.Pointer(periods)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TintersectsTgeoarrTgeoarr wraps MEOS C function tintersects_tgeoarr_tgeoarr.
func TintersectsTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer, periods unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tintersects_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)), (***C.SpanSet)(unsafe.Pointer(periods)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TtouchesTgeoarrTgeoarr wraps MEOS C function ttouches_tgeoarr_tgeoarr.
func TtouchesTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer, periods unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.ttouches_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)), (***C.SpanSet)(unsafe.Pointer(periods)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TdisjointTgeoarrTgeoarr wraps MEOS C function tdisjoint_tgeoarr_tgeoarr.
func TdisjointTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, count unsafe.Pointer, periods unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tdisjoint_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), (*C.int)(unsafe.Pointer(count)), (***C.SpanSet)(unsafe.Pointer(periods)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TdistanceTgeoGeo wraps MEOS C function tdistance_tgeo_geo.
func TdistanceTgeoGeo(temp *Temporal, gs *Geom, spheroid bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_tgeo_geo(temp._inner, gs._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTgeoTgeo wraps MEOS C function tdistance_tgeo_tgeo.
func TdistanceTgeoTgeo(temp1 *Temporal, temp2 *Temporal, spheroid bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_tgeo_tgeo(temp1._inner, temp2._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// NadSTBOXGeo wraps MEOS C function nad_stbox_geo.
func NadSTBOXGeo(box *STBox, gs *Geom, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_stbox_geo(box._inner, gs._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadSTBOXSTBOX wraps MEOS C function nad_stbox_stbox.
func NadSTBOXSTBOX(box1 *STBox, box2 *STBox, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_stbox_stbox(box1._inner, box2._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// STBOXSpatialDistance wraps MEOS C function stbox_spatial_distance.
func STBOXSpatialDistance(box1 *STBox, box2 *STBox) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_spatial_distance(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTgeoGeo wraps MEOS C function nad_tgeo_geo.
func NadTgeoGeo(temp *Temporal, gs *Geom, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tgeo_geo(temp._inner, gs._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTgeoSTBOX wraps MEOS C function nad_tgeo_stbox.
func NadTgeoSTBOX(temp *Temporal, box *STBox, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tgeo_stbox(temp._inner, box._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTgeoTgeo wraps MEOS C function nad_tgeo_tgeo.
func NadTgeoTgeo(temp1 *Temporal, temp2 *Temporal, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tgeo_tgeo(temp1._inner, temp2._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NaiTgeoGeo wraps MEOS C function nai_tgeo_geo.
func NaiTgeoGeo(temp *Temporal, gs *Geom, spheroid bool) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.nai_tgeo_geo(temp._inner, gs._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// NaiTgeoTgeo wraps MEOS C function nai_tgeo_tgeo.
func NaiTgeoTgeo(temp1 *Temporal, temp2 *Temporal, spheroid bool) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.nai_tgeo_tgeo(temp1._inner, temp2._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// ShortestlineTgeoGeo wraps MEOS C function shortestline_tgeo_geo.
func ShortestlineTgeoGeo(temp *Temporal, gs *Geom, spheroid bool) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.shortestline_tgeo_geo(temp._inner, gs._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// ShortestlineTgeoTgeo wraps MEOS C function shortestline_tgeo_tgeo.
func ShortestlineTgeoTgeo(temp1 *Temporal, temp2 *Temporal, spheroid bool) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.shortestline_tgeo_tgeo(temp1._inner, temp2._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// MindistanceTgeoarrTgeoarr wraps MEOS C function mindistance_tgeoarr_tgeoarr.
func MindistanceTgeoarrTgeoarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.mindistance_tgeoarr_tgeoarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2), C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TpointTcentroidCombinefn wraps MEOS C function tpoint_tcentroid_combinefn.
func TpointTcentroidCombinefn(state1 *SkipList, state2 *SkipList) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_tcentroid_combinefn(state1._inner, state2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TpointTcentroidFinalfn wraps MEOS C function tpoint_tcentroid_finalfn.
func TpointTcentroidFinalfn(state *SkipList) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_tcentroid_finalfn(state._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointTcentroidTransfn wraps MEOS C function tpoint_tcentroid_transfn.
func TpointTcentroidTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_tcentroid_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// STBOXExtentTransfn wraps MEOS C function stbox_extent_transfn.
func STBOXExtentTransfn(state *STBox, box *STBox) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_extent_transfn(state._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TspatialExtentTransfn wraps MEOS C function tspatial_extent_transfn.
func TspatialExtentTransfn(state *STBox, temp *Temporal) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tspatial_extent_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXGetSpaceTile wraps MEOS C function stbox_get_space_tile.
func STBOXGetSpaceTile(point *Geom, xsize float64, ysize float64, zsize float64, sorigin *Geom) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_get_space_tile(point._inner, C.double(xsize), C.double(ysize), C.double(zsize), sorigin._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXGetSpaceTimeTile wraps MEOS C function stbox_get_space_time_tile.
func STBOXGetSpaceTimeTile(point *Geom, t int64, xsize float64, ysize float64, zsize float64, duration *Interval, sorigin *Geom, torigin int64) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_get_space_time_tile(point._inner, C.TimestampTz(t), C.double(xsize), C.double(ysize), C.double(zsize), duration._inner, sorigin._inner, C.TimestampTz(torigin))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXGetTimeTile wraps MEOS C function stbox_get_time_tile.
func STBOXGetTimeTile(t int64, duration *Interval, torigin int64) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_get_time_tile(C.TimestampTz(t), duration._inner, C.TimestampTz(torigin))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXSpaceTiles wraps MEOS C function stbox_space_tiles.
func STBOXSpaceTiles(bounds *STBox, xsize float64, ysize float64, zsize float64, sorigin *Geom, border_inc bool, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_space_tiles(bounds._inner, C.double(xsize), C.double(ysize), C.double(zsize), sorigin._inner, C.bool(border_inc), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXSpaceTimeTiles wraps MEOS C function stbox_space_time_tiles.
func STBOXSpaceTimeTiles(bounds *STBox, xsize float64, ysize float64, zsize float64, duration *Interval, sorigin *Geom, torigin int64, border_inc bool, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_space_time_tiles(bounds._inner, C.double(xsize), C.double(ysize), C.double(zsize), duration._inner, sorigin._inner, C.TimestampTz(torigin), C.bool(border_inc), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// STBOXTimeTiles wraps MEOS C function stbox_time_tiles.
func STBOXTimeTiles(bounds *STBox, duration *Interval, torigin int64, border_inc bool, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_time_tiles(bounds._inner, duration._inner, C.TimestampTz(torigin), C.bool(border_inc), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TgeoSpaceSplit wraps MEOS C function tgeo_space_split.
func TgeoSpaceSplit(temp *Temporal, xsize float64, ysize float64, zsize float64, sorigin *Geom, bitmatrix bool, border_inc bool, space_bins unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_space_split(temp._inner, C.double(xsize), C.double(ysize), C.double(zsize), sorigin._inner, C.bool(bitmatrix), C.bool(border_inc), (***C.GSERIALIZED)(unsafe.Pointer(space_bins)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TgeoSpaceTimeSplit wraps MEOS C function tgeo_space_time_split.
func TgeoSpaceTimeSplit(temp *Temporal, xsize float64, ysize float64, zsize float64, duration *Interval, sorigin *Geom, torigin int64, bitmatrix bool, border_inc bool, space_bins unsafe.Pointer, time_bins unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_space_time_split(temp._inner, C.double(xsize), C.double(ysize), C.double(zsize), duration._inner, sorigin._inner, C.TimestampTz(torigin), C.bool(bitmatrix), C.bool(border_inc), (***C.GSERIALIZED)(unsafe.Pointer(space_bins)), (**C.TimestampTz)(unsafe.Pointer(time_bins)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// GeoClusterKmeans wraps MEOS C function geo_cluster_kmeans.
func GeoClusterKmeans(geoms unsafe.Pointer, ngeoms uint32, k uint32, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_cluster_kmeans((**C.GSERIALIZED)(unsafe.Pointer(geoms)), C.uint32_t(ngeoms), C.uint32_t(k), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// GeoClusterDbscan wraps MEOS C function geo_cluster_dbscan.
func GeoClusterDbscan(geoms unsafe.Pointer, ngeoms uint32, tolerance float64, minpoints int, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_cluster_dbscan((**C.GSERIALIZED)(unsafe.Pointer(geoms)), C.uint32_t(ngeoms), C.double(tolerance), C.int(minpoints), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// GeoClusterIntersecting wraps MEOS C function geo_cluster_intersecting.
func GeoClusterIntersecting(geoms unsafe.Pointer, ngeoms uint32, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_cluster_intersecting((**C.GSERIALIZED)(unsafe.Pointer(geoms)), C.uint32_t(ngeoms), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// GeoClusterWithin wraps MEOS C function geo_cluster_within.
func GeoClusterWithin(geoms unsafe.Pointer, ngeoms uint32, tolerance float64, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_cluster_within((**C.GSERIALIZED)(unsafe.Pointer(geoms)), C.uint32_t(ngeoms), C.double(tolerance), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}

