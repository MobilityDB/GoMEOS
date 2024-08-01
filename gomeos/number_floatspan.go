// collections/number/floatspan.go
package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#define gunion_span_float union_span_float
#define gunion_span_span union_span_span
#define gunion_spanset_span union_spanset_span
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type FloatSpan struct {
	_inner *C.Span
}

// ------------------------- Input ----------------------------------------
func NewFloatSpan(g_fs_in string) *FloatSpan {
	c_fs_in := C.CString(g_fs_in)
	defer C.free(unsafe.Pointer(c_fs_in))
	c_fs := C.floatspan_in(c_fs_in)
	g_fs := &FloatSpan{_inner: c_fs}
	return g_fs
}

// ------------------------- Output ----------------------------------------
func (g_fs FloatSpan) FloatSpanOut(max_decimal int) string {
	c_fs_out := C.floatspan_out(g_fs._inner, C.int(max_decimal))
	defer C.free(unsafe.Pointer(c_fs_out))
	g_fs_out := C.GoString(c_fs_out)
	return g_fs_out
}

// ------------------------- Conversions -----------------------------------
func (g_fs FloatSpan) ToSpanSet() FloatSpanSet {
	return FloatSpanSet{_inner: C.span_to_spanset(g_fs._inner)}
}

func (g_fs FloatSpan) ToIntSpan() IntSpan {
	return IntSpan{_inner: C.floatspan_to_intspan(g_fs._inner)}
}

// ------------------------- Accessors -------------------------------------
func (g_fs FloatSpan) Lower() float64 {
	return float64(C.floatspan_lower(g_fs._inner))
}

func (g_fs FloatSpan) Upper() float64 {
	return float64(C.floatspan_upper(g_fs._inner))
}

func (g_fs FloatSpan) Width() float32 {
	return float32(C.floatspan_width(g_fs._inner))
}

// ------------------------- Transformations -------------------------------
func (g_fs FloatSpan) ShiftScale(d float64, w float64) FloatSpan {
	modified := C.floatspan_shift_scale(g_fs._inner, C.double(d), C.double(w), C._Bool(d != 0), C._Bool(w != 0))
	return FloatSpan{_inner: modified}
}

func (g_fs FloatSpan) Shift(delta float64) FloatSpan {
	return g_fs.ShiftScale(delta, 0.0)
}

func (g_fs FloatSpan) Scale(width float64) FloatSpan {
	return g_fs.ShiftScale(0.0, width)
}

