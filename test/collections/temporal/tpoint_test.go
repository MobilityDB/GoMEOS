package temporal_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos"
	"github.com/alecthomas/assert/v2"
)

func TestTPointOut(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	tg := gomeos.TGeomPointIn("{POINT(1 1)@2022-10-01 00:00:00+00, POINT(2 2)@2022-10-02 00:00:00+00}", &gomeos.TGeomPointSeq{})
	assert.Equal(t, gomeos.TPointAsEWKT(tg, 5), "{POINT(1 1)@2022-10-01 00:00:00+00, POINT(2 2)@2022-10-02 00:00:00+00}")
}

func TestTPointToSTBox(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	tg := gomeos.TGeomPointIn("{POINT(1 1)@2022-10-01 00:00:00+00, POINT(2 2)@2022-10-02 00:00:00+00}", &gomeos.TGeomPointSeq{})
	stbox := gomeos.TPointToSTBox(tg)
	assert.Equal(t, gomeos.STBoxOut(stbox, 10), "STBOX XT(((1,1),(2,2)),[2022-10-01 00:00:00+00, 2022-10-02 00:00:00+00])")
}

func TestTPointGeoAsWKT(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	tg := gomeos.TGeomPointIn("POINT(1 1)@2022-10-01 00:00:00+00", &gomeos.TGeomPointSeq{})
	geom := gomeos.TPointTrajectory(tg)
	res := gomeos.GeoAsEWKT(geom, 10)
	assert.Equal(t, res, "POINT(1 1)")
}
