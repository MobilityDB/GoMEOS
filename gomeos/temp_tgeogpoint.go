package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// ------------------------- TGeogPointInst ---------------------------
type TGeogPointInst struct {
	_inner *C.Temporal
}

func NewTGeogPointInst(tgmpi_in string) *TGeogPointInst {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tgeogpoint_in(c_tgmpi_in)
	g_tgmpi := &TGeogPointInst{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TGeogPointInst) TPointOut(maxdd int) string {
	c_tgmpi_out := C.tpoint_as_text(tgmpi._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

// ------------------------- TGeogPointSeq ---------------------------
type TGeogPointSeq struct {
	_inner *C.Temporal
}

func NewTGeogPointSeq(tgmpi_in string) *TGeogPointSeq {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tgeogpoint_in(c_tgmpi_in)
	g_tgmpi := &TGeogPointSeq{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TGeogPointSeq) TPointOut(maxdd int) string {
	c_tgmpi_out := C.tpoint_as_text(tgmpi._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

// ------------------------- TGeogPointSeqSet ---------------------------
type TGeogPointSeqSet struct {
	_inner *C.Temporal
}

func NewTGeogPointSeqSet(tgmpi_in string) *TGeogPointSeqSet {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tgeogpoint_in(c_tgmpi_in)
	g_tgmpi := &TGeogPointSeqSet{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TGeogPointSeqSet) TPointOut(maxdd int) string {
	c_tgmpi_out := C.tpoint_as_text(tgmpi._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}
