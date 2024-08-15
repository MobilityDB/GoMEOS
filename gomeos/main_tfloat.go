package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"
import "unsafe"

// ------------------------- TFloatInst ---------------------------
type TFloatInst struct {
	_inner *C.Temporal
}

func NewTFloatInst(tf_in string) *TFloatInst {
	c_tf_in := C.CString(tf_in)
	defer C.free(unsafe.Pointer(c_tf_in))
	c_tf := C.tfloat_in(c_tf_in)
	g_tf := TFloatInst{_inner: c_tf}
	return &g_tf
}

func (tf *TFloatInst) TPointOut(maxdd int) string {
	c_tf_out := C.tfloat_out(tf._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tf_out))
	tf_out := C.GoString(c_tf_out)
	return tf_out
}

func (tf *TFloatInst) Inner() *C.Temporal {
	return tf._inner
}

func (tf *TFloatInst) Init(c_temp *C.Temporal) {
	tf._inner = c_temp
}

func (tf *TFloatInst) IsTFloat() bool {
	return true
}

func (tf *TFloatInst) String() string {
	return tf.TPointOut(10)
}

func (tf *TFloatInst) Type() string {
	return "TFloatInst"
}

// ------------------------- TFloatSeq ---------------------------
type TFloatSeq struct {
	_inner *C.Temporal
}

func NewTFloatSeq(tf_in string) *TFloatSeq {
	c_tf_in := C.CString(tf_in)
	defer C.free(unsafe.Pointer(c_tf_in))
	c_tf := C.tfloat_in(c_tf_in)
	g_tf := &TFloatSeq{_inner: c_tf}
	return g_tf
}

func (tf *TFloatSeq) TPointOut(maxdd int) string {
	c_tf_out := C.tfloat_out(tf._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tf_out))
	tf_out := C.GoString(c_tf_out)
	return tf_out
}

func (tf *TFloatSeq) Inner() *C.Temporal {
	return tf._inner
}

func (tf *TFloatSeq) Init(c_temp *C.Temporal) {
	tf._inner = c_temp
}

func (tf *TFloatSeq) IsTFloat() bool {
	return true
}

func (tf *TFloatSeq) String() string {
	return tf.TPointOut(10)
}

func (tf *TFloatSeq) Type() string {
	return "TFloatSeq"
}

func (tf *TFloatSeq) IsTSequence() bool {
	return true
}

// ------------------------- TFloatSeqSet ---------------------------
type TFloatSeqSet struct {
	_inner *C.Temporal
}

func NewTFloatSeqSet(tf_in string) *TFloatSeqSet {
	c_tf_in := C.CString(tf_in)
	defer C.free(unsafe.Pointer(c_tf_in))
	c_tf := C.tfloat_in(c_tf_in)
	g_tf := &TFloatSeqSet{_inner: c_tf}
	return g_tf
}

func (tf *TFloatSeqSet) TPointOut(maxdd int) string {
	c_tf_out := C.tfloat_out(tf._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tf_out))
	tf_out := C.GoString(c_tf_out)
	return tf_out
}

func (tf *TFloatSeqSet) Inner() *C.Temporal {
	return tf._inner
}

func (tf *TFloatSeqSet) Init(c_temp *C.Temporal) {
	tf._inner = c_temp
}

func (tf *TFloatSeqSet) IsTFloat() bool {
	return true
}

func (tf *TFloatSeqSet) String() string {
	return tf.TPointOut(10)
}

func (tf *TFloatSeqSet) Type() string {
	return "TFloatSeqSet"
}

// ------------------------- TFloat ---------------------------
func TFloatIn[TF TFloat](input string, output TF) TF {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_float := C.tfloat_in(c_str)
	output.Init(c_float)
	return output
}

func TFloatFromMFJSON[TF TFloat](input string, output TF) TF {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_float := C.tfloat_from_mfjson(c_str)
	output.Init(c_float)
	return output
}

func TFloatOut[TF TFloat](tf TF, maxdd int) string {
	c_float := C.tfloat_out(tf.Inner(), C.int(maxdd))
	defer C.free(unsafe.Pointer(c_float))
	float_out := C.GoString(c_float)
	return float_out
}
