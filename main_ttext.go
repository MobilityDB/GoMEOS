package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
#include "string.h"
*/
import "C"
import (
	"time"
	"unsafe"
)

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

// TTextFromBaseTemporal Return a temporal text from a text and the time frame of another temporal value
func TTextFromBaseTemporal[T1 Temporal, T2 Temporal](value string, base T1, output T2) T2 {
	c_temp := C.ttext_from_base_temp(C.cstring2text(C.CString(value)), base.Inner())
	output.Init(c_temp)
	return output
}

// TTextInstMake Return a temporal text instant from a text and a timestamptz
func TTextInstMake(value string, base time.Time) *TTextInst {
	res := C.ttextinst_make(C.cstring2text(C.CString(value)), DatetimeToTimestamptz(base))
	return &TTextInst{_inner: C.cast_tinstant_to_temporal(res)}
}

// TTextSeqFromBaseTstzset Return a temporal text discrete sequence from a text and a timestamptz set
func TTextSeqFromBaseTstzset(value string, base TsTzSet) *TTextSeq {
	res := C.ttextseq_from_base_tstzset(C.cstring2text(C.CString(value)), base._inner)
	return &TTextSeq{_inner: C.cast_tsequence_to_temporal(res)}
}

// TTextSeqFromBaseTstzspan Return a temporal text sequence from a text and a timestamptz span
func TTextSeqFromBaseTstzspan(value string, base TsTzSpan) *TTextSeq {
	res := C.ttextseq_from_base_tstzspan(C.cstring2text(C.CString(value)), base._inner)
	return &TTextSeq{_inner: C.cast_tsequence_to_temporal(res)}
}

// TTextSeqSetFromBaseTstzspanset Return a temporal text sequence set from a text and a timestamptz span set
func TTextSeqSetFromBaseTstzspanset(value string, base *TsTzSpanSet) *TTextSeqSet {
	res := C.ttextseqset_from_base_tstzspanset(C.cstring2text(C.CString(value)), base._inner)
	return &TTextSeqSet{_inner: C.cast_tsequenceset_to_temporal(res)}
}

// TTextIn Return a temporal text from its Well-Known Text (WKT) representation
func TTextIn[TF TText](input string, output TF) TF {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_text := C.ttext_in(c_str)
	output.Init(c_text)
	return output
}

// TTextFromMFJSON Return a temporal text from its MF-JSON representation
func TTextFromMFJSON[TF TText](input string, output TF) TF {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_text := C.ttext_from_mfjson(c_str)
	output.Init(c_text)
	return output
}

// TTextOut Return a temporal text from its Well-Known Text (WKT) representation
func TTextOut[TF TText](tf TF) string {
	c_text := C.ttext_out(tf.Inner())
	defer C.free(unsafe.Pointer(c_text))
	text_out := C.GoString(c_text)
	return text_out
}

// TTextStartValue Return the start value of a temporal text
func TTextStartValue[TT TText](tt TT) string {
	cValue := C.ttext_start_value(tt.Inner())
	return C.GoString(C.text2cstring(cValue))
}

// TTextEndValue Return the end value of a temporal text
func TTextEndValue[TT TText](tt TT) string {
	cValue := C.ttext_end_value(tt.Inner())
	return C.GoString(C.text2cstring(cValue))
}

// TTextMinValue Return the minimum value of a temporal text
func TTextMinValue[TT TText](tt TT) string {
	cValue := C.ttext_min_value(tt.Inner())
	return C.GoString(C.text2cstring(cValue))
}

// TTextMaxValue Return the maximum value of a temporal text
func TTextMaxValue[TT TText](tt TT) string {
	cValue := C.ttext_max_value(tt.Inner())
	return C.GoString(C.text2cstring(cValue))
}

// TTextValueAtTimestamp Return the value of a temporal text at a timestamptz
func TTextValueAtTimestamp[TT TText](tt TT, ts time.Time) string {
	ttextinst, _ := TemporalToTTextInst(TemporalAtTimestamptz(tt, ts))
	return TTextStartValue(ttextinst)
}

// TTextUpper Return a temporal text transformed to uppercase
func TTextUpper[TT TText](tt TT, output TT) TT {
	c_temp := C.ttext_upper(tt.Inner())
	output.Init(c_temp)
	return output
}

// TTextLower Return a temporal text transformed to lowercase
func TTextLower[TT TText](tt TT, output TT) TT {
	c_temp := C.ttext_lower(tt.Inner())
	output.Init(c_temp)
	return output
}

// AlwaysLtTTextText Return true if a temporal text is always less than a text
func AlwaysLtTTextText[TT TText](tt TT, value string) bool {
	return int(C.always_lt_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))) > 0
}

