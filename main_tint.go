package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

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

// TIntFromBaseTemporal Return a temporal int from a int and the time frame of another temporal value
func TIntFromBaseTemporal[T1 Temporal, T2 Temporal](value int, base T1, output T2) T2 {
	c_temp := C.tint_from_base_temp(C.int(value), base.Inner())
	output.Init(c_temp)
	return output
}

// TIntInstMake Return a temporal int instant from a int and a timestamptz
func TIntInstMake(value int, base time.Time) *TIntInst {
	res := C.tintinst_make(C.int(value), DatetimeToTimestamptz(base))
	return &TIntInst{_inner: C.cast_tinstant_to_temporal(res)}
}

// TIntSeqFromBaseTstzset Return a temporal int discrete sequence from a int and a timestamptz set
func TIntSeqFromBaseTstzset(value int, base TsTzSet) *TIntSeq {
	res := C.tintseq_from_base_tstzset(C.int(value), base._inner)
	return &TIntSeq{_inner: C.cast_tsequence_to_temporal(res)}
}

// TIntSeqFromBaseTstzspan Return a temporal int sequence from a int and a timestamptz span
func TIntSeqFromBaseTstzspan(value int, base TsTzSpan) *TIntSeq {
	res := C.tintseq_from_base_tstzspan(C.int(value), base._inner)
	return &TIntSeq{_inner: C.cast_tsequence_to_temporal(res)}
}

// TIntSeqSetFromBaseTstzspanset Return a temporal int sequence set from a int and a timestamptz span set
func TIntSeqSetFromBaseTstzspanset(value int, base *TsTzSpanSet) *TIntSeqSet {
	res := C.tintseqset_from_base_tstzspanset(C.int(value), base._inner)
	return &TIntSeqSet{_inner: C.cast_tsequenceset_to_temporal(res)}
}

// TIntIn Return a temporal integer from its Well-Known Text (WKT) representation
func TIntIn[TI TInt](input string, output TI) TI {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_int := C.tint_in(c_str)
	output.Init(c_int)
	return output
}

// TIntFromMFJSON Return a temporal integer from its MF-JSON representation
func TIntFromMFJSON[TI TInt](input string, output TI) TI {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_int := C.tint_from_mfjson(c_str)
	output.Init(c_int)
	return output
}

// TIntOut Return a temporal integer from its Well-Known Text (WKT) representation
func TIntOut[TI TInt](ti TI) string {
	c_int := C.tint_out(ti.Inner())
	defer C.free(unsafe.Pointer(c_int))
	int_out := C.GoString(c_int)
	return int_out
}

// TIntToTFloat Return a temporal integer converted to a temporal float
func TIntToTFloat[TI TInt, TF TFloat](ti TI, output TF) TF {
	c_temp := C.tint_to_tfloat(ti.Inner())
	output.Init(c_temp)
	return output
}

// TIntStartValue Return the start value of a temporal int
func TIntStartValue[TB TInt](tb TB) int {
	cValue := C.tint_start_value(tb.Inner())
	return int(cValue)
}

// TIntEndValue Return the end value of a temporal int
func TIntEndValue[TB TInt](tb TB) int {
	cValue := C.tint_end_value(tb.Inner())
	return int(cValue)
}

// TIntValues Return the array of base values of a temporal integer
func TIntValues[TF TInt](tf TF) ([]int, error) {
	var count C.int

	cValues := C.tint_values(tf.Inner(), &count)
	if cValues == nil {
		return nil, fmt.Errorf("failed to retrieve bool values")
	}
	length := int(count)
	values := unsafe.Slice((*C.int)(cValues), length)
	// Convert the C bool values to Go bool values
	goValues := make([]int, length)
	for i := 0; i < length; i++ {
		goValues[i] = int(values[i])
	}
	return goValues, nil
}

// TIntMinValue Return the minimum value of a temporal int
func TIntMinValue[TB TInt](tb TB) int {
	cValue := C.tint_min_value(tb.Inner())
	return int(cValue)
}

// TIntMaxValue Return the maximum value of a temporal int
func TIntMaxValue[TB TInt](tb TB) int {
	cValue := C.tint_max_value(tb.Inner())
	return int(cValue)
}

// TIntValueAtTimestamp Return the value of a temporal int at a timestamptz
func TIntValueAtTimestamp[TF TInt](tf TF, ts time.Time) int {
	tintinst, _ := TemporalToTIntInst(TemporalAtTimestamptz(tf, ts))
	return TIntStartValue(tintinst)
}

