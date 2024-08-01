package gomeos

/*
#include "meos.h"
#include "meos_catalog.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"time"
	"unsafe"

	"github.com/leekchan/timeutil"
)

type TsTzSet struct {
	_inner *C.Set
}

// ------------------------- Input ----------------------------------------
func NewTsTzSet(g_tts_in string) *TsTzSet {
	c_tts_in := C.CString(g_tts_in)
	defer C.free(unsafe.Pointer(c_tts_in))
	c_tts := C.tstzset_in(c_tts_in)
	g_tts := &TsTzSet{_inner: c_tts}
	return g_tts
}

// ------------------------- Output ----------------------------------------
func (g_tts *TsTzSet) TsTzSetOut() string {
	c_tts_out := C.tstzset_out(g_tts._inner)
	defer C.free(unsafe.Pointer(c_tts_out))
	g_tts_out := C.GoString(c_tts_out)
	return g_tts_out
}

// ------------------------- Conversions -----------------------------------
// ------------------------- Accessors -------------------------------------

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
