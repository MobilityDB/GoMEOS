package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/

import "C"

// ------------------------- Accessors -------------------------------------
func TNumberToTBox[TN TNumber](tn TN) *TBox {
	tbox := C.tnumber_to_tbox(tn.Inner())
	return &TBox{_inner: tbox}
}

func TNumberIntegral[TN TNumber](tn TN) float64 {
	return float64(C.tnumber_integral(tn.Inner()))
}

func TnumberTwavg[TN TNumber](temp TN) float64 {
	res := C.tnumber_twavg(temp.Inner())
	return float64(res)
}

// ------------------------- TODO:Transformations -------------------------------

// ------------------------- Restrictions ----------------------------------

func TnumberAtSpan[TN TNumber, S Span](tn TN, span S) Temporal {
	c_temp := C.tnumber_at_span(tn.Inner(), span.Inner())
	return CreateTemporal(c_temp)
}

func TnumberAtSpanSet[TN TNumber, SS SpanSet](tn TN, spanset SS) Temporal {
	c_temp := C.tnumber_at_spanset(tn.Inner(), spanset.Inner())
	return CreateTemporal(c_temp)
}

func TnumberAtTBox[TN TNumber](tn TN, tbox *TBox) Temporal {
	c_temp := C.tnumber_at_tbox(tn.Inner(), tbox._inner)
	return CreateTemporal(c_temp)
}

func TnumberMinusSpan[TN TNumber, S Span](tn TN, span S) Temporal {
	c_temp := C.tnumber_minus_span(tn.Inner(), span.Inner())
	return CreateTemporal(c_temp)
}

func TnumberMinusSpanSet[TN TNumber, SS SpanSet](tn TN, spanset SS) Temporal {
	c_temp := C.tnumber_minus_spanset(tn.Inner(), spanset.Inner())
	return CreateTemporal(c_temp)
}

func TnumberMinusTBox[TN TNumber](tn TN, tbox *TBox) Temporal {
	c_temp := C.tnumber_minus_tbox(tn.Inner(), tbox._inner)
	return CreateTemporal(c_temp)
}

// ------------------------- TODO:Position Operations ---------------------------

// ------------------------- TODO:Mathematical Operations -------------------------

// ------------------------- TODO:Distance Operations --------------------------
