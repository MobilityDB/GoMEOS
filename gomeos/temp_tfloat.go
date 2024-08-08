package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// ------------------------- TFloatInst ---------------------------
type TFloatInst struct {
	_inner *C.Temporal
}

func NewTFloatInst(tf_in string) TFloatInst {
	c_tf_in := C.CString(tf_in)
	defer C.free(unsafe.Pointer(c_tf_in))
	c_tf := C.tfloat_in(c_tf_in)
	g_tf := TFloatInst{_inner: c_tf}
	return g_tf
}

func (tf *TFloatInst) TPointOut(maxdd int) string {
	c_tf_out := C.tfloat_out(tf._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tf_out))
	tf_out := C.GoString(c_tf_out)
	return tf_out
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

func (tf TFloatSeq) Inner() *C.Temporal {
	return tf._inner
}

type Temporal interface {
	Inner() *C.Temporal
}

func TnumberTwavg[T Temporal](temp T) float64 {
	res := C.tnumber_twavg(temp.Inner())
	return float64(res)
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
