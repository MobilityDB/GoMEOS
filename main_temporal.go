package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"

	"github.com/leekchan/timeutil"
)

// TemporalFromWKB Return a temporal value from its Well-Known Binary (WKB) representation
func TemporalFromWKB(wkb []byte) Temporal {
	// Get the size of the WKB byte slice
	size := C.size_t(len(wkb))

	// Convert the Go byte slice to a C uint8_t* pointer
	wkbPtr := (*C.uint8_t)(unsafe.Pointer(&wkb[0]))

	// Call the C function
	temporal := C.temporal_from_wkb(wkbPtr, size)

	// Return the pointer to the Temporal object
	return CreateTemporal(temporal)
}

// TemporalFromHexWKB Returns a temporal object from a hex-encoded WKB string.
func TemporalFromHexWKB[T Temporal](s string, output T) T {
	c_temp := C.temporal_from_hexwkb(C.CString(s))
	output.Init(c_temp)
	return output
}

// TemporalAsMFJSON Returns the temporal object as a MF-JSON string.
func TemporalAsMFJSON[T Temporal](temp T, with_bbox bool, flags int, precision int, srs string) string {
	c_srs := C.CString(srs)
	defer C.free(unsafe.Pointer(c_srs))
	c_mfjson := C.temporal_as_mfjson(temp.Inner(), C.bool(with_bbox), C.int(flags), C.int(precision), C.CString(srs))
	defer C.free(unsafe.Pointer(c_mfjson))
	mfjson_out := C.GoString(c_mfjson)
	return mfjson_out
}

// TemporalAsWKB Returns the temporal object as a hex-encoded WKB string.
func TemporalAsWKB[T Temporal](temp T) ([]byte, error) {
	var sizeOut C.size_t
	wkbPtr := C.temporal_as_wkb(temp.Inner(), C.uint8_t(4), &sizeOut)
	if wkbPtr == nil {
		return nil, fmt.Errorf("failed to convert temporal to WKB")
	}
	defer C.free(unsafe.Pointer(wkbPtr)) // Free the memory allocated by C
	// Convert the C uint8_t* to a Go []byte slice
	length := int(sizeOut)
	wkb := C.GoBytes(unsafe.Pointer(wkbPtr), C.int(length))
	return wkb, nil
}

// TemporalAsHexWKB Return the hex-encoded ASCII Well-Known Binary (HexWKB) representation of a temporal
func TemporalAsHexWKB[T Temporal](temp T) (string, error) {
	var sizeOut C.size_t
	wkbPtr := C.temporal_as_hexwkb(temp.Inner(), C.uint8_t(4), &sizeOut)
	if wkbPtr == nil {
		return "", fmt.Errorf("failed to convert temporal to HEXWKB")
	}
	defer C.free(unsafe.Pointer(wkbPtr)) // Free the memory allocated by C
	wkb := C.GoString(wkbPtr)
	return wkb, nil
}

// TemporalToTsTzSpan Return the bounding period of a temporal value
func TemporalToTsTzSpan[T Temporal](temp T) *TsTzSpan {
	return &TsTzSpan{_inner: C.temporal_to_tstzspan(temp.Inner())}
}

// TemporalInterpolation Return the string representation of the interpolation of a temporal value
func TemporalInterpolation[T Temporal](temp T) string {
	return C.GoString(C.temporal_interp(temp.Inner()))
}

// TemporalTime Return the time frame of a temporal value as a span set
func TemporalTime[T Temporal](temp T) *TsTzSpanSet {
	return &TsTzSpanSet{_inner: C.temporal_time(temp.Inner())}
}

// TemporalDuration Return the duration of a temporal value
func TemporalDuration[T Temporal](temp T, ignore_gaps bool) timeutil.Timedelta {
	return IntervalToTimeDelta(*C.temporal_duration(temp.Inner(), C.bool(ignore_gaps)))
}

