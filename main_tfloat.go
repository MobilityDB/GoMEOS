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

func (tf *TFloatInst) IsTInstant() bool {
	return true
}

func (tf *TFloatInst) IsTNumber() bool {
	return true
}

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

func (tf *TFloatSeq) IsTNumber() bool {
	return true
}

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

func (tf *TFloatSeqSet) IsTNumber() bool {
	return true
}

// TFloatFromBaseTemporal Return a temporal float from a float and the time frame of another temporal value
func TFloatFromBaseTemporal[T1 Temporal, T2 Temporal](value float64, base T1, output T2) T2 {
	c_temp := C.tfloat_from_base_temp(C.double(value), base.Inner())
	output.Init(c_temp)
	return output
}

// TFloatInstMake Return a temporal float instant from a float and a timestamptz
func TFloatInstMake(value float64, base time.Time) *TFloatInst {
	res := C.tfloatinst_make(C.double(value), DatetimeToTimestamptz(base))
	return &TFloatInst{_inner: C.cast_tinstant_to_temporal(res)}
}

// TFloatSeqFromBaseTstzset Return a temporal float discrete sequence from a float and a timestamptz set
func TFloatSeqFromBaseTstzset(value float64, base TsTzSet) *TFloatSeq {
	res := C.tfloatseq_from_base_tstzset(C.double(value), base._inner)
	return &TFloatSeq{_inner: C.cast_tsequence_to_temporal(res)}
}

// TFloatSeqFromBaseTstzspan Return a temporal float sequence from a float and a timestamptz span
func TFloatSeqFromBaseTstzspan(value float64, base TsTzSpan, interp Interpolation) *TFloatSeq {
	res := C.tfloatseq_from_base_tstzspan(C.double(value), base._inner, C.interpType(interp))
	return &TFloatSeq{_inner: C.cast_tsequence_to_temporal(res)}
}

// TFloatSeqSetFromBaseTstzspanset Return a temporal float sequence set from a float and a timestamptz span set
func TFloatSeqSetFromBaseTstzspanset(value float64, base *TsTzSpanSet, interp Interpolation) *TFloatSeqSet {
	res := C.tfloatseqset_from_base_tstzspanset(C.double(value), base._inner, C.interpType(interp))
	return &TFloatSeqSet{_inner: C.cast_tsequenceset_to_temporal(res)}
}

// TFloatIn Return a temporal float from its Well-Known Text (WKT) representation
func TFloatIn[TF TFloat](input string, output TF) TF {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_float := C.tfloat_in(c_str)
	output.Init(c_float)
	return output
}

// TFloatFromMFJSON Return a temporal float from its MF-JSON representation
func TFloatFromMFJSON[TF TFloat](input string, output TF) TF {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_float := C.tfloat_from_mfjson(c_str)
	output.Init(c_float)
	return output
}

// TFloatOut Return a temporal float from its Well-Known Text (WKT) representation
func TFloatOut[TF TFloat](tf TF, maxdd int) string {
	c_float := C.tfloat_out(tf.Inner(), C.int(maxdd))
	defer C.free(unsafe.Pointer(c_float))
	float_out := C.GoString(c_float)
	return float_out
}

// TFloatToTInt Return a temporal float converted to a temporal integer
func TFloatToTInt[TF TFloat, TI TInt](tf TF, ti TI) (TI, error) {
	interp := TemporalInterpolation(tf)
	if interp == "LINEAR" {
		return ti, fmt.Errorf("cannot convert a temporal float with linear interpolation to a temporal integer")
	}
	c_temp := C.tfloat_to_tint(tf.Inner())
	ti.Init(c_temp)
	return ti, nil
}

// TFloatStartValue Return the start value of a temporal float
func TFloatStartValue[TB TFloat](tb TB) float64 {
	cValue := C.tfloat_start_value(tb.Inner())
	return float64(cValue)
}

// TFloatEndValue Return the end value of a temporal float
func TFloatEndValue[TB TFloat](tb TB) float64 {
	cValue := C.tfloat_end_value(tb.Inner())
	return float64(cValue)
}

