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
	"fmt"
	"time"
	"unsafe"

	"github.com/leekchan/timeutil"
)

type TsTzSpanSet struct {
	_inner *C.SpanSet
}

func NewTsTzSpanSet(g_tts_in string) *TsTzSpanSet {
	c_tts_in := C.CString(g_tts_in)
	defer C.free(unsafe.Pointer(c_tts_in))
	c_tts := C.tstzspanset_in(c_tts_in)
	g_tts := &TsTzSpanSet{_inner: c_tts}
	return g_tts
}

func (g_tts *TsTzSpanSet) TsTzSpanSetOut() string {
	c_tts_out := C.tstzspanset_out(g_tts._inner)
	defer C.free(unsafe.Pointer(c_tts_out))
	g_tts_out := C.GoString(c_tts_out)
	return g_tts_out
}

func (g_tts *TsTzSpanSet) ToSpan() *TsTzSpan {
	c_ds := C.spanset_span(g_tts._inner)
	return &TsTzSpan{_inner: c_ds}
}

func (g_tts *TsTzSpanSet) ToTsTzSpan() *TsTzSpan {
	return g_tts.ToSpan()
}

func (g_tts *TsTzSpanSet) Duration(ignore_gap bool) timeutil.Timedelta {
	interval := C.tstzspanset_duration(g_tts._inner, C.bool(ignore_gap))
	delta := IntervalToTimeDelta(*interval)
	return delta
}

func (g_tts *TsTzSpanSet) NumTimestamps() int {
	return int(C.tstzspanset_num_timestamps(g_tts._inner))
}

func (g_tts *TsTzSpanSet) StartTimestamp() time.Time {
	start_ts := C.tstzspanset_start_timestamptz(g_tts._inner)
	return TimestamptzToDatetime(start_ts)
}

func (g_tts *TsTzSpanSet) EndTimestamp() time.Time {
	start_ts := C.tstzspanset_end_timestamptz(g_tts._inner)
	return TimestamptzToDatetime(start_ts)
}

func (g_tts *TsTzSpanSet) TimestampN(n int) time.Time {
	res := C.malloc(C.sizeof_int)
	defer C.free(unsafe.Pointer(res)) // Ensure memory is freed.
	success := C.tstzspanset_timestamptz_n(g_tts._inner, C.int(n+1), (*C.TimestampTz)(res))
	if success {
		result := *(*C.TimestampTz)(res)
		return TimestamptzToDatetime(result)
	} else {
		return time.Time{}
	}
}

func (g_tts *TsTzSpanSet) Timestamps() []time.Time {
	nums := g_tts.NumTimestamps()
	timestamps := make([]time.Time, nums)
	for i := 0; i < nums; i++ {
		timestamps[i] = g_tts.TimestampN(i)
	}
	return timestamps
}

func (g_tts *TsTzSpanSet) NumSpans() int {
	return int(C.spanset_num_spans(g_tts._inner))
}

func (g_tts *TsTzSpanSet) StartSpan() TsTzSpan {
	return TsTzSpan{_inner: C.spanset_start_span(g_tts._inner)}
}

func (g_tts *TsTzSpanSet) EndSpan() TsTzSpan {
	return TsTzSpan{_inner: C.spanset_end_span(g_tts._inner)}
}

func (g_tts *TsTzSpanSet) SpanN(n int) TsTzSpan {
	return TsTzSpan{_inner: C.spanset_span_n(g_tts._inner, C.int(n+1))}
}

func (g_tts *TsTzSpanSet) Spans() []TsTzSpan {
	nums := g_tts.NumSpans()
	spans := make([]TsTzSpan, nums)
	for i := 0; i < nums; i++ {
		spans[i] = g_tts.SpanN(i)
	}
	return spans
}

func (g_tts *TsTzSpanSet) ShiftScale(shift interface{}, duration interface{}) (*TsTzSpanSet, error) {
	if shift == nil && duration == nil {
		return nil, fmt.Errorf("shift and duration must not be both nil")
	}
	var shift_delta, duration_delta int
	var shift_in, duration_in timeutil.Timedelta
	var shift_interval, duration_interval C.Interval
	switch s := shift.(type) {
	case timeutil.Timedelta:
		shift_delta = 1
		shift_in = s
		shift_interval = TimeDeltaToInterval(shift_in)
	case nil:
		shift_delta = 0
	default:
		return &TsTzSpanSet{}, fmt.Errorf("operation not supported with type %T", shift)
	}

	switch d := duration.(type) {
	case timeutil.Timedelta:
		duration_delta = 1
		duration_in = d
		duration_interval = TimeDeltaToInterval(duration_in)
	case nil:
		duration_delta = 0
	default:
		return &TsTzSpanSet{}, fmt.Errorf("operation not supported with type %T", duration)
	}

	if (shift_delta == 1) && (duration_delta == 1) {
		tss := C.tstzspanset_shift_scale(g_tts._inner, &shift_interval, &duration_interval)
		return &TsTzSpanSet{_inner: tss}, nil
	} else if shift_delta == 0 {
		tss := C.tstzspanset_shift_scale(g_tts._inner, nil, &duration_interval)
		return &TsTzSpanSet{_inner: tss}, nil
	} else {
		tss := C.tstzspanset_shift_scale(g_tts._inner, &shift_interval, nil)
		return &TsTzSpanSet{_inner: tss}, nil
	}
}

func (g_tts *TsTzSpanSet) Shift(delta interface{}) (*TsTzSpanSet, error) {
	return g_tts.ShiftScale(delta, nil)
}

func (g_tts *TsTzSpanSet) Scale(duration interface{}) (*TsTzSpanSet, error) {
	return g_tts.ShiftScale(nil, duration)
}