// TemporalNumInstants Returns the number of instants
func TemporalNumInstants[T Temporal](temp T) int {
	return int(C.temporal_num_instants(temp.Inner()))
}

// TemporalInstantN Return a copy of the n-th instant of a temporal value
func TemporalInstantN[T Temporal, TI TInstant](temp T, inst TI, n int) TI {
	c_inst := C.temporal_instant_n(temp.Inner(), C.int(n)+1)
	c_temp := C.cast_tinstant_to_temporal(c_inst)
	inst.Init(c_temp)
	return inst
}

// TemporalInstants Return a copy of the distinct instants of a temporal value
func TemporalInstants[T Temporal, TI TInstant](temp T, insts []TI) []TI {
	nums := TemporalNumInstants(temp)
	output_insts := make([]TI, nums)
	for i := 0; i < nums; i++ {
		output_insts[i] = TemporalInstantN(temp, insts[i], i)
	}
	return output_insts
}

// TemporalStartInstant Returns the first instant
func TemporalStartInstant[T Temporal, TI TInstant](temp T, inst TI) TI {
	c_inst := C.temporal_start_instant(temp.Inner())
	c_temp := C.cast_tinstant_to_temporal(c_inst)
	inst.Init(c_temp)
	return inst
}

// TemporalEndInstant Return a copy of the end instant of a temporal value
func TemporalEndInstant[T Temporal, TI TInstant](temp T, inst TI) TI {
	c_inst := C.temporal_end_instant(temp.Inner())
	c_temp := C.cast_tinstant_to_temporal(c_inst)
	inst.Init(c_temp)
	return inst
}

// TemporalMinInstant Return a copy of the instant with minimum base value of a temporal value
func TemporalMinInstant[T Temporal, TI TInstant](temp T, inst TI) TI {
	c_inst := C.temporal_min_instant(temp.Inner())
	c_temp := C.cast_tinstant_to_temporal(c_inst)
	inst.Init(c_temp)
	return inst
}

// TemporalMaxInstant Return a copy of the instant with maximum base value of a temporal value
func TemporalMaxInstant[T Temporal, TI TInstant](temp T, inst TI) TI {
	c_inst := C.temporal_max_instant(temp.Inner())
	c_temp := C.cast_tinstant_to_temporal(c_inst)
	inst.Init(c_temp)
	return inst
}

// TemporalNumTimestamps Return the number of distinct timestamps of a temporal value
func TemporalNumTimestamps[T Temporal](temp T) int {
	return int(C.temporal_num_timestamps(temp.Inner()))
}

// TemporalStartTimestamptz Return the start timestamp of a temporal value
func TemporalStartTimestamptz[T Temporal](temp T) time.Time {
	return TimestamptzToDatetime(C.temporal_start_timestamptz(temp.Inner()))
}

// TemporalEndTimestamptz Return the end timestamptz of a temporal value
func TemporalEndTimestamptz[T Temporal](temp T) time.Time {
	return TimestamptzToDatetime(C.temporal_end_timestamptz(temp.Inner()))
}

// TemporalTimestampN Return the n-th distinct timestamp of a temporal value in the last argument
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

// TemporalHash Return the 32-bit hash value of a temporal value
func TemporalHash[T Temporal](temp T) int {
	return int(C.temporal_hash(temp.Inner()))
}

// TemporalTimestamps Return the array of timestamps of a temporal value
func TemporalTimestamps[T Temporal](temp T) []time.Time {
	var times []time.Time
	nums := TemporalNumTimestamps(temp)
	for i := 0; i < nums; i++ {
		times = append(times, TemporalTimestampN(temp, i))
	}
	return times
}

// TemporalSetInterp Return a temporal value transformed to a given interpolation
func TemporalSetInterp[T Temporal](temp T, interp Interpolation, output T) T {
	new_temp := C.temporal_set_interp(temp.Inner(), C.interpType(interp))
	output.Init(new_temp)
	return output
}

