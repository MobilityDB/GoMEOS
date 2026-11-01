package gomeos

// Every generated type reads and writes each format the catalog gives it under
// the portable name of the SQL function, <Class>From<Format> and As<Format>,
// and prints as its asText, a spatial value's points written as WKT. The
// examples reach the generated package as #ExampleTemporalFromPointer of
// example_pointer_boundary_test.go does, through the pointer both share.

import (
	"fmt"
	"unsafe"

	"github.com/MobilityDB/GoMEOS/types"
)

func ExampleTGeomPointFromText() {
	temp, err := types.TGeomPointFromText("[Point(1 1)@2024-12-06, Point(2 2)@2024-12-07]")
	fmt.Println(temp, err)
	// Output:
	// [POINT(1 1)@2024-12-06 00:00:00+00, POINT(2 2)@2024-12-07 00:00:00+00] <nil>
}

func ExampleTGeomPointFromEWKT() {
	point, _ := types.TGeomPointFromEWKT("SRID=4326;Point(1 1)@2024-12-06")
	ewkt, err := point.AsEWKT(15)
	fmt.Println(ewkt, err)
	// Output:
	// SRID=4326;POINT(1 1)@2024-12-06 00:00:00+00 <nil>
}

func ExampleTGeomPointFromHexEWKB() {
	point, _ := types.TGeomPointFromEWKT("SRID=4326;Point(1 1)@2024-12-06")
	var size uint64 // the size_t the writer states the length in
	hex, _ := point.AsHexEWKB(0, unsafe.Pointer(&size))
	read, err := types.TGeomPointFromHexEWKB(hex)
	fmt.Println(read, err)
	// Output:
	// POINT(1 1)@2024-12-06 00:00:00+00 <nil>
}

func ExampleCbuffer_String() {
	cbuffer, _ := types.CbufferFromText("Cbuffer(Point(1 1),0.5)")
	fmt.Println(cbuffer)
	// Output:
	// Cbuffer(POINT(1 1),0.5)
}

func ExampleFloatSpan_String() {
	span, _ := types.FloatSpanFromText("[8, 10]")
	fmt.Println(span)
	// Output:
	// [8, 10]
}
