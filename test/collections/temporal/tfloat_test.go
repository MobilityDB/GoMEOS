package temporal_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos"
	"github.com/alecthomas/assert/v2"
)

func createTFloat() gomeos.TFloatInst {
	return gomeos.NewTFloatInst("0.12344@2000-01-01")
}

func TestNewCreateTFloat(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createTFloat()
	assert.Equal(t, g_is.TPointOut(5), "0.12344@2000-01-01 00:00:00+00")
	// // gomeos.MeosFinalize()
}