// TFloatValues Return the array of base values of a temporal float
func TFloatValues[TF TFloat](tf TF) ([]float64, error) {
	var count C.int

	cValues := C.tfloat_values(tf.Inner(), &count)
	if cValues == nil {
		return nil, fmt.Errorf("failed to retrieve bool values")
	}
	length := int(count)
	values := unsafe.Slice((*C.double)(cValues), length)
	// Convert the C bool values to Go bool values
	goValues := make([]float64, length)
	for i := 0; i < length; i++ {
		goValues[i] = float64(values[i])
	}
	return goValues, nil
}

// TFloatMinValue Return the minimum value of a temporal float
func TFloatMinValue[TB TFloat](tb TB) float64 {
	cValue := C.tfloat_min_value(tb.Inner())
	return float64(cValue)
}

// TFloatMaxValue Return the maximum value of a temporal float
func TFloatMaxValue[TB TFloat](tb TB) float64 {
	cValue := C.tfloat_max_value(tb.Inner())
	return float64(cValue)
}

// AlwaysLtTFloatFloat Return true if a temporal float is always less than a float
func AlwaysLtTFloatFloat[TF TFloat](tf TF, value float64) bool {
	return int(C.always_lt_tfloat_float(tf.Inner(), C.double(value))) > 0
}

// AlwaysLeTFloatFloat Return true if a temporal float is always less than or equal to a float
func AlwaysLeTFloatFloat[TF TFloat](tf TF, value float64) bool {
	return int(C.always_le_tfloat_float(tf.Inner(), C.double(value))) > 0
}

// AlwaysEqTFloatFloat Return true if a temporal float is always equal to a float
func AlwaysEqTFloatFloat[TF TFloat](tf TF, value float64) bool {
	return int(C.always_eq_tfloat_float(tf.Inner(), C.double(value))) > 0
}

// AlwaysNeTFloatFloat Return true if a temporal float is always not equal to a float
func AlwaysNeTFloatFloat[TF TFloat](tf TF, value float64) bool {
	return int(C.always_ne_tfloat_float(tf.Inner(), C.double(value))) > 0
}

// AlwaysGeTFloatFloat Return true if a temporal float is always greater than or equal to a float
func AlwaysGeTFloatFloat[TF TFloat](tf TF, value float64) bool {
	return int(C.always_ge_tfloat_float(tf.Inner(), C.double(value))) > 0
}

// AlwaysGtTFloatFloat Return true if a temporal float is always greater than a float
func AlwaysGtTFloatFloat[TF TFloat](tf TF, value float64) bool {
	return int(C.always_gt_tfloat_float(tf.Inner(), C.double(value))) > 0
}

// EverLtTFloatFloat Return true if a temporal float is ever less than a float
func EverLtTFloatFloat[TF TFloat](tf TF, value float64) bool {
	return int(C.ever_lt_tfloat_float(tf.Inner(), C.double(value))) > 0
}

// EverLeTFloatFloat Return true if a temporal float is ever less than or equal to a float
func EverLeTFloatFloat[TF TFloat](tf TF, value float64) bool {
	return int(C.ever_le_tfloat_float(tf.Inner(), C.double(value))) > 0
}

// EverEqTFloatFloat Return true if a temporal float is ever equal to a float
func EverEqTFloatFloat[TF TFloat](tf TF, value float64) bool {
	return int(C.ever_eq_tfloat_float(tf.Inner(), C.double(value))) > 0
}

// EverNeTFloatFloat Return true if a temporal float is ever not equal to a float
func EverNeTFloatFloat[TF TFloat](tf TF, value float64) bool {
	return int(C.ever_ne_tfloat_float(tf.Inner(), C.double(value))) > 0
}

// EverGeTFloatFloat Return true if a temporal float is ever greater than or equal to a float
func EverGeTFloatFloat[TF TFloat](tf TF, value float64) bool {
	return int(C.ever_ge_tfloat_float(tf.Inner(), C.double(value))) > 0
}

