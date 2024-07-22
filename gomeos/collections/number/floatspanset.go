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
import (
	"unsafe"
)

type FloatSpanSet struct {
	_inner *C.SpanSet
}

func NewFloatSpanSet(g_fss_in string) FloatSpanSet {
	c_fss_in := C.CString(g_fss_in)
	defer C.free(unsafe.Pointer(c_fss_in))
	c_fss := C.floatspanset_in(c_fss_in)
	g_fss := FloatSpanSet{_inner: c_fss}
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
func (g_fss FloatSpanSet) FloatSpanSetOut(max_decimal int) string {
	c_fss_out := C.floatspanset_out(g_fss._inner, C.int(max_decimal))
	defer C.free(unsafe.Pointer(c_fss_out))
	g_fss_out := C.GoString(c_fss_out)
	return g_fss_out
}
