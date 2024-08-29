package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"

// TNumberToTBox Return a temporal number converted to a temporal box
func TNumberToTBox[TN TNumber](tn TN) *TBox {
	return &TBox{_inner: C.tnumber_to_tbox(tn.Inner())}
}

// TNumberIntegral Return the integral (area under the curve) of a temporal number
func TNumberIntegral[TN TNumber](tn TN) float64 {
	return float64(C.tnumber_integral(tn.Inner()))
}

// TNumberTwavg Return the time-weighted average of a temporal number
func TNumberTwavg[TN TNumber](temp TN) float64 {
	res := C.tnumber_twavg(temp.Inner())
	return float64(res)
}

// TNumberAtSpan Return a temporal value restricted to a span of base values
func TNumberAtSpan[TN TNumber, S Span](tn TN, span S) Temporal {
	c_temp := C.tnumber_at_span(tn.Inner(), span.Inner())
	return CreateTemporal(c_temp)
}

// TNumberAtSpanSet Return a temporal value restricted to an array of spans of base values
func TNumberAtSpanSet[TN TNumber, SS SpanSet](tn TN, spanset SS) Temporal {
	c_temp := C.tnumber_at_spanset(tn.Inner(), spanset.Inner())
	return CreateTemporal(c_temp)
}

// TNumberAtTBox Return a temporal number restricted to a temporal box
func TNumberAtTBox[TN TNumber](tn TN, tbox *TBox) Temporal {
	c_temp := C.tnumber_at_tbox(tn.Inner(), tbox._inner)
	return CreateTemporal(c_temp)
}

// TNumberMinusSpan Return a temporal value restricted to the complement of a span of base values
func TNumberMinusSpan[TN TNumber, S Span](tn TN, span S) Temporal {
	c_temp := C.tnumber_minus_span(tn.Inner(), span.Inner())
	return CreateTemporal(c_temp)
}

// TNumberMinusSpanSet Return a temporal value restricted to the complement of an array of spans of base values
func TNumberMinusSpanSet[TN TNumber, SS SpanSet](tn TN, spanset SS) Temporal {
	c_temp := C.tnumber_minus_spanset(tn.Inner(), spanset.Inner())
	return CreateTemporal(c_temp)
}

// TNumberMinusTBox Return a temporal number restricted to a temporal box
func TNumberMinusTBox[TN TNumber](tn TN, tbox *TBox) Temporal {
	c_temp := C.tnumber_minus_tbox(tn.Inner(), tbox._inner)
	return CreateTemporal(c_temp)
}

// LeftTNumberTBox Return true if a temporal number is to the left of a temporal box
func LeftTNumberTBox[TN TNumber](tn TN, tbox *TBox) bool {
	return bool(C.left_tnumber_tbox(tn.Inner(), tbox._inner))
}

// LeftTNumberNumspan Return true if a temporal number is to the left of a number span
func LeftTNumberNumspan[TN TNumber, NS NumSpan](tn TN, ns NS) bool {
	return bool(C.left_tnumber_numspan(tn.Inner(), ns.Inner()))
}

// LeftTNumberTNumber Return true if the first temporal number is to the left of the second one
func LeftTNumberTNumber[TN1 TNumber, TN2 TNumber](tn TN1, tn2 TNumber) bool {
	return bool(C.left_tnumber_tnumber(tn.Inner(), tn2.Inner()))
}

// OverLeftTNumberNumspan returns true if a temporal number is overleft of a number span.
func OverLeftTNumberNumspan[TN TNumber, NS NumSpan](tn TN, ns NS) bool {
	return bool(C.overleft_tnumber_numspan(tn.Inner(), ns.Inner()))
}

// OverLeftTNumberTBox returns true if a temporal number is overleft of a temporal box.
func OverLeftTNumberTBox[TN TNumber](tn TN, tbox *TBox) bool {
	return bool(C.overleft_tnumber_tbox(tn.Inner(), tbox._inner))
}

// OverLeftTNumberTNumber returns true if the first temporal number is overleft of the second one.
func OverLeftTNumberTNumber[TN1 TNumber, TN2 TNumber](tn1 TN1, tn2 TN2) bool {
	return bool(C.overleft_tnumber_tnumber(tn1.Inner(), tn2.Inner()))
}

