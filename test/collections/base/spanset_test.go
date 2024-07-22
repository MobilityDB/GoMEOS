package base_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos/collections/base"
	"github.com/alecthomas/assert/v2"
)

func createIntSpanSet() *base.IntSpanSet {
	g_iss_in := "{[1,3]}"
	g_iss := base.NewIntSpanSet(g_iss_in)
	return g_iss
}

func createIntSpanSet2() *base.IntSpanSet {
	g_iss_in := "{[1,3], [9,100]}"
	g_iss := base.NewIntSpanSet(g_iss_in)
	return g_iss
}

func createFloatSpanSet() *base.FloatSpanSet {
	g_fss_in := "{[1.5,3.6]}"
	g_fss := base.NewFloatSpanSet(g_fss_in)
	return g_fss
}

func TestNewSpanSet(t *testing.T) {
	g_iss := createIntSpanSet()
	g_fss := createFloatSpanSet()
	assert.NotEqual(t, g_iss, nil)
	assert.NotEqual(t, g_fss, nil)
}

func TestSpanSetOut(t *testing.T) {
	g_iss := createIntSpanSet()
	g_iss_out := g_iss.IntSpanSetOut()
	g_fss := createFloatSpanSet()
	g_fss_out := g_fss.FloatSpanSetOut(10)
	assert.Equal(t, g_iss_out, "{[1, 4)}")
	assert.Equal(t, g_fss_out, "{[1.5, 3.6]}")
}

func TestToSpan(t *testing.T) {
	g_iss := createIntSpanSet()
	g_fss := createFloatSpanSet()
	g_is := g_iss.ToSpan()
	g_fs := g_fss.ToSpan()
	g_is_out := g_is.IntSpanOut()
	g_fs_out := g_fs.FloatSpanOut(10)
	assert.Equal(t, g_is_out, "[1, 4)")
	assert.Equal(t, g_fs_out, "[1.5, 3.6]")
}

func TestConversion(t *testing.T) {
	g_iss := createIntSpanSet2()
	cfss := g_iss.ToFloatSpanSet()
	c_fss_out := cfss.FloatSpanSetOut(10)
	assert.Equal(t, c_fss_out, "{[1, 3], [9, 100]}")

	g_fss := createFloatSpanSet()
	ciss := g_fss.ToIntSpanSet()
	c_iss_out := ciss.IntSpanSetOut()
	assert.Equal(t, c_iss_out, "{[1, 4)}")
}
