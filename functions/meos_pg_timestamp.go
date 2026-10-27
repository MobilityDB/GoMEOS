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

// AddTimestampInterval wraps MEOS C function add_timestamp_interval.
func AddTimestampInterval(ts int64, interv *Interval) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.add_timestamp_interval(C.Timestamp(ts), interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// AddTimestamptzInterval wraps MEOS C function add_timestamptz_interval.
func AddTimestamptzInterval(tstz int64, interv *Interval) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.add_timestamptz_interval(C.TimestampTz(tstz), interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// AddTimestamptzIntervalAtZone wraps MEOS C function add_timestamptz_interval_at_zone.
func AddTimestamptzIntervalAtZone(tstz int64, interv *Interval, zone string) (_r0 int64, _err error) {
	_c_zone := C.cstring_to_text(C.CString(zone))
	defer C.free(unsafe.Pointer(_c_zone))
	C.meos_errno_reset()
	_cret := C.add_timestamptz_interval_at_zone(C.TimestampTz(tstz), interv._inner, _c_zone)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// CmpTimestampTimestamp wraps MEOS C function cmp_timestamp_timestamp.
func CmpTimestampTimestamp(ts1 int64, ts2 int64) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.cmp_timestamp_timestamp(C.Timestamp(ts1), C.Timestamp(ts2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// CmpTimestampTimestamptz wraps MEOS C function cmp_timestamp_timestamptz.
func CmpTimestampTimestamptz(ts int64, tstz int64) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.cmp_timestamp_timestamptz(C.Timestamp(ts), C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// CmpTimestamptzTimestamp wraps MEOS C function cmp_timestamptz_timestamp.
func CmpTimestamptzTimestamp(tstz int64, ts int64) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.cmp_timestamptz_timestamp(C.TimestampTz(tstz), C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// EqTimestampTimestamp wraps MEOS C function eq_timestamp_timestamp.
func EqTimestampTimestamp(ts1 int64, ts2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_timestamp_timestamp(C.Timestamp(ts1), C.Timestamp(ts2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqTimestampTimestamptz wraps MEOS C function eq_timestamp_timestamptz.
func EqTimestampTimestamptz(ts int64, tstz int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_timestamp_timestamptz(C.Timestamp(ts), C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqTimestamptzTimestamp wraps MEOS C function eq_timestamptz_timestamp.
func EqTimestamptzTimestamp(tstz int64, ts int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_timestamptz_timestamp(C.TimestampTz(tstz), C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqTimestamptzTimestamptz wraps MEOS C function eq_timestamptz_timestamptz.
func EqTimestamptzTimestamptz(tstz1 int64, tstz2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_timestamptz_timestamptz(C.TimestampTz(tstz1), C.TimestampTz(tstz2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TODO float8_to_timestamptz: unsupported param float8
// func Float8ToTimestamptz(...) { /* not yet handled by codegen */ }


// GtTimestampTimestamp wraps MEOS C function gt_timestamp_timestamp.
func GtTimestampTimestamp(ts1 int64, ts2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_timestamp_timestamp(C.Timestamp(ts1), C.Timestamp(ts2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtTimestampTimestamptz wraps MEOS C function gt_timestamp_timestamptz.
func GtTimestampTimestamptz(ts int64, tstz int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_timestamp_timestamptz(C.Timestamp(ts), C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtTimestamptzTimestamp wraps MEOS C function gt_timestamptz_timestamp.
func GtTimestamptzTimestamp(tstz int64, ts int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_timestamptz_timestamp(C.TimestampTz(tstz), C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeTimestampTimestamp wraps MEOS C function ge_timestamp_timestamp.
func GeTimestampTimestamp(ts1 int64, ts2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_timestamp_timestamp(C.Timestamp(ts1), C.Timestamp(ts2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeTimestampTimestamptz wraps MEOS C function ge_timestamp_timestamptz.
func GeTimestampTimestamptz(ts int64, tstz int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_timestamp_timestamptz(C.Timestamp(ts), C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeTimestamptzTimestamp wraps MEOS C function ge_timestamptz_timestamp.
func GeTimestamptzTimestamp(tstz int64, ts int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_timestamptz_timestamp(C.TimestampTz(tstz), C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeTimestampTimestamp wraps MEOS C function le_timestamp_timestamp.
func LeTimestampTimestamp(ts1 int64, ts2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_timestamp_timestamp(C.Timestamp(ts1), C.Timestamp(ts2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeTimestampTimestamptz wraps MEOS C function le_timestamp_timestamptz.
func LeTimestampTimestamptz(ts int64, tstz int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_timestamp_timestamptz(C.Timestamp(ts), C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeTimestamptzTimestamp wraps MEOS C function le_timestamptz_timestamp.
func LeTimestamptzTimestamp(tstz int64, ts int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_timestamptz_timestamp(C.TimestampTz(tstz), C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtTimestampTimestamp wraps MEOS C function lt_timestamp_timestamp.
func LtTimestampTimestamp(ts1 int64, ts2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_timestamp_timestamp(C.Timestamp(ts1), C.Timestamp(ts2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtTimestampTimestamptz wraps MEOS C function lt_timestamp_timestamptz.
func LtTimestampTimestamptz(ts int64, tstz int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_timestamp_timestamptz(C.Timestamp(ts), C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtTimestamptzTimestamp wraps MEOS C function lt_timestamptz_timestamp.
func LtTimestamptzTimestamp(tstz int64, ts int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_timestamptz_timestamp(C.TimestampTz(tstz), C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// MinusTimestampInterval wraps MEOS C function minus_timestamp_interval.
func MinusTimestampInterval(ts int64, interv *Interval) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_timestamp_interval(C.Timestamp(ts), interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// MinusTimestampTimestamp wraps MEOS C function minus_timestamp_timestamp.
func MinusTimestampTimestamp(ts1 int64, ts2 int64) (_r0 *Interval, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_timestamp_timestamp(C.Timestamp(ts1), C.Timestamp(ts2))
	if _err = meosError(); _err != nil {
		return
	}
	return &Interval{_inner: _cret}, nil
}


// MinusTimestamptzInterval wraps MEOS C function minus_timestamptz_interval.
func MinusTimestamptzInterval(tstz int64, interv *Interval) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_timestamptz_interval(C.TimestampTz(tstz), interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// MinusTimestamptzIntervalAtZone wraps MEOS C function minus_timestamptz_interval_at_zone.
func MinusTimestamptzIntervalAtZone(tstz int64, interv *Interval, zone string) (_r0 int64, _err error) {
	_c_zone := C.cstring_to_text(C.CString(zone))
	defer C.free(unsafe.Pointer(_c_zone))
	C.meos_errno_reset()
	_cret := C.minus_timestamptz_interval_at_zone(C.TimestampTz(tstz), interv._inner, _c_zone)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// MinusTimestamptzTimestamptz wraps MEOS C function minus_timestamptz_timestamptz.
func MinusTimestamptzTimestamptz(tstz1 int64, tstz2 int64) (_r0 *Interval, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_timestamptz_timestamptz(C.TimestampTz(tstz1), C.TimestampTz(tstz2))
	if _err = meosError(); _err != nil {
		return
	}
	return &Interval{_inner: _cret}, nil
}


// NeTimestampTimestamp wraps MEOS C function ne_timestamp_timestamp.
func NeTimestampTimestamp(ts1 int64, ts2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_timestamp_timestamp(C.Timestamp(ts1), C.Timestamp(ts2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NeTimestampTimestamptz wraps MEOS C function ne_timestamp_timestamptz.
func NeTimestampTimestamptz(ts int64, tstz int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_timestamp_timestamptz(C.Timestamp(ts), C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NeTimestamptzTimestamp wraps MEOS C function ne_timestamptz_timestamp.
func NeTimestamptzTimestamp(tstz int64, ts int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_timestamptz_timestamp(C.TimestampTz(tstz), C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimestampAge wraps MEOS C function timestamp_age.
func TimestampAge(ts1 int64, ts2 int64) (_r0 *Interval, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamp_age(C.Timestamp(ts1), C.Timestamp(ts2))
	if _err = meosError(); _err != nil {
		return
	}
	return &Interval{_inner: _cret}, nil
}


// TimestampAtLocal wraps MEOS C function timestamp_at_local.
func TimestampAtLocal(ts int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamp_at_local(C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimestampBin wraps MEOS C function timestamp_bin.
func TimestampBin(ts int64, stride *Interval, origin int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamp_bin(C.Timestamp(ts), stride._inner, C.Timestamp(origin))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TODO timestamp_extract: unsupported return type Numeric
// func TimestampExtract(...) { /* not yet handled by codegen */ }


// TimestampHash wraps MEOS C function timestamp_hash.
func TimestampHash(ts int64) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamp_hash(C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// TimestampHashExtended wraps MEOS C function timestamp_hash_extended.
func TimestampHashExtended(tstz int64, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamp_hash_extended(C.TimestampTz(tstz), C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// TimestampIsFinite wraps MEOS C function timestamp_is_finite.
func TimestampIsFinite(ts int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamp_is_finite(C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimestampIzone wraps MEOS C function timestamp_izone.
func TimestampIzone(ts int64, zone *Interval) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamp_izone(C.Timestamp(ts), zone._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimeOfDay wraps MEOS C function time_of_day.
func TimeOfDay() (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.time_of_day()
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TimestampLarger wraps MEOS C function timestamp_larger.
func TimestampLarger(ts1 int64, ts2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamp_larger(C.Timestamp(ts1), C.Timestamp(ts2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TODO timestamp_make: unsupported param float8
// func TimestampMake(...) { /* not yet handled by codegen */ }


// TimestampOverlaps wraps MEOS C function timestamp_overlaps.
func TimestampOverlaps(ts1 int64, te1 int64, ts2 int64, te2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamp_overlaps(C.Timestamp(ts1), C.Timestamp(te1), C.Timestamp(ts2), C.Timestamp(te2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TODO timestamp_part: unsupported return type float8
// func TimestampPart(...) { /* not yet handled by codegen */ }


// TimestampScale wraps MEOS C function timestamp_scale.
func TimestampScale(ts int64, typmod int32) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamp_scale(C.Timestamp(ts), C.int32(typmod))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimestampSmaller wraps MEOS C function timestamp_smaller.
func TimestampSmaller(ts1 int64, ts2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamp_smaller(C.Timestamp(ts1), C.Timestamp(ts2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimestampToTimestamptz wraps MEOS C function timestamp_to_timestamptz.
func TimestampToTimestamptz(ts int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamp_to_timestamptz(C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimestampTrunc wraps MEOS C function timestamp_trunc.
func TimestampTrunc(ts int64, units string) (_r0 int64, _err error) {
	_c_units := C.cstring_to_text(C.CString(units))
	defer C.free(unsafe.Pointer(_c_units))
	C.meos_errno_reset()
	_cret := C.timestamp_trunc(C.Timestamp(ts), _c_units)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimestampZone wraps MEOS C function timestamp_zone.
func TimestampZone(ts int64, zone string) (_r0 int64, _err error) {
	_c_zone := C.cstring_to_text(C.CString(zone))
	defer C.free(unsafe.Pointer(_c_zone))
	C.meos_errno_reset()
	_cret := C.timestamp_zone(C.Timestamp(ts), _c_zone)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimestamptzAge wraps MEOS C function timestamptz_age.
func TimestamptzAge(tstz1 int64, tstz2 int64) (_r0 *Interval, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_age(C.TimestampTz(tstz1), C.TimestampTz(tstz2))
	if _err = meosError(); _err != nil {
		return
	}
	return &Interval{_inner: _cret}, nil
}


// TimestamptzAtLocal wraps MEOS C function timestamptz_at_local.
func TimestamptzAtLocal(tstz int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_at_local(C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimestamptzBin wraps MEOS C function timestamptz_bin.
func TimestamptzBin(tstz int64, stride *Interval, origin int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_bin(C.TimestampTz(tstz), stride._inner, C.TimestampTz(origin))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TODO timestamptz_extract: unsupported return type Numeric
// func TimestamptzExtract(...) { /* not yet handled by codegen */ }


// TimestamptzHash wraps MEOS C function timestamptz_hash.
func TimestamptzHash(tstz int64) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_hash(C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// TimestamptzHashExtended wraps MEOS C function timestamptz_hash_extended.
func TimestamptzHashExtended(tstz int64, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_hash_extended(C.TimestampTz(tstz), C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// TimestamptzIzone wraps MEOS C function timestamptz_izone.
func TimestamptzIzone(tstz int64, zone *Interval) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_izone(C.TimestampTz(tstz), zone._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TODO timestamptz_make: unsupported param float8
// func TimestamptzMake(...) { /* not yet handled by codegen */ }


// TODO timestamptz_make_at_timezone: unsupported param float8
// func TimestamptzMakeAtTimezone(...) { /* not yet handled by codegen */ }


// TimestamptzOverlaps wraps MEOS C function timestamptz_overlaps.
func TimestamptzOverlaps(ts1 int64, te1 int64, ts2 int64, te2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_overlaps(C.TimestampTz(ts1), C.TimestampTz(te1), C.TimestampTz(ts2), C.TimestampTz(te2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TODO timestamptz_part: unsupported return type float8
// func TimestamptzPart(...) { /* not yet handled by codegen */ }


// TimestamptzScale wraps MEOS C function timestamptz_scale.
func TimestamptzScale(tstz int64, typmod int32) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_scale(C.TimestampTz(tstz), C.int32(typmod))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimestamptzShift wraps MEOS C function timestamptz_shift.
func TimestamptzShift(tstz int64, interv *Interval) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_shift(C.TimestampTz(tstz), interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimestamptzToTimestamp wraps MEOS C function timestamptz_to_timestamp.
func TimestamptzToTimestamp(tstz int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_to_timestamp(C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimestamptzTrunc wraps MEOS C function timestamptz_trunc.
func TimestamptzTrunc(tstz int64, units string) (_r0 int64, _err error) {
	_c_units := C.cstring_to_text(C.CString(units))
	defer C.free(unsafe.Pointer(_c_units))
	C.meos_errno_reset()
	_cret := C.timestamptz_trunc(C.TimestampTz(tstz), _c_units)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimestamptzTruncZone wraps MEOS C function timestamptz_trunc_zone.
func TimestamptzTruncZone(tstz int64, units string, zone string) (_r0 int64, _err error) {
	_c_units := C.cstring_to_text(C.CString(units))
	defer C.free(unsafe.Pointer(_c_units))
	_c_zone := C.cstring_to_text(C.CString(zone))
	defer C.free(unsafe.Pointer(_c_zone))
	C.meos_errno_reset()
	_cret := C.timestamptz_trunc_zone(C.TimestampTz(tstz), _c_units, _c_zone)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimestamptzZone wraps MEOS C function timestamptz_zone.
func TimestamptzZone(tstz int64, zone string) (_r0 int64, _err error) {
	_c_zone := C.cstring_to_text(C.CString(zone))
	defer C.free(unsafe.Pointer(_c_zone))
	C.meos_errno_reset()
	_cret := C.timestamptz_zone(C.TimestampTz(tstz), _c_zone)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}

