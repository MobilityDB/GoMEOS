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

// AddInt16Int16 wraps MEOS C function add_int16_int16.
func AddInt16Int16(num1 int16, num2 int16) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.add_int16_int16(C.int16(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// AddInt16Int32 wraps MEOS C function add_int16_int32.
func AddInt16Int32(num1 int16, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.add_int16_int32(C.int16(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// AddInt16Int64 wraps MEOS C function add_int16_int64.
func AddInt16Int64(num1 int16, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.add_int16_int64(C.int16(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// AddInt32Int16 wraps MEOS C function add_int32_int16.
func AddInt32Int16(num1 int32, num2 int16) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.add_int32_int16(C.int32(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// AddInt32Int32 wraps MEOS C function add_int32_int32.
func AddInt32Int32(num1 int32, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.add_int32_int32(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// AddInt32Int64 wraps MEOS C function add_int32_int64.
func AddInt32Int64(num1 int32, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.add_int32_int64(C.int32(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// AddInt64Int16 wraps MEOS C function add_int64_int16.
func AddInt64Int16(num1 int64, num2 int16) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.add_int64_int16(C.int64_t(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// AddInt64Int32 wraps MEOS C function add_int64_int32.
func AddInt64Int32(num1 int64, num2 int32) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.add_int64_int32(C.int64_t(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// AddInt64Int64 wraps MEOS C function add_int64_int64.
func AddInt64Int64(num1 int64, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.add_int64_int64(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// BoolToInt32 wraps MEOS C function bool_to_int32.
func BoolToInt32(b bool) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.bool_to_int32(C.bool(b))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// DivInt16Int16 wraps MEOS C function div_int16_int16.
func DivInt16Int16(num1 int16, num2 int16) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.div_int16_int16(C.int16(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// DivInt16Int32 wraps MEOS C function div_int16_int32.
func DivInt16Int32(num1 int16, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.div_int16_int32(C.int16(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// DivInt16Int64 wraps MEOS C function div_int16_int64.
func DivInt16Int64(num1 int16, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.div_int16_int64(C.int16(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// DivInt32Int16 wraps MEOS C function div_int32_int16.
func DivInt32Int16(num1 int32, num2 int16) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.div_int32_int16(C.int32(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// DivInt32Int32 wraps MEOS C function div_int32_int32.
func DivInt32Int32(num1 int32, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.div_int32_int32(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// DivInt32Int64 wraps MEOS C function div_int32_int64.
func DivInt32Int64(num1 int32, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.div_int32_int64(C.int32(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// DivInt64Int16 wraps MEOS C function div_int64_int16.
func DivInt64Int16(num1 int64, num2 int16) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.div_int64_int16(C.int64_t(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// DivInt64Int32 wraps MEOS C function div_int64_int32.
func DivInt64Int32(num1 int64, num2 int32) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.div_int64_int32(C.int64_t(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// DivInt64Int64 wraps MEOS C function div_int64_int64.
func DivInt64Int64(num1 int64, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.div_int64_int64(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// EqInt16Int16 wraps MEOS C function eq_int16_int16.
func EqInt16Int16(num1 int16, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_int16_int16(C.int16(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqInt16Int32 wraps MEOS C function eq_int16_int32.
func EqInt16Int32(num1 int16, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_int16_int32(C.int16(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqInt16Int64 wraps MEOS C function eq_int16_int64.
func EqInt16Int64(num1 int16, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_int16_int64(C.int16(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqInt32Int16 wraps MEOS C function eq_int32_int16.
func EqInt32Int16(num1 int32, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_int32_int16(C.int32(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqInt32Int32 wraps MEOS C function eq_int32_int32.
func EqInt32Int32(num1 int32, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_int32_int32(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqInt32Int64 wraps MEOS C function eq_int32_int64.
func EqInt32Int64(num1 int32, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_int32_int64(C.int32(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqInt64Int16 wraps MEOS C function eq_int64_int16.
func EqInt64Int16(num1 int64, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_int64_int16(C.int64_t(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqInt64Int32 wraps MEOS C function eq_int64_int32.
func EqInt64Int32(num1 int64, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_int64_int32(C.int64_t(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqInt64Int64 wraps MEOS C function eq_int64_int64.
func EqInt64Int64(num1 int64, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_int64_int64(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TODO float4_to_int64: unsupported param float4
// func Float4ToInt64(...) { /* not yet handled by codegen */ }


// TODO float8_to_int64: unsupported param float8
// func Float8ToInt64(...) { /* not yet handled by codegen */ }


// GeInt16Int16 wraps MEOS C function ge_int16_int16.
func GeInt16Int16(num1 int16, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_int16_int16(C.int16(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeInt16Int32 wraps MEOS C function ge_int16_int32.
func GeInt16Int32(num1 int16, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_int16_int32(C.int16(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeInt16Int64 wraps MEOS C function ge_int16_int64.
func GeInt16Int64(num1 int16, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_int16_int64(C.int16(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeInt32Int16 wraps MEOS C function ge_int32_int16.
func GeInt32Int16(num1 int32, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_int32_int16(C.int32(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeInt32Int32 wraps MEOS C function ge_int32_int32.
func GeInt32Int32(num1 int32, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_int32_int32(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeInt32Int64 wraps MEOS C function ge_int32_int64.
func GeInt32Int64(num1 int32, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_int32_int64(C.int32(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeInt64Int16 wraps MEOS C function ge_int64_int16.
func GeInt64Int16(num1 int64, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_int64_int16(C.int64_t(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeInt64Int32 wraps MEOS C function ge_int64_int32.
func GeInt64Int32(num1 int64, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_int64_int32(C.int64_t(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeInt64Int64 wraps MEOS C function ge_int64_int64.
func GeInt64Int64(num1 int64, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_int64_int64(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtInt16Int16 wraps MEOS C function gt_int16_int16.
func GtInt16Int16(num1 int16, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_int16_int16(C.int16(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtInt16Int32 wraps MEOS C function gt_int16_int32.
func GtInt16Int32(num1 int16, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_int16_int32(C.int16(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtInt16Int64 wraps MEOS C function gt_int16_int64.
func GtInt16Int64(num1 int16, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_int16_int64(C.int16(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtInt32Int16 wraps MEOS C function gt_int32_int16.
func GtInt32Int16(num1 int32, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_int32_int16(C.int32(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtInt32Int32 wraps MEOS C function gt_int32_int32.
func GtInt32Int32(num1 int32, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_int32_int32(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtInt32Int64 wraps MEOS C function gt_int32_int64.
func GtInt32Int64(num1 int32, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_int32_int64(C.int32(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtInt64Int16 wraps MEOS C function gt_int64_int16.
func GtInt64Int16(num1 int64, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_int64_int16(C.int64_t(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtInt64Int32 wraps MEOS C function gt_int64_int32.
func GtInt64Int32(num1 int64, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_int64_int32(C.int64_t(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtInt64Int64 wraps MEOS C function gt_int64_int64.
func GtInt64Int64(num1 int64, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_int64_int64(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// Int16Abs wraps MEOS C function int16_abs.
func Int16Abs(num int16) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_abs(C.int16(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// Int16And wraps MEOS C function int16_and.
func Int16And(num1 int16, num2 int16) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_and(C.int16(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// Int16Hash wraps MEOS C function int16_hash.
func Int16Hash(val int16) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_hash(C.int16(val))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// Int16HashExtended wraps MEOS C function int16_hash_extended.
func Int16HashExtended(val int16, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_hash_extended(C.int16(val), C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// Int16In wraps MEOS C function int16_in.
func Int16In(str string) (_r0 int16, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.int16_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// Int16Larger wraps MEOS C function int16_larger.
func Int16Larger(num1 int16, num2 int16) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_larger(C.int16(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// Int16Mod wraps MEOS C function int16_mod.
func Int16Mod(num1 int16, num2 int16) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_mod(C.int16(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// Int16Not wraps MEOS C function int16_not.
func Int16Not(num int16) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_not(C.int16(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// Int16Or wraps MEOS C function int16_or.
func Int16Or(num1 int16, num2 int16) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_or(C.int16(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// Int16Out wraps MEOS C function int16_out.
func Int16Out(num int16) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_out(C.int16(num))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// Int16Shl wraps MEOS C function int16_shl.
func Int16Shl(num1 int16, num2 int32) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_shl(C.int16(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// Int16Shr wraps MEOS C function int16_shr.
func Int16Shr(num1 int16, num2 int32) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_shr(C.int16(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// Int16Smaller wraps MEOS C function int16_smaller.
func Int16Smaller(num1 int16, num2 int16) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_smaller(C.int16(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// Int16ToInt32 wraps MEOS C function int16_to_int32.
func Int16ToInt32(num int16) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_to_int32(C.int16(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int16ToInt64 wraps MEOS C function int16_to_int64.
func Int16ToInt64(num int16) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_to_int64(C.int16(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// Int16Uminus wraps MEOS C function int16_uminus.
func Int16Uminus(num int16) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_uminus(C.int16(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// Int16Uplus wraps MEOS C function int16_uplus.
func Int16Uplus(num int16) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_uplus(C.int16(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// Int16Xor wraps MEOS C function int16_xor.
func Int16Xor(num1 int16, num2 int16) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_xor(C.int16(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// Int32Abs wraps MEOS C function int32_abs.
func Int32Abs(num int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_abs(C.int32(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int32And wraps MEOS C function int32_and.
func Int32And(num1 int32, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_and(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int32Cmp wraps MEOS C function int32_cmp.
func Int32Cmp(l int32, r int32) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_cmp(C.int32(l), C.int32(r))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// Int32Gcd wraps MEOS C function int32_gcd.
func Int32Gcd(num1 int32, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_gcd(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int32Hash wraps MEOS C function int32_hash.
func Int32Hash(val int32) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_hash(C.int32(val))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// Int32HashExtended wraps MEOS C function int32_hash_extended.
func Int32HashExtended(val int32, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_hash_extended(C.int32(val), C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// Int32In wraps MEOS C function int32_in.
func Int32In(str string) (_r0 int32, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.int32_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int32Inc wraps MEOS C function int32_inc.
func Int32Inc(num int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_inc(C.int32(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int32Larger wraps MEOS C function int32_larger.
func Int32Larger(num1 int32, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_larger(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int32Lcm wraps MEOS C function int32_lcm.
func Int32Lcm(num1 int32, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_lcm(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int32Mod wraps MEOS C function int32_mod.
func Int32Mod(num1 int32, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_mod(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int32Not wraps MEOS C function int32_not.
func Int32Not(num int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_not(C.int32(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int32Or wraps MEOS C function int32_or.
func Int32Or(num1 int32, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_or(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int32Out wraps MEOS C function int32_out.
func Int32Out(num int32) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_out(C.int32(num))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// Int32Shl wraps MEOS C function int32_shl.
func Int32Shl(num1 int32, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_shl(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int32Shr wraps MEOS C function int32_shr.
func Int32Shr(num1 int32, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_shr(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int32Smaller wraps MEOS C function int32_smaller.
func Int32Smaller(num1 int32, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_smaller(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int32ToBool wraps MEOS C function int32_to_bool.
func Int32ToBool(num int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_to_bool(C.int32(num))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// Int32ToInt16 wraps MEOS C function int32_to_int16.
func Int32ToInt16(num int32) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_to_int16(C.int32(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// Int32ToInt64 wraps MEOS C function int32_to_int64.
func Int32ToInt64(num int32) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_to_int64(C.int32(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// Int32Uminus wraps MEOS C function int32_uminus.
func Int32Uminus(num int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_uminus(C.int32(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int32Uplus wraps MEOS C function int32_uplus.
func Int32Uplus(num int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_uplus(C.int32(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int32Xor wraps MEOS C function int32_xor.
func Int32Xor(num1 int32, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_xor(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int64Abs wraps MEOS C function int64_abs.
func Int64Abs(num int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_abs(C.int64_t(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// Int64And wraps MEOS C function int64_and.
func Int64And(num1 int64, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_and(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// Int64Cmp wraps MEOS C function int64_cmp.
func Int64Cmp(l int64, r int64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_cmp(C.int64_t(l), C.int64_t(r))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// Int64Dec wraps MEOS C function int64_dec.
func Int64Dec(num int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_dec(C.int64_t(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// Int64Gcd wraps MEOS C function int64_gcd.
func Int64Gcd(num1 int64, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_gcd(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// Int64Hash wraps MEOS C function int64_hash.
func Int64Hash(num int64) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_hash(C.int64_t(num))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// Int64HashExtended wraps MEOS C function int64_hash_extended.
func Int64HashExtended(num int64, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_hash_extended(C.int64_t(num), C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// Int64In wraps MEOS C function int64_in.
func Int64In(str string) (_r0 int64, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.int64_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// Int64Inc wraps MEOS C function int64_inc.
func Int64Inc(num int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_inc(C.int64_t(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// Int64Larger wraps MEOS C function int64_larger.
func Int64Larger(num1 int64, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_larger(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// Int64Lcm wraps MEOS C function int64_lcm.
func Int64Lcm(num1 int64, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_lcm(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// Int64Mod wraps MEOS C function int64_mod.
func Int64Mod(num1 int64, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_mod(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// Int64Not wraps MEOS C function int64_not.
func Int64Not(num int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_not(C.int64_t(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// Int64Or wraps MEOS C function int64_or.
func Int64Or(num1 int64, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_or(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// Int64Out wraps MEOS C function int64_out.
func Int64Out(num int64) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_out(C.int64_t(num))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// Int64Shl wraps MEOS C function int64_shl.
func Int64Shl(num1 int64, num2 int32) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_shl(C.int64_t(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// Int64Shr wraps MEOS C function int64_shr.
func Int64Shr(num1 int64, num2 int32) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_shr(C.int64_t(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// Int64Smaller wraps MEOS C function int64_smaller.
func Int64Smaller(num1 int64, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_smaller(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TODO int64_to_float4: unsupported return type float4
// func Int64ToFloat4(...) { /* not yet handled by codegen */ }


// TODO int64_to_float8: unsupported return type float8
// func Int64ToFloat8(...) { /* not yet handled by codegen */ }


// Int64ToInt16 wraps MEOS C function int64_to_int16.
func Int64ToInt16(num int64) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_to_int16(C.int64_t(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// Int64ToInt32 wraps MEOS C function int64_to_int32.
func Int64ToInt32(num int64) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_to_int32(C.int64_t(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Int64Uminus wraps MEOS C function int64_uminus.
func Int64Uminus(num int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_uminus(C.int64_t(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// Int64Uplus wraps MEOS C function int64_uplus.
func Int64Uplus(num int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_uplus(C.int64_t(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// Int64Xor wraps MEOS C function int64_xor.
func Int64Xor(num1 int64, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_xor(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// LeInt16Int16 wraps MEOS C function le_int16_int16.
func LeInt16Int16(num1 int16, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_int16_int16(C.int16(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeInt16Int32 wraps MEOS C function le_int16_int32.
func LeInt16Int32(num1 int16, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_int16_int32(C.int16(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeInt16Int64 wraps MEOS C function le_int16_int64.
func LeInt16Int64(num1 int16, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_int16_int64(C.int16(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeInt32Int16 wraps MEOS C function le_int32_int16.
func LeInt32Int16(num1 int32, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_int32_int16(C.int32(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeInt32Int32 wraps MEOS C function le_int32_int32.
func LeInt32Int32(num1 int32, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_int32_int32(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeInt32Int64 wraps MEOS C function le_int32_int64.
func LeInt32Int64(num1 int32, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_int32_int64(C.int32(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeInt64Int16 wraps MEOS C function le_int64_int16.
func LeInt64Int16(num1 int64, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_int64_int16(C.int64_t(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeInt64Int32 wraps MEOS C function le_int64_int32.
func LeInt64Int32(num1 int64, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_int64_int32(C.int64_t(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeInt64Int64 wraps MEOS C function le_int64_int64.
func LeInt64Int64(num1 int64, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_int64_int64(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtInt16Int16 wraps MEOS C function lt_int16_int16.
func LtInt16Int16(num1 int16, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_int16_int16(C.int16(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtInt16Int32 wraps MEOS C function lt_int16_int32.
func LtInt16Int32(num1 int16, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_int16_int32(C.int16(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtInt16Int64 wraps MEOS C function lt_int16_int64.
func LtInt16Int64(num1 int16, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_int16_int64(C.int16(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtInt32Int16 wraps MEOS C function lt_int32_int16.
func LtInt32Int16(num1 int32, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_int32_int16(C.int32(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtInt32Int32 wraps MEOS C function lt_int32_int32.
func LtInt32Int32(num1 int32, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_int32_int32(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtInt32Int64 wraps MEOS C function lt_int32_int64.
func LtInt32Int64(num1 int32, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_int32_int64(C.int32(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtInt64Int16 wraps MEOS C function lt_int64_int16.
func LtInt64Int16(num1 int64, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_int64_int16(C.int64_t(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtInt64Int32 wraps MEOS C function lt_int64_int32.
func LtInt64Int32(num1 int64, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_int64_int32(C.int64_t(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtInt64Int64 wraps MEOS C function lt_int64_int64.
func LtInt64Int64(num1 int64, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_int64_int64(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// MinusInt16Int16 wraps MEOS C function minus_int16_int16.
func MinusInt16Int16(num1 int16, num2 int16) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_int16_int16(C.int16(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// MinusInt16Int32 wraps MEOS C function minus_int16_int32.
func MinusInt16Int32(num1 int16, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_int16_int32(C.int16(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// MinusInt16Int64 wraps MEOS C function minus_int16_int64.
func MinusInt16Int64(num1 int16, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_int16_int64(C.int16(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// MinusInt32Int16 wraps MEOS C function minus_int32_int16.
func MinusInt32Int16(num1 int32, num2 int16) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_int32_int16(C.int32(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// MinusInt32Int32 wraps MEOS C function minus_int32_int32.
func MinusInt32Int32(num1 int32, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_int32_int32(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// MinusInt32Int64 wraps MEOS C function minus_int32_int64.
func MinusInt32Int64(num1 int32, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_int32_int64(C.int32(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// MinusInt64Int16 wraps MEOS C function minus_int64_int16.
func MinusInt64Int16(num1 int64, num2 int16) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_int64_int16(C.int64_t(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// MinusInt64Int32 wraps MEOS C function minus_int64_int32.
func MinusInt64Int32(num1 int64, num2 int32) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_int64_int32(C.int64_t(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// MinusInt64Int64 wraps MEOS C function minus_int64_int64.
func MinusInt64Int64(num1 int64, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_int64_int64(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// MulInt16Int16 wraps MEOS C function mul_int16_int16.
func MulInt16Int16(num1 int16, num2 int16) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_int16_int16(C.int16(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// MulInt16Int32 wraps MEOS C function mul_int16_int32.
func MulInt16Int32(num1 int16, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_int16_int32(C.int16(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// MulInt16Int64 wraps MEOS C function mul_int16_int64.
func MulInt16Int64(num1 int16, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_int16_int64(C.int16(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// MulInt32Int16 wraps MEOS C function mul_int32_int16.
func MulInt32Int16(num1 int32, num2 int16) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_int32_int16(C.int32(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// MulInt32Int32 wraps MEOS C function mul_int32_int32.
func MulInt32Int32(num1 int32, num2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_int32_int32(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// MulInt32Int64 wraps MEOS C function mul_int32_int64.
func MulInt32Int64(num1 int32, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_int32_int64(C.int32(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// MulInt64Int16 wraps MEOS C function mul_int64_int16.
func MulInt64Int16(num1 int64, num2 int16) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_int64_int16(C.int64_t(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// MulInt64Int32 wraps MEOS C function mul_int64_int32.
func MulInt64Int32(num1 int64, num2 int32) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_int64_int32(C.int64_t(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// MulInt64Int64 wraps MEOS C function mul_int64_int64.
func MulInt64Int64(num1 int64, num2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_int64_int64(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// NeInt16Int16 wraps MEOS C function ne_int16_int16.
func NeInt16Int16(num1 int16, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_int16_int16(C.int16(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NeInt16Int32 wraps MEOS C function ne_int16_int32.
func NeInt16Int32(num1 int16, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_int16_int32(C.int16(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NeInt16Int64 wraps MEOS C function ne_int16_int64.
func NeInt16Int64(num1 int16, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_int16_int64(C.int16(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NeInt32Int16 wraps MEOS C function ne_int32_int16.
func NeInt32Int16(num1 int32, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_int32_int16(C.int32(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NeInt32Int32 wraps MEOS C function ne_int32_int32.
func NeInt32Int32(num1 int32, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_int32_int32(C.int32(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NeInt32Int64 wraps MEOS C function ne_int32_int64.
func NeInt32Int64(num1 int32, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_int32_int64(C.int32(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NeInt64Int16 wraps MEOS C function ne_int64_int16.
func NeInt64Int16(num1 int64, num2 int16) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_int64_int16(C.int64_t(num1), C.int16(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NeInt64Int32 wraps MEOS C function ne_int64_int32.
func NeInt64Int32(num1 int64, num2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_int64_int32(C.int64_t(num1), C.int32(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NeInt64Int64 wraps MEOS C function ne_int64_int64.
func NeInt64Int64(num1 int64, num2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_int64_int64(C.int64_t(num1), C.int64_t(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}