// AlwaysLeTTextText Return true if a temporal text is always less than or equal to a text
func AlwaysLeTTextText[TT TText](tt TT, value string) bool {
	return int(C.always_le_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))) > 0
}

// AlwaysEqTTextText Return true if a temporal text is always equal to a text
func AlwaysEqTTextText[TT TText](tt TT, value string) bool {
	return int(C.always_eq_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))) > 0
}

// AlwaysNeTTextText Return true if a temporal text is always not equal to a text
func AlwaysNeTTextText[TT TText](tt TT, value string) bool {
	return int(C.always_ne_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))) > 0
}

// AlwaysGeTTextText Return true if a temporal text is always greater than or equal to a text
func AlwaysGeTTextText[TT TText](tt TT, value string) bool {
	return int(C.always_ge_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))) > 0
}

// AlwaysGtTTextText Return true if a temporal text is always greater than a text
func AlwaysGtTTextText[TT TText](tt TT, value string) bool {
	return int(C.always_gt_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))) > 0
}

// EverLtTTextText Return true if a temporal text is ever less than a text
func EverLtTTextText[TT TText](tt TT, value string) bool {
	return int(C.ever_lt_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))) > 0
}

// EverLeTTextText Return true if a temporal text is ever less than or equal to a text
func EverLeTTextText[TT TText](tt TT, value string) bool {
	return int(C.ever_le_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))) > 0
}

// EverEqTTextText Return true if a temporal text is ever equal to a text
func EverEqTTextText[TT TText](tt TT, value string) bool {
	return int(C.ever_eq_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))) > 0
}

// EverNeTTextText Return true if a temporal text is ever not equal to a text
func EverNeTTextText[TT TText](tt TT, value string) bool {
	return int(C.ever_ne_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))) > 0
}

// EverGeTTextText Return true if a temporal text is ever greater than or equal to a text
func EverGeTTextText[TT TText](tt TT, value string) bool {
	return int(C.ever_ge_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))) > 0
}

// EverGtTTextText Return true if a temporal text is ever greater than a text
func EverGtTTextText[TT TText](tt TT, value string) bool {
	return int(C.ever_gt_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))) > 0
}

// TEqTTextText Return a temporal value that represents where the temporal integer is equal to a constant text.
func TEqTTextText[TT TText](tt TT, value string) Temporal {
	c_temp := C.teq_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))
	return CreateTemporal(c_temp)
}

// TNeTTextText Return a temporal value that represents where the temporal integer is not equal to a constant text.
func TNeTTextText[TT TText](tt TT, value string) Temporal {
	c_temp := C.tne_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))
	return CreateTemporal(c_temp)
}

// TLtTTextText Return a temporal value that represents where the temporal integer is less than a constant text.
func TLtTTextText[TT TText](tt TT, value string) Temporal {
	c_temp := C.tlt_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))
	return CreateTemporal(c_temp)
}

// TLeTTextText Return a temporal value that represents where the temporal integer is less than or equal to a constant text.
func TLeTTextText[TT TText](tt TT, value string) Temporal {
	c_temp := C.tle_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))
	return CreateTemporal(c_temp)
}

// TGeTTextText Return a temporal value that represents where the temporal integer is greater than or equal to a constant text.
func TGeTTextText[TT TText](tt TT, value string) Temporal {
	c_temp := C.tge_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))
	return CreateTemporal(c_temp)
}

// TGtTTextText Return a temporal value that represents where the temporal integer is greater than a constant text.
func TGtTTextText[TT TText](tt TT, value string) Temporal {
	c_temp := C.tgt_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))
	return CreateTemporal(c_temp)
}

// TTextAtValue Return a temporal int restricted to a string
func TTextAtValue[TT TText](tt TT, value string) Temporal {
	c_ttexts := C.ttext_at_value(tt.Inner(), C.cstring2text(C.CString(value)))
	return CreateTemporal(c_ttexts)
}

// TTextMinusValue Return a temporal int restricted to a string
func TTextMinusValue[TT TText](tt TT, value string) Temporal {
	c_ttexts := C.ttext_minus_value(tt.Inner(), C.cstring2text(C.CString(value)))
	return CreateTemporal(c_ttexts)
}

// Return the concatenation of a temporal text and a text
func TextcatTTextText[TT TText](tt TT, value string) Temporal {
	c_temp := C.textcat_ttext_text(tt.Inner(), C.cstring2text(C.CString(value)))
	return CreateTemporal(c_temp)
}

// TextcatTTextTText Return the concatenation of two temporal text values
func TextcatTTextTText[TT1 TText, TT2 TText](tt1 TT1, tt2 TT2) Temporal {
	c_temp := C.textcat_ttext_ttext(tt1.Inner(), tt2.Inner())
	return CreateTemporal(c_temp)
}
