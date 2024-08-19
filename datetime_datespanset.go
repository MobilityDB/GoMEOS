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

type DateSpanSet struct {
	_inner *C.SpanSet
}

// ------------------------- Input ----------------------------------------
func NewDateSpanSet(g_dss_in string) *DateSpanSet {
	c_dss_in := C.CString(g_dss_in)
	defer C.free(unsafe.Pointer(c_dss_in))
	c_dss := C.datespanset_in(c_dss_in)
	g_dss := &DateSpanSet{_inner: c_dss}
	return g_dss
}

// ------------------------- Output ----------------------------------------
func (g_dss *DateSpanSet) DateSpanSetOut() string {
	if g_dss._inner == nil {
		return "Nil"
	} else {
		c_dss_out := C.datespanset_out(g_dss._inner)
		defer C.free(unsafe.Pointer(c_dss_out))
		g_dss_out := C.GoString(c_dss_out)
		return g_dss_out
	}

}

// ------------------------- Conversions -----------------------------------
func (g_dss *DateSpanSet) ToSpan() *DateSpan {
	c_ds := C.spanset_span(g_dss._inner)
	return &DateSpan{_inner: c_ds}
}

func (g_dss *DateSpanSet) ToTsTzSpanSet() *TsTzSpanSet {
	return &TsTzSpanSet{_inner: C.datespanset_to_tstzspanset(g_dss._inner)}
}

// ------------------------- Accessors -------------------------------------

func (g_dss *DateSpanSet) Duration(ignore_gap bool) timeutil.Timedelta {
	interval := C.datespanset_duration(g_dss._inner, C.bool(ignore_gap))
	microsecond := int(interval.time)
	day := int(interval.day)
	month := int(interval.month)
	dr := timeutil.Timedelta{
		Microseconds: time.Duration(microsecond),
		Days:         time.Duration(day) + time.Duration(month*30),
	}
	return dr
}

func (g_dss *DateSpanSet) NumDates() int {
	return int(C.datespanset_num_dates(g_dss._inner))
}

func (g_dss *DateSpanSet) StartDate() time.Time {
	start_date := C.datespanset_start_date(g_dss._inner)
	return DateADTToDate(start_date)
}

func (g_dss *DateSpanSet) EndDate() time.Time {
	end_date := C.datespanset_end_date(g_dss._inner)
	return DateADTToDate(end_date)
}

func (g_dss *DateSpanSet) DateN(n int) time.Time {
	res := C.malloc(C.sizeof_int)
	defer C.free(unsafe.Pointer(res)) // Ensure memory is freed.
	success := C.datespanset_date_n(g_dss._inner, C.int(n+1), (*C.DateADT)(res))
	if success {
		result := *(*C.DateADT)(res)
		return DateADTToDate(result)
	} else {
		return time.Time{}
	}
}

func (g_dss *DateSpanSet) Dates() []time.Time {
	nums := g_dss.NumSpans()
	dates := make([]time.Time, nums)
	for i := 0; i < nums; i++ {
		dates[i] = g_dss.DateN(i)
	}
	return dates
}

func (g_dss *DateSpanSet) NumSpans() int {
	return int(C.spanset_num_spans(g_dss._inner))
}

func (g_dss *DateSpanSet) StartSpan() DateSpan {
	return DateSpan{_inner: C.spanset_start_span(g_dss._inner)}
}

func (g_dss *DateSpanSet) EndSpan() DateSpan {
	return DateSpan{_inner: C.spanset_end_span(g_dss._inner)}
}

func (g_dss *DateSpanSet) SpanN(n int) DateSpan {
	return DateSpan{_inner: C.spanset_span_n(g_dss._inner, C.int(n+1))}
}

func (g_dss *DateSpanSet) Spans() []DateSpan {
	nums := g_dss.NumSpans()
	spans := make([]DateSpan, nums)
	for i := 0; i < nums; i++ {
		spans[i] = g_dss.SpanN(i)
	}
	return spans
}

// ------------------------- Transformations -------------------------------
func (g_ds *DateSpanSet) ShiftScale(shift interface{}, duration interface{}) (*DateSpanSet, error) {
	if shift == nil && duration == nil {
		return nil, fmt.Errorf("shift and duration must not be both nil")
	}
	var shift_in, duration_in int
	switch s := shift.(type) {
	case timeutil.Timedelta:
		shift_in = int(s.Days)
	case int:
		shift_in = s
	case nil:
		shift_in = 0
	default:
		return &DateSpanSet{_inner: nil}, fmt.Errorf("operation not supported with type %T", shift)
	}

	switch d := duration.(type) {
	case timeutil.Timedelta:
		duration_in = int(d.Days)
	case int:
		duration_in = d
	case nil:
		duration_in = 0
	default:
		return &DateSpanSet{_inner: nil}, fmt.Errorf("operation not supported with type %T", duration)
	}

	tss := C.datespanset_shift_scale(g_ds._inner, C.int(shift_in), C.int(duration_in), shift_in != 0, duration_in != 0)
	return &DateSpanSet{_inner: tss}, nil
}

func (g_ds *DateSpanSet) Shift(delta interface{}) (*DateSpanSet, error) {
	return g_ds.ShiftScale(delta, 0)
}

func (g_ds *DateSpanSet) Scale(duration interface{}) (*DateSpanSet, error) {
	return g_ds.ShiftScale(0, duration)
}

