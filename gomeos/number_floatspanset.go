// collections/number/floatspanset.go
package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#define gunion_spanset_float union_spanset_float
#define gunion_spanset_span union_spanset_span
#define gunion_spanset_spanset union_spanset_spanset
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type FloatSpanSet struct {
	_inner *C.SpanSet
}

// ------------------------- Input ----------------------------------------
func NewFloatSpanSet(g_fss_in string) *FloatSpanSet {
	c_fss_in := C.CString(g_fss_in)
	defer C.free(unsafe.Pointer(c_fss_in))
	c_fss := C.floatspanset_in(c_fss_in)
	g_fss := &FloatSpanSet{_inner: c_fss}
	return g_fss
}

// ------------------------- Output ----------------------------------------
/*
Return the string representation of the content of FloatSpanSet.

Returns:
	String

MEOS Functions:
	floatspanset_out
*/
func (g_fss *FloatSpanSet) FloatSpanSetOut(max_decimal int) string {
	c_fss_out := C.floatspanset_out(g_fss._inner, C.int(max_decimal))
	defer C.free(unsafe.Pointer(c_fss_out))
	g_fss_out := C.GoString(c_fss_out)
	return g_fss_out
}

// ------------------------- Conversions -----------------------------------
/*
Returns a span that encompasses _inner.

Returns:
	A new struct `FloatSpan` instance

MEOS Functions:
	spanset_span
*/
func (g_fss FloatSpanSet) ToSpan() FloatSpan {
	return FloatSpan{_inner: C.spanset_span(g_fss._inner)}
}

/*
Converts FloatSpanSet to an IntSpanSet instance.

Returns:

	A new IntSpanSet instance

MEOS Functions:

	floatspanset_to_intspanset
*/
func (g_fss FloatSpanSet) ToIntSpanSet() IntSpanSet {
	return IntSpanSet{_inner: C.floatspanset_to_intspanset(g_fss._inner)}
}

