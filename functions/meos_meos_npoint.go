package functions

/*
#include <stddef.h>
#include "meos.h"
#include "meos_catalog.h"
#include "meos_geo.h"
#include "meos_internal.h"
#include "meos_internal_geo.h"
#include "meos_npoint.h"
#include "meos_cbuffer.h"
#include "meos_pose.h"
#include "meos_rgeo.h"
#include "meos_h3.h"
#include "meos_quadbin.h"
#include "meos_json.h"
#include "meos_pointcloud.h"

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

// NpointAsEWKB wraps MEOS C function npoint_as_ewkb.
func NpointAsEWKB(np *Npoint, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_as_ewkb(np._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// NpointAsEWKT wraps MEOS C function npoint_as_ewkt.
func NpointAsEWKT(np *Npoint, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_as_ewkt(np._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// NpointAsHexewkb wraps MEOS C function npoint_as_hexewkb.
func NpointAsHexewkb(np *Npoint, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_as_hexewkb(np._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// NpointAsHexwkb wraps MEOS C function npoint_as_hexwkb.
func NpointAsHexwkb(np *Npoint, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_as_hexwkb(np._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// NpointAsText wraps MEOS C function npoint_as_text.
func NpointAsText(np *Npoint, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_as_text(np._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// NpointAsWKB wraps MEOS C function npoint_as_wkb.
func NpointAsWKB(np *Npoint, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_as_wkb(np._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// NpointFromHexwkb wraps MEOS C function npoint_from_hexwkb.
func NpointFromHexwkb(hexwkb string) (_r0 *Npoint, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.npoint_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return &Npoint{_inner: _cret}, nil
}


// NpointFromWKB wraps MEOS C function npoint_from_wkb.
func NpointFromWKB(wkb unsafe.Pointer, size uint) (_r0 *Npoint, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return &Npoint{_inner: _cret}, nil
}


// NpointIn wraps MEOS C function npoint_in.
func NpointIn(str string) (_r0 *Npoint, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.npoint_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Npoint{_inner: _cret}, nil
}


// NpointOut wraps MEOS C function npoint_out.
func NpointOut(np *Npoint, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_out(np._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// NsegmentAsHexewkb wraps MEOS C function nsegment_as_hexewkb.
func NsegmentAsHexewkb(ns *Nsegment, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_as_hexewkb(ns._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// NsegmentAsHexwkb wraps MEOS C function nsegment_as_hexwkb.
func NsegmentAsHexwkb(ns *Nsegment, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_as_hexwkb(ns._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// NsegmentAsWKB wraps MEOS C function nsegment_as_wkb.
func NsegmentAsWKB(ns *Nsegment, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_as_wkb(ns._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// NsegmentFromHexwkb wraps MEOS C function nsegment_from_hexwkb.
func NsegmentFromHexwkb(hexwkb string) (_r0 *Nsegment, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.nsegment_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return &Nsegment{_inner: _cret}, nil
}


// NsegmentFromWKB wraps MEOS C function nsegment_from_wkb.
func NsegmentFromWKB(wkb unsafe.Pointer, size uint) (_r0 *Nsegment, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return &Nsegment{_inner: _cret}, nil
}


// NsegmentIn wraps MEOS C function nsegment_in.
func NsegmentIn(str string) (_r0 *Nsegment, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.nsegment_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Nsegment{_inner: _cret}, nil
}


// NsegmentOut wraps MEOS C function nsegment_out.
func NsegmentOut(ns *Nsegment, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_out(ns._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// NsegmentAsEWKB wraps MEOS C function nsegment_as_ewkb.
func NsegmentAsEWKB(ns *Nsegment, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_as_ewkb(ns._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// NsegmentAsEWKT wraps MEOS C function nsegment_as_ewkt.
func NsegmentAsEWKT(ns *Nsegment, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_as_ewkt(ns._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// NsegmentAsText wraps MEOS C function nsegment_as_text.
func NsegmentAsText(ns *Nsegment, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_as_text(ns._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// NpointCopy wraps MEOS C function npoint_copy.
func NpointCopy(np *Npoint) (_r0 *Npoint, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_copy(np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Npoint{_inner: _cret}, nil
}


// NpointMake wraps MEOS C function npoint_make.
func NpointMake(rid int64, pos float64) (_r0 *Npoint, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_make(C.int64_t(rid), C.double(pos))
	if _err = meosError(); _err != nil {
		return
	}
	return &Npoint{_inner: _cret}, nil
}


// NsegmentCopy wraps MEOS C function nsegment_copy.
func NsegmentCopy(ns *Nsegment) (_r0 *Nsegment, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_copy(ns._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Nsegment{_inner: _cret}, nil
}


// NsegmentMake wraps MEOS C function nsegment_make.
func NsegmentMake(rid int64, pos1 float64, pos2 float64) (_r0 *Nsegment, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_make(C.int64_t(rid), C.double(pos1), C.double(pos2))
	if _err = meosError(); _err != nil {
		return
	}
	return &Nsegment{_inner: _cret}, nil
}


// GeompointToNpoint wraps MEOS C function geompoint_to_npoint.
func GeompointToNpoint(gs *Geom) (_r0 *Npoint, _err error) {
	C.meos_errno_reset()
	_cret := C.geompoint_to_npoint(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Npoint{_inner: _cret}, nil
}


// GeomToNsegment wraps MEOS C function geom_to_nsegment.
func GeomToNsegment(gs *Geom) (_r0 *Nsegment, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_to_nsegment(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Nsegment{_inner: _cret}, nil
}


// NpointToGeompoint wraps MEOS C function npoint_to_geompoint.
func NpointToGeompoint(np *Npoint) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_to_geompoint(np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// NpointToNsegment wraps MEOS C function npoint_to_nsegment.
func NpointToNsegment(np *Npoint) (_r0 *Nsegment, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_to_nsegment(np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Nsegment{_inner: _cret}, nil
}


// NpointToSTBOX wraps MEOS C function npoint_to_stbox.
func NpointToSTBOX(np *Npoint) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_to_stbox(np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// NsegmentToGeom wraps MEOS C function nsegment_to_geom.
func NsegmentToGeom(ns *Nsegment) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_to_geom(ns._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// NsegmentToSTBOX wraps MEOS C function nsegment_to_stbox.
func NsegmentToSTBOX(ns *Nsegment) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_to_stbox(ns._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// NpointHash wraps MEOS C function npoint_hash.
func NpointHash(np *Npoint) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_hash(np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// NpointHashExtended wraps MEOS C function npoint_hash_extended.
func NpointHashExtended(np *Npoint, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_hash_extended(np._inner, C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// NpointPosition wraps MEOS C function npoint_position.
func NpointPosition(np *Npoint) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_position(np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NpointRoute wraps MEOS C function npoint_route.
func NpointRoute(np *Npoint) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_route(np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// NsegmentEndPosition wraps MEOS C function nsegment_end_position.
func NsegmentEndPosition(ns *Nsegment) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_end_position(ns._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NsegmentRoute wraps MEOS C function nsegment_route.
func NsegmentRoute(ns *Nsegment) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_route(ns._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// NsegmentStartPosition wraps MEOS C function nsegment_start_position.
func NsegmentStartPosition(ns *Nsegment) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_start_position(ns._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// RouteExists wraps MEOS C function route_exists.
func RouteExists(rid int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.route_exists(C.int64_t(rid))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RouteGeom wraps MEOS C function route_geom.
func RouteGeom(rid int64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.route_geom(C.int64_t(rid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// RouteLength wraps MEOS C function route_length.
func RouteLength(rid int64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.route_length(C.int64_t(rid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NpointRound wraps MEOS C function npoint_round.
func NpointRound(np *Npoint, maxdd int) (_r0 *Npoint, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_round(np._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return &Npoint{_inner: _cret}, nil
}


// NsegmentRound wraps MEOS C function nsegment_round.
func NsegmentRound(ns *Nsegment, maxdd int) (_r0 *Nsegment, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_round(ns._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return &Nsegment{_inner: _cret}, nil
}


// GetSRIDWays wraps MEOS C function get_srid_ways.
func GetSRIDWays() (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.get_srid_ways()
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// NpointSRID wraps MEOS C function npoint_srid.
func NpointSRID(np *Npoint) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_srid(np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// NsegmentSRID wraps MEOS C function nsegment_srid.
func NsegmentSRID(ns *Nsegment) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_srid(ns._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// NpointTimestamptzToSTBOX wraps MEOS C function npoint_timestamptz_to_stbox.
func NpointTimestamptzToSTBOX(np *Npoint, t int64) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_timestamptz_to_stbox(np._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// NpointTstzspanToSTBOX wraps MEOS C function npoint_tstzspan_to_stbox.
func NpointTstzspanToSTBOX(np *Npoint, s *Span) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_tstzspan_to_stbox(np._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// NpointCmp wraps MEOS C function npoint_cmp.
func NpointCmp(np1 *Npoint, np2 *Npoint) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_cmp(np1._inner, np2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// NpointEq wraps MEOS C function npoint_eq.
func NpointEq(np1 *Npoint, np2 *Npoint) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_eq(np1._inner, np2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NpointGe wraps MEOS C function npoint_ge.
func NpointGe(np1 *Npoint, np2 *Npoint) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_ge(np1._inner, np2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NpointGt wraps MEOS C function npoint_gt.
func NpointGt(np1 *Npoint, np2 *Npoint) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_gt(np1._inner, np2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NpointLe wraps MEOS C function npoint_le.
func NpointLe(np1 *Npoint, np2 *Npoint) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_le(np1._inner, np2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NpointLt wraps MEOS C function npoint_lt.
func NpointLt(np1 *Npoint, np2 *Npoint) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_lt(np1._inner, np2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NpointNe wraps MEOS C function npoint_ne.
func NpointNe(np1 *Npoint, np2 *Npoint) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_ne(np1._inner, np2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NsegmentCmp wraps MEOS C function nsegment_cmp.
func NsegmentCmp(ns1 *Nsegment, ns2 *Nsegment) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_cmp(ns1._inner, ns2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// NsegmentEq wraps MEOS C function nsegment_eq.
func NsegmentEq(ns1 *Nsegment, ns2 *Nsegment) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_eq(ns1._inner, ns2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NsegmentGe wraps MEOS C function nsegment_ge.
func NsegmentGe(ns1 *Nsegment, ns2 *Nsegment) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_ge(ns1._inner, ns2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NsegmentGt wraps MEOS C function nsegment_gt.
func NsegmentGt(ns1 *Nsegment, ns2 *Nsegment) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_gt(ns1._inner, ns2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NsegmentLe wraps MEOS C function nsegment_le.
func NsegmentLe(ns1 *Nsegment, ns2 *Nsegment) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_le(ns1._inner, ns2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NsegmentLt wraps MEOS C function nsegment_lt.
func NsegmentLt(ns1 *Nsegment, ns2 *Nsegment) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_lt(ns1._inner, ns2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NsegmentNe wraps MEOS C function nsegment_ne.
func NsegmentNe(ns1 *Nsegment, ns2 *Nsegment) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.nsegment_ne(ns1._inner, ns2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NpointsetIn wraps MEOS C function npointset_in.
func NpointsetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.npointset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// NpointsetOut wraps MEOS C function npointset_out.
func NpointsetOut(s *Set, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.npointset_out(s._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// NpointsetMake wraps MEOS C function npointset_make.
func NpointsetMake(values *Npoint, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.npointset_make(values._inner, C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// NpointToSet wraps MEOS C function npoint_to_set.
func NpointToSet(np *Npoint) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_to_set(np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// NpointsetEndValue wraps MEOS C function npointset_end_value.
func NpointsetEndValue(s *Set) (_r0 *Npoint, _err error) {
	C.meos_errno_reset()
	_cret := C.npointset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Npoint{_inner: _cret}, nil
}


// NpointsetRoutes wraps MEOS C function npointset_routes.
func NpointsetRoutes(s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.npointset_routes(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// NpointsetStartValue wraps MEOS C function npointset_start_value.
func NpointsetStartValue(s *Set) (_r0 *Npoint, _err error) {
	C.meos_errno_reset()
	_cret := C.npointset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Npoint{_inner: _cret}, nil
}


// NpointsetValueN wraps MEOS C function npointset_value_n.
func NpointsetValueN(s *Set, n int) (_r0 bool, _r1 *Npoint, _err error) {
	var _out_result *C.Npoint
	C.meos_errno_reset()
	_cret := C.npointset_value_n(s._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Npoint{_inner: _out_result}, nil
}


// NpointsetValues wraps MEOS C function npointset_values.
func NpointsetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.npointset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// ContainedNpointSet wraps MEOS C function contained_npoint_set.
func ContainedNpointSet(np *Npoint, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_npoint_set(np._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSetNpoint wraps MEOS C function contains_set_npoint.
func ContainsSetNpoint(s *Set, np *Npoint) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_npoint(s._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// IntersectionNpointSet wraps MEOS C function intersection_npoint_set.
func IntersectionNpointSet(np *Npoint, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_npoint_set(np._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetNpoint wraps MEOS C function intersection_set_npoint.
func IntersectionSetNpoint(s *Set, np *Npoint) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_npoint(s._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusNpointSet wraps MEOS C function minus_npoint_set.
func MinusNpointSet(np *Npoint, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_npoint_set(np._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetNpoint wraps MEOS C function minus_set_npoint.
func MinusSetNpoint(s *Set, np *Npoint) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_npoint(s._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// NpointUnionTransfn wraps MEOS C function npoint_union_transfn.
func NpointUnionTransfn(state *Set, np *Npoint) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.npoint_union_transfn(state._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionNpointSet wraps MEOS C function union_npoint_set.
func UnionNpointSet(np *Npoint, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_npoint_set(np._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetNpoint wraps MEOS C function union_set_npoint.
func UnionSetNpoint(s *Set, np *Npoint) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_npoint(s._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TnpointIn wraps MEOS C function tnpoint_in.
func TnpointIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tnpoint_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnpointFromMFJSON wraps MEOS C function tnpoint_from_mfjson.
func TnpointFromMFJSON(mfjson string) (_r0 *Temporal, _err error) {
	_c_mfjson := C.CString(mfjson)
	defer C.free(unsafe.Pointer(_c_mfjson))
	C.meos_errno_reset()
	_cret := C.tnpoint_from_mfjson(_c_mfjson)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnpointOut wraps MEOS C function tnpoint_out.
func TnpointOut(temp *Temporal, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_out(temp._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TnpointinstMake wraps MEOS C function tnpointinst_make.
func TnpointinstMake(np *Npoint, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpointinst_make(np._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TnpointFromBaseTemp wraps MEOS C function tnpoint_from_base_temp.
func TnpointFromBaseTemp(np *Npoint, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_from_base_temp(np._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnpointseqFromBaseTstzset wraps MEOS C function tnpointseq_from_base_tstzset.
func TnpointseqFromBaseTstzset(np *Npoint, s *Set) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpointseq_from_base_tstzset(np._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TnpointseqFromBaseTstzspan wraps MEOS C function tnpointseq_from_base_tstzspan.
func TnpointseqFromBaseTstzspan(np *Npoint, s *Span, interp Interpolation) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpointseq_from_base_tstzspan(np._inner, s._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TnpointseqsetFromBaseTstzspanset wraps MEOS C function tnpointseqset_from_base_tstzspanset.
func TnpointseqsetFromBaseTstzspanset(np *Npoint, ss *SpanSet, interp Interpolation) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpointseqset_from_base_tstzspanset(np._inner, ss._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TgeompointToTnpoint wraps MEOS C function tgeompoint_to_tnpoint.
func TgeompointToTnpoint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeompoint_to_tnpoint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnpointToTgeompoint wraps MEOS C function tnpoint_to_tgeompoint.
func TnpointToTgeompoint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_to_tgeompoint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnpointCumulativeLength wraps MEOS C function tnpoint_cumulative_length.
func TnpointCumulativeLength(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_cumulative_length(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnpointEndValue wraps MEOS C function tnpoint_end_value.
func TnpointEndValue(temp *Temporal) (_r0 *Npoint, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Npoint{_inner: _cret}, nil
}


// TnpointLength wraps MEOS C function tnpoint_length.
func TnpointLength(temp *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_length(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TnpointPositions wraps MEOS C function tnpoint_positions.
func TnpointPositions(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_positions(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TnpointRoute wraps MEOS C function tnpoint_route.
func TnpointRoute(temp *Temporal) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_route(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TnpointRoutes wraps MEOS C function tnpoint_routes.
func TnpointRoutes(temp *Temporal) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_routes(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TnpointSpeed wraps MEOS C function tnpoint_speed.
func TnpointSpeed(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_speed(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnpointStartValue wraps MEOS C function tnpoint_start_value.
func TnpointStartValue(temp *Temporal) (_r0 *Npoint, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Npoint{_inner: _cret}, nil
}


// TnpointTrajectory wraps MEOS C function tnpoint_trajectory.
func TnpointTrajectory(temp *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_trajectory(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TnpointValue wraps MEOS C function tnpoint_value.
func TnpointValue(temp *Temporal) (_r0 *Npoint, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Npoint{_inner: _cret}, nil
}


// TnpointValueAtTimestamptz wraps MEOS C function tnpoint_value_at_timestamptz.
func TnpointValueAtTimestamptz(temp *Temporal, t int64, strict bool) (_r0 bool, _r1 *Npoint, _err error) {
	var _out_value *C.Npoint
	C.meos_errno_reset()
	_cret := C.tnpoint_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Npoint{_inner: _out_value}, nil
}


// TnpointValueN wraps MEOS C function tnpoint_value_n.
func TnpointValueN(temp *Temporal, n int) (_r0 bool, _r1 *Npoint, _err error) {
	var _out_result *C.Npoint
	C.meos_errno_reset()
	_cret := C.tnpoint_value_n(temp._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Npoint{_inner: _out_result}, nil
}


// TnpointValues wraps MEOS C function tnpoint_values.
func TnpointValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TnpointUnnest wraps MEOS C function tnpoint_unnest.
func TnpointUnnest(temp *Temporal, values unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_unnest(temp._inner, (***C.Npoint)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TnpointTwcentroid wraps MEOS C function tnpoint_twcentroid.
func TnpointTwcentroid(temp *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_twcentroid(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TnpointAtGeom wraps MEOS C function tnpoint_at_geom.
func TnpointAtGeom(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_at_geom(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnpointAtNpoint wraps MEOS C function tnpoint_at_npoint.
func TnpointAtNpoint(temp *Temporal, np *Npoint) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_at_npoint(temp._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnpointAtNpointset wraps MEOS C function tnpoint_at_npointset.
func TnpointAtNpointset(temp *Temporal, s *Set) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_at_npointset(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnpointAtSTBOX wraps MEOS C function tnpoint_at_stbox.
func TnpointAtSTBOX(temp *Temporal, box *STBox, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_at_stbox(temp._inner, box._inner, C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnpointMinusGeom wraps MEOS C function tnpoint_minus_geom.
func TnpointMinusGeom(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_minus_geom(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnpointMinusNpoint wraps MEOS C function tnpoint_minus_npoint.
func TnpointMinusNpoint(temp *Temporal, np *Npoint) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_minus_npoint(temp._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnpointMinusNpointset wraps MEOS C function tnpoint_minus_npointset.
func TnpointMinusNpointset(temp *Temporal, s *Set) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_minus_npointset(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnpointMinusSTBOX wraps MEOS C function tnpoint_minus_stbox.
func TnpointMinusSTBOX(temp *Temporal, box *STBox, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_minus_stbox(temp._inner, box._inner, C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTnpointNpoint wraps MEOS C function tdistance_tnpoint_npoint.
func TdistanceTnpointNpoint(temp *Temporal, np *Npoint) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_tnpoint_npoint(temp._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTnpointGeo wraps MEOS C function tdistance_tnpoint_geo.
func TdistanceTnpointGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_tnpoint_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTnpointTnpoint wraps MEOS C function tdistance_tnpoint_tnpoint.
func TdistanceTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_tnpoint_tnpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// NadTnpointGeo wraps MEOS C function nad_tnpoint_geo.
func NadTnpointGeo(temp *Temporal, gs *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tnpoint_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTnpointNpoint wraps MEOS C function nad_tnpoint_npoint.
func NadTnpointNpoint(temp *Temporal, np *Npoint) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tnpoint_npoint(temp._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTnpointSTBOX wraps MEOS C function nad_tnpoint_stbox.
func NadTnpointSTBOX(temp *Temporal, box *STBox) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tnpoint_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTnpointTnpoint wraps MEOS C function nad_tnpoint_tnpoint.
func NadTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tnpoint_tnpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NaiTnpointGeo wraps MEOS C function nai_tnpoint_geo.
func NaiTnpointGeo(temp *Temporal, gs *Geom) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.nai_tnpoint_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// NaiTnpointNpoint wraps MEOS C function nai_tnpoint_npoint.
func NaiTnpointNpoint(temp *Temporal, np *Npoint) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.nai_tnpoint_npoint(temp._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// NaiTnpointTnpoint wraps MEOS C function nai_tnpoint_tnpoint.
func NaiTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.nai_tnpoint_tnpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// ShortestlineTnpointGeo wraps MEOS C function shortestline_tnpoint_geo.
func ShortestlineTnpointGeo(temp *Temporal, gs *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.shortestline_tnpoint_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// ShortestlineTnpointNpoint wraps MEOS C function shortestline_tnpoint_npoint.
func ShortestlineTnpointNpoint(temp *Temporal, np *Npoint) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.shortestline_tnpoint_npoint(temp._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// ShortestlineTnpointTnpoint wraps MEOS C function shortestline_tnpoint_tnpoint.
func ShortestlineTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.shortestline_tnpoint_tnpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TnpointTcentroidTransfn wraps MEOS C function tnpoint_tcentroid_transfn.
func TnpointTcentroidTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tnpoint_tcentroid_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// AlwaysEqNpointTnpoint wraps MEOS C function always_eq_npoint_tnpoint.
func AlwaysEqNpointTnpoint(np *Npoint, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_npoint_tnpoint(np._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTnpointNpoint wraps MEOS C function always_eq_tnpoint_npoint.
func AlwaysEqTnpointNpoint(temp *Temporal, np *Npoint) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tnpoint_npoint(temp._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTnpointTnpoint wraps MEOS C function always_eq_tnpoint_tnpoint.
func AlwaysEqTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tnpoint_tnpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeNpointTnpoint wraps MEOS C function always_ne_npoint_tnpoint.
func AlwaysNeNpointTnpoint(np *Npoint, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_npoint_tnpoint(np._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTnpointNpoint wraps MEOS C function always_ne_tnpoint_npoint.
func AlwaysNeTnpointNpoint(temp *Temporal, np *Npoint) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tnpoint_npoint(temp._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTnpointTnpoint wraps MEOS C function always_ne_tnpoint_tnpoint.
func AlwaysNeTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tnpoint_tnpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqNpointTnpoint wraps MEOS C function ever_eq_npoint_tnpoint.
func EverEqNpointTnpoint(np *Npoint, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_npoint_tnpoint(np._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTnpointNpoint wraps MEOS C function ever_eq_tnpoint_npoint.
func EverEqTnpointNpoint(temp *Temporal, np *Npoint) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tnpoint_npoint(temp._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTnpointTnpoint wraps MEOS C function ever_eq_tnpoint_tnpoint.
func EverEqTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tnpoint_tnpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeNpointTnpoint wraps MEOS C function ever_ne_npoint_tnpoint.
func EverNeNpointTnpoint(np *Npoint, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_npoint_tnpoint(np._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTnpointNpoint wraps MEOS C function ever_ne_tnpoint_npoint.
func EverNeTnpointNpoint(temp *Temporal, np *Npoint) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tnpoint_npoint(temp._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTnpointTnpoint wraps MEOS C function ever_ne_tnpoint_tnpoint.
func EverNeTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tnpoint_tnpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TeqTnpointNpoint wraps MEOS C function teq_tnpoint_npoint.
func TeqTnpointNpoint(temp *Temporal, np *Npoint) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_tnpoint_npoint(temp._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTnpointNpoint wraps MEOS C function tne_tnpoint_npoint.
func TneTnpointNpoint(temp *Temporal, np *Npoint) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_tnpoint_npoint(temp._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// ContainedRidBigintTnpoint wraps MEOS C function contained_rid_bigint_tnpoint.
func ContainedRidBigintTnpoint(rid int64, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_rid_bigint_tnpoint(C.int64_t(rid), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedRidBigintsetTnpoint wraps MEOS C function contained_rid_bigintset_tnpoint.
func ContainedRidBigintsetTnpoint(s *Set, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_rid_bigintset_tnpoint(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedRidNpointTnpoint wraps MEOS C function contained_rid_npoint_tnpoint.
func ContainedRidNpointTnpoint(np *Npoint, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_rid_npoint_tnpoint(np._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedRidTnpointBigintset wraps MEOS C function contained_rid_tnpoint_bigintset.
func ContainedRidTnpointBigintset(temp *Temporal, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_rid_tnpoint_bigintset(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedRidTnpointTnpoint wraps MEOS C function contained_rid_tnpoint_tnpoint.
func ContainedRidTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_rid_tnpoint_tnpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsRidBigintsetTnpoint wraps MEOS C function contains_rid_bigintset_tnpoint.
func ContainsRidBigintsetTnpoint(s *Set, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_rid_bigintset_tnpoint(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsRidTnpointBigint wraps MEOS C function contains_rid_tnpoint_bigint.
func ContainsRidTnpointBigint(temp *Temporal, rid int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_rid_tnpoint_bigint(temp._inner, C.int64_t(rid))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsRidTnpointBigintset wraps MEOS C function contains_rid_tnpoint_bigintset.
func ContainsRidTnpointBigintset(temp *Temporal, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_rid_tnpoint_bigintset(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsRidTnpointNpoint wraps MEOS C function contains_rid_tnpoint_npoint.
func ContainsRidTnpointNpoint(temp *Temporal, np *Npoint) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_rid_tnpoint_npoint(temp._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsRidTnpointTnpoint wraps MEOS C function contains_rid_tnpoint_tnpoint.
func ContainsRidTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_rid_tnpoint_tnpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsRidBigintsetTnpoint wraps MEOS C function overlaps_rid_bigintset_tnpoint.
func OverlapsRidBigintsetTnpoint(s *Set, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_rid_bigintset_tnpoint(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsRidTnpointBigintset wraps MEOS C function overlaps_rid_tnpoint_bigintset.
func OverlapsRidTnpointBigintset(temp *Temporal, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_rid_tnpoint_bigintset(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsRidTnpointTnpoint wraps MEOS C function overlaps_rid_tnpoint_tnpoint.
func OverlapsRidTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_rid_tnpoint_tnpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameRidBigintTnpoint wraps MEOS C function same_rid_bigint_tnpoint.
func SameRidBigintTnpoint(rid int64, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_rid_bigint_tnpoint(C.int64_t(rid), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameRidBigintsetTnpoint wraps MEOS C function same_rid_bigintset_tnpoint.
func SameRidBigintsetTnpoint(s *Set, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_rid_bigintset_tnpoint(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameRidNpointTnpoint wraps MEOS C function same_rid_npoint_tnpoint.
func SameRidNpointTnpoint(np *Npoint, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_rid_npoint_tnpoint(np._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameRidTnpointBigint wraps MEOS C function same_rid_tnpoint_bigint.
func SameRidTnpointBigint(temp *Temporal, rid int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_rid_tnpoint_bigint(temp._inner, C.int64_t(rid))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameRidTnpointBigintset wraps MEOS C function same_rid_tnpoint_bigintset.
func SameRidTnpointBigintset(temp *Temporal, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_rid_tnpoint_bigintset(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameRidTnpointNpoint wraps MEOS C function same_rid_tnpoint_npoint.
func SameRidTnpointNpoint(temp *Temporal, np *Npoint) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_rid_tnpoint_npoint(temp._inner, np._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameRidTnpointTnpoint wraps MEOS C function same_rid_tnpoint_tnpoint.
func SameRidTnpointTnpoint(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_rid_tnpoint_tnpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}

