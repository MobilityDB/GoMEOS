package gomeos

/*
#include "meos.h"
#include "meos_catalog.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"

func TNumberAtSpan[TN TNumber, S Span](tn TN, span S) Temporal {
	c_temp := C.tnumber_at_span(tn.Inner(), span.Inner())
	return CreateTemporal(c_temp)
}

func TNumberAtSpanSet[TN TNumber, SS SpanSet](tn TN, spanset SS) Temporal {
	c_temp := C.tnumber_at_spanset(tn.Inner(), spanset.Inner())
	return CreateTemporal(c_temp)
}

func TNumberAtTBox[TN TNumber](tn TN, tbox *TBox) Temporal {
	c_temp := C.tnumber_at_tbox(tn.Inner(), tbox._inner)
	return CreateTemporal(c_temp)
}

func TNumberMinusSpan[TN TNumber, S Span](tn TN, span S) Temporal {
	c_temp := C.tnumber_minus_span(tn.Inner(), span.Inner())
	return CreateTemporal(c_temp)
}

func TNumberMinusSpanSet[TN TNumber, SS SpanSet](tn TN, spanset SS) Temporal {
	c_temp := C.tnumber_minus_spanset(tn.Inner(), spanset.Inner())
	return CreateTemporal(c_temp)
}

func TNumberMinusTBox[TN TNumber](tn TN, tbox *TBox) Temporal {
	c_temp := C.tnumber_minus_tbox(tn.Inner(), tbox._inner)
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
