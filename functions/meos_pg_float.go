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

// AddFloat4Float4 wraps MEOS C function add_float4_float4.
func AddFloat4Float4(num1 float32, num2 float32) (_r0 float32, _err error) {
	C.meos_errno_reset()
	_cret := C.add_float4_float4(C.float(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float32(_cret), nil
}


// AddFloat4Float8 wraps MEOS C function add_float4_float8.
func AddFloat4Float8(num1 float32, num2 float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.add_float4_float8(C.float(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// AddFloat8Float4 wraps MEOS C function add_float8_float4.
func AddFloat8Float4(num1 float64, num2 float32) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.add_float8_float4(C.double(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// AddFloat8Float8 wraps MEOS C function add_float8_float8.
func AddFloat8Float8(num1 float64, num2 float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.add_float8_float8(C.double(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DivFloat4Float4 wraps MEOS C function div_float4_float4.
func DivFloat4Float4(num1 float32, num2 float32) (_r0 float32, _err error) {
	C.meos_errno_reset()
	_cret := C.div_float4_float4(C.float(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float32(_cret), nil
}


// DivFloat4Float8 wraps MEOS C function div_float4_float8.
func DivFloat4Float8(num1 float32, num2 float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.div_float4_float8(C.float(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DivFloat8Float4 wraps MEOS C function div_float8_float4.
func DivFloat8Float4(num1 float64, num2 float32) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.div_float8_float4(C.double(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DivFloat8Float8 wraps MEOS C function div_float8_float8.
func DivFloat8Float8(num1 float64, num2 float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.div_float8_float8(C.double(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// EqFloat4Float4 wraps MEOS C function eq_float4_float4.
func EqFloat4Float4(num1 float32, num2 float32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_float4_float4(C.float(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqFloat4Float8 wraps MEOS C function eq_float4_float8.
func EqFloat4Float8(num1 float32, num2 float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_float4_float8(C.float(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqFloat8Float4 wraps MEOS C function eq_float8_float4.
func EqFloat8Float4(num1 float64, num2 float32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_float8_float4(C.double(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqFloat8Float8 wraps MEOS C function eq_float8_float8.
func EqFloat8Float8(num1 float64, num2 float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_float8_float8(C.double(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// Float4Abs wraps MEOS C function float4_abs.
func Float4Abs(num float32) (_r0 float32, _err error) {
	C.meos_errno_reset()
	_cret := C.float4_abs(C.float(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float32(_cret), nil
}


// Float4Cmp wraps MEOS C function float4_cmp.
func Float4Cmp(num1 float32, num2 float32) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.float4_cmp(C.float(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// Float4Hash wraps MEOS C function float4_hash.
func Float4Hash(num float32) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.float4_hash(C.float(num))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// Float4HashExtended wraps MEOS C function float4_hash_extended.
func Float4HashExtended(num float32, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.float4_hash_extended(C.float(num), C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// Float4In wraps MEOS C function float4_in.
func Float4In(num string) (_r0 float32, _err error) {
	_c_num := C.CString(num)
	defer C.free(unsafe.Pointer(_c_num))
	C.meos_errno_reset()
	_cret := C.float4_in(_c_num)
	if _err = meosError(); _err != nil {
		return
	}
	return float32(_cret), nil
}


// Float4Larger wraps MEOS C function float4_larger.
func Float4Larger(num1 float32, num2 float32) (_r0 float32, _err error) {
	C.meos_errno_reset()
	_cret := C.float4_larger(C.float(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float32(_cret), nil
}


// Float4Out wraps MEOS C function float4_out.
func Float4Out(num float32) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.float4_out(C.float(num))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// Float4Smaller wraps MEOS C function float4_smaller.
func Float4Smaller(num1 float32, num2 float32) (_r0 float32, _err error) {
	C.meos_errno_reset()
	_cret := C.float4_smaller(C.float(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float32(_cret), nil
}


// Float4ToFloat8 wraps MEOS C function float4_to_float8.
func Float4ToFloat8(num float32) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float4_to_float8(C.float(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float4ToInt16 wraps MEOS C function float4_to_int16.
func Float4ToInt16(num float32) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.float4_to_int16(C.float(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// Float4ToInt32 wraps MEOS C function float4_to_int32.
func Float4ToInt32(num float32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.float4_to_int32(C.float(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Float4Um wraps MEOS C function float4_um.
func Float4Um(num float32) (_r0 float32, _err error) {
	C.meos_errno_reset()
	_cret := C.float4_um(C.float(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float32(_cret), nil
}


// Float4Up wraps MEOS C function float4_up.
func Float4Up(num float32) (_r0 float32, _err error) {
	C.meos_errno_reset()
	_cret := C.float4_up(C.float(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float32(_cret), nil
}


// Float8Abs wraps MEOS C function float8_abs.
func Float8Abs(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_abs(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Acos wraps MEOS C function float8_acos.
func Float8Acos(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_acos(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Acosd wraps MEOS C function float8_acosd.
func Float8Acosd(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_acosd(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Acosh wraps MEOS C function float8_acosh.
func Float8Acosh(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_acosh(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Asin wraps MEOS C function float8_asin.
func Float8Asin(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_asin(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Asind wraps MEOS C function float8_asind.
func Float8Asind(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_asind(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Asinh wraps MEOS C function float8_asinh.
func Float8Asinh(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_asinh(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Atan wraps MEOS C function float8_atan.
func Float8Atan(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_atan(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Atan2 wraps MEOS C function float8_atan2.
func Float8Atan2(num1 float64, num2 float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_atan2(C.double(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Atan2d wraps MEOS C function float8_atan2d.
func Float8Atan2d(num1 float64, num2 float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_atan2d(C.double(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Atand wraps MEOS C function float8_atand.
func Float8Atand(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_atand(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Atanh wraps MEOS C function float8_atanh.
func Float8Atanh(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_atanh(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Cbrt wraps MEOS C function float8_cbrt.
func Float8Cbrt(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_cbrt(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Ceil wraps MEOS C function float8_ceil.
func Float8Ceil(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_ceil(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Cmp wraps MEOS C function float8_cmp.
func Float8Cmp(num1 float64, num2 float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_cmp(C.double(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// Float8Cos wraps MEOS C function float8_cos.
func Float8Cos(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_cos(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Cosd wraps MEOS C function float8_cosd.
func Float8Cosd(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_cosd(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Cosh wraps MEOS C function float8_cosh.
func Float8Cosh(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_cosh(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Cot wraps MEOS C function float8_cot.
func Float8Cot(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_cot(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Cotd wraps MEOS C function float8_cotd.
func Float8Cotd(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_cotd(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Degrees wraps MEOS C function float8_degrees.
func Float8Degrees(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_degrees(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Exp wraps MEOS C function float8_exp.
func Float8Exp(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_exp(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Floor wraps MEOS C function float8_floor.
func Float8Floor(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_floor(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Gamma wraps MEOS C function float8_gamma.
func Float8Gamma(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_gamma(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Hash wraps MEOS C function float8_hash.
func Float8Hash(num float64) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_hash(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// Float8HashExtended wraps MEOS C function float8_hash_extended.
func Float8HashExtended(num float64, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_hash_extended(C.double(num), C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// Float8In wraps MEOS C function float8_in.
func Float8In(str string) (_r0 float64, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.float8_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Larger wraps MEOS C function float8_larger.
func Float8Larger(num1 float64, num2 float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_larger(C.double(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Lgamma wraps MEOS C function float8_lgamma.
func Float8Lgamma(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_lgamma(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Ln wraps MEOS C function float8_ln.
func Float8Ln(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_ln(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Log10 wraps MEOS C function float8_log10.
func Float8Log10(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_log10(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Pi wraps MEOS C function float8_pi.
func Float8Pi() (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_pi()
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Pow wraps MEOS C function float8_pow.
func Float8Pow(num1 float64, num2 float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_pow(C.double(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Radians wraps MEOS C function float8_radians.
func Float8Radians(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_radians(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Rint wraps MEOS C function float8_rint.
func Float8Rint(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_rint(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Round wraps MEOS C function float8_round.
func Float8Round(num float64, maxdd int) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_round(C.double(num), C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Sign wraps MEOS C function float8_sign.
func Float8Sign(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_sign(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Sin wraps MEOS C function float8_sin.
func Float8Sin(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_sin(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Sind wraps MEOS C function float8_sind.
func Float8Sind(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_sind(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Sinh wraps MEOS C function float8_sinh.
func Float8Sinh(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_sinh(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Smaller wraps MEOS C function float8_smaller.
func Float8Smaller(num1 float64, num2 float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_smaller(C.double(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Sqrt wraps MEOS C function float8_sqrt.
func Float8Sqrt(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_sqrt(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Tan wraps MEOS C function float8_tan.
func Float8Tan(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_tan(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Tand wraps MEOS C function float8_tand.
func Float8Tand(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_tand(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Tanh wraps MEOS C function float8_tanh.
func Float8Tanh(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_tanh(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8ToFloat4 wraps MEOS C function float8_to_float4.
func Float8ToFloat4(num float64) (_r0 float32, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_to_float4(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float32(_cret), nil
}


// Float8ToInt16 wraps MEOS C function float8_to_int16.
func Float8ToInt16(num float64) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_to_int16(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// Float8ToInt32 wraps MEOS C function float8_to_int32.
func Float8ToInt32(num float64) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_to_int32(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// Float8Trunc wraps MEOS C function float8_trunc.
func Float8Trunc(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_trunc(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Um wraps MEOS C function float8_um.
func Float8Um(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_um(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8Up wraps MEOS C function float8_up.
func Float8Up(num float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_up(C.double(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Float8WidthBucket wraps MEOS C function float8_width_bucket.
func Float8WidthBucket(num float64, bound1 float64, bound2 float64, count int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_width_bucket(C.double(num), C.double(bound1), C.double(bound2), C.int32_t(count))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// GeFloat4Float4 wraps MEOS C function ge_float4_float4.
func GeFloat4Float4(num1 float32, num2 float32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_float4_float4(C.float(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeFloat4Float8 wraps MEOS C function ge_float4_float8.
func GeFloat4Float8(num1 float32, num2 float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_float4_float8(C.float(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeFloat8Float4 wraps MEOS C function ge_float8_float4.
func GeFloat8Float4(num1 float64, num2 float32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_float8_float4(C.double(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeFloat8Float8 wraps MEOS C function ge_float8_float8.
func GeFloat8Float8(num1 float64, num2 float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_float8_float8(C.double(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtFloat4Float4 wraps MEOS C function gt_float4_float4.
func GtFloat4Float4(num1 float32, num2 float32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_float4_float4(C.float(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtFloat4Float8 wraps MEOS C function gt_float4_float8.
func GtFloat4Float8(num1 float32, num2 float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_float4_float8(C.float(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtFloat8Float4 wraps MEOS C function gt_float8_float4.
func GtFloat8Float4(num1 float64, num2 float32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_float8_float4(C.double(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtFloat8Float8 wraps MEOS C function gt_float8_float8.
func GtFloat8Float8(num1 float64, num2 float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_float8_float8(C.double(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// Int16ToFloat4 wraps MEOS C function int16_to_float4.
func Int16ToFloat4(num int16) (_r0 float32, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_to_float4(C.int16_t(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float32(_cret), nil
}


// Int16ToFloat8 wraps MEOS C function int16_to_float8.
func Int16ToFloat8(num int16) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.int16_to_float8(C.int16_t(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Int32ToFloat4 wraps MEOS C function int32_to_float4.
func Int32ToFloat4(num int32) (_r0 float32, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_to_float4(C.int32_t(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float32(_cret), nil
}


// Int32ToFloat8 wraps MEOS C function int32_to_float8.
func Int32ToFloat8(num int32) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_to_float8(C.int32_t(num))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// LeFloat4Float4 wraps MEOS C function le_float4_float4.
func LeFloat4Float4(num1 float32, num2 float32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_float4_float4(C.float(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeFloat4Float8 wraps MEOS C function le_float4_float8.
func LeFloat4Float8(num1 float32, num2 float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_float4_float8(C.float(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeFloat8Float4 wraps MEOS C function le_float8_float4.
func LeFloat8Float4(num1 float64, num2 float32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_float8_float4(C.double(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeFloat8Float8 wraps MEOS C function le_float8_float8.
func LeFloat8Float8(num1 float64, num2 float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_float8_float8(C.double(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtFloat4Float4 wraps MEOS C function lt_float4_float4.
func LtFloat4Float4(num1 float32, num2 float32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_float4_float4(C.float(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtFloat4Float8 wraps MEOS C function lt_float4_float8.
func LtFloat4Float8(num1 float32, num2 float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_float4_float8(C.float(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtFloat8Float4 wraps MEOS C function lt_float8_float4.
func LtFloat8Float4(num1 float64, num2 float32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_float8_float4(C.double(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtFloat8Float8 wraps MEOS C function lt_float8_float8.
func LtFloat8Float8(num1 float64, num2 float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_float8_float8(C.double(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// MinusFloat4Float4 wraps MEOS C function minus_float4_float4.
func MinusFloat4Float4(num1 float32, num2 float32) (_r0 float32, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_float4_float4(C.float(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float32(_cret), nil
}


// MinusFloat4Float8 wraps MEOS C function minus_float4_float8.
func MinusFloat4Float8(num1 float32, num2 float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_float4_float8(C.float(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// MinusFloat8Float4 wraps MEOS C function minus_float8_float4.
func MinusFloat8Float4(num1 float64, num2 float32) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_float8_float4(C.double(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// MinusFloat8Float8 wraps MEOS C function minus_float8_float8.
func MinusFloat8Float8(num1 float64, num2 float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_float8_float8(C.double(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// MulFloat4Float4 wraps MEOS C function mul_float4_float4.
func MulFloat4Float4(num1 float32, num2 float32) (_r0 float32, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_float4_float4(C.float(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float32(_cret), nil
}


// MulFloat4Float8 wraps MEOS C function mul_float4_float8.
func MulFloat4Float8(num1 float32, num2 float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_float4_float8(C.float(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// MulFloat8Float4 wraps MEOS C function mul_float8_float4.
func MulFloat8Float4(num1 float64, num2 float32) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_float8_float4(C.double(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// MulFloat8Float8 wraps MEOS C function mul_float8_float8.
func MulFloat8Float8(num1 float64, num2 float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_float8_float8(C.double(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NeFloat4Float4 wraps MEOS C function ne_float4_float4.
func NeFloat4Float4(num1 float32, num2 float32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_float4_float4(C.float(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NeFloat4Float8 wraps MEOS C function ne_float4_float8.
func NeFloat4Float8(num1 float32, num2 float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_float4_float8(C.float(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NeFloat8Float4 wraps MEOS C function ne_float8_float4.
func NeFloat8Float4(num1 float64, num2 float32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_float8_float4(C.double(num1), C.float(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NeFloat8Float8 wraps MEOS C function ne_float8_float8.
func NeFloat8Float8(num1 float64, num2 float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_float8_float8(C.double(num1), C.double(num2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}

