// collections/number/intspan.go
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

type IntSpan struct {
	_inner *C.Span
}

func NewIntSpan(g_is_in string) IntSpan {
	c_is_in := C.CString(g_is_in)
	defer C.free(unsafe.Pointer(c_is_in))
	c_is := C.intspan_in(c_is_in)
	g_is := IntSpan{_inner: c_is}
	return g_is
}

func (g_is IntSpan) IntSpanOut() string {
	c_is_out := C.intspan_out(g_is._inner)
	defer C.free(unsafe.Pointer(c_is_out))
	g_is_out := C.GoString(c_is_out)
	return g_is_out
}
