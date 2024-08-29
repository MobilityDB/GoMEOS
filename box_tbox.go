package gomeos

/*
#include "meos.h"
#include "meos_catalog.h"
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

type TBox struct {
	_inner *C.TBox
}

// TBoxIn Return a temporal box from its Well-Known Text (WKT) representation
func TBoxIn(s string) *TBox {
	c_str := C.CString(s)
	defer C.free(unsafe.Pointer(c_str))
	c_tbox := C.tbox_in(c_str)
	tbox := &TBox{_inner: c_tbox}
	return tbox
}

// TBoxMake Return a temporal box from a number span and a timestamptz span
func TBoxMake[S Span](s S, ts *TsTzSpan) *TBox {
	tbox := C.tbox_make(s.Inner(), ts._inner)
	return &TBox{tbox}
}

// TBoxCopy Return a copy of a temporal box
func TBoxCopy(tbox *TBox) *TBox {
	return &TBox{tbox._inner}
}

// TBoxFromWKB Return a temporal box from its Well-Known Binary (WKB) representation
func TBoxFromWKB(wkb []byte) *TBox {
	// Get the size of the WKB byte slice
	size := C.size_t(len(wkb))

	// Convert the Go byte slice to a C uint8_t* pointer
	wkbPtr := (*C.uint8_t)(unsafe.Pointer(&wkb[0]))

	// Call the C function
	tbox := C.tbox_from_wkb(wkbPtr, size)
	return &TBox{tbox}
}

// TBoxFromHexWKB Return a temporal box from its hex-encoded ASCII Well-Known Binary (WKB) representation
func TBoxFromHexWKB(s string) *TBox {
	c_str := C.CString(s)
	defer C.free(unsafe.Pointer(c_str))
	c_tbox := C.tbox_from_hexwkb(c_str)
	tbox := &TBox{_inner: c_tbox}
	return tbox
}

// IntToTBox Return an integer converted to a temporal box
func IntToTBox(value int) *TBox {
	return &TBox{C.int_to_tbox(C.int(value))}
}

// FloatToTBox Return a float converted to a temporal box
func FloatToTBox(value float64) *TBox {
	return &TBox{C.float_to_tbox(C.double(value))}
}

// SetToTBox Return a number set converted to a temporal box
func SetToTBox[S Set](s S) *TBox {
	return &TBox{C.set_to_tbox(s.Inner())}
}

// SpanToTBox Return a number span converted to a temporal box
func SpanToTBox[S Span](s S) *TBox {
	return &TBox{C.span_to_tbox(s.Inner())}
}

// SpanSetToTBox Return a number spanset converted to a temporal box
func SpanSetToTBox[S SpanSet](s S) *TBox {
	return &TBox{C.spanset_to_tbox(s.Inner())}
}

// TimestamptzToTBox Return a timestamptz converted to a temporal box
func TimestamptzToTBox(ts time.Time) *TBox {
	return &TBox{C.timestamptz_to_tbox(DatetimeToTimestamptz(ts))}
}

// IntTimestamptzToTBox Return a temporal box from an integer and a timestamptz
func IntTimestamptzToTBox(value int, t time.Time) *TBox {
	return &TBox{_inner: C.int_timestamptz_to_tbox(C.int(value), DatetimeToTimestamptz(t))}
}

// FloatTsTzSpanToTBox Return a temporal box from a float and a timestamptz span
func FloatTsTzSpanToTBox[S Span](value float64, span S) *TBox {
	return &TBox{_inner: C.float_tstzspan_to_tbox(C.double(value), span.Inner())}
}

// FloatTimestamptzToTBox Return a temporal box from a float and a timestamptz
func FloatTimestamptzToTBox(value float64, t time.Time) *TBox {
	return &TBox{_inner: C.float_timestamptz_to_tbox(C.double(value), DatetimeToTimestamptz(t))}
}

// IntTsTzSpanToTBox Return a temporal box from an integer and a timestamptz span
func IntTsTzSpanToTBox[S Span](value int, span S) *TBox {
	return &TBox{_inner: C.int_tstzspan_to_tbox(C.int(value), span.Inner())}
}

// NumSpanTsTzSpanToTBox Return a temporal box from a number span and a timestamptz span
func NumSpanTsTzSpanToTBox[S NumSpan](numspan S, tstzspan *TsTzSpan) *TBox {
	return &TBox{_inner: C.numspan_tstzspan_to_tbox(numspan.Inner(), tstzspan._inner)}
}

// NumSpanTimestamptzToTBox Return a temporal box from a number span and a timestamptz
func NumSpanTimestamptzToTBox[S NumSpan](numspan S, t time.Time) *TBox {
	return &TBox{_inner: C.numspan_timestamptz_to_tbox(numspan.Inner(), DatetimeToTimestamptz(t))}
}

// TBoxOut Return the Well-Known Text (WKT) representation of a temporal box
func TBoxOut(tbox *TBox, max_decimals int) string {
	return C.GoString(C.tbox_out(tbox._inner, C.int(max_decimals)))
}

// TBoxAsWKB Returns the tbox object as a hex-encoded WKB string.
func TBoxAsWKB(tbox *TBox) ([]byte, error) {
	var sizeOut C.size_t
	wkbPtr := C.tbox_as_wkb(tbox._inner, C.uint8_t(4), &sizeOut)
	if wkbPtr == nil {
		return nil, fmt.Errorf("failed to convert tbox to WKB")
	}
	defer C.free(unsafe.Pointer(wkbPtr)) // Free the memory allocated by C
	// Convert the C uint8_t* to a Go []byte slice
	length := int(sizeOut)
	wkb := C.GoBytes(unsafe.Pointer(wkbPtr), C.int(length))
	return wkb, nil
}

// TBoxAsHexWKB Return the hex-encoded ASCII Well-Known Binary (HexWKB) representation of a tbox
func TBoxAsHexWKB(tbox *TBox) (string, error) {
	var sizeOut C.size_t
	wkbPtr := C.tbox_as_hexwkb(tbox._inner, C.uint8_t(4), &sizeOut)
	if wkbPtr == nil {
		return "", fmt.Errorf("failed to convert tbox to HEXWKB")
	}
	defer C.free(unsafe.Pointer(wkbPtr)) // Free the memory allocated by C
	wkb := C.GoString(wkbPtr)
	return wkb, nil
}

// TBoxToFloatSpan Return a temporal box converted to a float span
func TBoxToFloatSpan(tbox *TBox) *FloatSpan {
	return &FloatSpan{C.tbox_to_floatspan(tbox._inner)}
}

// TBoxToTsTzSpan Return a spatiotemporal box converted to a timestamptz span
func TBoxToTsTzSpan(tbox *TBox) *TsTzSpan {
	return &TsTzSpan{C.tbox_to_tstzspan(tbox._inner)}
}