// AlwaysLtTIntInt Return true if a temporal integer is always less than an integer
func AlwaysLtTIntInt[TI TInt](ti TI, value int) bool {
	return int(C.always_lt_tint_int(ti.Inner(), C.int(value))) > 0
}

// AlwaysLeTIntInt Return true if a temporal integer is always less than or equal to an integer
func AlwaysLeTIntInt[TI TInt](ti TI, value int) bool {
	return int(C.always_le_tint_int(ti.Inner(), C.int(value))) > 0
}

// AlwaysEqTIntInt Return true if a temporal integer is always equal to an integer
func AlwaysEqTIntInt[TI TInt](ti TI, value int) bool {
	return int(C.always_eq_tint_int(ti.Inner(), C.int(value))) > 0
}

// AlwaysNeTIntInt Return true if a temporal integer is always not equal to an integer
func AlwaysNeTIntInt[TI TInt](ti TI, value int) bool {
	return int(C.always_ne_tint_int(ti.Inner(), C.int(value))) > 0
}

// AlwaysGeTIntInt Return true if a temporal integer is always greater than or equal to an integer
func AlwaysGeTIntInt[TI TInt](ti TI, value int) bool {
	return int(C.always_ge_tint_int(ti.Inner(), C.int(value))) > 0
}

// AlwaysGtTIntInt Return true if a temporal integer is always greater than an integer
func AlwaysGtTIntInt[TI TInt](ti TI, value int) bool {
	return int(C.always_gt_tint_int(ti.Inner(), C.int(value))) > 0
}

// EverLtTIntInt Return true if a temporal integer is ever less than an integer
func EverLtTIntInt[TI TInt](ti TI, value int) bool {
	return int(C.ever_lt_tint_int(ti.Inner(), C.int(value))) > 0
}

// EverLeTIntInt Return true if a temporal integer is ever less than or equal to an integer
func EverLeTIntInt[TI TInt](ti TI, value int) bool {
	return int(C.ever_le_tint_int(ti.Inner(), C.int(value))) > 0
}

// EverEqTIntInt Return true if a temporal integer is ever equal to an integer
func EverEqTIntInt[TI TInt](ti TI, value int) bool {
	return int(C.ever_eq_tint_int(ti.Inner(), C.int(value))) > 0
}

// EverNeTIntInt Return true if a temporal integer is ever not equal to an integer
func EverNeTIntInt[TI TInt](ti TI, value int) bool {
	return int(C.ever_ne_tint_int(ti.Inner(), C.int(value))) > 0
}

// EverGeTIntInt Return true if a temporal integer is ever greater than or equal to an integer
func EverGeTIntInt[TI TInt](ti TI, value int) bool {
	return int(C.ever_ge_tint_int(ti.Inner(), C.int(value))) > 0
}

// EverGtTIntInt Return true if a temporal integer is ever greater than an integer
func EverGtTIntInt[TI TInt](ti TI, value int) bool {
	return int(C.ever_gt_tint_int(ti.Inner(), C.int(value))) > 0
}

// TEqTIntInt Return a temporal value that represents where the temporal integer is equal to a constant integer.
func TEqTIntInt[TI TInt](ti TI, value int) Temporal {
	c_temp := C.teq_tint_int(ti.Inner(), C.int(value))
	return CreateTemporal(c_temp)
}

// TNeTIntInt Return a temporal value that represents where the temporal integer is not equal to a constant integer.
func TNeTIntInt[TI TInt](ti TI, value int) Temporal {
	c_temp := C.tne_tint_int(ti.Inner(), C.int(value))
	return CreateTemporal(c_temp)
}

// TLtTIntInt Return a temporal value that represents where the temporal integer is less than a constant integer.
func TLtTIntInt[TI TInt](ti TI, value int) Temporal {
	c_temp := C.tlt_tint_int(ti.Inner(), C.int(value))
	return CreateTemporal(c_temp)
}

// TLeTIntInt Return a temporal value that represents where the temporal integer is less than or equal to a constant integer.
func TLeTIntInt[TI TInt](ti TI, value int) Temporal {
	c_temp := C.tle_tint_int(ti.Inner(), C.int(value))
	return CreateTemporal(c_temp)
}

// TGeTIntInt Return a temporal value that represents where the temporal integer is greater than or equal to a constant integer.
func TGeTIntInt[TI TInt](ti TI, value int) Temporal {
	c_temp := C.tge_tint_int(ti.Inner(), C.int(value))
	return CreateTemporal(c_temp)
}

