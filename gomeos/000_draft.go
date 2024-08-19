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