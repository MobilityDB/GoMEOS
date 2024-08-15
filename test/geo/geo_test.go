package number_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos"
	"github.com/alecthomas/assert/v2"
)

func createGeom() *gomeos.Geom {
	polygonWKT := "POLYGON((1 1,1 2,2 2,2 1,1 1))"
	point := gomeos.PgisGeometryIn(polygonWKT, -1)
	return point
}

func TestGeom(t *testing.T) {
	g := createGeom()
	assert.Equal(t, gomeos.GeoAsText(g, 5), "POLYGON((1 1,1 2,2 2,2 1,1 1))")
	assert.Equal(t, gomeos.GeoAsGeojson(g, 1, 6), "{\"type\":\"Polygon\",\"bbox\":[1.000000,1.000000,2.000000,2.000000],\"coordinates\":[[[1,1],[1,2],[2,2],[2,1],[1,1]]]}")
	assert.Equal(t, gomeos.GeoAsHexEwkb(g, "XDR"), "000000000300000001000000053FF00000000000003FF00000000000003FF000000000000040000000000000004000000000000000400000000000000040000000000000003FF00000000000003FF00000000000003FF0000000000000")
	assert.Equal(t, gomeos.GeoAsEWKT(g, 5), "POLYGON((1 1,1 2,2 2,2 1,1 1))")
}