// TemporalShiftTime Return a temporal value shifted by an interval
func TemporalShiftTime[T Temporal](temp T, delta timeutil.Timedelta, output T) T {
	interval := TimeDeltaToInterval(delta)
	shifted := C.temporal_shift_time(temp.Inner(), &interval)
	output.Init(shifted)
	return output
}

// TemporalScaleTime Return a temporal value scaled by an interval
func TemporalScaleTime[T Temporal](temp T, duration timeutil.Timedelta, output T) T {
	interval := TimeDeltaToInterval(duration)
	scaled := C.temporal_scale_time(temp.Inner(), &interval)
	output.Init(scaled)
	return output
}

// TemporalShiftScaleTime Return a temporal value shifted and/or scaled by two intervals
func TemporalShiftScaleTime[T Temporal](temp T, shift timeutil.Timedelta, duration timeutil.Timedelta, output T) T {
	s := TimeDeltaToInterval(shift)
	d := TimeDeltaToInterval(duration)
	scaled := C.temporal_shift_scale_time(temp.Inner(), &s, &d)
	output.Init(scaled)
	return output
}

// TemporalTSample Return a temporal value sampled according to period buckets
func TemporalTSample[T Temporal](temp T, duration timeutil.Timedelta, start time.Time, interp Interpolation) Temporal {
	s := DatetimeToTimestamptz(start)
	d := TimeDeltaToInterval(duration)
	res := C.temporal_tsample(temp.Inner(), &d, s, C.interpType(interp))
	return CreateTemporal(res)
}

// TemporalTPrecision Return a temporal value with the precision set to period buckets
func TemporalTPrecision[T Temporal](temp T, duration timeutil.Timedelta, start time.Time) Temporal {
	s := DatetimeToTimestamptz(start)
	d := TimeDeltaToInterval(duration)
	res := C.temporal_tprecision(temp.Inner(), &d, s)
	return CreateTemporal(res)
}

// TemporalAppendTInstant Append an instant to a temporal value
func TemporalAppendTInstant[T Temporal, TI TInstant](temp T, inst TI, max_dist float64, max_time timeutil.Timedelta, expand bool) Temporal {
	m := TimeDeltaToInterval(max_time)
	res := C.temporal_append_tinstant(temp.Inner(), C.cast_temporal_to_tinstant(inst.Inner()), C.double(max_dist), &m, C.bool(expand))
	return CreateTemporal(res)
}

// TemporalAppendTSequence Append a sequence to a temporal value
func TemporalAppendTSequence[T Temporal, TS TSequence](temp T, seq TS, expand bool) Temporal {
	res := C.temporal_append_tsequence(temp.Inner(), C.cast_temporal_to_tsequence(seq.Inner()), C.bool(expand))
	return CreateTemporal(res)
}

// TemporalInsert Insert the second temporal value into the first one
func TemporalInsert[T1 Temporal, T2 Temporal](temp1 T1, temp2 T2, connect bool) Temporal {
	res := C.temporal_insert(temp1.Inner(), temp2.Inner(), C.bool(connect))
	return CreateTemporal(res)
}

// TemporalUpdate Update the first temporal value with the second one
func TemporalUpdate[T1 Temporal, T2 Temporal](temp1 T1, temp2 T2, connect bool) Temporal {
	res := C.temporal_update(temp1.Inner(), temp2.Inner(), C.bool(connect))
	return CreateTemporal(res)
}

// TemporalDeleteTimestamptz Delete a timestamp from a temporal value
func TemporalDeleteTimestamptz[T Temporal](temp T, ts time.Time, connect bool) Temporal {
	res := C.temporal_delete_timestamptz(temp.Inner(), DatetimeToTimestamptz(ts), C.bool(connect))
	return CreateTemporal(res)
}

// TemporalDeleteTsTzSet Delete a timestamp set from a temporal value connecting the instants before and after the given timestamp, if any
func TemporalDeleteTsTzSet[T Temporal](temp T, set TsTzSet, connect bool) Temporal {
	res := C.temporal_delete_tstzset(temp.Inner(), set._inner, C.bool(connect))
	return CreateTemporal(res)
}

