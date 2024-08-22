package main

import (
	"fmt"

	gomeos "github.com/MobilityDB/GoMEOS"
)

func main() {
	// Initialize MEOS or similar library
	gomeos.MeosInitialize("UTC")

	// Input geometries in WKT format
	pointWKT := "POINT(1 1)"
	linestringWKT := "LINESTRING(1 1,2 2,1 1)"
	polygonWKT := "POLYGON((1 1,1 2,2 2,2 1,1 1))"

	// Read WKT into geometries
	point := gomeos.PgisGeometryIn(pointWKT, -1)
	linestring := gomeos.PgisGeometryIn(linestringWKT, -1)
	polygon := gomeos.PgisGeometryIn(polygonWKT, -1)

	// Convert geometries to WKT
	pointText := gomeos.GeoAsText(point, 6)
	linestringText := gomeos.GeoAsText(linestring, 6)
	polygonText := gomeos.GeoAsText(polygon, 6)

	// Revert generated WKT strings into geometries
	point1 := gomeos.PgisGeometryIn(pointText, -1)
	linestring1 := gomeos.PgisGeometryIn(linestringText, -1)
	polygon1 := gomeos.PgisGeometryIn(polygonText, -1)

	// Ensure that the reverted types are equal to the original ones
	if !gomeos.GeoSame(point, point1) {
		fmt.Printf("ERROR: Distinct input and output geometries in WKT\n%s\n%s", pointWKT, pointText)
	}
	if !gomeos.GeoSame(linestring, linestring1) {
		fmt.Printf("ERROR: Distinct input and output geometries in WKT\n%s\n%s", linestringWKT, linestringText)
	}
	if !gomeos.GeoSame(polygon, polygon1) {
		fmt.Printf("ERROR: Distinct input and output geometries in WKT\n%s\n%s", polygonWKT, polygonText)
	}

	// Convert geometries to GeoJSON
	pointGeoJSON := gomeos.GeoAsGeojson(point, 1, 6, "EPSG:4326")
	linestringGeoJSON := gomeos.GeoAsGeojson(linestring, 1, 6, "EPSG:4326")
	polygonGeoJSON := gomeos.GeoAsGeojson(polygon, 1, 6, "EPSG:4326")

	// Revert generated GeoJSON strings into geometries
	point2 := gomeos.GeoFromGeojson(pointGeoJSON)
	linestring2 := gomeos.GeoFromGeojson(linestringGeoJSON)
	polygon2 := gomeos.GeoFromGeojson(polygonGeoJSON)

	// Ensure that the reverted types are equal to the original ones
	if !gomeos.GeoSame(point, point2) {
		fmt.Printf("ERROR: Distinct input and output geometries in GeoJSON\n%s\n%s", pointWKT, pointGeoJSON)
	}
	if !gomeos.GeoSame(linestring, linestring2) {
		fmt.Printf("ERROR: Distinct input and output geometries in GeoJSON\n%s\n%s", linestringWKT, linestringGeoJSON)
	}
	if !gomeos.GeoSame(polygon, polygon2) {
		fmt.Printf("ERROR: Distinct input and output geometries in GeoJSON\n%s\n%s", polygonWKT, polygonGeoJSON)
	}

	// Convert geometries to HexEWKB
	pointHexEWKB := gomeos.GeoAsHexEwkb(point, "XDR")
	linestringHexEWKB := gomeos.GeoAsHexEwkb(linestring, "XDR")
	polygonHexEWKB := gomeos.GeoAsHexEwkb(polygon, "XDR")

	// Revert generated HexEWKB strings into geometries
	point3 := gomeos.GeometryFromHexEwkb(pointHexEWKB)
	linestring3 := gomeos.GeometryFromHexEwkb(linestringHexEWKB)
	polygon3 := gomeos.GeometryFromHexEwkb(polygonHexEWKB)

	// Ensure that the reverted types are equal to the original ones
	if !gomeos.GeoSame(point, point3) {
		fmt.Printf("ERROR: Distinct input and output geometries in HexEWKB\n%s\n%s", pointWKT, pointText)
	}
	if !gomeos.GeoSame(linestring, linestring3) {
		fmt.Printf("ERROR: Distinct input and output geometries in HexEWKB\n%s\n%s", linestringWKT, linestringText)
	}
	if !gomeos.GeoSame(polygon, polygon3) {
		fmt.Printf("ERROR: Distinct input and output geometries in HexEWKB\n%s\n%s", polygonWKT, polygonText)
	}

	// Print results
	fmt.Printf("\n--------\n| Point |\n--------\n\nWKT:\n----\n%s\n\nText:\n-----\n%s\n\nGeoJSON:\n--------\n%s\nHexWKB:\n-------\n%s\n",
		pointWKT, pointText, pointGeoJSON, pointHexEWKB)
	fmt.Printf("\n-------------\n| Linestring |\n-------------\n\nWKT:\n----\n%s\n\nText:\n-----\n%s\n\nGeoJSON:\n--------\n%s\nHexWKB:\n-------\n%s\n",
		linestringWKT, linestringText, linestringGeoJSON, linestringHexEWKB)
	fmt.Printf("\n----------\n| Polygon |\n----------\n\nWKT:\n----\n%s\n\nText:\n-----\n%s\n\nGeoJSON:\n--------\n%s\nHexWKB:\n-------\n%s\n",
		polygonWKT, polygonText, polygonGeoJSON, polygonHexEWKB)

	// Clean up allocated objects (In Go, the garbage collector handles this)
	// Finalize MEOS
	gomeos.MeosFinalize()
}
