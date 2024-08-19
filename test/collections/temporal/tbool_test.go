package temporal_test

import (
	"fmt"
	"testing"
	"time"

	gomeos "github.com/MobilityDB/GoMEOS"
	"github.com/alecthomas/assert/v2"
)

func createTBoolInst() *gomeos.TBoolInst {
	return gomeos.NewTBoolInst("TRUE@2022-10-01")
}

func createTBoolSeq() *gomeos.TBoolSeq {
	return gomeos.NewTBoolSeq("{FALSE@2022-10-01, FALSE@2022-10-02,TRUE@2022-10-03}")
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

func TestTBoolValues(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createTBoolSeq()
	res, _ := gomeos.TBoolValueSet(g_is)
	var bools = []bool{false, true}
	fmt.Println(res)
	assert.Equal(t, res, bools)

	// // gomeos.MeosFinalize()
}

func TestTBoolStartValue(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createTBoolSeq()
	res := gomeos.TBoolStartValue(g_is)
	fmt.Println(res)
	assert.Equal(t, res, true)

	// // gomeos.MeosFinalize()
}

func TestTBoolValueAtTimestamp(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createTBoolSeq()
	ts, _ := time.Parse("2006-01-02", "2022-10-01")
	res := gomeos.TBoolValueAtTimestamp(g_is, ts)
	fmt.Println(res)
	assert.Equal(t, res, false)

	// // gomeos.MeosFinalize()
}

func TestTBoolTemporalValueAtTimestamp(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createTBoolSeq()
	ts, _ := time.Parse("2006-01-02", "2022-10-01")
	res := gomeos.TemporalAtTimestamptz(g_is, ts)
	fmt.Println(res.String())
	// assert.Equal(t, res, true)

	// // gomeos.MeosFinalize()
}

func TestTBoolAtValue(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createTBoolSeq()
	res := gomeos.TBoolAtValue(g_is, false)
	assert.Equal(t, res.String(), "{f@2022-10-01 00:00:00+00, f@2022-10-02 00:00:00+00}")
}

func TestTBoolMinusValue(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createTBoolSeq()
	res := gomeos.TBoolMinusValue(g_is, false)
	assert.Equal(t, res.String(), "{t@2022-10-03 00:00:00+00}")
}
