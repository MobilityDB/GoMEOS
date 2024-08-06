package temporal_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos"
	"github.com/alecthomas/assert/v2"
)

func createTInt() *gomeos.TIntInst {
	return gomeos.NewTIntInst("1@2000-01-01")
}

func TestNewCreateTInt(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createTInt()
	assert.Equal(t, g_is.TPointOut(), "1@2000-01-01 00:00:00+00")
	// // gomeos.MeosFinalize()
}
