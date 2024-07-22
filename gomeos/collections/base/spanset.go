package base

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

type SpanSet interface {
	inner() *C.SpanSet
}

type IntSpanSet struct {
	_inner *C.SpanSet
}

type FloatSpanSet struct {
	_inner *C.SpanSet
}

func (iss *IntSpanSet) inner() *C.SpanSet {
	return iss._inner
}

func (fss *FloatSpanSet) inner() *C.SpanSet {
	return fss._inner
}

// ------------------------- Input ----------------------------------------
func NewIntSpanSet(g_iss_in string) *IntSpanSet {
	c_iss_in := C.CString(g_iss_in)
	defer C.free(unsafe.Pointer(c_iss_in))
	c_iss := C.intspanset_in(c_iss_in)
	g_iss := &IntSpanSet{_inner: c_iss}
	return g_iss
}

func NewFloatSpanSet(g_fss_in string) *FloatSpanSet {
	c_fss_in := C.CString(g_fss_in)
	defer C.free(unsafe.Pointer(c_fss_in))
	c_fss := C.floatspanset_in(c_fss_in)
	g_fss := &FloatSpanSet{_inner: c_fss}
	return g_fss
}

// ------------------------- Output ----------------------------------------
func (g_iss *IntSpanSet) IntSpanSetOut() string {
	c_iss_out := C.intspanset_out(g_iss._inner)
	defer C.free(unsafe.Pointer(c_iss_out))
	g_iss_out := C.GoString(c_iss_out)
	return g_iss_out
}

func (g_fss FloatSpanSet) FloatSpanSetOut(max_decimal int) string {
	c_fss_out := C.floatspanset_out(g_fss._inner, C.int(max_decimal))
	defer C.free(unsafe.Pointer(c_fss_out))
	g_fss_out := C.GoString(c_fss_out)
	return g_fss_out
}

// ------------------------- Conversions -----------------------------------
func (g_iss IntSpanSet) ToSpan() IntSpan {
	return IntSpan{_inner: C.spanset_span(g_iss._inner)}
}

func (g_fss FloatSpanSet) ToSpan() FloatSpan {
	return FloatSpan{_inner: C.spanset_span(g_fss._inner)}
}

func (g_iss IntSpanSet) ToFloatSpanSet() FloatSpanSet {
	return FloatSpanSet{_inner: C.intspanset_to_floatspanset(g_iss._inner)}
}

func (g_fss FloatSpanSet) ToIntSpanSet() IntSpanSet {
	return IntSpanSet{_inner: C.floatspanset_to_intspanset(g_fss._inner)}
}

// ------------------------- Accessors -------------------------------------

func NumSpans[SS SpanSet](ss SS) int {
	return int(C.spanset_num_spans(ss.inner()))
}