// TGtTIntInt Return a temporal value that represents where the temporal integer is greater than a constant integer.
func TGtTIntInt[TI TInt](ti TI, value int) Temporal {
	c_temp := C.tgt_tint_int(ti.Inner(), C.int(value))
	return CreateTemporal(c_temp)
}

// TIntAtValue Return a temporal int restricted to a int
func TIntAtValue[TI TInt](tf TI, value int) Temporal {
	c_tints := C.tint_at_value(tf.Inner(), C.int(value))
	return CreateTemporal(c_tints)
}

// TIntMinusValue Return a temporal int restricted to a int
func TIntMinusValue[TI TInt](tf TI, value int) Temporal {
	c_tints := C.tint_minus_value(tf.Inner(), C.int(value))
	return CreateTemporal(c_tints)
}

// NadTIntInt Return the nearest approach distance between a temporal number and a number
func NadTIntInt[TI TInt](tf TI, value int) float64 {
	return float64(C.nad_tint_int(tf.Inner(), C.int(value)))
}

// NadTIntTInt Return the nearest approach distance between two temporal integers
func NadTIntTInt[TI1 TInt, TI2 TInt](tf1 TI1, tf2 TI2) float64 {
	return float64(C.nad_tint_tint(tf1.Inner(), tf2.Inner()))
}

// TIntShiftValue Return a temporal integer whose value dimension is shifted by a value
func TIntShiftValue[TI TInt](ti TI, delta int, output TI) TI {
	c_temp := C.tint_shift_value(ti.Inner(), C.int(delta))
	output.Init(c_temp)
	return output
}

// TIntScaleValue Return a temporal number whose value dimension is scaled by a value
func TIntScaleValue[TI TInt](ti TI, width int, output TI) TI {
	c_temp := C.tint_scale_value(ti.Inner(), C.int(width))
	output.Init(c_temp)
	return output
}

// TIntShitScaleValue Return a temporal number whose value dimension is scaled by a value
func TIntShitScaleValue[TI TInt](ti TI, shift int, width int, output TI) TI {
	c_temp := C.tint_shift_scale_value(ti.Inner(), C.int(shift), C.int(width))
	output.Init(c_temp)
	return output
}

// AddTIntInt returns the temporal addition of a temporal integer and a constant integer.
func AddTIntInt[TI TInt](ti TI, value int) Temporal {
	c_temp := C.add_tint_int(ti.Inner(), C.int(value))
	return CreateTemporal(c_temp)
}

// AddIntTInt returns the temporal addition of a constant integer and a temporal integer.
func AddIntTInt[TI TInt](value int, ti TI) Temporal {
	c_temp := C.add_int_tint(C.int(value), ti.Inner())
	return CreateTemporal(c_temp)
}

// SubTIntInt returns the temporal subtraction of a temporal integer and a constant integer.
func SubTIntInt[TI TInt](ti TI, value int) Temporal {
	c_temp := C.sub_tint_int(ti.Inner(), C.int(value))
	return CreateTemporal(c_temp)
}

// SubIntTInt returns the temporal subtraction of a constant integer from a temporal integer.
func SubIntTInt[TI TInt](value int, ti TI) Temporal {
	c_temp := C.sub_int_tint(C.int(value), ti.Inner())
	return CreateTemporal(c_temp)
}

// MultTIntInt returns the temporal multiplication of a temporal integer and a constant integer.
func MultTIntInt[TI TInt](ti TI, value int) Temporal {
	c_temp := C.mult_tint_int(ti.Inner(), C.int(value))
	return CreateTemporal(c_temp)
}

// MultIntTInt returns the temporal multiplication of a constant integer and a temporal integer.
func MultIntTInt[TI TInt](value int, ti TI) Temporal {
	c_temp := C.mult_int_tint(C.int(value), ti.Inner())
	return CreateTemporal(c_temp)
}

// DivTIntInt returns the temporal division of a temporal integer by a constant integer.
func DivTIntInt[TI TInt](ti TI, value int) Temporal {
	c_temp := C.div_tint_int(ti.Inner(), C.int(value))
	return CreateTemporal(c_temp)
}

// DivIntTInt returns the temporal division of a constant integer by a temporal integer.
func DivIntTInt[TI TInt](value int, ti TI) Temporal {
	c_temp := C.div_int_tint(C.int(value), ti.Inner())
	return CreateTemporal(c_temp)
}
