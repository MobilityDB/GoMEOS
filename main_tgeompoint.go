package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"
import "unsafe"

type TGeomPointInst struct {
	_inner *C.Temporal
}

func NewTGeomPointInstInner(inner *C.Temporal) *TGeomPointInst {
	return &TGeomPointInst{_inner: inner}
}

func NewTGeomPointInst(tgmpi_in string) TGeomPointInst {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tgeompoint_in(c_tgmpi_in)
	g_tgmpi := TGeomPointInst{_inner: c_tgmpi}
	return g_tgmpi
}

func NewEmptyTGeomPointInst() TGeomPointInst {
	g_tgmpi := TGeomPointInst{_inner: nil}
	return g_tgmpi
}

func (tgmpi *TGeomPointInst) TPointOut(maxdd int) string {
	c_tgmpi_out := C.tpoint_as_text(tgmpi._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

func (tgmpi *TGeomPointInst) TInstantOut(maxdd int) string {
	c_tgmpi_out := C.tpoint_as_text(tgmpi._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

func (tgmpi *TGeomPointInst) IsTInstant() bool {
	return true
}

func (tb *TGeomPointInst) Init(c_temp *C.Temporal) {
	tb._inner = c_temp
}

func (tb *TGeomPointInst) Inner() *C.Temporal {
	return tb._inner
}

func (tb *TGeomPointInst) IsTGeomPoint() bool {
	return true
}

func (tgmpi *TGeomPointInst) IsTPoint() bool {
	return true
}

func (tgmpi *TGeomPointInst) String() string {
	return tgmpi.TPointOut(5)
}

func (tgmpi *TGeomPointInst) Type() string {
	return "TGeomPointInst"
}

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

func NewTGeomPointSeqFromWKB(tgmpi_in string) *TGeomPointSeq {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.temporal_from_hexwkb(c_tgmpi_in)
	g_tgmpi := &TGeomPointSeq{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TGeomPointSeq) TPointOut(maxdd int) string {
	c_tgmpi_out := C.tpoint_as_text(tgmpi._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

func (tgmpi *TGeomPointSeq) Inner() *C.Temporal {
	return tgmpi._inner
}

func (tgmpi *TGeomPointSeq) Init(c_temp *C.Temporal) {
	tgmpi._inner = c_temp
}

func (tgmpi *TGeomPointSeq) NewInit(c_temp *C.Temporal) *TGeomPointSeq {
	return &TGeomPointSeq{_inner: c_temp}
}

func (tgmpi *TGeomPointSeq) IsTGeomPoint() bool {
	return true
}

func (tgmpi *TGeomPointSeq) IsTPoint() bool {
	return true
}

func (tgmpi *TGeomPointSeq) String() string {
	return tgmpi.TPointOut(10)
}

func (tgmpi *TGeomPointSeq) Type() string {
	return "TGeomPointSeq"
}

func (tgmpi *TGeomPointSeq) IsTSequence() bool {
	return true
}

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

func (tb *TGeomPointSeqSet) Init(c_temp *C.Temporal) {
	tb._inner = c_temp
}

func (tb *TGeomPointSeqSet) Inner() *C.Temporal {
	return tb._inner
}

func (tb *TGeomPointSeqSet) IsTGeomPoint() bool {
	return true
}

func (tgmpi *TGeomPointSeqSet) IsTPoint() bool {
	return true
}

func (tgmpi *TGeomPointSeqSet) String() string {
	return tgmpi.TPointOut(10)
}

func (tgmpi *TGeomPointSeqSet) Type() string {
	return "TGeomPointSeqSet"
}

func TGeomPointIn[TG TGeomPoint](input string, output TG) TG {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_geompoint := C.tgeompoint_in(c_str)
	output.Init(c_geompoint)
	return output
}

func TGeomPointFromMFJSON[TG TGeomPoint](input string, output TG) TG {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_geompoint := C.tgeompoint_from_mfjson(c_str)
	output.Init(c_geompoint)
	return output
}
