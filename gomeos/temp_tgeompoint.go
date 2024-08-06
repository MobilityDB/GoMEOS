package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// ------------------------- TGeomPointInst ---------------------------
type TGeomPointInst struct {
	_inner *C.Temporal
}

func NewTGeomPointInst(tgmpi_in string) *TGeomPointInst {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tgeompoint_in(c_tgmpi_in)
	g_tgmpi := &TGeomPointInst{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TGeomPointInst) TPointOut(maxdd int) string {
	c_tgmpi_out := C.tpoint_as_text(tgmpi._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

// ------------------------- TGeomPointSeq ---------------------------
type TGeomPointSeq struct {
	_inner *C.Temporal
}

func NewTGeomPointSeq(tgmpi_in string) *TGeomPointSeq {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tgeompoint_in(c_tgmpi_in)
	g_tgmpi := &TGeomPointSeq{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TGeomPointSeq) TPointOut(maxdd int) string {
	c_tgmpi_out := C.tpoint_as_text(tgmpi._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

// ------------------------- TGeomPointSeqSet ---------------------------
type TGeomPointSeqSet struct {
	_inner *C.Temporal
}

func NewTGeomPointSeqSet(tgmpi_in string) *TGeomPointSeqSet {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tgeompoint_in(c_tgmpi_in)
	g_tgmpi := &TGeomPointSeqSet{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TGeomPointSeqSet) TPointOut(maxdd int) string {
	c_tgmpi_out := C.tpoint_as_text(tgmpi._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}
