package base_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos/collections/base"
	"github.com/alecthomas/assert/v2"
)

func createIntSpan() *base.IntSpan {
	g_is_in := "[1,3]"
	g_is := base.NewIntSpan(g_is_in)
	return g_is
}

func createFloatSpan() *base.FloatSpan {
	g_fs_in := "[1.5,3.6]"
	g_fs := base.NewFloatSpan(g_fs_in)
	return g_fs
}

func TestNewSpan(t *testing.T) {
	g_iss := createIntSpan()
	g_fss := createFloatSpan()
	assert.NotEqual(t, g_iss, nil)
	assert.NotEqual(t, g_fss, nil)
}

func TestSpanOut(t *testing.T) {
	g_iss := createIntSpan()
	g_iss_out := g_iss.IntSpanOut()
	g_fss := createFloatSpan()
	g_fss_out := g_fss.FloatSpanOut(10)
	assert.Equal(t, g_iss_out, "[1, 4)")
	assert.Equal(t, g_fss_out, "[1.5, 3.6]")
}