// ------------------------- Accessors -------------------------------------
/*
Returns the number of spans in FloatSpanSet.
Returns:
	An int

MEOS Functions:
	spanset_num_spans
*/
func (g_fss FloatSpanSet) NumSpans() int {
	return int(C.spanset_num_spans(g_fss._inner))
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

	floatspanset_width
*/
func (g_fss FloatSpanSet) Width(ignore_gap bool) float64 {
	ig := C._Bool(ignore_gap)
	return float64(C.floatspanset_width(g_fss._inner, ig))
}

/*
Returns the first span in spanset.
Returns:

	A FloatSpan instance

MEOS Functions:

	spanset_start_span
*/
func (g_fss FloatSpanSet) StartSpan() FloatSpan {
	return FloatSpan{_inner: C.spanset_start_span(g_fss._inner)}
}

/*
Returns the last span in FloatSpanSet.
Returns:

	A FloatSpan instance

MEOS Functions:

	spanset_end_span
*/
func (g_fss FloatSpanSet) EndSpan() FloatSpan {
	return FloatSpan{_inner: C.spanset_end_span(g_fss._inner)}
}

/*
Returns the n-th span in FloatSpanSet.
Returns:

	A FloatSpan instance

MEOS Functions:

	spanset_span_n
*/
func (g_fss FloatSpanSet) SpanN(n int) FloatSpan {
	return FloatSpan{_inner: C.spanset_span_n(g_fss._inner, C.int(n))}
}

/*
Returns the list of spans in FloatSpanSet.
Returns:

	A FloatSpan instance

MEOS Functions:

	spanset_spans
*/
func (g_fss FloatSpanSet) Spans() []FloatSpan {
	nums := g_fss.NumSpans()
	spans := make([]FloatSpan, nums)
	for i := 1; i < nums+1; i++ {
		spans[i-1] = g_fss.SpanN(i)
	}
	return spans
}

// ------------------------- Transformations -------------------------------

/*
Return a new “FloatSpanSet“ with the lower and upper bounds shifted by
“delta“.

Args:

	delta: The value to shift by

Returns:

	A new ``FloatSpanSet`` instance

MEOS Functions:

	floatspanset_shift_scale
*/
func (g_fss FloatSpanSet) ShiftScale(d float64, w float64) FloatSpanSet {
	modified := C.floatspanset_shift_scale(g_fss._inner, C.double(d), C.double(w), C._Bool(d != 0), C._Bool(w != 0))
	return FloatSpanSet{_inner: modified}
}

/*
Return a new “FloatSpanSet“ with the lower and upper bounds shifted by
“delta“.

Args:

	delta: The value to shift by

Returns:

	A new ``FloatSpanSet`` instance

MEOS Functions:

	floatspanset_shift_scale
*/
func (g_fss FloatSpanSet) Shift(delta float64) FloatSpanSet {
	return g_fss.ShiftScale(delta, 0)
}

/*
Return a new “FloatSpanSet“ with the lower and upper bounds scaled so
that the width is “width“.

Args:

	width: The new width

Returns:

	A new ``FloatSpanSet`` instance

MEOS Functions:

	floatspanset_shift_scale
*/
func (g_fss FloatSpanSet) Scale(width float64) FloatSpanSet {
	return g_fss.ShiftScale(0, width)
}

// ------------------------- Topological Operations --------------------------------
func (g_fss *FloatSpanSet) IsAdjacent(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.adjacent_spanset_float(g_fss._inner, C.double(o))), nil
	case float64:
		return bool(C.adjacent_spanset_float(g_fss._inner, C.double(o))), nil
	case *FloatSpan:
		return bool(C.adjacent_spanset_span(g_fss._inner, o._inner)), nil
	case *FloatSpanSet:
		return bool(C.adjacent_spanset_spanset(g_fss._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fss *FloatSpanSet) Contains(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.contains_spanset_float(g_fss._inner, C.double(o))), nil
	case *FloatSpan:
		return bool(C.contains_spanset_span(g_fss._inner, o._inner)), nil
	case *FloatSpanSet:
		return bool(C.contains_spanset_spanset(g_fss._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fss *FloatSpanSet) IsSame(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.spanset_eq(g_fss._inner, C.float_to_spanset(C.double(o)))), nil
	case *FloatSpan:
		return bool(C.span_eq(g_fss.ToSpan()._inner, o._inner)), nil
	case *FloatSpanSet:
		return bool(C.span_eq(g_fss.ToSpan()._inner, o.ToSpan()._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

// ------------------------- Position Operations ---------------------------

func (g_fss *FloatSpanSet) IsLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.left_spanset_float(g_fss._inner, C.double(o))), nil
	case *FloatSpan:
		return bool(C.left_spanset_span(g_fss._inner, o._inner)), nil
	case *FloatSpanSet:
		return bool(C.left_spanset_spanset(g_fss._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fss *FloatSpanSet) IsOverOrLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.overleft_spanset_float(g_fss._inner, C.double(o))), nil
	case *FloatSpan:
		return bool(C.overleft_spanset_span(g_fss._inner, o._inner)), nil
	case *FloatSpanSet:
		return bool(C.overleft_spanset_spanset(g_fss._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fss *FloatSpanSet) IsRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.right_spanset_float(g_fss._inner, C.double(o))), nil
	case *FloatSpan:
		return bool(C.right_spanset_span(g_fss._inner, o._inner)), nil
	case *FloatSpanSet:
		return bool(C.right_spanset_spanset(g_fss._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fss *FloatSpanSet) IsOverOrRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.overright_spanset_float(g_fss._inner, C.double(o))), nil
	case *FloatSpan:
		return bool(C.overright_spanset_span(g_fss._inner, o._inner)), nil
	case *FloatSpanSet:
		return bool(C.overright_spanset_spanset(g_fss._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

// ———————–– Distance Operations —————————
func (g_fss *FloatSpanSet) Distance(other interface{}) (float32, error) {
	switch o := other.(type) {
	case int:
		return float32(C.distance_spanset_float(g_fss._inner, C.double(float32(o)))), nil
	case float32:
		return float32(C.distance_spanset_float(g_fss._inner, C.double(o))), nil
	case float64:
		return float32(C.distance_spanset_float(g_fss._inner, C.double(o))), nil
	case *FloatSet:
		return float32(C.distance_floatspanset_floatspanset(g_fss._inner, o.ToSpanSet()._inner)), nil
	case *FloatSpan:
		return float32(C.distance_floatspanset_floatspan(g_fss._inner, o._inner)), nil
	case *FloatSpanSet:
		return float32(C.distance_floatspanset_floatspanset(g_fss._inner, o._inner)), nil
	default:
		return 0, fmt.Errorf("operation not supported with type %T", other)
	}
}

// ———————–– Set Operations ––––––––––––––––
func (g_fss *FloatSpanSet) Intersection(other interface{}) (*FloatSpanSet, error) {
	switch o := other.(type) {
	case float64:
		res := C.intersection_spanset_float(g_fss._inner, C.double(o))
		if res == nil {
			return nil, nil
		} else {
			return &FloatSpanSet{_inner: res}, nil
		}
	case *FloatSpan:
		res := C.intersection_spanset_span(g_fss._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &FloatSpanSet{_inner: res}, nil
		}
	case *FloatSpanSet:
		res := C.intersection_spanset_spanset(g_fss._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &FloatSpanSet{_inner: res}, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fss *FloatSpanSet) Mul(other interface{}) (*FloatSpanSet, error) {
	return g_fss.Intersection(other)
}

func (g_fss *FloatSpanSet) Minus(other interface{}) (*FloatSpanSet, error) {
	switch o := other.(type) {
	case float64:
		res := C.minus_spanset_float(g_fss._inner, C.double(o))
		if res == nil {
			return nil, nil
		} else {
			return &FloatSpanSet{_inner: res}, nil
		}
	case *FloatSpan:
		res := C.minus_spanset_span(g_fss._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &FloatSpanSet{_inner: res}, nil
		}
	case *FloatSpanSet:
		res := C.minus_spanset_spanset(g_fss._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &FloatSpanSet{_inner: res}, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fss *FloatSpanSet) Sub(other interface{}) (*FloatSpanSet, error) {
	return g_fss.Minus(other)
}

func (g_fss *FloatSpanSet) Union(other interface{}) (*FloatSpanSet, error) {
	switch o := other.(type) {
	case float64:
		res := C.gunion_spanset_float(g_fss._inner, C.double(o))
		if res == nil {
			return nil, nil
		} else {
			return &FloatSpanSet{_inner: res}, nil
		}
	case *FloatSpan:
		res := C.gunion_spanset_span(g_fss._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &FloatSpanSet{_inner: res}, nil
		}
	case *FloatSpanSet:
		res := C.gunion_spanset_spanset(g_fss._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &FloatSpanSet{_inner: res}, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fss *FloatSpanSet) Add(other interface{}) (*FloatSpanSet, error) {
	return g_fss.Union(other)
}
