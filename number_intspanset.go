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
	"unsafe"
)

type IntSpanSet struct {
	_inner *C.SpanSet
}

func NewIntSpanSet(g_iss_in string) *IntSpanSet {
	c_iss_in := C.CString(g_iss_in)
	defer C.free(unsafe.Pointer(c_iss_in))
	c_iss := C.intspanset_in(c_iss_in)
	g_iss := &IntSpanSet{_inner: c_iss}
	return g_iss
}

/*
Return the string representation of the content of IntSpanSet.

Returns:

	String

MEOS Functions:

	intspanset_out
*/
func (g_iss *IntSpanSet) IntSpanSetOut() string {
	c_iss_out := C.intspanset_out(g_iss._inner)
	defer C.free(unsafe.Pointer(c_iss_out))
	g_iss_out := C.GoString(c_iss_out)
	return g_iss_out
}

/*
Returns a span that encompasses _inner.

Returns:

	A new struct `IntSpan` instance

MEOS Functions:

	spanset_span
*/
func (g_iss IntSpanSet) ToSpan() IntSpan {
	return IntSpan{_inner: C.spanset_span(g_iss._inner)}
}

/*
Converts IntSpanSet to a FloatSpanSet instance.

Returns:

	A new FloatSpanSet instance

MEOS Functions:

	intspanset_to_floatspanset
*/
func (g_iss IntSpanSet) ToFloatSpanSet() FloatSpanSet {
	return FloatSpanSet{_inner: C.intspanset_to_floatspanset(g_iss._inner)}
}

/*
Returns the number of spans in IntSpanSet.
Returns:

	An int

MEOS Functions:

	spanset_num_spans
*/
func (g_iss IntSpanSet) NumSpans() int {
	return int(C.spanset_num_spans(g_iss._inner))
}

/*
Returns the width of the spanset. By default, i.e., when the second
argument is False, the function takes into account the gaps within,
i.e., returns the sum of the widths of the spans within.
Otherwise, the function returns the width of the spanset ignoring
any gap, i.e., the width from the lower bound of the first span to
the upper bound of the last span.

Parameters:

	ignore_gaps: Whether to take into account potential gaps in
	the spanset.

Returns:

	A `float` representing the duration of the spanset

MEOS Functions:

	intspanset_width
*/
func (g_iss IntSpanSet) Width(ignore_gap bool) int {
	ig := C._Bool(ignore_gap)
	return int(C.intspanset_width(g_iss._inner, ig))
}

/*
Returns the first span in spanset.
Returns:

	A IntSpan instance

MEOS Functions:

	spanset_start_span
*/
func (g_iss IntSpanSet) StartSpan() IntSpan {
	return IntSpan{_inner: C.spanset_start_span(g_iss._inner)}
}

/*
Returns the last span in IntSpanSet.
Returns:

	A IntSpan instance

MEOS Functions:

	spanset_end_span
*/
func (g_iss IntSpanSet) EndSpan() IntSpan {
	return IntSpan{_inner: C.spanset_end_span(g_iss._inner)}
}

/*
Returns the n-th span in IntSpanSet.
Returns:

	A IntSpan instance

MEOS Functions:

	spanset_span_n
*/
func (g_iss IntSpanSet) SpanN(n int) IntSpan {
	return IntSpan{_inner: C.spanset_span_n(g_iss._inner, C.int(n+1))}
}

/*
Returns the list of spans in IntSpanSet.
Returns:

	A IntSpan instance

MEOS Functions:

	spanset_spans
*/
func (g_iss IntSpanSet) Spans() []IntSpan {
	nums := g_iss.NumSpans()
	spans := make([]IntSpan, nums)
	for i := 0; i < nums; i++ {
		spans[i] = g_iss.SpanN(i)
	}
	return spans
}

/*
Return a new “IntSpanSet“ with the lower and upper bounds shifted by
“delta“.

Args:

	delta: The value to shift by

Returns:

	A new ``IntSpanSet`` instance

MEOS Functions:

	intspanset_shift_scale
*/
func (g_iss IntSpanSet) ShiftScale(d int, w int) IntSpanSet {
	modified := C.intspanset_shift_scale(g_iss._inner, C.int(d), C.int(w), C._Bool(d != 0), C._Bool(w != 0))
	return IntSpanSet{_inner: modified}
}

/*
Return a new “IntSpanSet“ with the lower and upper bounds shifted by
“delta“.

Args:

	delta: The value to shift by

Returns:

	A new ``IntSpanSet`` instance

MEOS Functions:

	intspanset_shift_scale
*/
func (g_iss IntSpanSet) Shift(delta int) IntSpanSet {
	return g_iss.ShiftScale(delta, 0)

}

/*
Return a new “IntSpanSet“ with the lower and upper bounds scaled so
that the width is “width“.

Args:

	width: The new width

Returns:

	A new ``IntSpanSet`` instance

MEOS Functions:

	intspanset_shift_scale
*/
func (g_iss IntSpanSet) Scale(width int) IntSpanSet {
	return g_iss.ShiftScale(0, width)
}

