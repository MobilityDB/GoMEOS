package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"
import "unsafe"

// ------------------------- TTextInst ---------------------------
type TTextInst struct {
	_inner *C.Temporal
}

func (tb *TTextInst) Init(c_temp *C.Temporal) {
	tb._inner = c_temp
}

func (tb *TTextInst) Inner() *C.Temporal {
	return tb._inner
}

func (tb *TTextInst) IsTText() bool {
	return true
}

func (tb *TTextInst) TTextOut() string {
	c_tf_out := C.ttext_out(tb._inner)
	defer C.free(unsafe.Pointer(c_tf_out))
	tf_out := C.GoString(c_tf_out)
	return tf_out
}

func (tb *TTextInst) String() string {
	return tb.TTextOut()
}

func (tb *TTextInst) Type() string {
	return "TTextInst"
}

// ------------------------- TTextSeq ---------------------------
type TTextSeq struct {
	_inner *C.Temporal
}

func (tb *TTextSeq) Init(c_temp *C.Temporal) {
	tb._inner = c_temp
}

func (tb *TTextSeq) Inner() *C.Temporal {
	return tb._inner
}

func (tb *TTextSeq) IsTText() bool {
	return true
}

func (tb *TTextSeq) TTextOut() string {
	c_tf_out := C.ttext_out(tb._inner)
	defer C.free(unsafe.Pointer(c_tf_out))
	tf_out := C.GoString(c_tf_out)
	return tf_out
}

func (tb *TTextSeq) String() string {
	return tb.TTextOut()
}

func (tb *TTextSeq) Type() string {
	return "TTextSeq"
}

func (tb *TTextSeq) IsTSequence() bool {
	return true
}

// ------------------------- TTextSeqSet ---------------------------
type TTextSeqSet struct {
	_inner *C.Temporal
}

func (tb *TTextSeqSet) Init(c_temp *C.Temporal) {
	tb._inner = c_temp
}

func (tb *TTextSeqSet) Inner() *C.Temporal {
	return tb._inner
}

func (tb *TTextSeqSet) IsTText() bool {
	return true
}

func (tb *TTextSeqSet) TTextOut() string {
	c_tf_out := C.ttext_out(tb._inner)
	defer C.free(unsafe.Pointer(c_tf_out))
	tf_out := C.GoString(c_tf_out)
	return tf_out
}

func (tb *TTextSeqSet) String() string {
	return tb.TTextOut()
}

func (tb *TTextSeqSet) Type() string {
	return "TTextSeqSet"
}

// ------------------------- TText ---------------------------
func TTextIn[TT TText](input string, output TT) TT {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_text := C.ttext_in(c_str)
	output.Init(c_text)
	return output
}

func TTextFromMFJSON[TT TText](input string, output TT) TT {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_text := C.ttext_from_mfjson(c_str)
	output.Init(c_text)
	return output
}

func TTextOut[TT TText](tt TT) string {
	c_text := C.ttext_out(tt.Inner())
	defer C.free(unsafe.Pointer(c_text))
	text_out := C.GoString(c_text)
	return text_out
}
