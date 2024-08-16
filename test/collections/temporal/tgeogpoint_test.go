package temporal_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos"
	"github.com/alecthomas/assert/v2"
)

func createTGeogPointInst() *gomeos.TGeogPointInst {
	return gomeos.NewTGeogPointInst("POINT(1 1)@2000-01-01")
}

func createTGeogPointInst2() *gomeos.TGeogPointInst {
	return gomeos.NewTGeogPointInst("POINT(1 2)@2000-01-02")
}

func createTGeogPointInst3() *gomeos.TGeogPointInst {
	return gomeos.NewTGeogPointInst("POINT(1 20)@2000-01-05")
}

func TestNewCreateTGeogPointInst(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createTGeogPointInst()
	assert.Equal(t, g_is.TPointOut(5), "POINT(1 1)@2000-01-01 00:00:00+00")
	// // gomeos.MeosFinalize()
}

// func TestSequenceMake(t *testing.T) {
// 	gomeos.MeosInitialize("UTC")
// 	g1 := createTGeogPointInst()
// 	g2 := createTGeogPointInst2()
// 	g3 := createTGeogPointInst3()
// 	var gs []gomeos.TGeogPointInst
// 	gs = append(gs, g1)
// 	gs = append(gs, g2)
// 	gs = append(gs, g3)
// 	gseq := gomeos.TGeogPointSeqMake(gs, len(gs), true, true, "STEP", true)
// 	assert.Equal(t, gseq.TPointOut(5), "Interp=Step;[POINT(1 1)@2000-01-01 00:00:00+00, POINT(1 2)@2000-01-02 00:00:00+00, POINT(1 20)@2000-01-05 00:00:00+00]")
// 	gomeos.MeosFinalize()
// }

// func TestTGeogPointInOut(t *testing.T) {
// 	gomeos.MeosInitialize("UTC")
// 	tg := gomeos.TGeogPointIn("{POINT(1 1)@2022-10-01 00:00:00+00, POINT(2 2)@2022-10-02 00:00:00+00}", &gomeos.TGeogPointSeq{})
// 	assert.Equal(t, gomeos.TPointOut(tg, 5), "{0101000020E6100000000000000000F03F000000000000F03F@2022-10-01 00:00:00+00, 0101000020E610000000000000000000400000000000000040@2022-10-02 00:00:00+00}")
// }
