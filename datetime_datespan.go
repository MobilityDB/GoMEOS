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

// DateSpan wraps the C Span type
type DateSpan struct {
	_inner *C.Span
}

// ------------------------- Input ----------------------------------------
func NewDateSpan(g_ds_in string) *DateSpan {
	c_ds_in := C.CString(g_ds_in)
	defer C.free(unsafe.Pointer(c_ds_in))
	c_ds := C.datespan_in(c_ds_in)
	g_ds := &DateSpan{_inner: c_ds}
	return g_ds
}

// ------------------------- Output ----------------------------------------
func (g_ds *DateSpan) DateSpanOut() string {
	if g_ds._inner == nil {
		return "Nil"
	} else {
		c_ds_out := C.datespan_out(g_ds._inner)
		defer C.free(unsafe.Pointer(c_ds_out))
		g_ds_out := C.GoString(c_ds_out)
		return g_ds_out
	}

}

// ------------------------- Conversions -----------------------------------

func (g_ds *DateSpan) ToSpanSet() *DateSpanSet {
	c_ds := C.span_to_spanset(g_ds._inner)
	return &DateSpanSet{_inner: c_ds}
}

func (g_ds *DateSpan) ToTstzspan() *TsTzSpan {
	c_ds := C.datespan_to_tstzspan(g_ds._inner)
	return &TsTzSpan{_inner: c_ds}
}

// ------------------------- Accessors -------------------------------------
func (g_ds *DateSpan) Lower() time.Time {
	return DateADTToDate(C.datespan_lower(g_ds._inner))
}

func (g_ds *DateSpan) Upper() time.Time {
	return DateADTToDate(C.datespan_upper(g_ds._inner))
}

func (g_ds *DateSpan) Duration() timeutil.Timedelta {
	interval := C.datespan_duration(g_ds._inner)
	microsecond := int(interval.time)
	day := int(interval.day)
	month := int(interval.month)
	dr := timeutil.Timedelta{
		Microseconds: time.Duration(microsecond),
		Days:         time.Duration(day) + time.Duration(month*30),
	}
	return dr
}

// ------------------------- Transformations -------------------------------
func (g_ds *DateSpan) ShiftScale(shift interface{}, duration interface{}) (*DateSpan, error) {
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
		return &DateSpan{_inner: nil}, fmt.Errorf("operation not supported with type %T", shift)
	}

	switch d := duration.(type) {
	case timeutil.Timedelta:
		duration_in = int(d.Days)
	case int:
		duration_in = d
	case nil:
		duration_in = 0
	default:
		return &DateSpan{_inner: nil}, fmt.Errorf("operation not supported with type %T", duration)
	}

	tss := C.datespan_shift_scale(g_ds._inner, C.int(shift_in), C.int(duration_in), shift_in != 0, duration_in != 0)
	return &DateSpan{_inner: tss}, nil
}

func (g_ds *DateSpan) Shift(delta interface{}) (*DateSpan, error) {
	return g_ds.ShiftScale(delta, 0)
}

func (g_ds *DateSpan) Scale(duration interface{}) (*DateSpan, error) {
	return g_ds.ShiftScale(0, duration)
}

