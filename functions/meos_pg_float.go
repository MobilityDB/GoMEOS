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

// TODO add_float4_float4: unsupported return type float4
// func AddFloat4Float4(...) { /* not yet handled by codegen */ }


// TODO add_float4_float8: unsupported return type float8
// func AddFloat4Float8(...) { /* not yet handled by codegen */ }


// TODO add_float8_float4: unsupported return type float8
// func AddFloat8Float4(...) { /* not yet handled by codegen */ }


// TODO add_float8_float8: unsupported return type float8
// func AddFloat8Float8(...) { /* not yet handled by codegen */ }


// TODO div_float4_float4: unsupported return type float4
// func DivFloat4Float4(...) { /* not yet handled by codegen */ }


// TODO div_float4_float8: unsupported return type float8
// func DivFloat4Float8(...) { /* not yet handled by codegen */ }


// TODO div_float8_float4: unsupported return type float8
// func DivFloat8Float4(...) { /* not yet handled by codegen */ }


// TODO div_float8_float8: unsupported return type float8
// func DivFloat8Float8(...) { /* not yet handled by codegen */ }


// TODO eq_float4_float4: unsupported param float4
// func EqFloat4Float4(...) { /* not yet handled by codegen */ }


// TODO eq_float4_float8: unsupported param float4
// func EqFloat4Float8(...) { /* not yet handled by codegen */ }


// TODO eq_float8_float4: unsupported param float8
// func EqFloat8Float4(...) { /* not yet handled by codegen */ }


// TODO eq_float8_float8: unsupported param float8
// func EqFloat8Float8(...) { /* not yet handled by codegen */ }


// TODO float4_abs: unsupported return type float4
// func Float4Abs(...) { /* not yet handled by codegen */ }


// TODO float4_cmp: unsupported param float4
// func Float4Cmp(...) { /* not yet handled by codegen */ }


// TODO float4_hash: unsupported param float4
// func Float4Hash(...) { /* not yet handled by codegen */ }


// TODO float4_hash_extended: unsupported param float4
// func Float4HashExtended(...) { /* not yet handled by codegen */ }


// TODO float4_in: unsupported return type float4
// func Float4In(...) { /* not yet handled by codegen */ }


// TODO float4_larger: unsupported return type float4
// func Float4Larger(...) { /* not yet handled by codegen */ }


// TODO float4_out: unsupported param float4
// func Float4Out(...) { /* not yet handled by codegen */ }


// TODO float4_smaller: unsupported return type float4
// func Float4Smaller(...) { /* not yet handled by codegen */ }


// TODO float4_to_float8: unsupported return type float8
// func Float4ToFloat8(...) { /* not yet handled by codegen */ }


// TODO float4_to_int16: unsupported param float4
// func Float4ToInt16(...) { /* not yet handled by codegen */ }


// TODO float4_to_int32: unsupported param float4
// func Float4ToInt32(...) { /* not yet handled by codegen */ }


// TODO float4_um: unsupported return type float4
// func Float4Um(...) { /* not yet handled by codegen */ }


// TODO float4_up: unsupported return type float4
// func Float4Up(...) { /* not yet handled by codegen */ }


// TODO float8_abs: unsupported return type float8
// func Float8Abs(...) { /* not yet handled by codegen */ }


// TODO float8_acos: unsupported return type float8
// func Float8Acos(...) { /* not yet handled by codegen */ }


// TODO float8_acosd: unsupported return type float8
// func Float8Acosd(...) { /* not yet handled by codegen */ }


// TODO float8_acosh: unsupported return type float8
// func Float8Acosh(...) { /* not yet handled by codegen */ }


// TODO float8_asin: unsupported return type float8
// func Float8Asin(...) { /* not yet handled by codegen */ }


// TODO float8_asind: unsupported return type float8
// func Float8Asind(...) { /* not yet handled by codegen */ }


// TODO float8_asinh: unsupported return type float8
// func Float8Asinh(...) { /* not yet handled by codegen */ }


// TODO float8_atan: unsupported return type float8
// func Float8Atan(...) { /* not yet handled by codegen */ }


// TODO float8_atan2: unsupported return type float8
// func Float8Atan2(...) { /* not yet handled by codegen */ }


// TODO float8_atan2d: unsupported return type float8
// func Float8Atan2d(...) { /* not yet handled by codegen */ }


// TODO float8_atand: unsupported return type float8
// func Float8Atand(...) { /* not yet handled by codegen */ }


