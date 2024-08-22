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

type FloatSet struct {
	_inner *C.Set
}

func NewFloatSet(g_fs_in string) FloatSet {
	c_fs_in := C.CString(g_fs_in)
	defer C.free(unsafe.Pointer(c_fs_in))
	c_fs := C.floatset_in(c_fs_in)
	g_fs := FloatSet{_inner: c_fs}
	return g_fs
}

func (g_fs *FloatSet) FloatSetOut(max_decimal int) string {
	c_fs_out := C.floatset_out(g_fs._inner, C.int(max_decimal))
	defer C.free(unsafe.Pointer(c_fs_out))
	g_fs_out := C.GoString(c_fs_out)
	return g_fs_out
}

func (g_fs FloatSet) ToIntSet() IntSet {
	return IntSet{_inner: C.floatset_to_intset(g_fs._inner)}
}

func (g_fs FloatSet) ToSpanSet() FloatSpanSet {
	return FloatSpanSet{_inner: C.set_to_spanset(g_fs._inner)}
}

func (g_fs FloatSet) StartElement() float64 {
	return float64(C.floatset_start_value(g_fs._inner))
}

func (g_fs FloatSet) EndElement() float64 {
	return float64(C.floatset_end_value(g_fs._inner))
}

func (g_fs FloatSet) ElementN(n int) float64 {
	res := C.malloc(C.sizeof_double)
	success := C.floatset_value_n(g_fs._inner, C.int(n+1), (*C.double)(res))
	if success {
		result := *(*C.double)(res)
		return float64(result)
	} else {
		return 0.0
	}
}

func (g_fs FloatSet) NumElements() int {
	return int(C.set_num_values(g_fs._inner))
}

func (g_fs FloatSet) Elements() []float64 {
	nums := g_fs.NumElements()
	floats := make([]float64, nums)
	for i := 0; i < nums; i++ {
		floats[i] = g_fs.ElementN(i)
	}
	return floats
}

func (g_fs FloatSet) ShiftScale(d float64, w float64) FloatSet {
	modified := C.floatset_shift_scale(g_fs._inner, C.double(d), C.double(w), C._Bool(d != 0), C._Bool(w != 0))
	return FloatSet{_inner: modified}
}

func (g_fs FloatSet) Shift(delta float64) FloatSet {
	return g_fs.ShiftScale(delta, 0)

}

func (g_fs FloatSet) Scale(width float64) FloatSet {
	return g_fs.ShiftScale(0, width)
}

func (g_fs *FloatSet) Contains(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.contains_set_float(g_fs._inner, C.double(o))), nil
	case *FloatSet:
		return bool(C.contains_set_set(g_fs._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fs *FloatSet) IsLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.left_set_float(g_fs._inner, C.double(o))), nil
	case *FloatSet:
		return bool(C.left_set_set(g_fs._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fs *FloatSet) IsOverOrLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.overleft_set_float(g_fs._inner, C.double(o))), nil
	case *FloatSet:
		return bool(C.overleft_set_set(g_fs._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fs *FloatSet) IsRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.right_set_float(g_fs._inner, C.double(o))), nil
	case *FloatSet:
		return bool(C.right_set_set(g_fs._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fs *FloatSet) IsOverOrRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case float64:
		return bool(C.overright_set_float(g_fs._inner, C.double(o))), nil
	case *FloatSet:
		return bool(C.overright_set_set(g_fs._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fs *FloatSet) Intersection(other interface{}) (*FloatSet, error) {
	switch o := other.(type) {
	case float64:
		res := C.intersection_set_float(g_fs._inner, C.double(o))
		if res == nil {
			return nil, nil
		} else {
			return &FloatSet{_inner: res}, nil
		}
	case *FloatSet:
		res := C.intersection_set_set(g_fs._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &FloatSet{_inner: res}, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fs *FloatSet) Mul(other interface{}) (*FloatSet, error) {
	return g_fs.Intersection(other)
}

func (g_fs *FloatSet) Minus(other interface{}) (*FloatSet, error) {
	switch o := other.(type) {
	case float64:
		res := C.minus_set_float(g_fs._inner, C.double(o))
		if res == nil {
			return nil, nil
		} else {
			return &FloatSet{_inner: res}, nil
		}
	case *FloatSet:
		res := C.minus_set_set(g_fs._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &FloatSet{_inner: res}, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fs *FloatSet) Sub(other interface{}) (*FloatSet, error) {
	return g_fs.Minus(other)
}

func (g_fs *FloatSet) Union(other interface{}) (*FloatSet, error) {
	switch o := other.(type) {
	case float64:
		res := C.gunion_set_float(g_fs._inner, C.double(o))
		if res == nil {
			return nil, nil
		} else {
			return &FloatSet{_inner: res}, nil
		}
	case *FloatSet:
		res := C.gunion_set_set(g_fs._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &FloatSet{_inner: res}, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_fs *FloatSet) Add(other interface{}) (*FloatSet, error) {
	return g_fs.Union(other)
}

func (g_fs *FloatSet) Distance(other interface{}) (float64, error) {
	switch o := other.(type) {
	case float64:
		return float64(C.distance_set_float(g_fs._inner, C.double(o))), nil
	case *FloatSet:
		return float64(C.distance_floatset_floatset(g_fs._inner, o._inner)), nil
	case *FloatSpan:
		return float64(C.distance_floatspanset_floatspan(g_fs.ToSpanSet()._inner, o._inner)), nil
	case *FloatSpanSet:
		return float64(C.distance_floatspanset_floatspanset(g_fs.ToSpanSet()._inner, o._inner)), nil
	default:
		return 0.0, fmt.Errorf("operation not supported with type %T", other)
	}
}
