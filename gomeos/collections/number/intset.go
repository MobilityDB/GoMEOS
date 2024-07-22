// collections/number/intset.go
package number

/*
#cgo CFLAGS: -I/opt/homebrew/include
#cgo LDFLAGS: -L/opt/homebrew/lib -lmeos -Wl,-rpath,/opt/homebrew/lib
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type IntSet struct {
	_inner *C.Set
}

// ------------------------- Input ----------------------------------------
func NewIntSet(g_is_in string) *IntSet {
	c_is_in := C.CString(g_is_in)
	defer C.free(unsafe.Pointer(c_is_in))
	c_is := C.intset_in(c_is_in)
	g_is := &IntSet{_inner: c_is}
	return g_is
}

// ------------------------- Output ----------------------------------------
/*
Return the string representation of the content of IntSpanSet.

Returns:
	String

MEOS Functions:
	intspanset_out
*/
func (g_is *IntSet) IntSetOut() string {
	c_is_out := C.intset_out(g_is._inner)
	defer C.free(unsafe.Pointer(c_is_out))
	g_is_out := C.GoString(c_is_out)
	return g_is_out
}

func (g_is *IntSet) ToSpanSet() *IntSpanSet {
	return &IntSpanSet{_inner: C.set_to_spanset(g_is._inner)}
}
