package temporal_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos"
	"github.com/alecthomas/assert/v2"
)

func createTGeogPointInst() *gomeos.TGeogPointInst {
	return gomeos.NewTGeogPointInst("POINT(1 1)@2000-01-01")
}

func TestNewCreateTGeogPointInst(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createTGeogPointInst()
	assert.Equal(t, g_is.TPointOut(5), "POINT(1 1)@2000-01-01 00:00:00+00")
	// // gomeos.MeosFinalize()
}
