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

type STBox struct {
	_inner *C.STBox
}

func GeoToSTBox(geom *Geom) *STBox {
	return &STBox{_inner: C.geo_to_stbox(geom._inner)}
}

func TimestamptzToSTBox(t time.Time) *STBox {
	return &STBox{_inner: C.timestamptz_to_stbox(DatetimeToTimestamptz(t))}
}

func TsTzSetToSTBox(tstzset *TsTzSet) *STBox {
	return &STBox{_inner: C.tstzset_to_stbox(tstzset._inner)}
}

func TsTzSpanToSTBox(tstzspan *TsTzSpan) *STBox {
	return &STBox{_inner: C.tstzspan_to_stbox(tstzspan._inner)}
}

func TsTzSpanSetToSTBox(tstzspanset *TsTzSpanSet) *STBox {
	return &STBox{_inner: C.tstzspanset_to_stbox(tstzspanset._inner)}
}

func STBoxIn(s string) *STBox {
	return &STBox{_inner: C.stbox_in(C.CString(s))}
}

func STBoxMake(hasx bool, hasz bool, geodetic bool, srid int, xmin float64, xmax float64, ymin float64, ymax float64, zmin float64, zmax float64, tstzspan *TsTzSpan) *STBox {
	return &STBox{_inner: C.stbox_make(C.bool(hasx), C.bool(hasz), C.bool(geodetic), C.int32(srid), C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.double(zmin), C.double(zmax), tstzspan._inner)}
}

func STBoxCopy(stbox *STBox) *STBox {
	return &STBox{_inner: C.stbox_copy(stbox._inner)}
}

// STBoxFromWKB Return a temporal stbox from its Well-Known Binary (WKB) representation
func STBoxFromWKB(wkb []byte) *STBox {
	// Get the size of the WKB byte slice
	size := C.size_t(len(wkb))

	// Convert the Go byte slice to a C uint8_t* pointer
	wkbPtr := (*C.uint8_t)(unsafe.Pointer(&wkb[0]))

	// Call the C function
	stbox := C.stbox_from_wkb(wkbPtr, size)
	return &STBox{stbox}
}

// STBoxFromHexWKB Return a temporal stbox from its hex-encoded ASCII Well-Known Binary (WKB) representation
func STBoxFromHexWKB(s string) *STBox {
	c_str := C.CString(s)
	defer C.free(unsafe.Pointer(c_str))
	c_stbox := C.stbox_from_hexwkb(c_str)
	stbox := &STBox{_inner: c_stbox}
	return stbox
}

func GeoTimestamptzToSTBox(geom *Geom, t time.Time) *STBox {
	return &STBox{_inner: C.geo_timestamptz_to_stbox(geom._inner, DatetimeToTimestamptz(t))}
}

func STBoxOut(stbox *STBox, max_decimals int) string {
	return C.GoString(C.stbox_out(stbox._inner, C.int(max_decimals)))
}

// STBoxAsWKB Returns the stbox object as a hex-encoded WKB string.
func STBoxAsWKB(stbox *STBox) ([]byte, error) {
	var sizeOut C.size_t
	wkbPtr := C.stbox_as_wkb(stbox._inner, C.uint8_t(4), &sizeOut)
	if wkbPtr == nil {
		return nil, fmt.Errorf("failed to convert stbox to WKB")
	}
	defer C.free(unsafe.Pointer(wkbPtr)) // Free the memory allocated by C
	// Convert the C uint8_t* to a Go []byte slice
	length := int(sizeOut)
	wkb := C.GoBytes(unsafe.Pointer(wkbPtr), C.int(length))
	return wkb, nil
}

// STBoxAsHexWKB Return the hex-encoded ASCII Well-Known Binary (HexWKB) representation of a stbox
func STBoxAsHexWKB(stbox *STBox) (string, error) {
	var sizeOut C.size_t
	wkbPtr := C.stbox_as_hexwkb(stbox._inner, C.uint8_t(4), &sizeOut)
	if wkbPtr == nil {
		return "", fmt.Errorf("failed to convert tbox to HEXWKB")
	}
	defer C.free(unsafe.Pointer(wkbPtr)) // Free the memory allocated by C
	wkb := C.GoString(wkbPtr)
	return wkb, nil
}

func STBoxToGeo(stbox *STBox) *Geom {
	return &Geom{_inner: C.stbox_to_geo(stbox._inner)}
}

func STBoxToTsTzSpan(stbox *STBox) *TsTzSpan {
	return &TsTzSpan{_inner: C.stbox_to_tstzspan(stbox._inner)}
}

func STBoxHasX(stbox *STBox) bool {
	return bool(C.stbox_hasx(stbox._inner))
}

func STBoxHasZ(stbox *STBox) bool {
	return bool(C.stbox_hasz(stbox._inner))
}

func STBoxHasT(stbox *STBox) bool {
	return bool(C.stbox_hast(stbox._inner))
}

func STBoxIsGeodetic(stbox *STBox) bool {
	return bool(C.stbox_isgeodetic(stbox._inner))
}