// TemporalDeleteTsTzSpan Delete a timestamptz span from a temporal value
func TemporalDeleteTsTzSpan[T Temporal](temp T, span TsTzSpan, connect bool) Temporal {
	res := C.temporal_delete_tstzspan(temp.Inner(), span._inner, C.bool(connect))
	return CreateTemporal(res)
}

// TemporalDeleteTsTzSpanSet Delete a timestamptz span set from a temporal value
func TemporalDeleteTsTzSpanSet[T Temporal](temp T, spanset TsTzSpanSet, connect bool) Temporal {
	res := C.temporal_delete_tstzspanset(temp.Inner(), spanset._inner, C.bool(connect))
	return CreateTemporal(res)
}

// TemporalAtTimestamptz Return a temporal value restricted to a timestamptz
func TemporalAtTimestamptz[T Temporal](temp T, ts time.Time) Temporal {
	c_temp := C.temporal_at_timestamptz(temp.Inner(), DatetimeToTimestamptz(ts))
	res := CreateTemporal(c_temp)
	return res
}

// TemporalAtTsTzSet Return a temporal value restricted to a timestamptz set
func TemporalAtTsTzSet[T Temporal](temp T, tstzset TsTzSet) Temporal {
	c_temp := C.temporal_at_tstzset(temp.Inner(), tstzset._inner)
	res := CreateTemporal(c_temp)
	return res
}

// TemporalAtTsTzSpan Return a temporal value restricted to a timestamptz span
func TemporalAtTsTzSpan[T Temporal](temp T, tstzspan TsTzSpan) Temporal {
	c_temp := C.temporal_at_tstzspan(temp.Inner(), tstzspan._inner)
	res := CreateTemporal(c_temp)
	return res
}

// TemporalAtTsTzSpanSet Return a temporal value restricted to a timestamptz span set
func TemporalAtTsTzSpanSet[T Temporal](temp T, tstzspanset TsTzSpanSet) Temporal {
	c_temp := C.temporal_at_tstzspanset(temp.Inner(), tstzspanset._inner)
	res := CreateTemporal(c_temp)
	return res
}

// TemporalAtMax Return a temporal value restricted to its maximum base value
func TemporalAtMax[T Temporal](temp T) Temporal {
	c_temp := C.temporal_at_max(temp.Inner())
	res := CreateTemporal(c_temp)
	return res
}

// TemporalAtMin Return a temporal value restricted to its minimum base value
func TemporalAtMin[T Temporal](temp T) Temporal {
	c_temp := C.temporal_at_min(temp.Inner())
	res := CreateTemporal(c_temp)
	return res
}

// TemporalAtValues Return a temporal value restricted to a set of values
func TemporalAtValues[T Temporal, S Set](temp T, set S) Temporal {
	c_temp := C.temporal_at_values(temp.Inner(), set.Inner())
	res := CreateTemporal(c_temp)
	return res
}

// TemporalMinusTimestamptz Return a temporal value restricted to the complement of a timestamptz
func TemporalMinusTimestamptz[T Temporal](temp T, ts time.Time) Temporal {
	c_temp := C.temporal_minus_timestamptz(temp.Inner(), DatetimeToTimestamptz(ts))
	res := CreateTemporal(c_temp)
	return res
}

// TemporalMinusTsTzSet Return a temporal value restricted to the complement of a timestamptz set
func TemporalMinusTsTzSet[T Temporal](temp T, tstzset TsTzSet) Temporal {
	c_temp := C.temporal_minus_tstzset(temp.Inner(), tstzset._inner)
	res := CreateTemporal(c_temp)
	return res
}

