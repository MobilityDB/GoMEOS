package gomeos

import (
	"fmt"
)

func createGeom() *Geom {
	polygonWKT := "POLYGON((1 1,1 2,2 2,2 1,1 1))"
	point := PgisGeometryIn(polygonWKT, -1)
	return point
}

func ExampleGeoAsText() {
	g := createGeom()
	result := GeoAsText(g, 5)
	fmt.Println(result)
	// Output:
	// POLYGON((1 1,1 2,2 2,2 1,1 1))
}

func ExampleGeoAsGeojson() {
	g := createGeom()
	result := GeoAsGeojson(g, 1, 6, "EPSG:4326")
	fmt.Println(result)
	// Output:
	// {"type":"Polygon","crs":{"type":"name","properties":{"name":"EPSG:4326"}},"bbox":[1.000000,1.000000,2.000000,2.000000],"coordinates":[[[1,1],[1,2],[2,2],[2,1],[1,1]]]}
}

func ExampleGeoAsHexEwkb() {
	g := createGeom()
	result := GeoAsHexEwkb(g, "XDR")
	fmt.Println(result)
	// Output:
	// 000000000300000001000000053FF00000000000003FF00000000000003FF000000000000040000000000000004000000000000000400000000000000040000000000000003FF00000000000003FF00000000000003FF0000000000000
}

func ExampleGeoAsEWKT() {
	g := createGeom()
	result := GeoAsEWKT(g, 5)
	fmt.Println(result)
	// Output:
	// POLYGON((1 1,1 2,2 2,2 1,1 1))
}
