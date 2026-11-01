package gomeos

// Every generated type reads and writes each format the catalog gives it under
// the portable name of the SQL function, <Class>From<Format> and As<Format>,
// and prints as its asText, a spatial value's points written as WKT. The
// examples reach the generated package as #ExampleTemporalFromPointer of
// example_pointer_boundary_test.go does, through the pointer both share.

import (
	"bytes"
	"fmt"

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
	hex, _ := point.AsHexEWKB(0)
	read, err := types.TGeomPointFromHexEWKB(hex)
	fmt.Println(read, err)
	// Output:
	// POINT(1 1)@2024-12-06 00:00:00+00 <nil>
}

func Example_sequenceAtGeometry() {
	// A sequence carries the methods of its leaf, and a geometry argument is
	// a *Geo: TGeomPointSeq reaches TPoint.AtGeom through TGeomPoint, whose
	// restriction MEOS answers as a sequence set. It reads the value as
	// #ExampleTGeomPointFromText does.
	temp, _ := types.TGeomPointFromText("[Point(0 0)@2024-12-06, Point(2 0)@2024-12-08]")
	seq := types.TGeomPointSeqFromPointer(temp.Pointer())
	box, _ := types.GeometryFromText("Polygon((0 -1,1 -1,1 1,0 1,0 -1))", 0)
	at, err := seq.AtGeom(&box.Geo)
	fmt.Println(types.TGeomPointFromPointer(at.Pointer()), err)
	// Output:
	// {[POINT(0 0)@2024-12-06 00:00:00+00, POINT(1 0)@2024-12-07 00:00:00+00]} <nil>
}

func Example_cbufferString() {
	cbuffer, _ := types.CbufferFromText("Cbuffer(Point(1 1),0.5)")
	fmt.Println(cbuffer)
	// Output:
	// Cbuffer(POINT(1 1),0.5)
}

func Example_floatSpanString() {
	span, _ := types.FloatSpanFromText("[8, 10]")
	fmt.Println(span)
	// Output:
	// [8, 10]
}

func Example_temporalEWKB() {
	point, _ := types.TGeomPointFromEWKT("SRID=4326;Point(1 1)@2024-12-06")
	wkb, err := point.AsEWKB(0)
	read, _ := types.TGeomPointFromEWKB(wkb)
	again, _ := read.AsEWKB(0)
	fmt.Println(len(wkb) > 0, err, bytes.Equal(wkb, again), read)
	// Output:
	// true <nil> true POINT(1 1)@2024-12-06 00:00:00+00
}