// TemporalMinusTsTzSpan Return a temporal value restricted to the complement of a timestamptz span
func TemporalMinusTsTzSpan[T Temporal](temp T, tstzspan TsTzSpan) Temporal {
	c_temp := C.temporal_minus_tstzspan(temp.Inner(), tstzspan._inner)
	res := CreateTemporal(c_temp)
	return res
}

// TemporalMinusTsTzSpanSet Return a temporal value restricted to the complement of a timestamptz span set
func TemporalMinusTsTzSpanSet[T Temporal](temp T, tstzspanset TsTzSpanSet) Temporal {
	c_temp := C.temporal_minus_tstzspanset(temp.Inner(), tstzspanset._inner)
	res := CreateTemporal(c_temp)
	return res
}

// TemporalMinusMax Return a temporal value restricted to the complement of its maximum base value
func TemporalMinusMax[T Temporal](temp T) Temporal {
	c_temp := C.temporal_at_max(temp.Inner())
	res := CreateTemporal(c_temp)
	return res
}

// TemporalMinusMin Return a temporal value restricted to the complement of its minimum base value
func TemporalMinusMin[T Temporal](temp T) Temporal {
	c_temp := C.temporal_minus_min(temp.Inner())
	res := CreateTemporal(c_temp)
	return res
}

// AdjacentTemporalTemporal Return true if the time spans of two temporal values are adjacent
func AdjacentTemporalTemporal[T1 Temporal, T2 Temporal](temp1 T1, temp2 T2) bool {
	return bool(C.adjacent_temporal_temporal(temp1.Inner(), temp2.Inner()))
}

// AdjacentTemporalTsTzSpan Return true if the time span of a temporal value and a timestamptz span are adjacent
func AdjacentTemporalTsTzSpan[T Temporal](temp T, tsTzSpan TsTzSpan) bool {
	return bool(C.adjacent_temporal_tstzspan(temp.Inner(), tsTzSpan._inner))
}

// ContainedTemporalTemporal checks if one temporal value is contained within another
func ContainedTemporalTemporal[T1, T2 Temporal](temp1 T1, temp2 T2) bool {
	return bool(C.contained_temporal_temporal(temp1.Inner(), temp2.Inner()))
}

// ContainedTemporalTsTzSpan checks if a temporal value is contained within a timestamptz span
func ContainedTemporalTsTzSpan[T Temporal](temp T, tsTzSpan TsTzSpan) bool {
	return bool(C.contained_temporal_tstzspan(temp.Inner(), tsTzSpan._inner))
}

// ContainsTemporalTsTzSpan checks if a temporal value contains a timestamptz span
func ContainsTemporalTsTzSpan[T Temporal](temp T, tsTzSpan TsTzSpan) bool {
	return bool(C.contains_temporal_tstzspan(temp.Inner(), tsTzSpan._inner))
}

// ContainsTemporalTemporal checks if one temporal value contains another
func ContainsTemporalTemporal[T1, T2 Temporal](temp1 T1, temp2 T2) bool {
	return bool(C.contains_temporal_temporal(temp1.Inner(), temp2.Inner()))
}

// OverlapsTemporalTemporal checks if two temporal values overlap
func OverlapsTemporalTemporal[T1, T2 Temporal](temp1 T1, temp2 T2) bool {
	return bool(C.overlaps_temporal_temporal(temp1.Inner(), temp2.Inner()))
}

// OverlapsTemporalTsTzSpan checks if a temporal value overlaps with a timestamptz span
func OverlapsTemporalTsTzSpan[T Temporal](temp T, tsTzSpan TsTzSpan) bool {
	return bool(C.overlaps_temporal_tstzspan(temp.Inner(), tsTzSpan._inner))
}

// SameTemporalTemporal checks if two temporal values are the same
func SameTemporalTemporal[T1, T2 Temporal](temp1 T1, temp2 T2) bool {
	return bool(C.same_temporal_temporal(temp1.Inner(), temp2.Inner()))
}

