package gomeos

/*
#include "meos.h"
#include "meos_catalog.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"
import (
	"time"
	"unsafe"

	"github.com/leekchan/timeutil"
)

// ------------------------- Input ---------------------------
// func TemporalFromWKB[T Temporal](input []byte, output T) T {
// 	c_wkb := (*C.uint8_t)(unsafe.Pointer(&input[0]))
// 	c_size := C.size_t(len(input))
// 	c_temporal := C.temporal_from_wkb(c_wkb, c_size)
// 	output.Init(c_temporal)
// 	return output
// }

// func TemporalFromHexWKB[T Temporal](input string, output T) T {
// 	c_hexwkb := C.CString(input)
// 	defer C.free(unsafe.Pointer(c_hexwkb))
// 	c_temporal := C.temporal_from_hexwkb(c_hexwkb)
// 	output.Init(c_temporal)
// 	return output
// }

// ------------------------- Input ----------------------------------
func TemporalFromHexWKB[T Temporal](s string, output T) T {
	c_temp := C.temporal_from_hexwkb(C.CString(s))
	output.Init(c_temp)
	return output
}

// ------------------------- Output ---------------------------
func TemporalAsMFJSON[T Temporal](temp T, with_bbox bool, flags int, precision int, srs string) string {
	c_srs := C.CString(srs)
	defer C.free(unsafe.Pointer(c_srs))
	c_mfjson := C.temporal_as_mfjson(temp.Inner(), C.bool(with_bbox), C.int(flags), C.int(precision), C.CString(srs))
	defer C.free(unsafe.Pointer(c_mfjson))
	mfjson_out := C.GoString(c_mfjson)
	return mfjson_out
}

// ------------------------- Accessors -------------------------------------
func TemporalToTsTzSpan[T Temporal](temp T) *TsTzSpan {
	return &TsTzSpan{_inner: C.temporal_to_tstzspan(temp.Inner())}
}

func TemporalInterpolation[T Temporal](temp T) string {
	return C.GoString(C.temporal_interp(temp.Inner()))
}

func TemporalTime[T Temporal](temp T) *TsTzSpanSet {
	return &TsTzSpanSet{_inner: C.temporal_time(temp.Inner())}
}

func TemporalDuration[T Temporal](temp T, ignore_gaps bool) timeutil.Timedelta {
	return IntervalToTimeDelta(*C.temporal_duration(temp.Inner(), C.bool(ignore_gaps)))
}

func TemporalNumInstants[T Temporal](temp T) int {
	return int(C.temporal_num_instants(temp.Inner()))
}

func TemporalInstantN[T Temporal, TI TInstant](temp T, inst TI, n int) TI {
	c_inst := C.temporal_instant_n(temp.Inner(), C.int(n)+1)
	c_temp := C.cast_tinstant_to_temporal(c_inst)
	inst.Init(c_temp)
	return inst
}

func TemporalInstants[T Temporal, TI TInstant](temp T, insts []TI) []TI {
	nums := TemporalNumInstants(temp)
	output_insts := make([]TI, nums)
	for i := 0; i < nums; i++ {
		output_insts[i] = TemporalInstantN(temp, insts[i], i)
	}
	return output_insts
}

func TemporalStartInstant[T Temporal, TI TInstant](temp T, inst TI) TI {
	c_inst := C.temporal_start_instant(temp.Inner())
	c_temp := C.cast_tinstant_to_temporal(c_inst)
	inst.Init(c_temp)
	return inst
}

func TemporalEndInstant[T Temporal, TI TInstant](temp T, inst TI) TI {
	c_inst := C.temporal_end_instant(temp.Inner())
	c_temp := C.cast_tinstant_to_temporal(c_inst)
	inst.Init(c_temp)
	return inst
}

func TemporalMinInstant[T Temporal, TI TInstant](temp T, inst TI) TI {
	c_inst := C.temporal_min_instant(temp.Inner())
	c_temp := C.cast_tinstant_to_temporal(c_inst)
	inst.Init(c_temp)
	return inst
}

func TemporalMaxInstant[T Temporal, TI TInstant](temp T, inst TI) TI {
	c_inst := C.temporal_max_instant(temp.Inner())
	c_temp := C.cast_tinstant_to_temporal(c_inst)
	inst.Init(c_temp)
	return inst
}

func TemporalNumTimestamps[T Temporal](temp T) int {
	return int(C.temporal_num_timestamps(temp.Inner()))
}

func TemporalStartTimestamptz[T Temporal](temp T) time.Time {
	return TimestamptzToDatetime(C.temporal_start_timestamptz(temp.Inner()))
}

func TemporalEndTimestamptz[T Temporal](temp T) time.Time {
	return TimestamptzToDatetime(C.temporal_end_timestamptz(temp.Inner()))
}

func TemporalTimestampN[T Temporal](temp T, n int) time.Time {
	res := C.malloc(C.sizeof_int)
	defer C.free(unsafe.Pointer(res)) // Ensure memory is freed.
	success := C.temporal_timestamptz_n(temp.Inner(), C.int(n+1), (*C.TimestampTz)(res))
	if success {
		result := *(*C.TimestampTz)(res)
		return TimestamptzToDatetime(result)
	} else {
		return time.Time{}
	}
}

func TemporalHash[T Temporal](temp T) int {
	return int(C.temporal_hash(temp.Inner()))
}

func TemporalTimestamps[T Temporal](temp T) []time.Time {
	var times []time.Time
	nums := TemporalNumTimestamps(temp)
	for i := 0; i < nums; i++ {
		times = append(times, TemporalTimestampN(temp, i))
	}
	return times
}

//TODO: TemporalSegments. How many segments does a temporal have?

// ------------------------- TODO:Transformations -------------------------------
// ------------------------- TODO:Modifications ---------------------------------
// ------------------------- Restrictions ----------------------------------
func TemporalAtTimestamptz[T Temporal](temp T, ts time.Time) Temporal {
	c_temp := C.temporal_at_timestamptz(temp.Inner(), DatetimeToTimestamptz(ts))
	res := CreateTemporal(c_temp)
	return res
}

func TemporalAtTsTzSet[T Temporal](temp T, tstzset TsTzSet) Temporal {
	c_temp := C.temporal_at_tstzset(temp.Inner(), tstzset._inner)
	res := CreateTemporal(c_temp)
	return res
}

func TemporalAtTsTzSpan[T Temporal](temp T, tstzspan TsTzSpan) Temporal {
	c_temp := C.temporal_at_tstzspan(temp.Inner(), tstzspan._inner)
	res := CreateTemporal(c_temp)
	return res
}

func TemporalAtTsTzSpanSet[T Temporal](temp T, tstzspanset TsTzSpanSet) Temporal {
	c_temp := C.temporal_at_tstzspanset(temp.Inner(), tstzspanset._inner)
	res := CreateTemporal(c_temp)
	return res
}

func TemporalAtMax[T Temporal](temp T) Temporal {
	c_temp := C.temporal_at_max(temp.Inner())
	res := CreateTemporal(c_temp)
	return res
}

func TemporalAtMin[T Temporal](temp T) Temporal {
	c_temp := C.temporal_at_min(temp.Inner())
	res := CreateTemporal(c_temp)
	return res
}

func TemporalAtValues[T Temporal, S Set](temp T, set S) Temporal {
	c_temp := C.temporal_at_values(temp.Inner(), set.Inner())
	res := CreateTemporal(c_temp)
	return res
}

func TemporalMinusTimestamptz[T Temporal](temp T, ts time.Time) Temporal {
	c_temp := C.temporal_minus_timestamptz(temp.Inner(), DatetimeToTimestamptz(ts))
	res := CreateTemporal(c_temp)
	return res
}

func TemporalMinusTsTzSet[T Temporal](temp T, tstzset TsTzSet) Temporal {
	c_temp := C.temporal_minus_tstzset(temp.Inner(), tstzset._inner)
	res := CreateTemporal(c_temp)
	return res
}

func TemporalMinusTsTzSpan[T Temporal](temp T, tstzspan TsTzSpan) Temporal {
	c_temp := C.temporal_minus_tstzspan(temp.Inner(), tstzspan._inner)
	res := CreateTemporal(c_temp)
	return res
}

func TemporalMinusTsTzSpanSet[T Temporal](temp T, tstzspanset TsTzSpanSet) Temporal {
	c_temp := C.temporal_minus_tstzspanset(temp.Inner(), tstzspanset._inner)
	res := CreateTemporal(c_temp)
	return res
}

func TemporalMinusMax[T Temporal](temp T) Temporal {
	c_temp := C.temporal_at_max(temp.Inner())
	res := CreateTemporal(c_temp)
	return res
}

func TemporalMinusMin[T Temporal](temp T) Temporal {
	c_temp := C.temporal_minus_min(temp.Inner())
	res := CreateTemporal(c_temp)
	return res
}

// ------------------------- TODO:Topological Operations ------------------------
// ------------------------- TODO:Position Operations ---------------------------
// ------------------------- Similarity Operations -------------------------
func TemporalFrechetDistance[T Temporal](temp1 T, temp2 T) float64 {
	return float64(C.temporal_frechet_distance(temp1.Inner(), temp2.Inner()))
}

func TemporalDyntimewarpDistance[T Temporal](temp1 T, temp2 T) float64 {
	return float64(C.temporal_dyntimewarp_distance(temp1.Inner(), temp2.Inner()))
}

func TemporalHausdorffDistance[T Temporal](temp1 T, temp2 T) float64 {
	return float64(C.temporal_hausdorff_distance(temp1.Inner(), temp2.Inner()))
}

// ------------------------- TODO:Split Operations ------------------------------

// func TemporalDuration[T Temporal](temp T, boundspan bool) timeutil.Timedelta {
// 	return IntervalToTimeDelta(*C.temporal_duration(temp.Inner(), C.bool(boundspan)))
// }

func TPointLength[T Temporal](temp T) float64 {
	return float64(C.tpoint_length(temp.Inner()))
}

func TPointSpeed[T Temporal, F TFloat](temp T, new_temp F) F {
	c_temp := C.tpoint_speed(temp.Inner())
	new_temp.Init(c_temp)
	return new_temp
}

func TpointAtGeomTime[T Temporal](temp T, new_temp T, geom *Geom) T {
	c_temp := C.tpoint_at_geom_time(temp.Inner(), geom._inner, nil, nil)
	new_temp.Init(c_temp)
	return new_temp
}

func TpointMinusGeomTime[T Temporal](temp T, new_temp T, geom *Geom) T {
	c_temp := C.tpoint_minus_geom_time(temp.Inner(), geom._inner, nil, nil)
	new_temp.Init(c_temp)
	return new_temp
}

func TemporalSimplifyDP[T Temporal](temp T, new_temp T, dist float64, syncdist bool) T {
	c_temp := C.temporal_simplify_dp(temp.Inner(), C.double(dist), C.bool(syncdist))
	new_temp.Init(c_temp)
	return new_temp
}

// func CreateTemporalFunc(*C.Temporal) func() {

// }
