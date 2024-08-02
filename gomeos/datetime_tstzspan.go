package gomeos

/*
#include "meos.h"
#include "meos_catalog.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"

	"github.com/leekchan/timeutil"
)

type TsTzSpan struct {
	_inner *C.Span
}

// ------------------------- Input ----------------------------------------
func NewTsTzSpan(g_tts_in string) *TsTzSpan {
	c_tts_in := C.CString(g_tts_in)
	defer C.free(unsafe.Pointer(c_tts_in))
	c_tts := C.tstzspan_in(c_tts_in)
	g_tts := &TsTzSpan{_inner: c_tts}
	return g_tts
}

// ------------------------- Output ----------------------------------------
func (g_tts *TsTzSpan) TsTzSpanOut() string {
	c_tts_out := C.tstzspan_out(g_tts._inner)
	defer C.free(unsafe.Pointer(c_tts_out))
	g_tts_out := C.GoString(c_tts_out)
	return g_tts_out
}

// ------------------------- Conversions -----------------------------------
func (g_tts *TsTzSpan) ToSpanSet() TsTzSpanSet {
	c_ds := C.span_to_spanset(g_tts._inner)
	return TsTzSpanSet{_inner: c_ds}
}

// ------------------------- Accessors -------------------------------------
func (g_tts *TsTzSpan) Duration() timeutil.Timedelta {
	duration := C.tstzspan_duration(g_tts._inner)
	return IntervalToTimeDelta(*duration)
}

func (g_tts *TsTzSpan) Lower() time.Time {
	s := C.tstzspan_lower(g_tts._inner)
	return TimestamptzToDatetime(s)
}

func (g_tts *TsTzSpan) Upper() time.Time {
	s := C.tstzspan_upper(g_tts._inner)
	return TimestamptzToDatetime(s)
}

// ------------------------- Transformations -------------------------------
func (g_tss *TsTzSpan) ShiftScale(shift interface{}, duration interface{}) (*TsTzSpan, error) {
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
		return &TsTzSpan{}, fmt.Errorf("operation not supported with type %T", shift)
	}

	switch d := duration.(type) {
	case timeutil.Timedelta:
		duration_delta = 1
		duration_in = d
		duration_interval = TimeDeltaToInterval(duration_in)
	case nil:
		duration_delta = 0
	default:
		return &TsTzSpan{}, fmt.Errorf("operation not supported with type %T", duration)
	}

	if (shift_delta == 1) && (duration_delta == 1) {
		tss := C.tstzspan_shift_scale(g_tss._inner, &shift_interval, &duration_interval)
		return &TsTzSpan{_inner: tss}, nil
	} else if shift_delta == 0 {
		tss := C.tstzspan_shift_scale(g_tss._inner, nil, &duration_interval)
		return &TsTzSpan{_inner: tss}, nil
	} else {
		tss := C.tstzspan_shift_scale(g_tss._inner, &shift_interval, nil)
		return &TsTzSpan{_inner: tss}, nil
	}
}

func (g_tss *TsTzSpan) Shift(delta interface{}) (*TsTzSpan, error) {
	return g_tss.ShiftScale(delta, nil)
}

func (g_tss *TsTzSpan) Scale(duration interface{}) (*TsTzSpan, error) {
	return g_tss.ShiftScale(nil, duration)
}
