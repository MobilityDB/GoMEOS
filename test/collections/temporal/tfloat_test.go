package temporal_test

import (
	"testing"

	gomeos "github.com/MobilityDB/GoMEOS"
	"github.com/alecthomas/assert/v2"
)

func createTFloat() *gomeos.TFloatInst {
	return gomeos.NewTFloatInst("0.12344@2000-01-01")
}

func TestNewCreateTFloat(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createTFloat()
	assert.Equal(t, g_is.TPointOut(5), "0.12344@2000-01-01 00:00:00+00")
	// // gomeos.MeosFinalize()
}

// func TestTFloatInOut(t *testing.T) {
// 	gomeos.MeosInitialize("UTC")
// 	tf := gomeos.TFloatIn("{1.5@2022-10-01 00:00:00+00, 2.5@2022-10-02 00:00:00+00}", &gomeos.TFloatSeq{})
// 	assert.Equal(t, gomeos.TFloatOut(tf, 2), "{1.5@2022-10-01 00:00:00+00, 2.5@2022-10-02 00:00:00+00}")
// }

func TestTFloatToTInt(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createTFloat()
	res, _ := gomeos.TFloatToTInt(g_is)
	assert.Equal(t, res.String(), "0@2000-01-01 00:00:00+00")
	// // gomeos.MeosFinalize()
}
