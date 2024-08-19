package temporal_test

import (
	"testing"

	gomeos "github.com/MobilityDB/GoMEOS"
	"github.com/alecthomas/assert/v2"
)

func createTInt2() *gomeos.TIntInst {
	return gomeos.NewTIntInst("1@2000-01-01")
}

func TestNewCreateTInt(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createTInt()
	assert.Equal(t, g_is.TIntOut(), "1@2000-01-01 00:00:00+00")
	// // gomeos.MeosFinalize()
}

// func TestTIntInOut(t *testing.T) {
// 	gomeos.MeosInitialize("UTC")
// 	ti := gomeos.TIntIn("{1@2022-10-01 00:00:00+00, 2@2022-10-02 00:00:00+00}", &gomeos.TIntSeq{})
// 	assert.Equal(t, gomeos.TIntOut(ti), "{1@2022-10-01 00:00:00+00, 2@2022-10-02 00:00:00+00}")
// }