// ------------------------- Topological Operations ------------------------
func (g_dss *DateSpanSet) Contains(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return bool(C.contains_spanset_date(g_dss._inner, DateToDateADT(o))), nil
	case *DateSpan:
		return bool(C.contains_spanset_span(g_dss._inner, o._inner)), nil
	case *DateSpanSet:
		return bool(C.contains_spanset_spanset(g_dss._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_dss *DateSpanSet) Overlaps(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return g_dss.Contains(o)
	case *DateSpan:
		return bool(C.overlaps_spanset_span(g_dss._inner, o._inner)), nil
	case *DateSpanSet:
		return bool(C.overlaps_spanset_spanset(g_dss._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_dss *DateSpanSet) IsAdjacent(other interface{}) (bool, error) {
	switch o := other.(type) {
	case *DateSpan:
		return bool(C.adjacent_spanset_span(g_dss._inner, o._inner)), nil
	case *DateSpanSet:
		return bool(C.adjacent_spanset_spanset(g_dss._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

// ------------------------- Position Operations ---------------------------
func (g_ds *DateSpanSet) IsLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return bool(C.before_spanset_date(g_ds._inner, DateToDateADT(o))), nil
	case *DateSpan:
		return bool(C.left_spanset_span(g_ds._inner, o._inner)), nil
	case *DateSpanSet:
		return bool(C.left_spanset_spanset(g_ds._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_ds *DateSpanSet) IsOverOrLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return bool(C.overbefore_spanset_date(g_ds._inner, DateToDateADT(o))), nil
	case *DateSpan:
		return bool(C.overleft_spanset_span(g_ds._inner, o._inner)), nil
	case *DateSpanSet:
		return bool(C.overleft_spanset_spanset(g_ds._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_ds *DateSpanSet) IsRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return bool(C.after_spanset_date(g_ds._inner, DateToDateADT(o))), nil
	case *DateSpan:
		return bool(C.right_spanset_span(g_ds._inner, o._inner)), nil
	case *DateSpanSet:
		return bool(C.right_spanset_spanset(g_ds._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_ds *DateSpanSet) IsOverOrRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return bool(C.overafter_spanset_date(g_ds._inner, DateToDateADT(o))), nil
	case *DateSpan:
		return bool(C.overright_spanset_span(g_ds._inner, o._inner)), nil
	case *DateSpanSet:
		return bool(C.overright_spanset_spanset(g_ds._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

// ------------------------- Distance Operations ---------------------------
func (g_dss *DateSpanSet) Distance(other interface{}) (timeutil.Timedelta, error) {
	switch o := other.(type) {
	case time.Time:
		days := int(C.distance_spanset_date(g_dss._inner, DateToDateADT(o)))
		return timeutil.Timedelta{Days: time.Duration(days)}, nil
	case *DateSet:
		return g_dss.Distance(o.ToSpanSet())
	case *DateSpan:
		days := int(C.distance_datespanset_datespan(g_dss._inner, o._inner))
		return timeutil.Timedelta{Days: time.Duration(days)}, nil
	case *DateSpanSet:
		days := int(C.distance_datespanset_datespanset(g_dss._inner, o._inner))
		return timeutil.Timedelta{Days: time.Duration(days)}, nil
	default:
		return timeutil.Timedelta{}, fmt.Errorf("operation not supported with type %T", other)
	}
}

// ------------------------- Set Operations --------------------------------
func (g_dss *DateSpanSet) Intersection(other interface{}) (*DateSpanSet, error) {
	switch o := other.(type) {
	case time.Time:
		res := C.intersection_spanset_date(g_dss._inner, DateToDateADT(o))
		return &DateSpanSet{_inner: res}, nil
	case *DateSpan:
		res := C.intersection_spanset_span(g_dss._inner, o._inner)
		return &DateSpanSet{_inner: res}, nil
	case *DateSpanSet:
		res := C.intersection_spanset_spanset(g_dss._inner, o._inner)
		return &DateSpanSet{_inner: res}, nil
	default:
		return &DateSpanSet{_inner: nil}, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_dss *DateSpanSet) Mul(other interface{}) (*DateSpanSet, error) {
	return g_dss.Intersection(other)
}

func (g_dss *DateSpanSet) Minus(other interface{}) (*DateSpanSet, error) {
	switch o := other.(type) {
	case time.Time:
		res := C.minus_spanset_date(g_dss._inner, DateToDateADT(o))
		return &DateSpanSet{_inner: res}, nil
	case *DateSpan:
		res := C.minus_spanset_span(g_dss._inner, o._inner)
		return &DateSpanSet{_inner: res}, nil
	case *DateSpanSet:
		res := C.minus_spanset_spanset(g_dss._inner, o._inner)
		return &DateSpanSet{_inner: res}, nil
	default:
		return &DateSpanSet{_inner: nil}, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_dss *DateSpanSet) Sub(other interface{}) (*DateSpanSet, error) {
	return g_dss.Minus(other)
}

func (g_dss *DateSpanSet) Union(other interface{}) (*DateSpanSet, error) {
	switch o := other.(type) {
	case time.Time:
		res := C.gunion_spanset_date(g_dss._inner, DateToDateADT(o))
		return &DateSpanSet{_inner: res}, nil
	case *DateSpan:
		res := C.gunion_spanset_span(g_dss._inner, o._inner)
		return &DateSpanSet{_inner: res}, nil
	case *DateSpanSet:
		res := C.gunion_spanset_spanset(g_dss._inner, o._inner)
		return &DateSpanSet{_inner: res}, nil
	default:
		return &DateSpanSet{_inner: nil}, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_dss *DateSpanSet) Add(other interface{}) (*DateSpanSet, error) {
	return g_dss.Union(other)
}
