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

// CbufferAsEWKB wraps MEOS C function cbuffer_as_ewkb.
func CbufferAsEWKB(cb *Cbuffer, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_as_ewkb(cb._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// CbufferAsEWKT wraps MEOS C function cbuffer_as_ewkt.
func CbufferAsEWKT(cb *Cbuffer, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_as_ewkt(cb._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// CbufferAsHexewkb wraps MEOS C function cbuffer_as_hexewkb.
func CbufferAsHexewkb(cb *Cbuffer, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_as_hexewkb(cb._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// CbufferAsHexwkb wraps MEOS C function cbuffer_as_hexwkb.
func CbufferAsHexwkb(cb *Cbuffer, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_as_hexwkb(cb._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// CbufferAsText wraps MEOS C function cbuffer_as_text.
func CbufferAsText(cb *Cbuffer, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_as_text(cb._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// CbufferAsWKB wraps MEOS C function cbuffer_as_wkb.
func CbufferAsWKB(cb *Cbuffer, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_as_wkb(cb._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// CbufferFromHexwkb wraps MEOS C function cbuffer_from_hexwkb.
func CbufferFromHexwkb(hexwkb string) (_r0 *Cbuffer, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.cbuffer_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return &Cbuffer{_inner: _cret}, nil
}


// CbufferFromWKB wraps MEOS C function cbuffer_from_wkb.
func CbufferFromWKB(wkb unsafe.Pointer, size uint) (_r0 *Cbuffer, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return &Cbuffer{_inner: _cret}, nil
}


// CbufferIn wraps MEOS C function cbuffer_in.
func CbufferIn(str string) (_r0 *Cbuffer, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.cbuffer_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Cbuffer{_inner: _cret}, nil
}


// CbufferOut wraps MEOS C function cbuffer_out.
func CbufferOut(cb *Cbuffer, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_out(cb._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// CbufferCopy wraps MEOS C function cbuffer_copy.
func CbufferCopy(cb *Cbuffer) (_r0 *Cbuffer, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_copy(cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Cbuffer{_inner: _cret}, nil
}


// CbufferMake wraps MEOS C function cbuffer_make.
func CbufferMake(point *Geom, radius float64) (_r0 *Cbuffer, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_make(point._inner, C.double(radius))
	if _err = meosError(); _err != nil {
		return
	}
	return &Cbuffer{_inner: _cret}, nil
}


// CbufferToGeom wraps MEOS C function cbuffer_to_geom.
func CbufferToGeom(cb *Cbuffer) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_to_geom(cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// CbufferToSTBOX wraps MEOS C function cbuffer_to_stbox.
func CbufferToSTBOX(cb *Cbuffer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_to_stbox(cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// CbufferarrToGeom wraps MEOS C function cbufferarr_to_geom.
func CbufferarrToGeom(cbarr unsafe.Pointer, count int) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.cbufferarr_to_geom((**C.Cbuffer)(unsafe.Pointer(cbarr)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeomToCbuffer wraps MEOS C function geom_to_cbuffer.
func GeomToCbuffer(gs *Geom) (_r0 *Cbuffer, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_to_cbuffer(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Cbuffer{_inner: _cret}, nil
}


// CbufferHash wraps MEOS C function cbuffer_hash.
func CbufferHash(cb *Cbuffer) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_hash(cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// CbufferHashExtended wraps MEOS C function cbuffer_hash_extended.
func CbufferHashExtended(cb *Cbuffer, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_hash_extended(cb._inner, C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// CbufferMemSize wraps MEOS C function cbuffer_mem_size.
func CbufferMemSize(cb *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_mem_size(cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// CbufferPoint wraps MEOS C function cbuffer_point.
func CbufferPoint(cb *Cbuffer) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_point(cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// CbufferRadius wraps MEOS C function cbuffer_radius.
func CbufferRadius(cb *Cbuffer) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_radius(cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// CbufferRound wraps MEOS C function cbuffer_round.
func CbufferRound(cb *Cbuffer, maxdd int) (_r0 *Cbuffer, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_round(cb._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return &Cbuffer{_inner: _cret}, nil
}


// CbufferarrRound wraps MEOS C function cbufferarr_round.
func CbufferarrRound(cbarr unsafe.Pointer, count int, maxdd int) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.cbufferarr_round((**C.Cbuffer)(unsafe.Pointer(cbarr)), C.int(count), C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// CbufferSetSRID wraps MEOS C function cbuffer_set_srid.
func CbufferSetSRID(cb *Cbuffer, srid int32) (_r0 *Cbuffer, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_set_srid(cb._inner, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Cbuffer{_inner: _cret}, nil
}


// CbufferSRID wraps MEOS C function cbuffer_srid.
func CbufferSRID(cb *Cbuffer) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_srid(cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// CbufferTransform wraps MEOS C function cbuffer_transform.
func CbufferTransform(cb *Cbuffer, srid int32) (_r0 *Cbuffer, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_transform(cb._inner, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Cbuffer{_inner: _cret}, nil
}


// CbufferTransformPipeline wraps MEOS C function cbuffer_transform_pipeline.
func CbufferTransformPipeline(cb *Cbuffer, pipelinestr string, srid int32, is_forward bool) (_r0 *Cbuffer, _err error) {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	C.meos_errno_reset()
	_cret := C.cbuffer_transform_pipeline(cb._inner, _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	if _err = meosError(); _err != nil {
		return
	}
	return &Cbuffer{_inner: _cret}, nil
}


// ContainsCbufferCbuffer wraps MEOS C function contains_cbuffer_cbuffer.
func ContainsCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_cbuffer_cbuffer(cb1._inner, cb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// CoversCbufferCbuffer wraps MEOS C function covers_cbuffer_cbuffer.
func CoversCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.covers_cbuffer_cbuffer(cb1._inner, cb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// DisjointCbufferCbuffer wraps MEOS C function disjoint_cbuffer_cbuffer.
func DisjointCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.disjoint_cbuffer_cbuffer(cb1._inner, cb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// DwithinCbufferCbuffer wraps MEOS C function dwithin_cbuffer_cbuffer.
func DwithinCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer, dist float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.dwithin_cbuffer_cbuffer(cb1._inner, cb2._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// IntersectsCbufferCbuffer wraps MEOS C function intersects_cbuffer_cbuffer.
func IntersectsCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.intersects_cbuffer_cbuffer(cb1._inner, cb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TouchesCbufferCbuffer wraps MEOS C function touches_cbuffer_cbuffer.
func TouchesCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.touches_cbuffer_cbuffer(cb1._inner, cb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// CbufferTstzspanToSTBOX wraps MEOS C function cbuffer_tstzspan_to_stbox.
func CbufferTstzspanToSTBOX(cb *Cbuffer, s *Span) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_tstzspan_to_stbox(cb._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// CbufferTimestamptzToSTBOX wraps MEOS C function cbuffer_timestamptz_to_stbox.
func CbufferTimestamptzToSTBOX(cb *Cbuffer, t int64) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_timestamptz_to_stbox(cb._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// DistanceCbufferCbuffer wraps MEOS C function distance_cbuffer_cbuffer.
func DistanceCbufferCbuffer(cb1 *Cbuffer, cb2 *Cbuffer) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_cbuffer_cbuffer(cb1._inner, cb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistanceCbufferGeo wraps MEOS C function distance_cbuffer_geo.
func DistanceCbufferGeo(cb *Cbuffer, gs *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_cbuffer_geo(cb._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistanceCbufferSTBOX wraps MEOS C function distance_cbuffer_stbox.
func DistanceCbufferSTBOX(cb *Cbuffer, box *STBox) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_cbuffer_stbox(cb._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadCbufferSTBOX wraps MEOS C function nad_cbuffer_stbox.
func NadCbufferSTBOX(cb *Cbuffer, box *STBox) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_cbuffer_stbox(cb._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// CbufferCmp wraps MEOS C function cbuffer_cmp.
func CbufferCmp(cb1 *Cbuffer, cb2 *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_cmp(cb1._inner, cb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// CbufferEq wraps MEOS C function cbuffer_eq.
func CbufferEq(cb1 *Cbuffer, cb2 *Cbuffer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_eq(cb1._inner, cb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// CbufferGe wraps MEOS C function cbuffer_ge.
func CbufferGe(cb1 *Cbuffer, cb2 *Cbuffer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_ge(cb1._inner, cb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// CbufferGt wraps MEOS C function cbuffer_gt.
func CbufferGt(cb1 *Cbuffer, cb2 *Cbuffer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_gt(cb1._inner, cb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// CbufferLe wraps MEOS C function cbuffer_le.
func CbufferLe(cb1 *Cbuffer, cb2 *Cbuffer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_le(cb1._inner, cb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// CbufferLt wraps MEOS C function cbuffer_lt.
func CbufferLt(cb1 *Cbuffer, cb2 *Cbuffer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_lt(cb1._inner, cb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// CbufferNe wraps MEOS C function cbuffer_ne.
func CbufferNe(cb1 *Cbuffer, cb2 *Cbuffer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_ne(cb1._inner, cb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// CbuffersetIn wraps MEOS C function cbufferset_in.
func CbuffersetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.cbufferset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// CbuffersetOut wraps MEOS C function cbufferset_out.
func CbuffersetOut(s *Set, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.cbufferset_out(s._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// CbuffersetMake wraps MEOS C function cbufferset_make.
func CbuffersetMake(values *Cbuffer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.cbufferset_make(values._inner, C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// CbufferToSet wraps MEOS C function cbuffer_to_set.
func CbufferToSet(cb *Cbuffer) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_to_set(cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// CbuffersetEndValue wraps MEOS C function cbufferset_end_value.
func CbuffersetEndValue(s *Set) (_r0 *Cbuffer, _err error) {
	C.meos_errno_reset()
	_cret := C.cbufferset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Cbuffer{_inner: _cret}, nil
}


// CbuffersetStartValue wraps MEOS C function cbufferset_start_value.
func CbuffersetStartValue(s *Set) (_r0 *Cbuffer, _err error) {
	C.meos_errno_reset()
	_cret := C.cbufferset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Cbuffer{_inner: _cret}, nil
}


// CbuffersetValueN wraps MEOS C function cbufferset_value_n.
func CbuffersetValueN(s *Set, n int) (_r0 bool, _r1 *Cbuffer, _err error) {
	var _out_result *C.Cbuffer
	C.meos_errno_reset()
	_cret := C.cbufferset_value_n(s._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Cbuffer{_inner: _out_result}, nil
}


// CbuffersetValues wraps MEOS C function cbufferset_values.
func CbuffersetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.cbufferset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// CbufferUnionTransfn wraps MEOS C function cbuffer_union_transfn.
func CbufferUnionTransfn(state *Set, cb *Cbuffer) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.cbuffer_union_transfn(state._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// ContainedCbufferSet wraps MEOS C function contained_cbuffer_set.
func ContainedCbufferSet(cb *Cbuffer, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_cbuffer_set(cb._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSetCbuffer wraps MEOS C function contains_set_cbuffer.
func ContainsSetCbuffer(s *Set, cb *Cbuffer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_cbuffer(s._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// IntersectionCbufferSet wraps MEOS C function intersection_cbuffer_set.
func IntersectionCbufferSet(cb *Cbuffer, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_cbuffer_set(cb._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetCbuffer wraps MEOS C function intersection_set_cbuffer.
func IntersectionSetCbuffer(s *Set, cb *Cbuffer) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_cbuffer(s._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusCbufferSet wraps MEOS C function minus_cbuffer_set.
func MinusCbufferSet(cb *Cbuffer, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_cbuffer_set(cb._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetCbuffer wraps MEOS C function minus_set_cbuffer.
func MinusSetCbuffer(s *Set, cb *Cbuffer) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_cbuffer(s._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionCbufferSet wraps MEOS C function union_cbuffer_set.
func UnionCbufferSet(cb *Cbuffer, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_cbuffer_set(cb._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetCbuffer wraps MEOS C function union_set_cbuffer.
func UnionSetCbuffer(s *Set, cb *Cbuffer) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_cbuffer(s._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TcbufferIn wraps MEOS C function tcbuffer_in.
func TcbufferIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tcbuffer_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcbufferFromMFJSON wraps MEOS C function tcbuffer_from_mfjson.
func TcbufferFromMFJSON(mfjson string) (_r0 *Temporal, _err error) {
	_c_mfjson := C.CString(mfjson)
	defer C.free(unsafe.Pointer(_c_mfjson))
	C.meos_errno_reset()
	_cret := C.tcbuffer_from_mfjson(_c_mfjson)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcbufferinstMake wraps MEOS C function tcbufferinst_make.
func TcbufferinstMake(cb *Cbuffer, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbufferinst_make(cb._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TcbufferMake wraps MEOS C function tcbuffer_make.
func TcbufferMake(tpoint *Temporal, tfloat *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_make(tpoint._inner, tfloat._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcbufferFromBaseTemp wraps MEOS C function tcbuffer_from_base_temp.
func TcbufferFromBaseTemp(cb *Cbuffer, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_from_base_temp(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcbufferseqFromBaseTstzset wraps MEOS C function tcbufferseq_from_base_tstzset.
func TcbufferseqFromBaseTstzset(cb *Cbuffer, s *Set) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbufferseq_from_base_tstzset(cb._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TcbufferseqFromBaseTstzspan wraps MEOS C function tcbufferseq_from_base_tstzspan.
func TcbufferseqFromBaseTstzspan(cb *Cbuffer, s *Span, interp Interpolation) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbufferseq_from_base_tstzspan(cb._inner, s._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TcbufferseqsetFromBaseTstzspanset wraps MEOS C function tcbufferseqset_from_base_tstzspanset.
func TcbufferseqsetFromBaseTstzspanset(cb *Cbuffer, ss *SpanSet, interp Interpolation) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbufferseqset_from_base_tstzspanset(cb._inner, ss._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TcbufferEndValue wraps MEOS C function tcbuffer_end_value.
func TcbufferEndValue(temp *Temporal) (_r0 *Cbuffer, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Cbuffer{_inner: _cret}, nil
}


// TcbufferPoints wraps MEOS C function tcbuffer_points.
func TcbufferPoints(temp *Temporal) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_points(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TcbufferRadius wraps MEOS C function tcbuffer_radius.
func TcbufferRadius(temp *Temporal) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_radius(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TcbufferTraversedArea wraps MEOS C function tcbuffer_traversed_area.
func TcbufferTraversedArea(temp *Temporal, unary_union bool) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_traversed_area(temp._inner, C.bool(unary_union))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TcbufferConvexHull wraps MEOS C function tcbuffer_convex_hull.
func TcbufferConvexHull(temp *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_convex_hull(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TcbufferStartValue wraps MEOS C function tcbuffer_start_value.
func TcbufferStartValue(temp *Temporal) (_r0 *Cbuffer, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Cbuffer{_inner: _cret}, nil
}


// TcbufferValue wraps MEOS C function tcbuffer_value.
func TcbufferValue(temp *Temporal) (_r0 *Cbuffer, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Cbuffer{_inner: _cret}, nil
}


// TcbufferValueAtTimestamptz wraps MEOS C function tcbuffer_value_at_timestamptz.
func TcbufferValueAtTimestamptz(temp *Temporal, t int64, strict bool) (_r0 bool, _r1 *Cbuffer, _err error) {
	var _out_value *C.Cbuffer
	C.meos_errno_reset()
	_cret := C.tcbuffer_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Cbuffer{_inner: _out_value}, nil
}


// TcbufferValueN wraps MEOS C function tcbuffer_value_n.
func TcbufferValueN(temp *Temporal, n int) (_r0 bool, _r1 *Cbuffer, _err error) {
	var _out_result *C.Cbuffer
	C.meos_errno_reset()
	_cret := C.tcbuffer_value_n(temp._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Cbuffer{_inner: _out_result}, nil
}


// TcbufferValues wraps MEOS C function tcbuffer_values.
func TcbufferValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TcbufferUnnest wraps MEOS C function tcbuffer_unnest.
func TcbufferUnnest(temp *Temporal, values unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_unnest(temp._inner, (***C.Cbuffer)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TcbufferToTfloat wraps MEOS C function tcbuffer_to_tfloat.
func TcbufferToTfloat(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_to_tfloat(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcbufferToTgeompoint wraps MEOS C function tcbuffer_to_tgeompoint.
func TcbufferToTgeompoint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_to_tgeompoint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeometryToTcbuffer wraps MEOS C function tgeometry_to_tcbuffer.
func TgeometryToTcbuffer(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeometry_to_tcbuffer(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcbufferExpand wraps MEOS C function tcbuffer_expand.
func TcbufferExpand(temp *Temporal, dist float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_expand(temp._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcbufferRotate wraps MEOS C function tcbuffer_rotate.
func TcbufferRotate(temp *Temporal, angle float64, x0 float64, y0 float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_rotate(temp._inner, C.double(angle), C.double(x0), C.double(y0))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcbufferRotateZ wraps MEOS C function tcbuffer_rotate_z.
func TcbufferRotateZ(temp *Temporal, angle float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_rotate_z(temp._inner, C.double(angle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcbufferTranslate wraps MEOS C function tcbuffer_translate.
func TcbufferTranslate(temp *Temporal, deltax float64, deltay float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_translate(temp._inner, C.double(deltax), C.double(deltay))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcbufferAtCbuffer wraps MEOS C function tcbuffer_at_cbuffer.
func TcbufferAtCbuffer(temp *Temporal, cb *Cbuffer) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_at_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcbufferAtGeom wraps MEOS C function tcbuffer_at_geom.
func TcbufferAtGeom(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_at_geom(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcbufferAtSTBOX wraps MEOS C function tcbuffer_at_stbox.
func TcbufferAtSTBOX(temp *Temporal, box *STBox, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_at_stbox(temp._inner, box._inner, C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcbufferMinusCbuffer wraps MEOS C function tcbuffer_minus_cbuffer.
func TcbufferMinusCbuffer(temp *Temporal, cb *Cbuffer) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_minus_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcbufferMinusGeom wraps MEOS C function tcbuffer_minus_geom.
func TcbufferMinusGeom(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_minus_geom(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcbufferMinusSTBOX wraps MEOS C function tcbuffer_minus_stbox.
func TcbufferMinusSTBOX(temp *Temporal, box *STBox, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcbuffer_minus_stbox(temp._inner, box._inner, C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTcbufferCbuffer wraps MEOS C function tdistance_tcbuffer_cbuffer.
func TdistanceTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTcbufferGeo wraps MEOS C function tdistance_tcbuffer_geo.
func TdistanceTcbufferGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTcbufferTcbuffer wraps MEOS C function tdistance_tcbuffer_tcbuffer.
func TdistanceTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// NadTcbufferCbuffer wraps MEOS C function nad_tcbuffer_cbuffer.
func NadTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTcbufferGeo wraps MEOS C function nad_tcbuffer_geo.
func NadTcbufferGeo(temp *Temporal, gs *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTcbufferSTBOX wraps MEOS C function nad_tcbuffer_stbox.
func NadTcbufferSTBOX(temp *Temporal, box *STBox) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tcbuffer_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTcbufferTcbuffer wraps MEOS C function nad_tcbuffer_tcbuffer.
func NadTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// MindistanceTcbufferTcbuffer wraps MEOS C function mindistance_tcbuffer_tcbuffer.
func MindistanceTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal, threshold float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.mindistance_tcbuffer_tcbuffer(temp1._inner, temp2._inner, C.double(threshold))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// MindistanceTcbufferarrTcbufferarr wraps MEOS C function mindistance_tcbufferarr_tcbufferarr.
func MindistanceTcbufferarrTcbufferarr(arr1 unsafe.Pointer, count1 int, arr2 unsafe.Pointer, count2 int) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.mindistance_tcbufferarr_tcbufferarr((**C.Temporal)(unsafe.Pointer(arr1)), C.int(count1), (**C.Temporal)(unsafe.Pointer(arr2)), C.int(count2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NaiTcbufferCbuffer wraps MEOS C function nai_tcbuffer_cbuffer.
func NaiTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.nai_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// NaiTcbufferGeo wraps MEOS C function nai_tcbuffer_geo.
func NaiTcbufferGeo(temp *Temporal, gs *Geom) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.nai_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// NaiTcbufferTcbuffer wraps MEOS C function nai_tcbuffer_tcbuffer.
func NaiTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.nai_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// ShortestlineTcbufferCbuffer wraps MEOS C function shortestline_tcbuffer_cbuffer.
func ShortestlineTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.shortestline_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// ShortestlineTcbufferGeo wraps MEOS C function shortestline_tcbuffer_geo.
func ShortestlineTcbufferGeo(temp *Temporal, gs *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.shortestline_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// ShortestlineTcbufferTcbuffer wraps MEOS C function shortestline_tcbuffer_tcbuffer.
func ShortestlineTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.shortestline_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// AlwaysEqCbufferTcbuffer wraps MEOS C function always_eq_cbuffer_tcbuffer.
func AlwaysEqCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTcbufferCbuffer wraps MEOS C function always_eq_tcbuffer_cbuffer.
func AlwaysEqTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTcbufferTcbuffer wraps MEOS C function always_eq_tcbuffer_tcbuffer.
func AlwaysEqTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeCbufferTcbuffer wraps MEOS C function always_ne_cbuffer_tcbuffer.
func AlwaysNeCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTcbufferCbuffer wraps MEOS C function always_ne_tcbuffer_cbuffer.
func AlwaysNeTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTcbufferTcbuffer wraps MEOS C function always_ne_tcbuffer_tcbuffer.
func AlwaysNeTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqCbufferTcbuffer wraps MEOS C function ever_eq_cbuffer_tcbuffer.
func EverEqCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTcbufferCbuffer wraps MEOS C function ever_eq_tcbuffer_cbuffer.
func EverEqTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTcbufferTcbuffer wraps MEOS C function ever_eq_tcbuffer_tcbuffer.
func EverEqTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeCbufferTcbuffer wraps MEOS C function ever_ne_cbuffer_tcbuffer.
func EverNeCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTcbufferCbuffer wraps MEOS C function ever_ne_tcbuffer_cbuffer.
func EverNeTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTcbufferTcbuffer wraps MEOS C function ever_ne_tcbuffer_tcbuffer.
func EverNeTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TeqCbufferTcbuffer wraps MEOS C function teq_cbuffer_tcbuffer.
func TeqCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTcbufferCbuffer wraps MEOS C function teq_tcbuffer_cbuffer.
func TeqTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneCbufferTcbuffer wraps MEOS C function tne_cbuffer_tcbuffer.
func TneCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTcbufferCbuffer wraps MEOS C function tne_tcbuffer_cbuffer.
func TneTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// AcontainsCbufferTcbuffer wraps MEOS C function acontains_cbuffer_tcbuffer.
func AcontainsCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acontains_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AcontainsGeoTcbuffer wraps MEOS C function acontains_geo_tcbuffer.
func AcontainsGeoTcbuffer(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acontains_geo_tcbuffer(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AcontainsTcbufferCbuffer wraps MEOS C function acontains_tcbuffer_cbuffer.
func AcontainsTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acontains_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AcontainsTcbufferGeo wraps MEOS C function acontains_tcbuffer_geo.
func AcontainsTcbufferGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acontains_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AcontainsTcbufferTcbuffer wraps MEOS C function acontains_tcbuffer_tcbuffer.
func AcontainsTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acontains_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AcoversCbufferTcbuffer wraps MEOS C function acovers_cbuffer_tcbuffer.
func AcoversCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acovers_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AcoversGeoTcbuffer wraps MEOS C function acovers_geo_tcbuffer.
func AcoversGeoTcbuffer(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acovers_geo_tcbuffer(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AcoversTcbufferCbuffer wraps MEOS C function acovers_tcbuffer_cbuffer.
func AcoversTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acovers_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AcoversTcbufferGeo wraps MEOS C function acovers_tcbuffer_geo.
func AcoversTcbufferGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acovers_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AcoversTcbufferTcbuffer wraps MEOS C function acovers_tcbuffer_tcbuffer.
func AcoversTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acovers_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdisjointCbufferTcbuffer wraps MEOS C function adisjoint_cbuffer_tcbuffer.
func AdisjointCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adisjoint_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdisjointGeoTcbuffer wraps MEOS C function adisjoint_geo_tcbuffer.
func AdisjointGeoTcbuffer(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adisjoint_geo_tcbuffer(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdisjointTcbufferGeo wraps MEOS C function adisjoint_tcbuffer_geo.
func AdisjointTcbufferGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adisjoint_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdisjointTcbufferCbuffer wraps MEOS C function adisjoint_tcbuffer_cbuffer.
func AdisjointTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adisjoint_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdisjointTcbufferTcbuffer wraps MEOS C function adisjoint_tcbuffer_tcbuffer.
func AdisjointTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adisjoint_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdwithinCbufferTcbuffer wraps MEOS C function adwithin_cbuffer_tcbuffer.
func AdwithinCbufferTcbuffer(cb *Cbuffer, temp *Temporal, dist float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adwithin_cbuffer_tcbuffer(cb._inner, temp._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdwithinGeoTcbuffer wraps MEOS C function adwithin_geo_tcbuffer.
func AdwithinGeoTcbuffer(gs *Geom, temp *Temporal, dist float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adwithin_geo_tcbuffer(gs._inner, temp._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdwithinTcbufferGeo wraps MEOS C function adwithin_tcbuffer_geo.
func AdwithinTcbufferGeo(temp *Temporal, gs *Geom, dist float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adwithin_tcbuffer_geo(temp._inner, gs._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdwithinTcbufferCbuffer wraps MEOS C function adwithin_tcbuffer_cbuffer.
func AdwithinTcbufferCbuffer(temp *Temporal, cb *Cbuffer, dist float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adwithin_tcbuffer_cbuffer(temp._inner, cb._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdwithinTcbufferTcbuffer wraps MEOS C function adwithin_tcbuffer_tcbuffer.
func AdwithinTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal, dist float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adwithin_tcbuffer_tcbuffer(temp1._inner, temp2._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AintersectsCbufferTcbuffer wraps MEOS C function aintersects_cbuffer_tcbuffer.
func AintersectsCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.aintersects_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AintersectsGeoTcbuffer wraps MEOS C function aintersects_geo_tcbuffer.
func AintersectsGeoTcbuffer(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.aintersects_geo_tcbuffer(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AintersectsTcbufferGeo wraps MEOS C function aintersects_tcbuffer_geo.
func AintersectsTcbufferGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.aintersects_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AintersectsTcbufferCbuffer wraps MEOS C function aintersects_tcbuffer_cbuffer.
func AintersectsTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.aintersects_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AintersectsTcbufferTcbuffer wraps MEOS C function aintersects_tcbuffer_tcbuffer.
func AintersectsTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.aintersects_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AtouchesCbufferTcbuffer wraps MEOS C function atouches_cbuffer_tcbuffer.
func AtouchesCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.atouches_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AtouchesGeoTcbuffer wraps MEOS C function atouches_geo_tcbuffer.
func AtouchesGeoTcbuffer(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.atouches_geo_tcbuffer(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AtouchesTcbufferGeo wraps MEOS C function atouches_tcbuffer_geo.
func AtouchesTcbufferGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.atouches_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AtouchesTcbufferCbuffer wraps MEOS C function atouches_tcbuffer_cbuffer.
func AtouchesTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.atouches_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AtouchesTcbufferTcbuffer wraps MEOS C function atouches_tcbuffer_tcbuffer.
func AtouchesTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.atouches_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcontainsCbufferTcbuffer wraps MEOS C function econtains_cbuffer_tcbuffer.
func EcontainsCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.econtains_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcontainsGeoTcbuffer wraps MEOS C function econtains_geo_tcbuffer.
func EcontainsGeoTcbuffer(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.econtains_geo_tcbuffer(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcontainsTcbufferCbuffer wraps MEOS C function econtains_tcbuffer_cbuffer.
func EcontainsTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.econtains_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcontainsTcbufferGeo wraps MEOS C function econtains_tcbuffer_geo.
func EcontainsTcbufferGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.econtains_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcontainsTcbufferTcbuffer wraps MEOS C function econtains_tcbuffer_tcbuffer.
func EcontainsTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.econtains_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcoversCbufferTcbuffer wraps MEOS C function ecovers_cbuffer_tcbuffer.
func EcoversCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ecovers_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcoversGeoTcbuffer wraps MEOS C function ecovers_geo_tcbuffer.
func EcoversGeoTcbuffer(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ecovers_geo_tcbuffer(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcoversTcbufferCbuffer wraps MEOS C function ecovers_tcbuffer_cbuffer.
func EcoversTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ecovers_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcoversTcbufferGeo wraps MEOS C function ecovers_tcbuffer_geo.
func EcoversTcbufferGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ecovers_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcoversTcbufferTcbuffer wraps MEOS C function ecovers_tcbuffer_tcbuffer.
func EcoversTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ecovers_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdisjointCbufferTcbuffer wraps MEOS C function edisjoint_cbuffer_tcbuffer.
func EdisjointCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edisjoint_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdisjointGeoTcbuffer wraps MEOS C function edisjoint_geo_tcbuffer.
func EdisjointGeoTcbuffer(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edisjoint_geo_tcbuffer(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdisjointTcbufferGeo wraps MEOS C function edisjoint_tcbuffer_geo.
func EdisjointTcbufferGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edisjoint_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdisjointTcbufferCbuffer wraps MEOS C function edisjoint_tcbuffer_cbuffer.
func EdisjointTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edisjoint_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdisjointTcbufferTcbuffer wraps MEOS C function edisjoint_tcbuffer_tcbuffer.
func EdisjointTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edisjoint_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdwithinCbufferTcbuffer wraps MEOS C function edwithin_cbuffer_tcbuffer.
func EdwithinCbufferTcbuffer(cb *Cbuffer, temp *Temporal, dist float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edwithin_cbuffer_tcbuffer(cb._inner, temp._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdwithinGeoTcbuffer wraps MEOS C function edwithin_geo_tcbuffer.
func EdwithinGeoTcbuffer(gs *Geom, temp *Temporal, dist float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edwithin_geo_tcbuffer(gs._inner, temp._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdwithinTcbufferGeo wraps MEOS C function edwithin_tcbuffer_geo.
func EdwithinTcbufferGeo(temp *Temporal, gs *Geom, dist float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edwithin_tcbuffer_geo(temp._inner, gs._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdwithinTcbufferCbuffer wraps MEOS C function edwithin_tcbuffer_cbuffer.
func EdwithinTcbufferCbuffer(temp *Temporal, cb *Cbuffer, dist float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edwithin_tcbuffer_cbuffer(temp._inner, cb._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdwithinTcbufferTcbuffer wraps MEOS C function edwithin_tcbuffer_tcbuffer.
func EdwithinTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal, dist float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edwithin_tcbuffer_tcbuffer(temp1._inner, temp2._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EintersectsCbufferTcbuffer wraps MEOS C function eintersects_cbuffer_tcbuffer.
func EintersectsCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.eintersects_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EintersectsGeoTcbuffer wraps MEOS C function eintersects_geo_tcbuffer.
func EintersectsGeoTcbuffer(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.eintersects_geo_tcbuffer(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EintersectsTcbufferGeo wraps MEOS C function eintersects_tcbuffer_geo.
func EintersectsTcbufferGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.eintersects_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EintersectsTcbufferCbuffer wraps MEOS C function eintersects_tcbuffer_cbuffer.
func EintersectsTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.eintersects_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EintersectsTcbufferTcbuffer wraps MEOS C function eintersects_tcbuffer_tcbuffer.
func EintersectsTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.eintersects_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EtouchesCbufferTcbuffer wraps MEOS C function etouches_cbuffer_tcbuffer.
func EtouchesCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.etouches_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EtouchesGeoTcbuffer wraps MEOS C function etouches_geo_tcbuffer.
func EtouchesGeoTcbuffer(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.etouches_geo_tcbuffer(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EtouchesTcbufferGeo wraps MEOS C function etouches_tcbuffer_geo.
func EtouchesTcbufferGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.etouches_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EtouchesTcbufferCbuffer wraps MEOS C function etouches_tcbuffer_cbuffer.
func EtouchesTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.etouches_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EtouchesTcbufferTcbuffer wraps MEOS C function etouches_tcbuffer_tcbuffer.
func EtouchesTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.etouches_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TcontainsCbufferTcbuffer wraps MEOS C function tcontains_cbuffer_tcbuffer.
func TcontainsCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcontains_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcontainsGeoTcbuffer wraps MEOS C function tcontains_geo_tcbuffer.
func TcontainsGeoTcbuffer(gs *Geom, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcontains_geo_tcbuffer(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcontainsTcbufferGeo wraps MEOS C function tcontains_tcbuffer_geo.
func TcontainsTcbufferGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcontains_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcontainsTcbufferCbuffer wraps MEOS C function tcontains_tcbuffer_cbuffer.
func TcontainsTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcontains_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcontainsTcbufferTcbuffer wraps MEOS C function tcontains_tcbuffer_tcbuffer.
func TcontainsTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcontains_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcoversCbufferTcbuffer wraps MEOS C function tcovers_cbuffer_tcbuffer.
func TcoversCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcovers_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcoversGeoTcbuffer wraps MEOS C function tcovers_geo_tcbuffer.
func TcoversGeoTcbuffer(gs *Geom, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcovers_geo_tcbuffer(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcoversTcbufferGeo wraps MEOS C function tcovers_tcbuffer_geo.
func TcoversTcbufferGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcovers_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcoversTcbufferCbuffer wraps MEOS C function tcovers_tcbuffer_cbuffer.
func TcoversTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcovers_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TcoversTcbufferTcbuffer wraps MEOS C function tcovers_tcbuffer_tcbuffer.
func TcoversTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tcovers_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdwithinCbufferTcbuffer wraps MEOS C function tdwithin_cbuffer_tcbuffer.
func TdwithinCbufferTcbuffer(cb *Cbuffer, temp *Temporal, dist float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdwithin_cbuffer_tcbuffer(cb._inner, temp._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdwithinGeoTcbuffer wraps MEOS C function tdwithin_geo_tcbuffer.
func TdwithinGeoTcbuffer(gs *Geom, temp *Temporal, dist float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdwithin_geo_tcbuffer(gs._inner, temp._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdwithinTcbufferGeo wraps MEOS C function tdwithin_tcbuffer_geo.
func TdwithinTcbufferGeo(temp *Temporal, gs *Geom, dist float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdwithin_tcbuffer_geo(temp._inner, gs._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdwithinTcbufferCbuffer wraps MEOS C function tdwithin_tcbuffer_cbuffer.
func TdwithinTcbufferCbuffer(temp *Temporal, cb *Cbuffer, dist float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdwithin_tcbuffer_cbuffer(temp._inner, cb._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdwithinTcbufferTcbuffer wraps MEOS C function tdwithin_tcbuffer_tcbuffer.
func TdwithinTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal, dist float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdwithin_tcbuffer_tcbuffer(temp1._inner, temp2._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdisjointCbufferTcbuffer wraps MEOS C function tdisjoint_cbuffer_tcbuffer.
func TdisjointCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdisjoint_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdisjointGeoTcbuffer wraps MEOS C function tdisjoint_geo_tcbuffer.
func TdisjointGeoTcbuffer(gs *Geom, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdisjoint_geo_tcbuffer(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdisjointTcbufferGeo wraps MEOS C function tdisjoint_tcbuffer_geo.
func TdisjointTcbufferGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdisjoint_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdisjointTcbufferCbuffer wraps MEOS C function tdisjoint_tcbuffer_cbuffer.
func TdisjointTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdisjoint_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdisjointTcbufferTcbuffer wraps MEOS C function tdisjoint_tcbuffer_tcbuffer.
func TdisjointTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdisjoint_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintersectsCbufferTcbuffer wraps MEOS C function tintersects_cbuffer_tcbuffer.
func TintersectsCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tintersects_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintersectsGeoTcbuffer wraps MEOS C function tintersects_geo_tcbuffer.
func TintersectsGeoTcbuffer(gs *Geom, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tintersects_geo_tcbuffer(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintersectsTcbufferGeo wraps MEOS C function tintersects_tcbuffer_geo.
func TintersectsTcbufferGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tintersects_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintersectsTcbufferCbuffer wraps MEOS C function tintersects_tcbuffer_cbuffer.
func TintersectsTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tintersects_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintersectsTcbufferTcbuffer wraps MEOS C function tintersects_tcbuffer_tcbuffer.
func TintersectsTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tintersects_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TtouchesGeoTcbuffer wraps MEOS C function ttouches_geo_tcbuffer.
func TtouchesGeoTcbuffer(gs *Geom, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ttouches_geo_tcbuffer(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TtouchesTcbufferGeo wraps MEOS C function ttouches_tcbuffer_geo.
func TtouchesTcbufferGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ttouches_tcbuffer_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TtouchesCbufferTcbuffer wraps MEOS C function ttouches_cbuffer_tcbuffer.
func TtouchesCbufferTcbuffer(cb *Cbuffer, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ttouches_cbuffer_tcbuffer(cb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TtouchesTcbufferCbuffer wraps MEOS C function ttouches_tcbuffer_cbuffer.
func TtouchesTcbufferCbuffer(temp *Temporal, cb *Cbuffer) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ttouches_tcbuffer_cbuffer(temp._inner, cb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TtouchesTcbufferTcbuffer wraps MEOS C function ttouches_tcbuffer_tcbuffer.
func TtouchesTcbufferTcbuffer(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ttouches_tcbuffer_tcbuffer(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}

