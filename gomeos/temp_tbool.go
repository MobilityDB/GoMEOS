package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// ------------------------- TBoolInst ---------------------------
type TBoolInst struct {
	_inner *C.Temporal
}

func NewTBoolInst(tgmpi_in string) *TBoolInst {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tbool_in(c_tgmpi_in)
	g_tgmpi := &TBoolInst{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TBoolInst) TBoolOut() string {
	c_tgmpi_out := C.tbool_out(tgmpi._inner)
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

// ------------------------- TBoolSeq ---------------------------
type TBoolSeq struct {
	_inner *C.Temporal
}

func NewTBoolSeq(tgmpi_in string) *TBoolSeq {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tbool_in(c_tgmpi_in)
	g_tgmpi := &TBoolSeq{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TBoolSeq) TBoolOut() string {
	c_tgmpi_out := C.tbool_out(tgmpi._inner)
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

// ------------------------- TBoolSeqSet ---------------------------
type TBoolSeqSet struct {
	_inner *C.Temporal
}

func NewTBoolSeqSet(tgmpi_in string) *TBoolSeqSet {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tbool_in(c_tgmpi_in)
	g_tgmpi := &TBoolSeqSet{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TBoolSeqSet) TBoolOut() string {
	c_tgmpi_out := C.tbool_out(tgmpi._inner)
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}
