package gomeos

/*
#include "meos.h"
#include "meos_catalog.h"
#include <stdio.h>
#include <stdlib.h>
#define gunion_set_set union_set_set
#define gunion_set_date union_set_date
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"

	"github.com/leekchan/timeutil"
)

type DateSet struct {
	_inner *C.Set
}

// ------------------------- Input ----------------------------------------
func NewDateSet(g_ds_in string) *DateSet {
	c_ds_in := C.CString(g_ds_in)
	defer C.free(unsafe.Pointer(c_ds_in))
	c_ds := C.dateset_in(c_ds_in)
	g_ds := &DateSet{_inner: c_ds}
	return g_ds
}

func FalseDateSet() *DateSet {
	c_ds_in := C.CString("{0001-01-01}")
	defer C.free(unsafe.Pointer(c_ds_in))
	c_ds := C.dateset_in(c_ds_in)
	g_ds := &DateSet{_inner: c_ds}
	return g_ds
}

// ------------------------- Output ----------------------------------------
func (g_ds *DateSet) DateSetOut() string {
	if g_ds._inner == nil {
		return "Nil"
	} else {
		c_ds_out := C.dateset_out(g_ds._inner)
		defer C.free(unsafe.Pointer(c_ds_out))
		g_ds_out := C.GoString(c_ds_out)
		return g_ds_out
	}
}

// TODO: remove library duplicate(fix the warnings)
// ------------------------- Conversions -----------------------------------
func (g_ds *DateSet) ToSpan() *DateSpan {
	c_ds := C.set_to_span(g_ds._inner)
	return &DateSpan{_inner: c_ds}
}

func (g_ds *DateSet) ToSpanSet() *DateSpanSet {
	c_ds := C.set_to_spanset(g_ds._inner)
	return &DateSpanSet{_inner: c_ds}
}

// ------------------------- Accessors -------------------------------------

func (g_ds *DateSet) Duration() timeutil.Timedelta {
	return g_ds.ToSpan().Duration()
}

func (g_ds *DateSet) StartElement() time.Time {
	parsedDate := DateADTToDate(C.dateset_start_value(g_ds._inner))
	return parsedDate
}

func (g_ds *DateSet) EndElement() time.Time {
	parsedDate := DateADTToDate(C.dateset_end_value(g_ds._inner))
	return parsedDate
}

func (g_ds *DateSet) ElementN(n int) time.Time {
	res := C.malloc(C.sizeof_int)
	defer C.free(unsafe.Pointer(res)) // Ensure memory is freed.
	success := C.dateset_value_n(g_ds._inner, C.int(n+1), (*C.DateADT)(res))
	if success {
		result := *(*C.DateADT)(res)
		return DateADTToDate(result)
	} else {
		return time.Time{}
	}
}

func (g_ds *DateSet) NumElements() int {
	return int(C.set_num_values(g_ds._inner))
}

func (g_ds *DateSet) Elements() []time.Time {
	nums := g_ds.NumElements()
	dates := make([]time.Time, nums)
	for i := 0; i < nums; i++ {
		dates[i] = g_ds.ElementN(i)
	}
	return dates
}

// ------------------------- Transformations -------------------------------
func (g_ds *DateSet) ShiftScale(shift interface{}, duration interface{}) (*DateSet, error) {
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
		return FalseDateSet(), fmt.Errorf("operation not supported with type %T", shift)
	}

	switch d := duration.(type) {
	case timeutil.Timedelta:
		duration_in = int(d.Days)
	case int:
		duration_in = d
	case nil:
		duration_in = 0
	default:
		return FalseDateSet(), fmt.Errorf("operation not supported with type %T", duration)
	}

	tss := C.dateset_shift_scale(g_ds._inner, C.int(shift_in), C.int(duration_in), shift_in != 0, duration_in != 0)
	return &DateSet{_inner: tss}, nil
}

func (g_ds *DateSet) Shift(delta interface{}) (*DateSet, error) {
	return g_ds.ShiftScale(delta, 0)
}

func (g_ds *DateSet) Scale(duration interface{}) (*DateSet, error) {
	return g_ds.ShiftScale(0, duration)
}

// ------------------------- Topological Operations ------------------------
func (g_ds *DateSet) Contains(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return bool(C.contains_set_date(g_ds._inner, DateToDateADT(o))), nil
	case *DateSet:
		return bool(C.contains_set_set(g_ds._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_ds *DateSet) Overlaps(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return bool(C.contains_set_date(g_ds._inner, DateToDateADT(o))), nil
	case *DateSet:
		return bool(C.overlaps_set_set(g_ds._inner, o._inner)), nil
	case *DateSpan:
		return g_ds.ToSpan().IsAdjacent(o)
	case *DateSpanSet:
		return g_ds.ToSpanSet().IsAdjacent(o)
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

// ------------------------- Position Operations ---------------------------
func (g_ds *DateSet) IsLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return bool(C.before_set_date(g_ds._inner, DateToDateADT(o))), nil
	case *DateSet:
		return bool(C.overlaps_set_set(g_ds._inner, o._inner)), nil
	case *DateSpan:
		return g_ds.ToSpan().IsLeft(o)
	case *DateSpanSet:
		return g_ds.ToSpan().IsLeft(o)
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_ds *DateSet) IsOverOrLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return bool(C.overbefore_set_date(g_ds._inner, DateToDateADT(o))), nil
	case *DateSet:
		return bool(C.overleft_set_set(g_ds._inner, o._inner)), nil
	case *DateSpan:
		return g_ds.ToSpan().IsOverOrLeft(o)
	case *DateSpanSet:
		return g_ds.ToSpan().IsOverOrLeft(o)
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_ds *DateSet) IsRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return bool(C.after_set_date(g_ds._inner, DateToDateADT(o))), nil
	case *DateSet:
		return bool(C.right_set_set(g_ds._inner, o._inner)), nil
	case *DateSpan:
		return g_ds.ToSpan().IsRight(o)
	case *DateSpanSet:
		return g_ds.ToSpan().IsRight(o)
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_ds *DateSet) IsOverOrRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case time.Time:
		return bool(C.overafter_set_date(g_ds._inner, DateToDateADT(o))), nil
	case *DateSet:
		return bool(C.overright_set_set(g_ds._inner, o._inner)), nil
	case *DateSpan:
		return g_ds.ToSpan().IsOverOrRight(o)
	case *DateSpanSet:
		return g_ds.ToSpan().IsOverOrRight(o)
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

// ------------------------- Distance Operations ---------------------------
func (g_ds *DateSet) Distance(other interface{}) (timeutil.Timedelta, error) {
	switch o := other.(type) {
	case time.Time:
		days := int(C.distance_set_date(g_ds._inner, DateToDateADT(o)))
		return timeutil.Timedelta{Days: time.Duration(days)}, nil
	case *DateSet:
		days := int(C.distance_dateset_dateset(g_ds._inner, o._inner))
		return timeutil.Timedelta{Days: time.Duration(days)}, nil
	case *DateSpan:
		return g_ds.ToSpanSet().Distance(other)
	case *DateSpanSet:
		return g_ds.ToSpanSet().Distance(other)
	default:
		return timeutil.Timedelta{}, fmt.Errorf("operation not supported with type %T", other)
	}
}

// ------------------------- Set Operations --------------------------------
func (g_ds *DateSet) intersection(other interface{}) (Dates, error) {
	switch o := other.(type) {
	case time.Time:
		res := C.intersection_set_date(g_ds._inner, DateToDateADT(o))
		return &DateSet{_inner: res}, nil
	case *DateSet:
		res := C.intersection_set_set(g_ds._inner, o._inner)
		return &DateSet{_inner: res}, nil
	case *DateSpan:
		return g_ds.ToSpanSet().Intersection(o)
	case *DateSpanSet:
		return g_ds.ToSpanSet().Intersection(o)
	default:
		return &DateSet{_inner: nil}, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_ds *DateSet) Intersection(other interface{}) (interface{}, error) {
	switch o := other.(type) {
	case time.Time:
		res := C.intersection_set_date(g_ds._inner, DateToDateADT(o))
		return &DateSet{_inner: res}, nil
	case *DateSet:
		res := C.intersection_set_set(g_ds._inner, o._inner)
		return &DateSet{_inner: res}, nil
	case *DateSpan:
		return g_ds.ToSpanSet().Intersection(o)
	case *DateSpanSet:
		return g_ds.ToSpanSet().Intersection(o)
	default:
		return &DateSet{_inner: nil}, fmt.Errorf("operation not supported with type %T", other)
	}
}

// type DatesGenerics interface {
// 	*DateSet | *DateSpan | *DateSpanSet
// }

// // IntersectionAndUnwrap combines the Intersection method and type assertion into one function.
// func Intersection[T DatesGenerics](g_ds *DateSet, other interface{}) (T, error) {
// 	var output T

// 	// Perform the intersection based on the type of `other`
// 	var dates Dates
// 	switch o := other.(type) {
// 	case time.Time:
// 		res := C.intersection_set_date(g_ds._inner, DateToDateADT(o))
// 		dates = &DateSet{_inner: res}
// 	case *DateSet:
// 		res := C.intersection_set_set(g_ds._inner, o._inner)
// 		dates = &DateSet{_inner: res}
// 	case *DateSpan:
// 		// Perform the intersection with DateSpan and DateSet
// 		spanSet, err := g_ds.ToSpanSet().Intersection(o)
// 		if err != nil {
// 			return output, err
// 		}
// 		dates = spanSet
// 	case *DateSpanSet:
// 		// Perform the intersection with DateSpanSet
// 		spanSet, err := g_ds.ToSpanSet().Intersection(o)
// 		if err != nil {
// 			return output, err
// 		}
// 		dates = spanSet
// 	default:
// 		return output, fmt.Errorf("operation not supported with type %T", other)
// 	}

// 	// Assert the type to the requested concrete type
// 	if result, ok := dates.(T); ok {
// 		return result, nil
// 	}

// 	return output, fmt.Errorf("unexpected type: %T", dates)
// }

// func (g_ds *DateSet) Mul(other interface{}) (Dates, error) {
// 	return g_ds.Intersection(other)
// }

func (g_ds *DateSet) Minus(other interface{}) (Dates, error) {
	switch o := other.(type) {
	case time.Time:
		res := C.minus_set_date(g_ds._inner, DateToDateADT(o))
		return &DateSet{_inner: res}, nil
	case *DateSet:
		res := C.minus_set_set(g_ds._inner, o._inner)
		return &DateSet{_inner: res}, nil
	case *DateSpan:
		return g_ds.ToSpanSet().Minus(o)
	case *DateSpanSet:
		return g_ds.ToSpanSet().Minus(o)
	default:
		return &DateSet{_inner: nil}, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_ds *DateSet) Sub(other interface{}) (Dates, error) {
	return g_ds.Minus(other)
}

func (g_ds *DateSet) Union(other interface{}) (Dates, error) {
	switch o := other.(type) {
	case time.Time:
		res := C.gunion_set_date(g_ds._inner, DateToDateADT(o))
		return &DateSet{_inner: res}, nil
	case *DateSet:
		res := C.gunion_set_set(g_ds._inner, o._inner)
		return &DateSet{_inner: res}, nil
	case *DateSpan:
		return g_ds.ToSpanSet().Union(o)
	case *DateSpanSet:
		return g_ds.ToSpanSet().Union(o)
	default:
		return &DateSet{_inner: nil}, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_ds *DateSet) Add(other interface{}) (Dates, error) {
	return g_ds.Union(other)
}
