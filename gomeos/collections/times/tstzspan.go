package times

/*
#cgo CFLAGS: -I/opt/homebrew/include
#cgo LDFLAGS: -L/opt/homebrew/lib -lmeos -Wl,-rpath,/opt/homebrew/lib
#include "meos.h"
#include "meos_catalog.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type TsTzSpan struct {
	_inner *C.Span
}

// ------------------------- Input ----------------------------------------
func NewTsTzSpan(g_tts_in string) *TsTzSpan {
	c_tts_in := C.CString(g_tts_in)
	defer C.free(unsafe.Pointer(c_tts_in))
	c_tts := C.tstzspan_in(c_tts_in)
	g_tts := &TsTzSpan{_inner: c_tts}
	return g_tts
}

// ------------------------- Output ----------------------------------------
func (g_tts *TsTzSpan) TsTzSpanOut() string {
	c_tts_out := C.tstzspan_out(g_tts._inner)
	defer C.free(unsafe.Pointer(c_tts_out))
	g_tts_out := C.GoString(c_tts_out)
	return g_tts_out
}
