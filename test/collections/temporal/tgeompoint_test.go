package temporal_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos"
	"github.com/alecthomas/assert/v2"
)

func createTGeomPointInst() gomeos.TGeomPointInst {
	return gomeos.NewTGeomPointInst("POINT(1 1)@2000-01-01")
}

func TestNewCreateTGeomPointInst(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createTGeomPointInst()
	assert.Equal(t, g_is.TPointOut(5), "POINT(1 1)@2000-01-01 00:00:00+00")
	// // gomeos.MeosFinalize()
}

func TestTGeomPointInOut(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	tg := gomeos.TGeomPointIn("{POINT(1 1)@2022-10-01 00:00:00+00, POINT(2 2)@2022-10-02 00:00:00+00}", &gomeos.TGeomPointSeq{})
	assert.Equal(t, gomeos.TPointOut(tg, 5), "{0101000000000000000000F03F000000000000F03F@2022-10-01 00:00:00+00, 010100000000000000000000400000000000000040@2022-10-02 00:00:00+00}")
}