// ------------------------- Topological Operations --------------------------------
func (g_fs *FloatSpan) IsAdjacent(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.adjacent_span_float(g_fs._inner, C.double(o))), nil
	case *FloatSpan:
		return bool(C.adjacent_span_span(g_fs._inner, o._inner)), nil
	case *FloatSpanSet:
		return bool(C.adjacent_spanset_span(o._inner, g_fs._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fs *FloatSpan) Contains(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.contains_span_float(g_fs._inner, C.double(o))), nil
	case *FloatSpan:
		return bool(C.contains_span_span(g_fs._inner, o._inner)), nil
	case *FloatSpanSet:
		return bool(C.contains_span_spanset(g_fs._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fs *FloatSpan) IsSame(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.span_eq(g_fs._inner, C.float_to_span(C.double(o)))), nil
	case *FloatSpan:
		return bool(C.span_eq(g_fs._inner, o._inner)), nil
	case *FloatSpanSet:
		return bool(C.span_eq(g_fs._inner, o.ToSpan()._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

// ------------------------- Position Operations ---------------------------
func (g_fs *FloatSpan) IsLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.left_span_float(g_fs._inner, C.double(o))), nil
	case *FloatSpan:
		return bool(C.left_span_span(g_fs._inner, o._inner)), nil
	case *FloatSpanSet:
		return bool(C.left_span_spanset(g_fs._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fs *FloatSpan) IsOverOrLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.overleft_span_float(g_fs._inner, C.double(o))), nil
	case *FloatSpan:
		return bool(C.overleft_span_span(g_fs._inner, o._inner)), nil
	case *FloatSpanSet:
		return bool(C.overleft_span_spanset(g_fs._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fs *FloatSpan) IsRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.right_span_float(g_fs._inner, C.double(o))), nil
	case *FloatSpan:
		return bool(C.right_span_span(g_fs._inner, o._inner)), nil
	case *FloatSpanSet:
		return bool(C.right_span_spanset(g_fs._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fs *FloatSpan) IsOverOrRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.overright_span_float(g_fs._inner, C.double(o))), nil
	case *FloatSpan:
		return bool(C.overright_span_span(g_fs._inner, o._inner)), nil
	case *FloatSpanSet:
		return bool(C.overright_span_spanset(g_fs._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

// ------------------------- Distance Operations ---------------------------

func (g_fs *FloatSpan) Distance(other interface{}) (float64, error) {
	switch o := other.(type) {
	case int:
		return float64(C.distance_span_float(g_fs._inner, C.double(o))), nil
	case float64:
		return float64(C.distance_span_float(g_fs._inner, C.double(o))), nil
	case float32:
		return float64(C.distance_span_float(g_fs._inner, C.double(o))), nil
	case *FloatSpan:
		return float64(C.distance_floatspan_floatspan(g_fs._inner, o._inner)), nil
	case *FloatSpanSet:
		return float64(C.distance_floatspanset_floatspan(o._inner, g_fs._inner)), nil
	default:
		return 0.0, fmt.Errorf("operation not supported with type %T", other)
	}
}

// ------------------------- Set Operations --------------------------------
func (g_fs *FloatSpan) Intersection(other interface{}) (*FloatSpan, error) {
	switch o := other.(type) {
	case float64:
		res := C.intersection_span_float(g_fs._inner, C.double(o))
		if res == nil {
			return nil, nil
		} else {
			return &FloatSpan{_inner: res}, nil
		}
	case *FloatSpan:
		res := C.intersection_span_span(g_fs._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &FloatSpan{_inner: res}, nil
		}
	case *FloatSpanSet:
		res := C.intersection_spanset_span(o._inner, g_fs._inner)
		if res == nil {
			return nil, nil
		} else {
			// In Pymeos, it is FloatSpanSet, but in Go, it's FloatSpan because I need to declare the type of output
			fss := FloatSpanSet{_inner: res}
			fs := fss.ToSpan()
			return &fs, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fs *FloatSpan) Mul(other interface{}) (*FloatSpan, error) {
	return g_fs.Intersection(other)
}

func (g_fs *FloatSpan) Minus(other interface{}) (*FloatSpanSet, error) {
	switch o := other.(type) {
	case float64:
		res := C.minus_span_float(g_fs._inner, C.double(o))
		if res == nil {
			return nil, nil
		} else {
			return &FloatSpanSet{_inner: res}, nil
		}
	case *FloatSpan:
		res := C.minus_span_span(g_fs._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &FloatSpanSet{_inner: res}, nil
		}
	case *FloatSpanSet:
		res := C.minus_spanset_span(o._inner, g_fs._inner)
		if res == nil {
			return nil, nil
		} else {
			return &FloatSpanSet{_inner: res}, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fs *FloatSpan) Sub(other interface{}) (*FloatSpanSet, error) {
	return g_fs.Minus(other)
}

func (g_fs *FloatSpan) Union(other interface{}) (*FloatSpanSet, error) {
	switch o := other.(type) {
	case float64:
		res := C.gunion_span_float(g_fs._inner, C.double(o))
		if res == nil {
			return nil, nil
		} else {
			return &FloatSpanSet{_inner: res}, nil
		}
	case *FloatSpan:
		res := C.gunion_span_span(g_fs._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &FloatSpanSet{_inner: res}, nil
		}
	case *FloatSpanSet:
		res := C.gunion_spanset_span(o._inner, g_fs._inner)
		if res == nil {
			return nil, nil
		} else {
			return &FloatSpanSet{_inner: res}, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fs *FloatSpan) Add(other interface{}) (*FloatSpanSet, error) {
	return g_fs.Union(other)
}