// TODO float8_atanh: unsupported return type float8
// func Float8Atanh(...) { /* not yet handled by codegen */ }


// TODO float8_cbrt: unsupported return type float8
// func Float8Cbrt(...) { /* not yet handled by codegen */ }


// TODO float8_ceil: unsupported return type float8
// func Float8Ceil(...) { /* not yet handled by codegen */ }


// TODO float8_cmp: unsupported param float8
// func Float8Cmp(...) { /* not yet handled by codegen */ }


// TODO float8_cos: unsupported return type float8
// func Float8Cos(...) { /* not yet handled by codegen */ }


// TODO float8_cosd: unsupported return type float8
// func Float8Cosd(...) { /* not yet handled by codegen */ }


// TODO float8_cosh: unsupported return type float8
// func Float8Cosh(...) { /* not yet handled by codegen */ }


// TODO float8_cot: unsupported return type float8
// func Float8Cot(...) { /* not yet handled by codegen */ }


// TODO float8_cotd: unsupported return type float8
// func Float8Cotd(...) { /* not yet handled by codegen */ }


// TODO float8_degrees: unsupported return type float8
// func Float8Degrees(...) { /* not yet handled by codegen */ }


// TODO float8_exp: unsupported return type float8
// func Float8Exp(...) { /* not yet handled by codegen */ }


// TODO float8_floor: unsupported return type float8
// func Float8Floor(...) { /* not yet handled by codegen */ }


// TODO float8_gamma: unsupported return type float8
// func Float8Gamma(...) { /* not yet handled by codegen */ }


// TODO float8_hash: unsupported param float8
// func Float8Hash(...) { /* not yet handled by codegen */ }


// TODO float8_hash_extended: unsupported param float8
// func Float8HashExtended(...) { /* not yet handled by codegen */ }


// TODO float8_in: unsupported return type float8
// func Float8In(...) { /* not yet handled by codegen */ }


// TODO float8_larger: unsupported return type float8
// func Float8Larger(...) { /* not yet handled by codegen */ }


// TODO float8_lgamma: unsupported return type float8
// func Float8Lgamma(...) { /* not yet handled by codegen */ }


// TODO float8_ln: unsupported return type float8
// func Float8Ln(...) { /* not yet handled by codegen */ }


// TODO float8_log10: unsupported return type float8
// func Float8Log10(...) { /* not yet handled by codegen */ }


// TODO float8_pi: unsupported return type float8
// func Float8Pi(...) { /* not yet handled by codegen */ }


// TODO float8_pow: unsupported return type float8
// func Float8Pow(...) { /* not yet handled by codegen */ }


// TODO float8_radians: unsupported return type float8
// func Float8Radians(...) { /* not yet handled by codegen */ }


// TODO float8_rint: unsupported return type float8
// func Float8Rint(...) { /* not yet handled by codegen */ }


// TODO float8_round: unsupported return type float8
// func Float8Round(...) { /* not yet handled by codegen */ }


// TODO float8_sign: unsupported return type float8
// func Float8Sign(...) { /* not yet handled by codegen */ }


// TODO float8_sin: unsupported return type float8
// func Float8Sin(...) { /* not yet handled by codegen */ }


// TODO float8_sind: unsupported return type float8
// func Float8Sind(...) { /* not yet handled by codegen */ }


// TODO float8_sinh: unsupported return type float8
// func Float8Sinh(...) { /* not yet handled by codegen */ }


// TODO float8_smaller: unsupported return type float8
// func Float8Smaller(...) { /* not yet handled by codegen */ }


// TODO float8_sqrt: unsupported return type float8
// func Float8Sqrt(...) { /* not yet handled by codegen */ }


// TODO float8_tan: unsupported return type float8
// func Float8Tan(...) { /* not yet handled by codegen */ }


// TODO float8_tand: unsupported return type float8
// func Float8Tand(...) { /* not yet handled by codegen */ }


// TODO float8_tanh: unsupported return type float8
// func Float8Tanh(...) { /* not yet handled by codegen */ }


// TODO float8_to_float4: unsupported return type float4
// func Float8ToFloat4(...) { /* not yet handled by codegen */ }


// TODO float8_to_int16: unsupported param float8
// func Float8ToInt16(...) { /* not yet handled by codegen */ }


// TODO float8_to_int32: unsupported param float8
// func Float8ToInt32(...) { /* not yet handled by codegen */ }


// TODO float8_trunc: unsupported return type float8
// func Float8Trunc(...) { /* not yet handled by codegen */ }


