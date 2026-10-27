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

// AddDateInt wraps MEOS C function add_date_int.
func AddDateInt(date int32, days int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.add_date_int(C.DateADT(date), C.int32(days))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// AddDateInterval wraps MEOS C function add_date_interval.
func AddDateInterval(date int32, interv *Interval) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.add_date_interval(C.DateADT(date), interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// CmpDateTimestamp wraps MEOS C function cmp_date_timestamp.
func CmpDateTimestamp(date int32, ts int64) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.cmp_date_timestamp(C.DateADT(date), C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// CmpDateDate wraps MEOS C function cmp_date_date.
func CmpDateDate(date1 int32, date2 int32) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.cmp_date_date(C.DateADT(date1), C.DateADT(date2))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// CmpDateTimestamptz wraps MEOS C function cmp_date_timestamptz.
func CmpDateTimestamptz(date int32, tstz int64) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.cmp_date_timestamptz(C.DateADT(date), C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// CmpTimestampDate wraps MEOS C function cmp_timestamp_date.
func CmpTimestampDate(ts int64, date int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.cmp_timestamp_date(C.Timestamp(ts), C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// CmpTimestamptzDate wraps MEOS C function cmp_timestamptz_date.
func CmpTimestamptzDate(tstz int64, date int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.cmp_timestamptz_date(C.TimestampTz(tstz), C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// TODO date_extract: unsupported return type Numeric
// func DateExtract(...) { /* not yet handled by codegen */ }


// DateHash wraps MEOS C function date_hash.
func DateHash(date int32) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.date_hash(C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// DateHashExtended wraps MEOS C function date_hash_extended.
func DateHashExtended(date int32, seed int64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.date_hash_extended(C.DateADT(date), C.int64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// DateIsFinite wraps MEOS C function date_is_finite.
func DateIsFinite(date int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.date_is_finite(C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// DateLarger wraps MEOS C function date_larger.
func DateLarger(date1 int32, date2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.date_larger(C.DateADT(date1), C.DateADT(date2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// DateMake wraps MEOS C function date_make.
func DateMake(year int, mon int, mday int) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.date_make(C.int(year), C.int(mon), C.int(mday))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// DateSmaller wraps MEOS C function date_smaller.
func DateSmaller(date1 int32, date2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.date_smaller(C.DateADT(date1), C.DateADT(date2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// DateTimeToTimestamp wraps MEOS C function date_time_to_timestamp.
func DateTimeToTimestamp(date int32, time int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.date_time_to_timestamp(C.DateADT(date), C.TimeADT(time))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// DateToTimestamp wraps MEOS C function date_to_timestamp.
func DateToTimestamp(date int32) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.date_to_timestamp(C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// DateToTimestamptz wraps MEOS C function date_to_timestamptz.
func DateToTimestamptz(date int32) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.date_to_timestamptz(C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// EqDateDate wraps MEOS C function eq_date_date.
func EqDateDate(date1 int32, date2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_date_date(C.DateADT(date1), C.DateADT(date2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqDateTimestamp wraps MEOS C function eq_date_timestamp.
func EqDateTimestamp(date int32, ts int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_date_timestamp(C.DateADT(date), C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqDateTimestamptz wraps MEOS C function eq_date_timestamptz.
func EqDateTimestamptz(date int32, tstz int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_date_timestamptz(C.DateADT(date), C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqTimestampDate wraps MEOS C function eq_timestamp_date.
func EqTimestampDate(ts int64, date int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_timestamp_date(C.Timestamp(ts), C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EqTimestamptzDate wraps MEOS C function eq_timestamptz_date.
func EqTimestamptzDate(tstz int64, date int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eq_timestamptz_date(C.TimestampTz(tstz), C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeDateDate wraps MEOS C function ge_date_date.
func GeDateDate(date1 int32, date2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_date_date(C.DateADT(date1), C.DateADT(date2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeDateTimestamp wraps MEOS C function ge_date_timestamp.
func GeDateTimestamp(date int32, ts int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_date_timestamp(C.DateADT(date), C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeDateTimestamptz wraps MEOS C function ge_date_timestamptz.
func GeDateTimestamptz(date int32, tstz int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_date_timestamptz(C.DateADT(date), C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeTimestampDate wraps MEOS C function ge_timestamp_date.
func GeTimestampDate(ts int64, date int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_timestamp_date(C.Timestamp(ts), C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeTimestamptzDate wraps MEOS C function ge_timestamptz_date.
func GeTimestamptzDate(tstz int64, date int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ge_timestamptz_date(C.TimestampTz(tstz), C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtDateDate wraps MEOS C function gt_date_date.
func GtDateDate(date1 int32, date2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_date_date(C.DateADT(date1), C.DateADT(date2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtDateTimestamp wraps MEOS C function gt_date_timestamp.
func GtDateTimestamp(date int32, ts int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_date_timestamp(C.DateADT(date), C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtDateTimestamptz wraps MEOS C function gt_date_timestamptz.
func GtDateTimestamptz(date int32, tstz int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_date_timestamptz(C.DateADT(date), C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtTimestampDate wraps MEOS C function gt_timestamp_date.
func GtTimestampDate(ts int64, date int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_timestamp_date(C.Timestamp(ts), C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GtTimestamptzDate wraps MEOS C function gt_timestamptz_date.
func GtTimestamptzDate(tstz int64, date int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.gt_timestamptz_date(C.TimestampTz(tstz), C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeDateDate wraps MEOS C function le_date_date.
func LeDateDate(date1 int32, date2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_date_date(C.DateADT(date1), C.DateADT(date2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeDateTimestamp wraps MEOS C function le_date_timestamp.
func LeDateTimestamp(date int32, ts int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_date_timestamp(C.DateADT(date), C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeDateTimestamptz wraps MEOS C function le_date_timestamptz.
func LeDateTimestamptz(date int32, tstz int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_date_timestamptz(C.DateADT(date), C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeTimestampDate wraps MEOS C function le_timestamp_date.
func LeTimestampDate(ts int64, date int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_timestamp_date(C.Timestamp(ts), C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeTimestamptzDate wraps MEOS C function le_timestamptz_date.
func LeTimestamptzDate(tstz int64, date int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.le_timestamptz_date(C.TimestampTz(tstz), C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtDateDate wraps MEOS C function lt_date_date.
func LtDateDate(date1 int32, date2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_date_date(C.DateADT(date1), C.DateADT(date2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtDateTimestamp wraps MEOS C function lt_date_timestamp.
func LtDateTimestamp(date int32, ts int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_date_timestamp(C.DateADT(date), C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtDateTimestamptz wraps MEOS C function lt_date_timestamptz.
func LtDateTimestamptz(date int32, tstz int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_date_timestamptz(C.DateADT(date), C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtTimestampDate wraps MEOS C function lt_timestamp_date.
func LtTimestampDate(ts int64, date int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_timestamp_date(C.Timestamp(ts), C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LtTimestamptzDate wraps MEOS C function lt_timestamptz_date.
func LtTimestamptzDate(tstz int64, date int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lt_timestamptz_date(C.TimestampTz(tstz), C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// MinusDateDate wraps MEOS C function minus_date_date.
func MinusDateDate(date1 int32, date2 int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_date_date(C.DateADT(date1), C.DateADT(date2))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// MinusDateInt wraps MEOS C function minus_date_int.
func MinusDateInt(date int32, days int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_date_int(C.DateADT(date), C.int32(days))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// MinusDateInterval wraps MEOS C function minus_date_interval.
func MinusDateInterval(date int32, span *Interval) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_date_interval(C.DateADT(date), span._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// NeDateDate wraps MEOS C function ne_date_date.
func NeDateDate(date1 int32, date2 int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_date_date(C.DateADT(date1), C.DateADT(date2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NeDateTimestamp wraps MEOS C function ne_date_timestamp.
func NeDateTimestamp(date int32, ts int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_date_timestamp(C.DateADT(date), C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NeDateTimestamptz wraps MEOS C function ne_date_timestamptz.
func NeDateTimestamptz(date int32, tstz int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_date_timestamptz(C.DateADT(date), C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NeTimestampDate wraps MEOS C function ne_timestamp_date.
func NeTimestampDate(ts int64, date int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_timestamp_date(C.Timestamp(ts), C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NeTimestamptzDate wraps MEOS C function ne_timestamptz_date.
func NeTimestamptzDate(tstz int64, date int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ne_timestamptz_date(C.TimestampTz(tstz), C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TimestampToDate wraps MEOS C function timestamp_to_date.
func TimestampToDate(ts int64) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamp_to_date(C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// TimestamptzToDate wraps MEOS C function timestamptz_to_date.
func TimestamptzToDate(tstz int64) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_to_date(C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}