// SameTemporalTsTzSpan checks if a temporal value is the same as a timestamptz span
func SameTemporalTsTzSpan[T Temporal](temp T, tsTzSpan TsTzSpan) bool {
	return bool(C.same_temporal_tstzspan(temp.Inner(), tsTzSpan._inner))
}

// AfterTemporalTsTzSpan checks if a temporal value is after a timestamptz span
func AfterTemporalTsTzSpan[T Temporal](temp T, tsTzSpan TsTzSpan) bool {
	return bool(C.after_temporal_tstzspan(temp.Inner(), tsTzSpan._inner))
}

// AfterTemporalTemporal checks if one temporal value is after another
func AfterTemporalTemporal[T1, T2 Temporal](temp1 T1, temp2 T2) bool {
	return bool(C.after_temporal_temporal(temp1.Inner(), temp2.Inner()))
}

// BeforeTemporalTsTzSpan checks if a temporal value is before a timestamptz span
func BeforeTemporalTsTzSpan[T Temporal](temp T, tsTzSpan TsTzSpan) bool {
	return bool(C.before_temporal_tstzspan(temp.Inner(), tsTzSpan._inner))
}

// BeforeTemporalTemporal checks if one temporal value is before another
func BeforeTemporalTemporal[T1, T2 Temporal](temp1 T1, temp2 T2) bool {
	return bool(C.before_temporal_temporal(temp1.Inner(), temp2.Inner()))
}

// OverAfterTemporalTsTzSpan checks if a temporal value is over and after a timestamptz span
func OverAfterTemporalTsTzSpan[T Temporal](temp T, tsTzSpan TsTzSpan) bool {
	return bool(C.overafter_temporal_tstzspan(temp.Inner(), tsTzSpan._inner))
}

// OverAfterTemporalTemporal checks if one temporal value is over and after another
func OverAfterTemporalTemporal[T1, T2 Temporal](temp1 T1, temp2 T2) bool {
	return bool(C.overafter_temporal_temporal(temp1.Inner(), temp2.Inner()))
}

// OverBeforeTemporalTsTzSpan checks if a temporal value is over and before a timestamptz span
func OverBeforeTemporalTsTzSpan[T Temporal](temp T, tsTzSpan TsTzSpan) bool {
	return bool(C.overbefore_temporal_tstzspan(temp.Inner(), tsTzSpan._inner))
}

// OverBeforeTemporalTemporal checks if one temporal value is over and before another
func OverBeforeTemporalTemporal[T1, T2 Temporal](temp1 T1, temp2 T2) bool {
	return bool(C.overbefore_temporal_temporal(temp1.Inner(), temp2.Inner()))
}

// TemporalFrechetDistance Return the Frechet distance between two temporal values
func TemporalFrechetDistance[T Temporal](temp1 T, temp2 T) float64 {
	return float64(C.temporal_frechet_distance(temp1.Inner(), temp2.Inner()))
}

// TemporalDyntimewarpDistance Return the Dynamic Time Warp distance between two temporal values
func TemporalDyntimewarpDistance[T Temporal](temp1 T, temp2 T) float64 {
	return float64(C.temporal_dyntimewarp_distance(temp1.Inner(), temp2.Inner()))
}

// TemporalHausdorffDistance Return the Hausdorf distance between two temporal values
func TemporalHausdorffDistance[T Temporal](temp1 T, temp2 T) float64 {
	return float64(C.temporal_hausdorff_distance(temp1.Inner(), temp2.Inner()))
}

// TemporalSimplifyDP Return a temporal float/point simplified using the Douglas-Peucker line simplification algorithm
func TemporalSimplifyDP[T Temporal](temp T, new_temp T, dist float64, syncdist bool) T {
	c_temp := C.temporal_simplify_dp(temp.Inner(), C.double(dist), C.bool(syncdist))
	new_temp.Init(c_temp)
	return new_temp
}

// AlwaysEqTemporalTemporal Return true if two temporal values are always equal
func AlwaysEqTemporalTemporal[T1 Temporal, T2 Temporal](temp1 T1, temp2 T2) bool {
	return int(C.always_eq_temporal_temporal(temp1.Inner(), temp2.Inner())) > 0
}

