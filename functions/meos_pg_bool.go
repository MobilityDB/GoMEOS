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

// BoolEq wraps MEOS C function bool_eq.
func BoolEq(arg1 bool, arg2 bool) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.bool_eq(C.bool(arg1), C.bool(arg2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BoolGe wraps MEOS C function bool_ge.
func BoolGe(arg1 bool, arg2 bool) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.bool_ge(C.bool(arg1), C.bool(arg2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BoolGt wraps MEOS C function bool_gt.
func BoolGt(arg1 bool, arg2 bool) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.bool_gt(C.bool(arg1), C.bool(arg2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BoolHash wraps MEOS C function bool_hash.
func BoolHash(arg bool) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.bool_hash(C.bool(arg))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// BoolHashExtended wraps MEOS C function bool_hash_extended.
func BoolHashExtended(arg bool, seed int64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.bool_hash_extended(C.bool(arg), C.int64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// BoolLe wraps MEOS C function bool_le.
func BoolLe(arg1 bool, arg2 bool) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.bool_le(C.bool(arg1), C.bool(arg2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BoolLt wraps MEOS C function bool_lt.
func BoolLt(arg1 bool, arg2 bool) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.bool_lt(C.bool(arg1), C.bool(arg2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BoolNe wraps MEOS C function bool_ne.
func BoolNe(arg1 bool, arg2 bool) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.bool_ne(C.bool(arg1), C.bool(arg2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BoolToText wraps MEOS C function bool_to_text.
func BoolToText(b bool) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.bool_to_text(C.bool(b))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}

