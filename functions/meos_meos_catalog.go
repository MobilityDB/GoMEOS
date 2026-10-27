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

// TemptypeSubtype wraps MEOS C function temptype_subtype.
func TemptypeSubtype(subtype TempSubtype) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.temptype_subtype(C.tempSubtype(subtype))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TemptypeSubtypeAll wraps MEOS C function temptype_subtype_all.
func TemptypeSubtypeAll(subtype TempSubtype) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.temptype_subtype_all(C.tempSubtype(subtype))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TempsubtypeName wraps MEOS C function tempsubtype_name.
func TempsubtypeName(subtype TempSubtype) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tempsubtype_name(C.tempSubtype(subtype))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// TempsubtypeFromString wraps MEOS C function tempsubtype_from_string.
func TempsubtypeFromString(str string, subtype unsafe.Pointer) (_r0 bool, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tempsubtype_from_string(_c_str, (*C.int16)(unsafe.Pointer(subtype)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// MeosoperName wraps MEOS C function meosoper_name.
func MeosoperName(oper MeosOper) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.meosoper_name(C.MeosOper(oper))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// MeosoperFromString wraps MEOS C function meosoper_from_string.
func MeosoperFromString(name string) (_r0 MeosOper, _err error) {
	_c_name := C.CString(name)
	defer C.free(unsafe.Pointer(_c_name))
	C.meos_errno_reset()
	_cret := C.meosoper_from_string(_c_name)
	if _err = meosError(); _err != nil {
		return
	}
	return MeosOper(_cret), nil
}


// InterptypeName wraps MEOS C function interptype_name.
func InterptypeName(interp Interpolation) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.interptype_name(C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// MeosTypeofHexwkb wraps MEOS C function meos_typeof_hexwkb.
func MeosTypeofHexwkb(hexwkb string) (_r0 MeosType, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.meos_typeof_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return MeosType(_cret), nil
}


// MeostypeName wraps MEOS C function meostype_name.
func MeostypeName(type_ MeosType) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.meostype_name(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// MeostypeFromString wraps MEOS C function meostype_from_string.
func MeostypeFromString(name string) (_r0 MeosType, _err error) {
	_c_name := C.CString(name)
	defer C.free(unsafe.Pointer(_c_name))
	C.meos_errno_reset()
	_cret := C.meostype_from_string(_c_name)
	if _err = meosError(); _err != nil {
		return
	}
	return MeosType(_cret), nil
}


// TemptypeBasetype wraps MEOS C function temptype_basetype.
func TemptypeBasetype(type_ MeosType) (_r0 MeosType, _err error) {
	C.meos_errno_reset()
	_cret := C.temptype_basetype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return MeosType(_cret), nil
}


// SettypeBasetype wraps MEOS C function settype_basetype.
func SettypeBasetype(type_ MeosType) (_r0 MeosType, _err error) {
	C.meos_errno_reset()
	_cret := C.settype_basetype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return MeosType(_cret), nil
}


// SpantypeBasetype wraps MEOS C function spantype_basetype.
func SpantypeBasetype(type_ MeosType) (_r0 MeosType, _err error) {
	C.meos_errno_reset()
	_cret := C.spantype_basetype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return MeosType(_cret), nil
}


// SpantypeSpansettype wraps MEOS C function spantype_spansettype.
func SpantypeSpansettype(type_ MeosType) (_r0 MeosType, _err error) {
	C.meos_errno_reset()
	_cret := C.spantype_spansettype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return MeosType(_cret), nil
}


// SpansettypeSpantype wraps MEOS C function spansettype_spantype.
func SpansettypeSpantype(type_ MeosType) (_r0 MeosType, _err error) {
	C.meos_errno_reset()
	_cret := C.spansettype_spantype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return MeosType(_cret), nil
}


// BasetypeSpantype wraps MEOS C function basetype_spantype.
func BasetypeSpantype(type_ MeosType) (_r0 MeosType, _err error) {
	C.meos_errno_reset()
	_cret := C.basetype_spantype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return MeosType(_cret), nil
}


// BasetypeSettype wraps MEOS C function basetype_settype.
func BasetypeSettype(type_ MeosType) (_r0 MeosType, _err error) {
	C.meos_errno_reset()
	_cret := C.basetype_settype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return MeosType(_cret), nil
}


// TnumberBasetype wraps MEOS C function tnumber_basetype.
func TnumberBasetype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_basetype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeoBasetype wraps MEOS C function geo_basetype.
func GeoBasetype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_basetype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// MeosBasetype wraps MEOS C function meos_basetype.
func MeosBasetype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.meos_basetype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AlphanumBasetype wraps MEOS C function alphanum_basetype.
func AlphanumBasetype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.alphanum_basetype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureAlphanumBasetype wraps MEOS C function ensure_alphanum_basetype.
func EnsureAlphanumBasetype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_alphanum_basetype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AlphanumTemptype wraps MEOS C function alphanum_temptype.
func AlphanumTemptype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.alphanum_temptype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimeType wraps MEOS C function time_type.
func TimeType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.time_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SetBasetype wraps MEOS C function set_basetype.
func SetBasetype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.set_basetype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SetType wraps MEOS C function set_type.
func SetType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.set_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NumsetType wraps MEOS C function numset_type.
func NumsetType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.numset_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureNumsetType wraps MEOS C function ensure_numset_type.
func EnsureNumsetType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_numset_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimesetType wraps MEOS C function timeset_type.
func TimesetType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.timeset_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SetSpantype wraps MEOS C function set_spantype.
func SetSpantype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.set_spantype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureSetSpantype wraps MEOS C function ensure_set_spantype.
func EnsureSetSpantype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_set_spantype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AlphanumsetType wraps MEOS C function alphanumset_type.
func AlphanumsetType(settype MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.alphanumset_type(C.MeosType(settype))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeosetType wraps MEOS C function geoset_type.
func GeosetType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geoset_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureGeosetType wraps MEOS C function ensure_geoset_type.
func EnsureGeosetType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_geoset_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpatialsetType wraps MEOS C function spatialset_type.
func SpatialsetType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.spatialset_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureSpatialsetType wraps MEOS C function ensure_spatialset_type.
func EnsureSpatialsetType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_spatialset_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PointcloudBasetype wraps MEOS C function pointcloud_basetype.
func PointcloudBasetype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pointcloud_basetype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PointcloudsetType wraps MEOS C function pointcloudset_type.
func PointcloudsetType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pointcloudset_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpointcloudTemptype wraps MEOS C function tpointcloud_temptype.
func TpointcloudTemptype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointcloud_temptype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureTpointcloudTemptype wraps MEOS C function ensure_tpointcloud_temptype.
func EnsureTpointcloudTemptype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_tpointcloud_temptype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanBasetype wraps MEOS C function span_basetype.
func SpanBasetype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_basetype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanCanonBasetype wraps MEOS C function span_canon_basetype.
func SpanCanonBasetype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_canon_basetype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanType wraps MEOS C function span_type.
func SpanType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TypeSpanBbox wraps MEOS C function type_span_bbox.
func TypeSpanBbox(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.type_span_bbox(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TypeBboxtype wraps MEOS C function type_bboxtype.
func TypeBboxtype(type_ MeosType) (_r0 MeosType, _err error) {
	C.meos_errno_reset()
	_cret := C.type_bboxtype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return MeosType(_cret), nil
}


// SpanTBOXType wraps MEOS C function span_tbox_type.
func SpanTBOXType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_tbox_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureSpanTBOXType wraps MEOS C function ensure_span_tbox_type.
func EnsureSpanTBOXType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_span_tbox_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NumspanBasetype wraps MEOS C function numspan_basetype.
func NumspanBasetype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.numspan_basetype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NumspanType wraps MEOS C function numspan_type.
func NumspanType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.numspan_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureNumspanType wraps MEOS C function ensure_numspan_type.
func EnsureNumspanType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_numspan_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimespanBasetype wraps MEOS C function timespan_basetype.
func TimespanBasetype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.timespan_basetype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimespanType wraps MEOS C function timespan_type.
func TimespanType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.timespan_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpansetType wraps MEOS C function spanset_type.
func SpansetType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NumspansetType wraps MEOS C function numspanset_type.
func NumspansetType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.numspanset_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureNumspansetType wraps MEOS C function ensure_numspanset_type.
func EnsureNumspansetType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_numspanset_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimespansetType wraps MEOS C function timespanset_type.
func TimespansetType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.timespanset_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureTimespansetType wraps MEOS C function ensure_timespanset_type.
func EnsureTimespansetType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_timespanset_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TemporalType wraps MEOS C function temporal_type.
func TemporalType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TemporalBasetype wraps MEOS C function temporal_basetype.
func TemporalBasetype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_basetype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TemptypeSupportsLinear wraps MEOS C function temptype_supports_linear.
func TemptypeSupportsLinear(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.temptype_supports_linear(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BasetypeByvalue wraps MEOS C function basetype_byvalue.
func BasetypeByvalue(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.basetype_byvalue(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BasetypeVarlength wraps MEOS C function basetype_varlength.
func BasetypeVarlength(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.basetype_varlength(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// MeostypeLength wraps MEOS C function meostype_length.
func MeostypeLength(type_ MeosType) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.meostype_length(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// TalphanumType wraps MEOS C function talphanum_type.
func TalphanumType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.talphanum_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TalphaType wraps MEOS C function talpha_type.
func TalphaType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.talpha_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TnumberType wraps MEOS C function tnumber_type.
func TnumberType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureTnumberType wraps MEOS C function ensure_tnumber_type.
func EnsureTnumberType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_tnumber_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TorderType wraps MEOS C function torder_type.
func TorderType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.torder_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureTorderType wraps MEOS C function ensure_torder_type.
func EnsureTorderType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_torder_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureTnumberBasetype wraps MEOS C function ensure_tnumber_basetype.
func EnsureTnumberBasetype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_tnumber_basetype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TnumberSpantype wraps MEOS C function tnumber_spantype.
func TnumberSpantype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_spantype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpatialBasetype wraps MEOS C function spatial_basetype.
func SpatialBasetype(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.spatial_basetype(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TspatialType wraps MEOS C function tspatial_type.
func TspatialType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tspatial_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureTspatialType wraps MEOS C function ensure_tspatial_type.
func EnsureTspatialType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_tspatial_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpointType wraps MEOS C function tpoint_type.
func TpointType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureTpointType wraps MEOS C function ensure_tpoint_type.
func EnsureTpointType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_tpoint_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TgeoType wraps MEOS C function tgeo_type.
func TgeoType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureTgeoType wraps MEOS C function ensure_tgeo_type.
func EnsureTgeoType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_tgeo_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TspatialBodyType wraps MEOS C function tspatial_body_type.
func TspatialBodyType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tspatial_body_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TgeoTypeAll wraps MEOS C function tgeo_type_all.
func TgeoTypeAll(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_type_all(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureTgeoTypeAll wraps MEOS C function ensure_tgeo_type_all.
func EnsureTgeoTypeAll(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_tgeo_type_all(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TgeometryType wraps MEOS C function tgeometry_type.
func TgeometryType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeometry_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureTgeometryType wraps MEOS C function ensure_tgeometry_type.
func EnsureTgeometryType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_tgeometry_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TgeodeticType wraps MEOS C function tgeodetic_type.
func TgeodeticType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeodetic_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureTgeodeticType wraps MEOS C function ensure_tgeodetic_type.
func EnsureTgeodeticType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_tgeodetic_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureTnumberTpointType wraps MEOS C function ensure_tnumber_tpoint_type.
func EnsureTnumberTpointType(type_ MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_tnumber_tpoint_type(C.MeosType(type_))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}