// EverEqTemporalTemporal Return true if two temporal values are ever equal
func EverEqTemporalTemporal[T1 Temporal, T2 Temporal](temp1 T1, temp2 T2) bool {
	return int(C.ever_eq_temporal_temporal(temp1.Inner(), temp2.Inner())) > 0
}

// TEqTemporalTemporal Return the temporal equality of two temporal values
func TEqTemporalTemporal[T1 Temporal, T2 Temporal](temp1 T1, temp2 T2) Temporal {
	c_temp := C.teq_temporal_temporal(temp1.Inner(), temp2.Inner())
	return CreateTemporal(c_temp)
}

// TNEqTemporalTemporal Return the temporal equality of two temporal values
func TNEqTemporalTemporal[T1 Temporal, T2 Temporal](temp1 T1, temp2 T2) Temporal {
	c_temp := C.tne_temporal_temporal(temp1.Inner(), temp2.Inner())
	return CreateTemporal(c_temp)
}

// AlwaysLtTemporalTemporal returns true if one temporal value is always less than another temporal value
func AlwaysLtTemporalTemporal[T Temporal](t1, t2 T) bool {
	return int(C.always_lt_temporal_temporal(t1.Inner(), t2.Inner())) > 0
}

// AlwaysLeTemporalTemporal returns true if one temporal value is always less than or equal to another temporal value
func AlwaysLeTemporalTemporal[T Temporal](t1, t2 T) bool {
	return int(C.always_le_temporal_temporal(t1.Inner(), t2.Inner())) > 0
}

// AlwaysNeTemporalTemporal returns true if one temporal value is always not equal to another temporal value
func AlwaysNeTemporalTemporal[T Temporal](t1, t2 T) bool {
	return int(C.always_ne_temporal_temporal(t1.Inner(), t2.Inner())) > 0
}

// AlwaysGeTemporalTemporal returns true if one temporal value is always greater than or equal to another temporal value
func AlwaysGeTemporalTemporal[T Temporal](t1, t2 T) bool {
	return int(C.always_ge_temporal_temporal(t1.Inner(), t2.Inner())) > 0
}

// AlwaysGtTemporalTemporal returns true if one temporal value is always greater than another temporal value
func AlwaysGtTemporalTemporal[T Temporal](t1, t2 T) bool {
	return int(C.always_gt_temporal_temporal(t1.Inner(), t2.Inner())) > 0
}

// EverLtTemporalTemporal returns true if one temporal value is ever less than another temporal value
func EverLtTemporalTemporal[T Temporal](t1, t2 T) bool {
	return int(C.ever_lt_temporal_temporal(t1.Inner(), t2.Inner())) > 0
}

// EverLeTemporalTemporal returns true if one temporal value is ever less than or equal to another temporal value
func EverLeTemporalTemporal[T Temporal](t1, t2 T) bool {
	return int(C.ever_le_temporal_temporal(t1.Inner(), t2.Inner())) > 0
}

// EverNeTemporalTemporal returns true if one temporal value is ever not equal to another temporal value
func EverNeTemporalTemporal[T Temporal](t1, t2 T) bool {
	return int(C.ever_ne_temporal_temporal(t1.Inner(), t2.Inner())) > 0
}

// EverGeTemporalTemporal returns true if one temporal value is ever greater than or equal to another temporal value
func EverGeTemporalTemporal[T Temporal](t1, t2 T) bool {
	return int(C.ever_ge_temporal_temporal(t1.Inner(), t2.Inner())) > 0
}

// EverGtTemporalTemporal returns true if one temporal value is ever greater than another temporal value
func EverGtTemporalTemporal[T Temporal](t1, t2 T) bool {
	return int(C.ever_gt_temporal_temporal(t1.Inner(), t2.Inner())) > 0
}
