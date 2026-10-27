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

// DateTimetzToTimestamptz wraps MEOS C function date_timetz_to_timestamptz.
func DateTimetzToTimestamptz(date int32, timetz unsafe.Pointer) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.date_timetz_to_timestamptz(C.DateADT(date), (*C.TimeTzADT)(unsafe.Pointer(timetz)))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// IntervalToTime wraps MEOS C function interval_to_time.
func IntervalToTime(interv *Interval) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.interval_to_time(interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// MinusTimeInterval wraps MEOS C function minus_time_interval.
func MinusTimeInterval(time int64, interv *Interval) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_time_interval(C.TimeADT(time), interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// MinusTimeTime wraps MEOS C function minus_time_time.
func MinusTimeTime(time1 int64, time2 int64) (_r0 *Interval, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_time_time(C.TimeADT(time1), C.TimeADT(time2))
	if _err = meosError(); _err != nil {
		return
	}
	return &Interval{_inner: _cret}, nil
}


// MinusTimetzInterval wraps MEOS C function minus_timetz_interval.
func MinusTimetzInterval(timetz unsafe.Pointer, interv *Interval) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_timetz_interval((*C.TimeTzADT)(unsafe.Pointer(timetz)), interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PlusTimeInterval wraps MEOS C function plus_time_interval.
func PlusTimeInterval(time int64, interv *Interval) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.plus_time_interval(C.TimeADT(time), interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// PlusTimetzInterval wraps MEOS C function plus_timetz_interval.
func PlusTimetzInterval(timetz unsafe.Pointer, interv *Interval) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.plus_timetz_interval((*C.TimeTzADT)(unsafe.Pointer(timetz)), interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TimeCmp wraps MEOS C function time_cmp.
func TimeCmp(time1 int64, time2 int64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.time_cmp(C.TimeADT(time1), C.TimeADT(time2))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TimeEq wraps MEOS C function time_eq.
func TimeEq(time1 int64, time2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.time_eq(C.TimeADT(time1), C.TimeADT(time2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TODO time_extract: unsupported return type Numeric
// func TimeExtract(...) { /* not yet handled by codegen */ }


// TimeGe wraps MEOS C function time_ge.
func TimeGe(time1 int64, time2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.time_ge(C.TimeADT(time1), C.TimeADT(time2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimeGt wraps MEOS C function time_gt.
func TimeGt(time1 int64, time2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.time_gt(C.TimeADT(time1), C.TimeADT(time2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimeHash wraps MEOS C function time_hash.
func TimeHash(time int64) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.time_hash(C.TimeADT(time))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// TimeHashExtended wraps MEOS C function time_hash_extended.
func TimeHashExtended(time int64, seed int32) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.time_hash_extended(C.TimeADT(time), C.int32(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// TimeLarger wraps MEOS C function time_larger.
func TimeLarger(time1 int64, time2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.time_larger(C.TimeADT(time1), C.TimeADT(time2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimeLe wraps MEOS C function time_le.
func TimeLe(time1 int64, time2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.time_le(C.TimeADT(time1), C.TimeADT(time2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimeLt wraps MEOS C function time_lt.
func TimeLt(time1 int64, time2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.time_lt(C.TimeADT(time1), C.TimeADT(time2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimeMake wraps MEOS C function time_make.
func TimeMake(tm_hour int, tm_min int, sec float64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.time_make(C.int(tm_hour), C.int(tm_min), C.double(sec))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimeNe wraps MEOS C function time_ne.
func TimeNe(time1 int64, time2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.time_ne(C.TimeADT(time1), C.TimeADT(time2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimeOverlaps wraps MEOS C function time_overlaps.
func TimeOverlaps(ts1 int64, te1 int64, ts2 int64, te2 int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.time_overlaps(C.TimeADT(ts1), C.TimeADT(te1), C.TimeADT(ts2), C.TimeADT(te2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TODO time_part: unsupported return type float8
// func TimePart(...) { /* not yet handled by codegen */ }


// TimeScale wraps MEOS C function time_scale.
func TimeScale(date int64, typmod int32) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.time_scale(C.TimeADT(date), C.int32(typmod))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimeSmaller wraps MEOS C function time_smaller.
func TimeSmaller(time1 int64, time2 int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.time_smaller(C.TimeADT(time1), C.TimeADT(time2))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimeToInterval wraps MEOS C function time_to_interval.
func TimeToInterval(time int64) (_r0 *Interval, _err error) {
	C.meos_errno_reset()
	_cret := C.time_to_interval(C.TimeADT(time))
	if _err = meosError(); _err != nil {
		return
	}
	return &Interval{_inner: _cret}, nil
}


// TimeToTimetz wraps MEOS C function time_to_timetz.
func TimeToTimetz(time int64) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.time_to_timetz(C.TimeADT(time))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TimestampToTime wraps MEOS C function timestamp_to_time.
func TimestampToTime(ts int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamp_to_time(C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimestamptzToTime wraps MEOS C function timestamptz_to_time.
func TimestamptzToTime(tztz int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_to_time(C.TimestampTz(tztz))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimestamptzToTimetz wraps MEOS C function timestamptz_to_timetz.
func TimestamptzToTimetz(tztz int64) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_to_timetz(C.TimestampTz(tztz))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TimetzAtLocal wraps MEOS C function timetz_at_local.
func TimetzAtLocal(timetz unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_at_local((*C.TimeTzADT)(unsafe.Pointer(timetz)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TimetzCmp wraps MEOS C function timetz_cmp.
func TimetzCmp(timetz1 unsafe.Pointer, timetz2 unsafe.Pointer) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_cmp((*C.TimeTzADT)(unsafe.Pointer(timetz1)), (*C.TimeTzADT)(unsafe.Pointer(timetz2)))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// TimetzEq wraps MEOS C function timetz_eq.
func TimetzEq(timetz1 unsafe.Pointer, timetz2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_eq((*C.TimeTzADT)(unsafe.Pointer(timetz1)), (*C.TimeTzADT)(unsafe.Pointer(timetz2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TODO timetz_extract: unsupported return type Numeric
// func TimetzExtract(...) { /* not yet handled by codegen */ }


// TimetzGe wraps MEOS C function timetz_ge.
func TimetzGe(timetz1 unsafe.Pointer, timetz2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_ge((*C.TimeTzADT)(unsafe.Pointer(timetz1)), (*C.TimeTzADT)(unsafe.Pointer(timetz2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimetzGt wraps MEOS C function timetz_gt.
func TimetzGt(timetz1 unsafe.Pointer, timetz2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_gt((*C.TimeTzADT)(unsafe.Pointer(timetz1)), (*C.TimeTzADT)(unsafe.Pointer(timetz2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimetzHash wraps MEOS C function timetz_hash.
func TimetzHash(timetz unsafe.Pointer) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_hash((*C.TimeTzADT)(unsafe.Pointer(timetz)))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// TimetzHashExtended wraps MEOS C function timetz_hash_extended.
func TimetzHashExtended(timetz unsafe.Pointer, seed int64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_hash_extended((*C.TimeTzADT)(unsafe.Pointer(timetz)), C.int64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// TimetzIn wraps MEOS C function timetz_in.
func TimetzIn(str string, typmod int32) (_r0 unsafe.Pointer, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.timetz_in(_c_str, C.int32(typmod))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TimetzIzone wraps MEOS C function timetz_izone.
func TimetzIzone(timetz unsafe.Pointer, zone *Interval) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_izone((*C.TimeTzADT)(unsafe.Pointer(timetz)), zone._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TimetzLarger wraps MEOS C function timetz_larger.
func TimetzLarger(timetz1 unsafe.Pointer, timetz2 unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_larger((*C.TimeTzADT)(unsafe.Pointer(timetz1)), (*C.TimeTzADT)(unsafe.Pointer(timetz2)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TimetzLe wraps MEOS C function timetz_le.
func TimetzLe(timetz1 unsafe.Pointer, timetz2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_le((*C.TimeTzADT)(unsafe.Pointer(timetz1)), (*C.TimeTzADT)(unsafe.Pointer(timetz2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimetzLt wraps MEOS C function timetz_lt.
func TimetzLt(timetz1 unsafe.Pointer, timetz2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_lt((*C.TimeTzADT)(unsafe.Pointer(timetz1)), (*C.TimeTzADT)(unsafe.Pointer(timetz2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimetzNe wraps MEOS C function timetz_ne.
func TimetzNe(timetz1 unsafe.Pointer, timetz2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_ne((*C.TimeTzADT)(unsafe.Pointer(timetz1)), (*C.TimeTzADT)(unsafe.Pointer(timetz2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimetzCopy wraps MEOS C function timetz_copy.
func TimetzCopy(timetz unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_copy((*C.TimeTzADT)(unsafe.Pointer(timetz)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TimetzOut wraps MEOS C function timetz_out.
func TimetzOut(timetz unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_out((*C.TimeTzADT)(unsafe.Pointer(timetz)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TimetzOverlaps wraps MEOS C function timetz_overlaps.
func TimetzOverlaps(ts1 unsafe.Pointer, te1 unsafe.Pointer, ts2 unsafe.Pointer, te2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_overlaps((*C.TimeTzADT)(unsafe.Pointer(ts1)), (*C.TimeTzADT)(unsafe.Pointer(te1)), (*C.TimeTzADT)(unsafe.Pointer(ts2)), (*C.TimeTzADT)(unsafe.Pointer(te2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TODO timetz_part: unsupported return type float8
// func TimetzPart(...) { /* not yet handled by codegen */ }


// TimetzScale wraps MEOS C function timetz_scale.
func TimetzScale(timetz unsafe.Pointer, typmod int32) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_scale((*C.TimeTzADT)(unsafe.Pointer(timetz)), C.int32(typmod))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TimetzSmaller wraps MEOS C function timetz_smaller.
func TimetzSmaller(timetz1 unsafe.Pointer, timetz2 unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_smaller((*C.TimeTzADT)(unsafe.Pointer(timetz1)), (*C.TimeTzADT)(unsafe.Pointer(timetz2)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TimetzToTime wraps MEOS C function timetz_to_time.
func TimetzToTime(timetz unsafe.Pointer) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timetz_to_time((*C.TimeTzADT)(unsafe.Pointer(timetz)))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimetzZone wraps MEOS C function timetz_zone.
func TimetzZone(timetz unsafe.Pointer, zone string) (_r0 unsafe.Pointer, _err error) {
	_c_zone := C.cstring_to_text(C.CString(zone))
	defer C.free(unsafe.Pointer(_c_zone))
	C.meos_errno_reset()
	_cret := C.timetz_zone((*C.TimeTzADT)(unsafe.Pointer(timetz)), _c_zone)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}