// TODO float8_um: unsupported return type float8
// func Float8Um(...) { /* not yet handled by codegen */ }


// TODO float8_up: unsupported return type float8
// func Float8Up(...) { /* not yet handled by codegen */ }


// TODO float8_width_bucket: unsupported param float8
// func Float8WidthBucket(...) { /* not yet handled by codegen */ }


// TODO ge_float4_float4: unsupported param float4
// func GeFloat4Float4(...) { /* not yet handled by codegen */ }


// TODO ge_float4_float8: unsupported param float4
// func GeFloat4Float8(...) { /* not yet handled by codegen */ }


// TODO ge_float8_float4: unsupported param float8
// func GeFloat8Float4(...) { /* not yet handled by codegen */ }


// TODO ge_float8_float8: unsupported param float8
// func GeFloat8Float8(...) { /* not yet handled by codegen */ }


// TODO gt_float4_float4: unsupported param float4
// func GtFloat4Float4(...) { /* not yet handled by codegen */ }


// TODO gt_float4_float8: unsupported param float4
// func GtFloat4Float8(...) { /* not yet handled by codegen */ }


// TODO gt_float8_float4: unsupported param float8
// func GtFloat8Float4(...) { /* not yet handled by codegen */ }


// TODO gt_float8_float8: unsupported param float8
// func GtFloat8Float8(...) { /* not yet handled by codegen */ }


// TODO int16_to_float4: unsupported return type float4
// func Int16ToFloat4(...) { /* not yet handled by codegen */ }


// TODO int16_to_float8: unsupported return type float8
// func Int16ToFloat8(...) { /* not yet handled by codegen */ }


// TODO int32_to_float4: unsupported return type float4
// func Int32ToFloat4(...) { /* not yet handled by codegen */ }


// TODO int32_to_float8: unsupported return type float8
// func Int32ToFloat8(...) { /* not yet handled by codegen */ }


// TODO le_float4_float4: unsupported param float4
// func LeFloat4Float4(...) { /* not yet handled by codegen */ }


// TODO le_float4_float8: unsupported param float4
// func LeFloat4Float8(...) { /* not yet handled by codegen */ }


// TODO le_float8_float4: unsupported param float8
// func LeFloat8Float4(...) { /* not yet handled by codegen */ }


// TODO le_float8_float8: unsupported param float8
// func LeFloat8Float8(...) { /* not yet handled by codegen */ }


// TODO lt_float4_float4: unsupported param float4
// func LtFloat4Float4(...) { /* not yet handled by codegen */ }


// TODO lt_float4_float8: unsupported param float4
// func LtFloat4Float8(...) { /* not yet handled by codegen */ }


// TODO lt_float8_float4: unsupported param float8
// func LtFloat8Float4(...) { /* not yet handled by codegen */ }


// TODO lt_float8_float8: unsupported param float8
// func LtFloat8Float8(...) { /* not yet handled by codegen */ }


// TODO minus_float4_float4: unsupported return type float4
// func MinusFloat4Float4(...) { /* not yet handled by codegen */ }


// TODO minus_float4_float8: unsupported return type float8
// func MinusFloat4Float8(...) { /* not yet handled by codegen */ }


// TODO minus_float8_float4: unsupported return type float8
// func MinusFloat8Float4(...) { /* not yet handled by codegen */ }


// TODO minus_float8_float8: unsupported return type float8
// func MinusFloat8Float8(...) { /* not yet handled by codegen */ }


// TODO mul_float4_float4: unsupported return type float4
// func MulFloat4Float4(...) { /* not yet handled by codegen */ }


// TODO mul_float4_float8: unsupported return type float8
// func MulFloat4Float8(...) { /* not yet handled by codegen */ }


// TODO mul_float8_float4: unsupported return type float8
// func MulFloat8Float4(...) { /* not yet handled by codegen */ }


// TODO mul_float8_float8: unsupported return type float8
// func MulFloat8Float8(...) { /* not yet handled by codegen */ }


// TODO ne_float4_float4: unsupported param float4
// func NeFloat4Float4(...) { /* not yet handled by codegen */ }


// TODO ne_float4_float8: unsupported param float4
// func NeFloat4Float8(...) { /* not yet handled by codegen */ }


// TODO ne_float8_float4: unsupported param float8
// func NeFloat8Float4(...) { /* not yet handled by codegen */ }


// TODO ne_float8_float8: unsupported param float8
// func NeFloat8Float8(...) { /* not yet handled by codegen */ }

