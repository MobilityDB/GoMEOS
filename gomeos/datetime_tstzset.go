package gomeos

/*
#include "meos.h"
#include "meos_catalog.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type TsTzSet struct {
	_inner *C.Set
}

// ------------------------- Input ----------------------------------------
func NewTsTzSet(g_tts_in string) *TsTzSet {
	c_tts_in := C.CString(g_tts_in)
	defer C.free(unsafe.Pointer(c_tts_in))
	c_tts := C.tstzset_in(c_tts_in)
	g_tts := &TsTzSet{_inner: c_tts}
	return g_tts
}

// ------------------------- Output ----------------------------------------
func (g_tts *TsTzSet) TsTzSetOut() string {
	c_tts_out := C.tstzset_out(g_tts._inner)
	defer C.free(unsafe.Pointer(c_tts_out))
	g_tts_out := C.GoString(c_tts_out)
	return g_tts_out
}