func (g_iss *IntSpanSet) IsAdjacent(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.adjacent_spanset_int(g_iss._inner, C.int(o))), nil
	case *IntSpan:
		return bool(C.adjacent_spanset_span(g_iss._inner, o._inner)), nil
	case *IntSpanSet:
		return bool(C.adjacent_spanset_spanset(g_iss._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_iss *IntSpanSet) Contains(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.contains_spanset_int(g_iss._inner, C.int(o))), nil
	case *IntSpan:
		return bool(C.contains_spanset_span(g_iss._inner, o._inner)), nil
	case *IntSpanSet:
		return bool(C.contains_spanset_spanset(g_iss._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_iss *IntSpanSet) IsSame(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.spanset_eq(g_iss._inner, C.int_to_spanset(C.int(o)))), nil
	case *IntSpan:
		return bool(C.span_eq(g_iss.ToSpan()._inner, o._inner)), nil
	case *IntSpanSet:
		return bool(C.span_eq(g_iss.ToSpan()._inner, o.ToSpan()._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_iss *IntSpanSet) IsLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.left_spanset_int(g_iss._inner, C.int(o))), nil
	case *IntSpan:
		return bool(C.left_spanset_span(g_iss._inner, o._inner)), nil
	case *IntSpanSet:
		return bool(C.left_spanset_spanset(g_iss._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_iss *IntSpanSet) IsOverOrLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.overleft_spanset_int(g_iss._inner, C.int(o))), nil
	case *IntSpan:
		return bool(C.overleft_spanset_span(g_iss._inner, o._inner)), nil
	case *IntSpanSet:
		return bool(C.overleft_spanset_spanset(g_iss._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_iss *IntSpanSet) IsRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.right_spanset_int(g_iss._inner, C.int(o))), nil
	case *IntSpan:
		return bool(C.right_spanset_span(g_iss._inner, o._inner)), nil
	case *IntSpanSet:
		return bool(C.right_spanset_spanset(g_iss._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_iss *IntSpanSet) IsOverOrRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.overright_spanset_int(g_iss._inner, C.int(o))), nil
	case *IntSpan:
		return bool(C.overright_spanset_span(g_iss._inner, o._inner)), nil
	case *IntSpanSet:
		return bool(C.overright_spanset_spanset(g_iss._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_iss *IntSpanSet) Distance(other interface{}) (int, error) {
	switch o := other.(type) {
	case int:
		return int(C.distance_spanset_int(g_iss._inner, C.int(o))), nil
	case *IntSet:
		return int(C.distance_intspanset_intspanset(g_iss._inner, o.ToSpanSet()._inner)), nil
	case *IntSpan:
		return int(C.distance_intspanset_intspan(g_iss._inner, o._inner)), nil
	case *IntSpanSet:
		return int(C.distance_intspanset_intspanset(g_iss._inner, o._inner)), nil
	default:
		return 0, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_iss *IntSpanSet) Intersection(other interface{}) (*IntSpanSet, error) {
	switch o := other.(type) {
	case int:
		res := C.intersection_spanset_int(g_iss._inner, C.int(o))
		if res == nil {
			return nil, nil
		} else {
			return &IntSpanSet{_inner: res}, nil
		}
	case *IntSpan:
		res := C.intersection_spanset_span(g_iss._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &IntSpanSet{_inner: res}, nil
		}
	case *IntSpanSet:
		res := C.intersection_spanset_spanset(g_iss._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &IntSpanSet{_inner: res}, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_iss *IntSpanSet) Mul(other interface{}) (*IntSpanSet, error) {
	return g_iss.Intersection(other)
}

func (g_iss *IntSpanSet) Minus(other interface{}) (*IntSpanSet, error) {
	switch o := other.(type) {
	case int:
		res := C.minus_spanset_int(g_iss._inner, C.int(o))
		if res == nil {
			return nil, nil
		} else {
			return &IntSpanSet{_inner: res}, nil
		}
	case *IntSpan:
		res := C.minus_spanset_span(g_iss._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &IntSpanSet{_inner: res}, nil
		}
	case *IntSpanSet:
		res := C.minus_spanset_spanset(g_iss._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &IntSpanSet{_inner: res}, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_iss *IntSpanSet) Sub(other interface{}) (*IntSpanSet, error) {
	return g_iss.Minus(other)
}

func (g_iss *IntSpanSet) Union(other interface{}) (*IntSpanSet, error) {
	switch o := other.(type) {
	case int:
		res := C.gunion_spanset_int(g_iss._inner, C.int(o))
		if res == nil {
			return nil, nil
		} else {
			return &IntSpanSet{_inner: res}, nil
		}
	case *IntSpan:
		res := C.gunion_spanset_span(g_iss._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &IntSpanSet{_inner: res}, nil
		}
	case *IntSpanSet:
		res := C.gunion_spanset_spanset(g_iss._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &IntSpanSet{_inner: res}, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_iss *IntSpanSet) Add(other interface{}) (*IntSpanSet, error) {
	return g_iss.Union(other)
}
