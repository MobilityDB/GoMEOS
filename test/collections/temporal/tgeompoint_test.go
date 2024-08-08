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
