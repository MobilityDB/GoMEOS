package base

/*
#cgo CFLAGS: -I/opt/homebrew/include
#cgo LDFLAGS: -L/opt/homebrew/lib -lmeos -Wl,-rpath,/opt/homebrew/lib
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type Span interface {
	IsSpanSet() bool
}

type IntSpan struct {
	_inner *C.Span
}

type FloatSpan struct {
	_inner *C.Span
}

func (iss *IntSpan) IsSpan() bool {
	return true
}

func (fss *FloatSpan) IsSpan() bool {
	return true
}

// ------------------------- Input ----------------------------------------
func NewIntSpan(g_is_in string) *IntSpan {
	c_is_in := C.CString(g_is_in)
	defer C.free(unsafe.Pointer(c_is_in))
	c_is := C.intspan_in(c_is_in)
	g_is := &IntSpan{_inner: c_is}
	return g_is
}

func NewFloatSpan(g_fs_in string) *FloatSpan {
	c_fs_in := C.CString(g_fs_in)
	defer C.free(unsafe.Pointer(c_fs_in))
	c_fs := C.floatspan_in(c_fs_in)
	g_fs := &FloatSpan{_inner: c_fs}
	return g_fs
}

// ------------------------- Output ----------------------------------------
func (g_is *IntSpan) IntSpanOut() string {
	c_is_out := C.intspan_out(g_is._inner)
	defer C.free(unsafe.Pointer(c_is_out))
	g_is_out := C.GoString(c_is_out)
	return g_is_out
}

func (g_fs *FloatSpan) FloatSpanOut(max_decimal int) string {
	c_fs_out := C.floatspan_out(g_fs._inner, C.int(max_decimal))
	defer C.free(unsafe.Pointer(c_fs_out))
	g_fs_out := C.GoString(c_fs_out)
	return g_fs_out
}
