package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
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

func (tf *TIntInst) TPointOut() string {
	c_tf_out := C.tint_out(tf._inner)
	defer C.free(unsafe.Pointer(c_tf_out))
	tf_out := C.GoString(c_tf_out)
	return tf_out
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

func (tf *TIntSeq) TPointOut() string {
	c_tf_out := C.tint_out(tf._inner)
	defer C.free(unsafe.Pointer(c_tf_out))
	tf_out := C.GoString(c_tf_out)
	return tf_out
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

func (tf *TIntSeqSet) TPointOut() string {
	c_tf_out := C.tint_out(tf._inner)
	defer C.free(unsafe.Pointer(c_tf_out))
	tf_out := C.GoString(c_tf_out)
	return tf_out
}
