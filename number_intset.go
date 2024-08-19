// collections/number/intset.go
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

type IntSet struct {
	_inner *C.Set
}

// ------------------------- Input ----------------------------------------
func NewIntSet(g_is_in string) IntSet {
	c_is_in := C.CString(g_is_in)
	defer C.free(unsafe.Pointer(c_is_in))
	c_is := C.intset_in(c_is_in)
	g_is := IntSet{_inner: c_is}
	return g_is
}

// ------------------------- Output ----------------------------------------
func (g_is *IntSet) IntSetOut() string {
	c_is_out := C.intset_out(g_is._inner)
	defer C.free(unsafe.Pointer(c_is_out))
	g_is_out := C.GoString(c_is_out)
	return g_is_out
}

// ------------------------- Conversions -----------------------------------
func (g_is IntSet) ToFloatSet() FloatSet {
	return FloatSet{_inner: C.intset_to_floatset(g_is._inner)}
}

func (g_is IntSet) ToSpanSet() IntSpanSet {
	return IntSpanSet{_inner: C.set_to_spanset(g_is._inner)}
}

// ------------------------- Accessors -------------------------------------
func (g_is IntSet) StartElement() int {
	return int(C.intset_start_value(g_is._inner))
}

func (g_is IntSet) EndElement() int {
	return int(C.intset_end_value(g_is._inner))
}

func (g_is IntSet) ElementN(n int) int {
	res := C.malloc(C.sizeof_int)
	success := C.intset_value_n(g_is._inner, C.int(n+1), (*C.int)(res))
	if success {
		result := *(*C.int)(res)
		return int(result)
	} else {
		return 0
	}
}

func (g_is IntSet) NumElements() int {
	return int(C.set_num_values(g_is._inner))
}

func (g_is IntSet) Elements() []int {
	nums := g_is.NumElements()
	ints := make([]int, nums)
	for i := 0; i < nums; i++ {
		ints[i] = g_is.ElementN(i)
	}
	return ints
}

// ------------------------- Transformations -------------------------------

func (g_is IntSet) ShiftScale(d int, w int) IntSet {
	modified := C.intset_shift_scale(g_is._inner, C.int(d), C.int(w), C._Bool(d != 0), C._Bool(w != 0))
	return IntSet{_inner: modified}
}

func (g_is IntSet) Shift(delta int) IntSet {
	return g_is.ShiftScale(delta, 0)

}

func (g_is IntSet) Scale(width int) IntSet {
	return g_is.ShiftScale(0, width)
}

// ------------------------- Topological Operations --------------------------------

func (g_is *IntSet) Contains(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.contains_set_int(g_is._inner, C.int(o))), nil
	case *IntSet:
		return bool(C.contains_set_set(g_is._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

// ------------------------- Position Operations ---------------------------
func (g_is *IntSet) IsLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.left_set_int(g_is._inner, C.int(o))), nil
	case *IntSet:
		return bool(C.left_set_set(g_is._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_is *IntSet) IsOverOrLeft(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.overleft_set_int(g_is._inner, C.int(o))), nil
	case *IntSet:
		return bool(C.overleft_set_set(g_is._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_is *IntSet) IsRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.right_set_int(g_is._inner, C.int(o))), nil
	case *IntSet:
		return bool(C.right_set_set(g_is._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_is *IntSet) IsOverOrRight(other interface{}) (bool, error) {
	switch o := other.(type) {
	case int:
		return bool(C.overright_set_int(g_is._inner, C.int(o))), nil
	case *IntSet:
		return bool(C.overright_set_set(g_is._inner, o._inner)), nil
	default:
		return false, fmt.Errorf("operation not supported with type %T", other)
	}
}

// ------------------------- Set Operations --------------------------------
func (g_is *IntSet) Intersection(other interface{}) (*IntSet, error) {
	switch o := other.(type) {
	case int:
		res := C.intersection_set_int(g_is._inner, C.int(o))
		if res == nil {
			return nil, nil
		} else {
			return &IntSet{_inner: res}, nil
		}
	case *IntSet:
		res := C.intersection_set_set(g_is._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &IntSet{_inner: res}, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_is *IntSet) Mul(other interface{}) (*IntSet, error) {
	return g_is.Intersection(other)
}

func (g_is *IntSet) Minus(other interface{}) (*IntSet, error) {
	switch o := other.(type) {
	case int:
		res := C.minus_set_int(g_is._inner, C.int(o))
		if res == nil {
			return nil, nil
		} else {
			return &IntSet{_inner: res}, nil
		}
	case *IntSet:
		res := C.minus_set_set(g_is._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &IntSet{_inner: res}, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_is *IntSet) Sub(other interface{}) (*IntSet, error) {
	return g_is.Minus(other)
}

func (g_is *IntSet) Union(other interface{}) (*IntSet, error) {
	switch o := other.(type) {
	case int:
		res := C.gunion_set_int(g_is._inner, C.int(o))
		if res == nil {
			return nil, nil
		} else {
			return &IntSet{_inner: res}, nil
		}
	case *IntSet:
		res := C.gunion_set_set(g_is._inner, o._inner)
		if res == nil {
			return nil, nil
		} else {
			return &IntSet{_inner: res}, nil
		}
	default:
		return nil, fmt.Errorf("operation not supported with type %T", other)
	}
}

func (g_is *IntSet) Add(other interface{}) (*IntSet, error) {
	return g_is.Union(other)
}

// ------------------------- Distance Operations --------------------------------
func (g_is *IntSet) Distance(other interface{}) (int, error) {
	switch o := other.(type) {
	case int:
		return int(C.distance_set_int(g_is._inner, C.int(o))), nil
	case *IntSet:
		return int(C.distance_intset_intset(g_is._inner, o._inner)), nil
	case *IntSpan:
		return int(C.distance_intspanset_intspan(g_is.ToSpanSet()._inner, o._inner)), nil
	case *IntSpanSet:
		return int(C.distance_intspanset_intspanset(g_is.ToSpanSet()._inner, o._inner)), nil
	default:
		return 0, fmt.Errorf("operation not supported with type %T", other)
	}
}
