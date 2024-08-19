package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"
import "unsafe"

// ------------------------- TIntInst ---------------------------
type TIntInst struct {
	_inner *C.Temporal
}

func NewTIntInst(tf_in string) *TIntInst {
	c_tf_in := C.CString(tf_in)
	defer C.free(unsafe.Pointer(c_tf_in))
	c_tf := C.tint_in(c_tf_in)
	g_tf := &TIntInst{_inner: c_tf}
	return g_tf
}

func (tf *TIntInst) TIntOut() string {
	c_tf_out := C.tint_out(tf._inner)
	defer C.free(unsafe.Pointer(c_tf_out))
	tf_out := C.GoString(c_tf_out)
	return tf_out
}

func (tb *TIntInst) Init(c_temp *C.Temporal) {
	tb._inner = c_temp
}

func (tb *TIntInst) Inner() *C.Temporal {
	return tb._inner
}

func (tb *TIntInst) IsTInt() bool {
	return true
}

func (tb *TIntInst) String() string {
	return tb.TIntOut()
}

func (tb *TIntInst) Type() string {
	return "TIntInst"
}

func (tb *TIntInst) IsTNumber() bool {
	return true
}

// ------------------------- TIntSeq ---------------------------
type TIntSeq struct {
	_inner *C.Temporal
}

func NewTIntSeq(tf_in string) *TIntSeq {
	c_tf_in := C.CString(tf_in)
	defer C.free(unsafe.Pointer(c_tf_in))
	c_tf := C.tint_in(c_tf_in)
	g_tf := &TIntSeq{_inner: c_tf}
	return g_tf
}

func (tf *TIntSeq) TIntOut() string {
	c_tf_out := C.tint_out(tf._inner)
	defer C.free(unsafe.Pointer(c_tf_out))
	tf_out := C.GoString(c_tf_out)
	return tf_out
}

func (tb *TIntSeq) Init(c_temp *C.Temporal) {
	tb._inner = c_temp
}

func (tb *TIntSeq) Inner() *C.Temporal {
	return tb._inner
}

func (tb *TIntSeq) IsTInt() bool {
	return true
}

func (tb *TIntSeq) String() string {
	return tb.TIntOut()
}

func (tb *TIntSeq) Type() string {
	return "TIntSeq"
}

func (tb *TIntSeq) IsTSequence() bool {
	return true
}

func (tb *TIntSeq) IsTNumber() bool {
	return true
}

// ------------------------- TIntSeqSet ---------------------------
type TIntSeqSet struct {
	_inner *C.Temporal
}

func NewTIntSeqSet(tf_in string) *TIntSeqSet {
	c_tf_in := C.CString(tf_in)
	defer C.free(unsafe.Pointer(c_tf_in))
	c_tf := C.tint_in(c_tf_in)
	g_tf := &TIntSeqSet{_inner: c_tf}
	return g_tf
}

func (tf *TIntSeqSet) TIntOut() string {
	c_tf_out := C.tint_out(tf._inner)
	defer C.free(unsafe.Pointer(c_tf_out))
	tf_out := C.GoString(c_tf_out)
	return tf_out
}

func (tb *TIntSeqSet) Init(c_temp *C.Temporal) {
	tb._inner = c_temp
}

func (tb *TIntSeqSet) Inner() *C.Temporal {
	return tb._inner
}

func (tb *TIntSeqSet) IsTInt() bool {
	return true
}

func (tb *TIntSeqSet) String() string {
	return tb.TIntOut()
}

func (tb *TIntSeqSet) Type() string {
	return "TIntSeqSet"
}

func (tb *TIntSeqSet) IsTNumber() bool {
	return true
}

// ------------------------- TInt ---------------------------
func TIntIn[TI TInt](input string, output TI) TI {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_int := C.tint_in(c_str)
	output.Init(c_int)
	return output
}

func TIntFromMFJSON[TI TInt](input string, output TI) TI {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_int := C.tint_from_mfjson(c_str)
	output.Init(c_int)
	return output
}

func TIntOut[TI TInt](ti TI) string {
	c_int := C.tint_out(ti.Inner())
	defer C.free(unsafe.Pointer(c_int))
	int_out := C.GoString(c_int)
	return int_out
}
