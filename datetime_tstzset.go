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

type TsTzSet struct {
	_inner *C.Set
}

func NewTsTzSet(g_tts_in string) *TsTzSet {
	c_tts_in := C.CString(g_tts_in)
	defer C.free(unsafe.Pointer(c_tts_in))
	c_tts := C.tstzset_in(c_tts_in)
	g_tts := &TsTzSet{_inner: c_tts}
	return g_tts
}

func (g_tts *TsTzSet) TsTzSetOut() string {
	c_tts_out := C.tstzset_out(g_tts._inner)
	defer C.free(unsafe.Pointer(c_tts_out))
	g_tts_out := C.GoString(c_tts_out)
	return g_tts_out
}

func (g_tts *TsTzSet) Duration() timeutil.Timedelta {
	span := C.set_to_span(g_tts._inner)
	duration := C.tstzspan_duration(span)
	return IntervalToTimeDelta(*duration)
}

func (g_tts *TsTzSet) StartElement() time.Time {
	s := C.tstzset_start_value(g_tts._inner)
	return TimestamptzToDatetime(s)
}

func (g_tts *TsTzSet) EndElement() time.Time {
	s := C.tstzset_end_value(g_tts._inner)
	return TimestamptzToDatetime(s)
}

func (g_tss *TsTzSet) ElementN(n int) time.Time {
	res := C.malloc(C.sizeof_int)
	defer C.free(unsafe.Pointer(res)) // Ensure memory is freed.
	success := C.tstzset_value_n(g_tss._inner, C.int(n+1), (*C.TimestampTz)(res))
	if success {
		result := *(*C.TimestampTz)(res)
		return TimestamptzToDatetime(result)
	} else {
		return time.Time{}
	}
}

func (g_tss *TsTzSet) NumElements() int {
	return int(C.set_num_values(g_tss._inner))
}

func (g_tss *TsTzSet) Elements() []time.Time {
	nums := g_tss.NumElements()
	dates := make([]time.Time, nums)
	for i := 0; i < nums; i++ {
		dates[i] = g_tss.ElementN(i)
	}
	return dates
}

func (g_tss *TsTzSet) ShiftScale(shift interface{}, duration interface{}) (*TsTzSet, error) {
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
		return &TsTzSet{}, fmt.Errorf("operation not supported with type %T", shift)
	}

	switch d := duration.(type) {
	case timeutil.Timedelta:
		duration_delta = 1
		duration_in = d
		duration_interval = TimeDeltaToInterval(duration_in)
	case nil:
		duration_delta = 0
	default:
		return &TsTzSet{}, fmt.Errorf("operation not supported with type %T", duration)
	}

	if (shift_delta == 1) && (duration_delta == 1) {
		tss := C.tstzset_shift_scale(g_tss._inner, &shift_interval, &duration_interval)
		return &TsTzSet{_inner: tss}, nil
	} else if shift_delta == 0 {
		tss := C.tstzset_shift_scale(g_tss._inner, nil, &duration_interval)
		return &TsTzSet{_inner: tss}, nil
	} else {
		tss := C.tstzset_shift_scale(g_tss._inner, &shift_interval, nil)
		return &TsTzSet{_inner: tss}, nil
	}
}

func (g_tss *TsTzSet) Shift(delta interface{}) (*TsTzSet, error) {
	return g_tss.ShiftScale(delta, nil)
}

func (g_tss *TsTzSet) Scale(duration interface{}) (*TsTzSet, error) {
	return g_tss.ShiftScale(nil, duration)
}
