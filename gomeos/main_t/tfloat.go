// collections/number/floatset.go
package main_t

/*
#cgo CFLAGS: -I/opt/homebrew/include
#cgo LDFLAGS: -L/opt/homebrew/lib -lmeos -Wl,-rpath,/opt/homebrew/lib
#include "meos.h"
#include "meos_catalog.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

type TFloat struct {
	_inner *C.Temporal
}

// ------------------------- Input ----------------------------------------
func NewTFloat(g_tf_in string) *TFloat {
	c_tf_in := C.CString(g_tf_in)
	defer C.free(unsafe.Pointer(c_tf_in))
	c_tf := C.tfloat_in(c_tf_in)
	g_tf := &TFloat{_inner: c_tf}
	return g_tf
}

// ------------------------- Output ----------------------------------------
func (g_tf *TFloat) TFloatOut(max_decimal int) string {
	c_tf_out := C.tfloat_out(g_tf._inner, C.int(max_decimal))
	defer C.free(unsafe.Pointer(c_tf_out))
	g_tf_out := C.GoString(c_tf_out)
	return g_tf_out
}

// func createTFloat() *TFloat {
// 	return NewTFloat("[1@2020-03-01, 10@2020-03-10]")
// }

// func main() {
// 	// C.meos_initialize(nil, nil)
// 	g_is := createTFloat()
// 	fmt.Println(g_is.TFloatOut(5))
// 	// C.meos_finalize()
// }
