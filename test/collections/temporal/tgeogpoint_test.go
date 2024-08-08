package temporal_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos"
	"github.com/alecthomas/assert/v2"
)

func createTGeogPointInst() gomeos.TGeogPointInst {
	return gomeos.NewTGeogPointInst("POINT(1 1)@2000-01-01")
}

func createTGeogPointInst2() gomeos.TGeogPointInst {
	return gomeos.NewTGeogPointInst("POINT(1 2)@2000-01-02")
}

func createTGeogPointInst3() gomeos.TGeogPointInst {
	return gomeos.NewTGeogPointInst("POINT(1 20)@2000-01-05")
}

func TestNewCreateTGeogPointInst(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createTGeogPointInst()
	assert.Equal(t, g_is.TPointOut(5), "POINT(1 1)@2000-01-01 00:00:00+00")
	// // gomeos.MeosFinalize()
}

func TestSequenceMake(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g1 := createTGeogPointInst()
	g2 := createTGeogPointInst2()
	g3 := createTGeogPointInst3()
	var gs []gomeos.TGeogPointInst
	gs = append(gs, g1)
	gs = append(gs, g2)
	gs = append(gs, g3)
	gseq := gomeos.TGeogPointSeqMake(gs, len(gs), true, true, "STEP", true)
	assert.Equal(t, gseq.TPointOut(5), "Interp=Step;[POINT(1 1)@2000-01-01 00:00:00+00, POINT(1 2)@2000-01-02 00:00:00+00, POINT(1 20)@2000-01-05 00:00:00+00]")
	gomeos.MeosFinalize()
}