// RightTNumberTBox Return true if a temporal number is to the right of a temporal box
func RightTNumberTBox[TN TNumber](tn TN, tbox *TBox) bool {
	return bool(C.right_tnumber_tbox(tn.Inner(), tbox._inner))
}

// RightTNumberNumspan Return true if a temporal number is to the right of a number span
func RightTNumberNumspan[TN TNumber, NS NumSpan](tn TN, ns NS) bool {
	return bool(C.right_tnumber_numspan(tn.Inner(), ns.Inner()))
}

// RightTNumberTNumber Return true if the first temporal number is to the right of the second one
func RightTNumberTNumber[TN1 TNumber, TN2 TNumber](tn TN1, tn2 TNumber) bool {
	return bool(C.right_tnumber_tnumber(tn.Inner(), tn2.Inner()))
}

// OverRightTNumberNumspan returns true if a temporal number is overright of a number span.
func OverRightTNumberNumspan[TN TNumber, NS NumSpan](tn TN, ns NS) bool {
	return bool(C.overright_tnumber_numspan(tn.Inner(), ns.Inner()))
}

// OverRightTNumberTBox returns true if a temporal number is overright of a temporal box.
func OverRightTNumberTBox[TN TNumber](tn TN, tbox *TBox) bool {
	return bool(C.overright_tnumber_tbox(tn.Inner(), tbox._inner))
}

// OverRightTNumberTNumber returns true if the first temporal number is overright of the second one.
func OverRightTNumberTNumber[TN1 TNumber, TN2 TNumber](tn1 TN1, tn2 TN2) bool {
	return bool(C.overright_tnumber_tnumber(tn1.Inner(), tn2.Inner()))
}

// AddTNumberTNumber Return the temporal addition of the temporal numbers
func AddTNumberTNumber[TN1 TNumber, TN2 TNumber](tn1 TN1, tn2 TN2) Temporal {
	c_temp := C.add_tnumber_tnumber(tn1.Inner(), tn2.Inner())
	return CreateTemporal(c_temp)
}

// SubTNumberTNumber returns the temporal subtraction of two temporal numbers.
func SubTNumberTNumber[TN1 TNumber, TN2 TNumber](tn1 TN1, tn2 TN2) Temporal {
	c_temp := C.sub_tnumber_tnumber(tn1.Inner(), tn2.Inner())
	return CreateTemporal(c_temp)
}

// MultTNumberTNumber returns the temporal multiplication of two temporal numbers.
func MultTNumberTNumber[TN1 TNumber, TN2 TNumber](tn1 TN1, tn2 TN2) Temporal {
	c_temp := C.mult_tnumber_tnumber(tn1.Inner(), tn2.Inner())
	return CreateTemporal(c_temp)
}

// DivTNumberTNumber returns the temporal division of two temporal numbers.
func DivTNumberTNumber[TN1 TNumber, TN2 TNumber](tn1 TN1, tn2 TN2) Temporal {
	c_temp := C.div_tnumber_tnumber(tn1.Inner(), tn2.Inner())
	return CreateTemporal(c_temp)
}

// TNumberAbs Return the absolute value of a temporal number
func TNumberAbs[TN TNumber](tn TN, output TN) TN {
	c_temp := C.tnumber_abs(tn.Inner())
	output.Init(c_temp)
	return output
}

// TNumberDeltaValue Return the delta value of a temporal number
func TNumberDeltaValue[TN TNumber](tn TN) Temporal {
	c_temp := C.tnumber_delta_value(tn.Inner())
	return CreateTemporal(c_temp)
}

// DistanceTNumberTNumber returns the temporal distance between two temporal numbers.
func DistanceTNumberTNumber[TN1 TNumber, TN2 TNumber](tn1 TN1, tn2 TN2) Temporal {
	c_temp := C.distance_tnumber_tnumber(tn1.Inner(), tn2.Inner())
	return CreateTemporal(c_temp)
}

// TNumberToSpan Return the value span of a temporal number
func TNumberToSpan[TN TNumber, S Span](tn TN, output S) S {
	span := C.tnumber_to_span(tn.Inner())
	output.Init(span)
	return output
}

// TNumberValueSpans Return the base values of a temporal number as a span set
func TNumberValueSpans[TN TNumber, SS SpanSet](tn TN, ss SS) SS {
	spanset := C.tnumber_valuespans(tn.Inner())
	ss.Init(spanset)
	return ss
}