// EverGtTFloatFloat Return true if a temporal float is ever greater than a float
func EverGtTFloatFloat[TF TFloat](tf TF, value float64) bool {
	return int(C.ever_gt_tfloat_float(tf.Inner(), C.double(value))) > 0
}

// TEqTFloatFloat Return a temporal value that represents where the temporal float is equal to a constant float.
func TEqTFloatFloat[TF TFloat](tf TF, value float64) Temporal {
	c_temp := C.teq_tfloat_float(tf.Inner(), C.double(value))
	return CreateTemporal(c_temp)
}

// TNeTFloatFloat Return a temporal value that represents where the temporal float is not equal to a constant float.
func TNeTFloatFloat[TF TFloat](tf TF, value float64) Temporal {
	c_temp := C.tne_tfloat_float(tf.Inner(), C.double(value))
	return CreateTemporal(c_temp)
}

// TLtTFloatFloat Return a temporal value that represents where the temporal float is less than a constant float.
func TLtTFloatFloat[TF TFloat](tf TF, value float64) Temporal {
	c_temp := C.tlt_tfloat_float(tf.Inner(), C.double(value))
	return CreateTemporal(c_temp)
}

// TLeTFloatFloat Return a temporal value that represents where the temporal float is less than or equal to a constant float.
func TLeTFloatFloat[TF TFloat](tf TF, value float64) Temporal {
	c_temp := C.tle_tfloat_float(tf.Inner(), C.double(value))
	return CreateTemporal(c_temp)
}

// TGeTFloatFloat Return a temporal value that represents where the temporal float is greater than or equal to a constant float.
func TGeTFloatFloat[TF TFloat](tf TF, value float64) Temporal {
	c_temp := C.tge_tfloat_float(tf.Inner(), C.double(value))
	return CreateTemporal(c_temp)
}

// TGtTFloatFloat Return a temporal value that represents where the temporal float is greater than a constant float.
func TGtTFloatFloat[TF TFloat](tf TF, value float64) Temporal {
	c_temp := C.tgt_tfloat_float(tf.Inner(), C.double(value))
	return CreateTemporal(c_temp)
}

// TFloatAtValue Return a temporal float restricted to a float
func TFloatAtValue[TF TFloat](tf TF, value float64) Temporal {
	c_tbools := C.tfloat_at_value(tf.Inner(), C.double(value))
	return CreateTemporal(c_tbools)
}

// TFloatMinusValue Return a temporal float restricted to a float
func TFloatMinusValue[TF TFloat](tf TF, value float64) Temporal {
	c_tbools := C.tfloat_minus_value(tf.Inner(), C.double(value))
	return CreateTemporal(c_tbools)
}

// TFloatValueAtTimestamp Return the value of a temporal float at a timestamptz
func TFloatValueAtTimestamp[TF TFloat](tf TF, ts time.Time) float64 {
	tfloatinst, _ := TemporalToTFloatInst(TemporalAtTimestamptz(tf, ts))
	return TFloatStartValue(tfloatinst)
}

// TFloatDerivative Return the derivative of a temporal number
func TFloatDerivative[TF TFloat](tf TF) Temporal {
	return CreateTemporal(C.tfloat_derivative(tf.Inner()))
}

// TFloatToDegrees Return a temporal number transformed from radians to degrees
func TFloatToDegrees[TF TFloat](tf TF, normalize bool) Temporal {
	return CreateTemporal(C.tfloat_degrees(tf.Inner(), C.bool(normalize)))
}

// TFloatToRadians Return a temporal number transformed from degrees to radians
func TFloatToRadians[TF TFloat](tf TF) Temporal {
	return CreateTemporal(C.tfloat_radians(tf.Inner()))
}

// TFloatRound Return a temporal float with the precision of the values set to a number of decimal places
func TFloatRound[TF TFloat](tf TF, max_decimals int) Temporal {
	return CreateTemporal(C.tfloat_round(tf.Inner(), C.int(max_decimals)))
}

