package gomeos

/*
#include "meos.h"
#include "meos_catalog.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type TsTzSpanSet struct {
	_inner *C.SpanSet
}

// ------------------------- Input ----------------------------------------
func NewTsTzSpanSet(g_tts_in string) *TsTzSpanSet {
	c_tts_in := C.CString(g_tts_in)
	defer C.free(unsafe.Pointer(c_tts_in))
	c_tts := C.tstzspanset_in(c_tts_in)
	g_tts := &TsTzSpanSet{_inner: c_tts}
	return g_tts
}

// ------------------------- Output ----------------------------------------
func (g_tts *TsTzSpanSet) TsTzSpanSetOut() string {
	c_tts_out := C.tstzspanset_out(g_tts._inner)
	defer C.free(unsafe.Pointer(c_tts_out))
	g_tts_out := C.GoString(c_tts_out)
	return g_tts_out
}
