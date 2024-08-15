package temporal_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos"
	"github.com/alecthomas/assert/v2"
)

func createTBoolInst() *gomeos.TBoolInst {
	return gomeos.NewTBoolInst("TRUE@2022-10-01")
}

func createTBoolSeq() *gomeos.TBoolSeq {
	return gomeos.NewTBoolSeq("{TRUE@2022-10-01, TRUE@2022-10-02}")
}

func TestNewCreateTBoolInst(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createTBoolInst()
	assert.Equal(t, g_is.TBoolOut(), "t@2022-10-01 00:00:00+00")
	// // gomeos.MeosFinalize()
}

func TestNewCreateTBoolSeq(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createTBoolSeq()
	assert.Equal(t, g_is.TBoolOut(), "{t@2022-10-01 00:00:00+00, t@2022-10-02 00:00:00+00}")
	// // gomeos.MeosFinalize()
}

// func TestTBoolInOut(t *testing.T) {
// 	gomeos.MeosInitialize("UTC")
// 	tb := gomeos.TBoolIn("{t@2022-10-01 00:00:00+00, t@2022-10-02 00:00:00+00}", &gomeos.TBoolSeq{})
// 	assert.Equal(t, gomeos.TBoolOut(tb), "{t@2022-10-01 00:00:00+00, t@2022-10-02 00:00:00+00}")
// }