// TFloatShiftValue Return a temporal integer whose value dimension is shifted by a value
func TFloatShiftValue[TF TFloat](ti TF, delta int, output TF) TF {
	c_temp := C.tfloat_shift_value(ti.Inner(), C.double(delta))
	output.Init(c_temp)
	return output
}

// TFloatScaleValue Return a temporal number whose value dimension is scaled by a value
func TFloatScaleValue[TF TFloat](ti TF, width int, output TF) TF {
	c_temp := C.tfloat_scale_value(ti.Inner(), C.double(width))
	output.Init(c_temp)
	return output
}

// TFloatShitScaleValue Return a temporal number whose value dimension is scaled by a value
func TFloatShitScaleValue[TF TFloat](ti TF, shift int, width int, output TF) TF {
	c_temp := C.tfloat_shift_scale_value(ti.Inner(), C.double(shift), C.double(width))
	output.Init(c_temp)
	return output
}

// AddTFloatFloat returns the temporal addition of a temporal float and a constant float.
func AddTFloatFloat[TF TFloat](tf TF, value float64) Temporal {
	c_temp := C.add_tfloat_float(tf.Inner(), C.double(value))
	return CreateTemporal(c_temp)
}

// AddFloatTFloat returns the temporal addition of a constant float and a temporal float.
func AddFloatTFloat[TF TFloat](value float64, tf TF) Temporal {
	c_temp := C.add_float_tfloat(C.double(value), tf.Inner())
	return CreateTemporal(c_temp)
}

// SubTFloatFloat returns the temporal subtraction of a temporal float and a constant float.
func SubTFloatFloat[TF TFloat](tf TF, value float64) Temporal {
	c_temp := C.sub_tfloat_float(tf.Inner(), C.double(value))
	return CreateTemporal(c_temp)
}

// SubFloatTFloat returns the temporal subtraction of a constant float from a temporal float.
func SubFloatTFloat[TF TFloat](value float64, tf TF) Temporal {
	c_temp := C.sub_float_tfloat(C.double(value), tf.Inner())
	return CreateTemporal(c_temp)
}

// MultTFloatFloat returns the temporal multiplication of a temporal float and a constant float.
func MultTFloatFloat[TF TFloat](tf TF, value float64) Temporal {
	c_temp := C.mult_tfloat_float(tf.Inner(), C.double(value))
	return CreateTemporal(c_temp)
}

// MultFloatTFloat returns the temporal multiplication of a constant float and a temporal float.
func MultFloatTFloat[TF TFloat](value float64, tf TF) Temporal {
	c_temp := C.mult_float_tfloat(C.double(value), tf.Inner())
	return CreateTemporal(c_temp)
}

// DivTFloatFloat returns the temporal division of a temporal float by a constant float.
func DivTFloatFloat[TF TFloat](tf TF, value float64) Temporal {
	c_temp := C.div_tfloat_float(tf.Inner(), C.double(value))
	return CreateTemporal(c_temp)
}

// DivFloatTFloat returns the temporal division of a constant float by a temporal float.
func DivFloatTFloat[TF TFloat](value float64, tf TF) Temporal {
	c_temp := C.div_float_tfloat(C.double(value), tf.Inner())
	return CreateTemporal(c_temp)
}

// DistanceTFloatFloat returns the temporal distance between a temporal float and a constant float.
func DistanceTFloatFloat[TF TFloat](tf TF, value float64) Temporal {
	c_temp := C.distance_tfloat_float(tf.Inner(), C.double(value))
	return CreateTemporal(c_temp)
}

// NADTFloatFloat returns the nearest approach distance between a temporal float and a constant float.
func NADTFloatFloat[TF TFloat](tf TF, value float64) float64 {
	c_temp := C.nad_tfloat_float(tf.Inner(), C.double(value))
	return float64(c_temp)
}

// NADTFloatTFloat returns the nearest approach distance between two temporal floats.
func NADTFloatTFloat[TF1 TFloat, TF2 TFloat](tf1 TF1, tf2 TF2) float64 {
	c_temp := C.nad_tfloat_tfloat(tf1.Inner(), tf2.Inner())
	return float64(c_temp)
}
