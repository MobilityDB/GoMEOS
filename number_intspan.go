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

type IntSpan struct {
	_inner *C.Span
}

func (g_fs *IntSpan) Inner() *C.Span {
	return g_fs._inner
}

func (g_fs *IntSpan) Init(c_span *C.Span) {
	g_fs._inner = c_span
}
func (g_fs *IntSpan) IsNumSpan() bool { return true }

func NewIntSpan(g_is_in string) *IntSpan {
	c_is_in := C.CString(g_is_in)
	defer C.free(unsafe.Pointer(c_is_in))
	c_is := C.intspan_in(c_is_in)
	g_is := &IntSpan{_inner: c_is}
	return g_is
}

func (g_is *IntSpan) IntSpanOut() string {
	c_is_out := C.intspan_out(g_is._inner)
	defer C.free(unsafe.Pointer(c_is_out))
	g_is_out := C.GoString(c_is_out)
	return g_is_out
}

func (g_is *IntSpan) ToSpanSet() IntSpanSet {
	return IntSpanSet{_inner: C.span_to_spanset(g_is._inner)}
}

func (g_is *IntSpan) ToFloatSpan() FloatSpan {
	return FloatSpan{_inner: C.intspan_to_floatspan(g_is._inner)}
}

func (g_is *IntSpan) Lower() int {
	return int(C.intspan_lower(g_is._inner))
}

func (g_is *IntSpan) Upper() int {
	return int(C.intspan_upper(g_is._inner))
}

func (g_is *IntSpan) Width() float32 {
	return float32(C.intspan_width(g_is._inner))
}

func (g_is *IntSpan) ShiftScale(d int, w int) IntSpan {
	modified := C.intspan_shift_scale(g_is._inner, C.int(d), C.int(w), C._Bool(d != 0), C._Bool(w != 0))
	return IntSpan{_inner: modified}
}

func (g_is *IntSpan) Shift(delta int) IntSpan {
	return g_is.ShiftScale(delta, 0)

}

func (g_is *IntSpan) Scale(width int) IntSpan {
	return g_is.ShiftScale(0, width)
}

func (g_is *IntSpan) IsAdjacent(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.adjacent_span_int(g_is._inner, C.int(o))), nil
	case *IntSpan:
		return bool(C.adjacent_span_span(g_is._inner, o._inner)), nil
	case *IntSpanSet:
		return bool(C.adjacent_spanset_span(o._inner, g_is._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_is *IntSpan) Contains(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.contains_span_int(g_is._inner, C.int(o))), nil
	case *IntSpan:
		return bool(C.contains_span_span(g_is._inner, o._inner)), nil
	case *IntSpanSet:
		return bool(C.contains_span_spanset(g_is._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_is *IntSpan) IsSame(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.span_eq(g_is._inner, C.int_to_span(C.int(o)))), nil
	case *IntSpan:
		return bool(C.span_eq(g_is._inner, o._inner)), nil
	case *IntSpanSet:
		return bool(C.span_eq(g_is._inner, o.ToSpan()._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_is *IntSpan) IsLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.left_span_int(g_is._inner, C.int(o))), nil
	case *IntSpan:
		return bool(C.left_span_span(g_is._inner, o._inner)), nil
	case *IntSpanSet:
		return bool(C.left_span_spanset(g_is._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_is *IntSpan) IsOverOrLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.overleft_span_int(g_is._inner, C.int(o))), nil
	case *IntSpan:
		return bool(C.overleft_span_span(g_is._inner, o._inner)), nil
	case *IntSpanSet:
		return bool(C.overleft_span_spanset(g_is._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_is *IntSpan) IsRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.right_span_int(g_is._inner, C.int(o))), nil
	case *IntSpan:
		return bool(C.right_span_span(g_is._inner, o._inner)), nil
	case *IntSpanSet:
		return bool(C.right_span_spanset(g_is._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_is *IntSpan) IsOverOrRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.overright_span_int(g_is._inner, C.int(o))), nil
	case *IntSpan:
		return bool(C.overright_span_span(g_is._inner, o._inner)), nil
	case *IntSpanSet:
		return bool(C.overright_span_spanset(g_is._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_is *IntSpan) Distance(other interface{}) (int, error) {
	switch o := other.(type) {
	case int:
		return int(C.distance_span_int(g_is._inner, C.int(o))), nil
	case *IntSet:
		return int(C.distance_intspanset_intspan(o.ToSpanSet()._inner, g_is._inner)), nil
	case *IntSpan:
		return int(C.distance_intspan_intspan(g_is._inner, o._inner)), nil
	case *IntSpanSet:
		return int(C.distance_intspanset_intspan(o._inner, g_is._inner)), nil
	default:
		return 0, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_is *IntSpan) Intersection(other interface{}) (*IntSpan, error) {
	switch o := other.(type) {
	case int:
		res := C.intersection_span_int(g_is._inner, C.int(o))
		if res == nil {
			return nil, nil
		} else {
			return &IntSpan{_inner: res}, nil
		}
	case *IntSpan:
		res := C.intersection_span_span(g_is._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &IntSpan{_inner: res}, nil
		}
	case *IntSpanSet:
		res := C.intersection_spanset_span(o._inner, g_is._inner)
		if res == nil {
			return nil, nil
		} else {
			// In Pymeos, it is IntSpanSet, but in Go, it's IntSpan because I need to declare the type of output
			iss := IntSpanSet{_inner: res}
			is := iss.ToSpan()
			return &is, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_is *IntSpan) Mul(other interface{}) (*IntSpan, error) {
	return g_is.Intersection(other)
}

func (g_is *IntSpan) Minus(other interface{}) (*IntSpanSet, error) {
	switch o := other.(type) {
	case int:
		res := C.minus_span_int(g_is._inner, C.int(o))
		if res == nil {
			return nil, nil
		} else {
			return &IntSpanSet{_inner: res}, nil
		}
	case *IntSpan:
		res := C.minus_span_span(g_is._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &IntSpanSet{_inner: res}, nil
		}
	case *IntSpanSet:
		res := C.minus_spanset_span(o._inner, g_is._inner)
		if res == nil {
			return nil, nil
		} else {
			return &IntSpanSet{_inner: res}, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_is *IntSpan) Sub(other interface{}) (*IntSpanSet, error) {
	return g_is.Minus(other)
}

func (g_is *IntSpan) Union(other interface{}) (*IntSpanSet, error) {
	switch o := other.(type) {
	case int:
		res := C.gunion_span_int(g_is._inner, C.int(o))
		if res == nil {
			return nil, nil
		} else {
			return &IntSpanSet{_inner: res}, nil
		}
	case *IntSpan:
		res := C.gunion_span_span(g_is._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &IntSpanSet{_inner: res}, nil
		}
	case *IntSpanSet:
		res := C.gunion_spanset_span(o._inner, g_is._inner)
		if res == nil {
			return nil, nil
		} else {
			return &IntSpanSet{_inner: res}, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_is *IntSpan) Add(other interface{}) (*IntSpanSet, error) {
	return g_is.Union(other)
}