// ------------------------- Topological Operations ------------------------
func (g_ds *DateSpan) IsAdjacent(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return bool(C.adjacent_span_date(g_ds._inner, DateToDateADT(o))), nil
	case *DateSpan:
		return bool(C.adjacent_span_span(g_ds._inner, o._inner)), nil
	case *DateSpanSet:
		return bool(C.adjacent_spanset_span(o._inner, g_ds._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_ds *DateSpan) Contains(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return bool(C.contains_span_date(g_ds._inner, DateToDateADT(o))), nil
	case *DateSet:
		return bool(C.contains_span_spanset(g_ds._inner, o.ToSpanSet()._inner)), nil
	case *DateSpan:
		return bool(C.contains_span_span(g_ds._inner, o._inner)), nil
	case *DateSpanSet:
		return bool(C.contains_span_spanset(g_ds._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_ds *DateSpan) Overlaps(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return g_ds.Contains(other)
	case *DateSet:
		return bool(C.overlaps_span_spanset(g_ds._inner, o.ToSpanSet()._inner)), nil
	case *DateSpan:
		return bool(C.overlaps_span_span(g_ds._inner, o._inner)), nil
	case *DateSpanSet:
		return bool(C.overlaps_span_spanset(g_ds._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

// ------------------------- Position Operations ---------------------------
func (g_ds *DateSpan) IsLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return bool(C.before_span_date(g_ds._inner, DateToDateADT(o))), nil
	case *DateSpan:
		return bool(C.left_span_span(g_ds._inner, o._inner)), nil
	case *DateSpanSet:
		return bool(C.left_span_spanset(g_ds._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_ds *DateSpan) IsOverOrLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return bool(C.overbefore_span_date(g_ds._inner, DateToDateADT(o))), nil
	case *DateSpan:
		return bool(C.overleft_span_span(g_ds._inner, o._inner)), nil
	case *DateSpanSet:
		return bool(C.overleft_span_spanset(g_ds._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_ds *DateSpan) IsRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return bool(C.after_span_date(g_ds._inner, DateToDateADT(o))), nil
	case *DateSpan:
		return bool(C.right_span_span(g_ds._inner, o._inner)), nil
	case *DateSpanSet:
		return bool(C.right_span_spanset(g_ds._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_ds *DateSpan) IsOverOrRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return bool(C.overafter_span_date(g_ds._inner, DateToDateADT(o))), nil
	case *DateSpan:
		return bool(C.overright_span_span(g_ds._inner, o._inner)), nil
	case *DateSpanSet:
		return bool(C.overright_span_spanset(g_ds._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

// ------------------------- Distance Operations ---------------------------
func (g_ds *DateSpan) Distance(other interface{}) (timeutil.Timedelta, error) {
	switch o := other.(type) {
	case time.Time:
		days := int(C.distance_span_date(g_ds._inner, DateToDateADT(o)))
		return timeutil.Timedelta{Days: time.Duration(days)}, nil
	case *DateSet:
		return g_ds.Distance(o.ToSpanSet())
	case *DateSpan:
		days := int(C.distance_datespan_datespan(g_ds._inner, o._inner))
		return timeutil.Timedelta{Days: time.Duration(days)}, nil
	case *DateSpanSet:
		days := int(C.distance_datespanset_datespan(o._inner, g_ds._inner))
		return timeutil.Timedelta{Days: time.Duration(days)}, nil
	default:
		return timeutil.Timedelta{}, fmt.Errorf("operation not supported with type %T", other)
	}
}

// ------------------------- Set Operations --------------------------------
func (g_ds *DateSpan) Intersection(other interface{}) (Dates, error) {
	switch o := other.(type) {
	case time.Time:
		res := C.intersection_span_date(g_ds._inner, DateToDateADT(o))
		return &DateSpan{_inner: res}, nil
	case *DateSet:
		return g_ds.Intersection(o.ToSpanSet())
	case *DateSpan:
		res := C.intersection_span_span(g_ds._inner, o._inner)
		return &DateSpan{_inner: res}, nil
	case *DateSpanSet:
		res := C.intersection_spanset_span(o._inner, g_ds._inner)
		return &DateSpanSet{_inner: res}, nil
	default:
		return &DateSpan{_inner: nil}, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_ds *DateSpan) Minus(other interface{}) (*DateSpanSet, error) {
	switch o := other.(type) {
	case time.Time:
		res := C.minus_span_date(g_ds._inner, DateToDateADT(o))
		return &DateSpanSet{_inner: res}, nil
	case *DateSet:
		return g_ds.Minus(o.ToSpanSet())
	case *DateSpan:
		res := C.minus_span_span(g_ds._inner, o._inner)
		return &DateSpanSet{_inner: res}, nil
	case *DateSpanSet:
		res := C.minus_span_spanset(g_ds._inner, o._inner)
		return &DateSpanSet{_inner: res}, nil
	default:
		return &DateSpanSet{_inner: nil}, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_ds *DateSpan) Union(other interface{}) (*DateSpanSet, error) {
	switch o := other.(type) {
	case time.Time:
		res := C.gunion_span_date(g_ds._inner, DateToDateADT(o))
		return &DateSpanSet{_inner: res}, nil
	case *DateSet:
		return g_ds.Union(o.ToSpanSet())
	case *DateSpan:
		res := C.gunion_span_span(g_ds._inner, o._inner)
		return &DateSpanSet{_inner: res}, nil
	case *DateSpanSet:
		res := C.gunion_span_spanset(g_ds._inner, o._inner)
		return &DateSpanSet{_inner: res}, nil
	default:
		return &DateSpanSet{_inner: nil}, fmt.Errorf("operation not supported with type %T", other)
	}
}
