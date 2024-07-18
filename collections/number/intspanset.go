// collections/number/intspanset.go
package number

/*
#cgo CFLAGS: -I/opt/homebrew/include
#cgo LDFLAGS: -L/opt/homebrew/lib -lmeos -Wl,-rpath,/opt/homebrew/lib
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

type IntSpanSet struct {
	_inner *C.SpanSet
}

func NewIntSpanSet(g_iss_in string) IntSpanSet {
	c_iss_in := C.CString(g_iss_in)
	defer C.free(unsafe.Pointer(c_iss_in))
	c_iss := C.intspanset_in(c_iss_in)
	g_iss := IntSpanSet{_inner: c_iss}
	return g_iss
}

// ------------------------- Output ----------------------------------------
/*
Return the string representation of the content of IntSpanSet.

Returns:
	String

MEOS Functions:
	intspanset_out
*/
func (g_iss IntSpanSet) IntSpanSetOut() string {
	c_iss_out := C.intspanset_out(g_iss._inner)
	defer C.free(unsafe.Pointer(c_iss_out))
	g_iss_out := C.GoString(c_iss_out)
	return g_iss_out
}

// ------------------------- Conversions -----------------------------------
/*
Returns a span that encompasses _inner.

Returns:
	A new struct `IntSpan` instance

MEOS Functions:
	spanset_span
*/
func (g_iss IntSpanSet) ToSpan() IntSpan {
	return IntSpan{_inner: C.spanset_span(g_iss._inner)}
}

/*
Converts IntSpanSet to a FloatSpanSet instance.

Returns:

	A new FloatSpanSet instance

MEOS Functions:

	intspanset_to_floatspanset
*/
func (g_iss IntSpanSet) ToFloatSpanSet() FloatSpanSet {
	return FloatSpanSet{_inner: C.intspanset_to_floatspanset(g_iss._inner)}
}

// ignore gap:no default parameter in Go
// float32?or others?

// ------------------------- Accessors -------------------------------------
/*
Returns the number of spans in IntSpanSet.
Returns:
	An int

MEOS Functions:
	spanset_num_spans
*/
func (g_iss IntSpanSet) NumSpans() int {
	return int(C.spanset_num_spans(g_iss._inner))
}

/*
Returns the width of the spanset. By default, i.e., when the second
argument is False, the function takes into account the gaps within,
i.e., returns the sum of the widths of the spans within.
Otherwise, the function returns the width of the spanset ignoring
any gap, i.e., the width from the lower bound of the first span to
the upper bound of the last span.

Parameters:

	ignore_gaps: Whether to take into account potential gaps in
	the spanset.

Returns:

	A `float` representing the duration of the spanset

MEOS Functions:

	intspanset_width
*/
func (g_iss IntSpanSet) Width() int {
	ignore_gap := C._Bool(false)
	return int(C.intspanset_width(g_iss._inner, ignore_gap))
}

/*
Returns the first span in spanset.
Returns:

	A IntSpan instance

MEOS Functions:

	spanset_start_span
*/
func (g_iss IntSpanSet) StartSpan() IntSpan {
	return IntSpan{_inner: C.spanset_start_span(g_iss._inner)}
}

/*
Returns the last span in IntSpanSet.
Returns:

	A IntSpan instance

MEOS Functions:

	spanset_end_span
*/
func (g_iss IntSpanSet) EndSpan() IntSpan {
	return IntSpan{_inner: C.spanset_end_span(g_iss._inner)}
}

/*
Returns the n-th span in IntSpanSet.
Returns:

	A IntSpan instance

MEOS Functions:

	spanset_span_n
*/
func (g_iss IntSpanSet) SpanN(n int) IntSpan {
	return IntSpan{_inner: C.spanset_span_n(g_iss._inner, C.int(n))}
}

/*
Returns the list of spans in IntSpanSet.
Returns:

	A IntSpan instance

MEOS Functions:

	spanset_spans
*/
func (g_iss IntSpanSet) Spans() []IntSpan {
	nums := g_iss.NumSpans()
	spans := make([]IntSpan, nums)
	for i := 1; i < nums+1; i++ {
		spans[i-1] = g_iss.SpanN(i)
	}
	return spans
}

// ------------------------- Transformations -------------------------------

/*
Return a new “IntSpanSet“ with the lower and upper bounds shifted by
“delta“.

Args:

	delta: The value to shift by

Returns:

	A new ``IntSpanSet`` instance

MEOS Functions:

	intspanset_shift_scale
*/
func (g_iss IntSpanSet) ShiftScale(delta int, width int) IntSpanSet {
	d := 0
	if delta != 0 {
		d = delta
	}

	w := 0
	if width != 0 {
		w = width
	}
	modified := C.intspanset_shift_scale(g_iss._inner, C.int(d), C.int(w), C._Bool(delta != 0), C._Bool(width != 0))
	return IntSpanSet{_inner: modified}
}

/*
Return a new “IntSpanSet“ with the lower and upper bounds shifted by
“delta“.

Args:

	delta: The value to shift by

Returns:

	A new ``IntSpanSet`` instance

MEOS Functions:

	intspanset_shift_scale
*/
func (g_iss IntSpanSet) Shift(delta int, width int) IntSpanSet {
	return g_iss.ShiftScale(delta, 0)

}

/*
Return a new “IntSpanSet“ with the lower and upper bounds scaled so
that the width is “width“.

Args:

	width: The new width

Returns:

	A new ``IntSpanSet`` instance

MEOS Functions:

	intspanset_shift_scale
*/
func (g_iss IntSpanSet) Scale(delta int, width int) IntSpanSet {
	return g_iss.ShiftScale(0, width)
}

// ------------------------- Topological Operations --------------------------------
