package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
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

func (tb *TBoolInst) Init(c_temp *C.Temporal) {
	tb._inner = c_temp
}

func (tb *TBoolInst) Inner() *C.Temporal {
	return tb._inner
}

func (tb *TBoolInst) IsTBool() bool {
	return true
}

func (tb *TBoolInst) String() string {
	return tb.TBoolOut()
}

func (tb *TBoolInst) Type() string {
	return "TBoolInst"
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

func (tb *TBoolSeq) Init(c_temp *C.Temporal) {
	tb._inner = c_temp
}

func (tb *TBoolSeq) Inner() *C.Temporal {
	return tb._inner
}

func (tb *TBoolSeq) IsTBool() bool {
	return true
}

func (tb *TBoolSeq) String() string {
	return tb.TBoolOut()
}

func (tb *TBoolSeq) Type() string {
	return "TBoolSeq"
}

func (tb *TBoolSeq) IsTSequence() bool {
	return true
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

func (tb *TBoolSeqSet) Init(c_temp *C.Temporal) {
	tb._inner = c_temp
}

func (tb *TBoolSeqSet) Inner() *C.Temporal {
	return tb._inner
}

func (tb *TBoolSeqSet) IsTBool() bool {
	return true
}

func (tb *TBoolSeqSet) String() string {
	return tb.TBoolOut()
}

func (tb *TBoolSeqSet) Type() string {
	return "TBoolSeqSet"
}

// ------------------------- TBool ---------------------------
func TBoolIn[TB TBool](input string, output TB) TB {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_bool := C.tbool_in(c_str)
	output.Init(c_bool)
	return output
}

func TBoolFromMFJSON[TB TBool](input string, output TB) TB {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_bool := C.tbool_from_mfjson(c_str)
	output.Init(c_bool)
	return output
}

func TBoolOut[TB TBool](tb TB) string {
	c_bool := C.tbool_out(tb.Inner())
	defer C.free(unsafe.Pointer(c_bool))
	bool_out := C.GoString(c_bool)
	return bool_out
}
