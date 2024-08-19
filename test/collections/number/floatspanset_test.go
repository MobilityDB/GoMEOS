package number_test

import (
	"testing"

	gomeos "github.com/MobilityDB/GoMEOS"
	"github.com/alecthomas/assert/v2"
)

func createFloatSpanSet() *gomeos.FloatSpanSet {
	g_fss_in := "{[1.1,3.3]}"
	g_fss := gomeos.NewFloatSpanSet(g_fss_in)
	return g_fss
}

func createFloatSpanSet2() *gomeos.FloatSpanSet {
	g_fss_in := "{[1.1,4.4), [9.9,100.0]}"
	g_fss := gomeos.NewFloatSpanSet(g_fss_in)
	return g_fss
}

func TestFSSInOutput(t *testing.T) {
	g_fss := createFloatSpanSet()
	g_fss_out := g_fss.FloatSpanSetOut(10)
	assert.Equal(t, g_fss_out, "{[1.1, 3.3]}")
}

func TestFSSToSpan(t *testing.T) {
	g_fss := createFloatSpanSet()
	g_fs := g_fss.ToSpan()
	g_fs_out := g_fs.FloatSpanOut(10)
	assert.Equal(t, g_fs_out, "[1.1, 3.3]")
}

func TestFSSConversion(t *testing.T) {
	g_fss := createFloatSpanSet2()
	ciss := g_fss.ToIntSpanSet()
	c_iss_out := ciss.IntSpanSetOut()
	assert.Equal(t, c_iss_out, "{[1, 4), [9, 101)}")
}

func TestFSSNumSpan(t *testing.T) {
	g_fss := createFloatSpanSet2()
	assert.Equal(t, g_fss.NumSpans(), 2)
}

func TestFSSWidth(t *testing.T) {
	g_fss := createFloatSpanSet2()
	assert.Equal(t, g_fss.Width(true), 98.9)
	assert.Equal(t, g_fss.Width(false), 93.39999999999999)
}

func TestFSSSpans(t *testing.T) {
	g_fss := createFloatSpanSet2()
	s := g_fss.StartSpan()
	assert.Equal(t, s.FloatSpanOut(10), "[1.1, 4.4)")

	e := g_fss.EndSpan()
	assert.Equal(t, e.FloatSpanOut(10), "[9.9, 100]")

	n := g_fss.SpanN(1)
	assert.Equal(t, n.FloatSpanOut(10), "[1.1, 4.4)")

	sp := g_fss.Spans()[0]
	assert.Equal(t, sp.FloatSpanOut(10), "[1.1, 4.4)")
}

func TestFSSShiftScale(t *testing.T) {
	g_fss := createFloatSpanSet2()
	g_shift := g_fss.Shift(2.0)
	assert.Equal(t, g_shift.FloatSpanSetOut(10), "{[3.1, 6.4), [11.9, 102]}")
	g_scale := g_fss.Scale(80.0)
	assert.Equal(t, g_scale.FloatSpanSetOut(10), "{[1.1, 3.7693629929), [8.2183013145, 81.1]}")
	g_ss := g_fss.ShiftScale(10.0, 70.0)
	assert.Equal(t, g_ss.FloatSpanSetOut(10), "{[11.1, 13.4356926188), [17.3285136502, 81.1]}")
}

func createFloatSpan() *gomeos.FloatSpan {
	g_fs_in := "[4.4,5.5]"
	g_fs := gomeos.NewFloatSpan(g_fs_in)
	return g_fs
}

func TestFSSAdjacent(t *testing.T) {
	g_fss := createFloatSpanSet()
	g_fs := createFloatSpan()
	res, err := g_fss.IsAdjacent(g_fs)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_fss.IsAdjacent(3.3)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	g_fss_2 := createFloatSpanSet2()
	res, err = g_fss.IsAdjacent(g_fss_2)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_fss.IsAdjacent("aaaa")
	assert.False(t, res)
	assert.NotEqual(t, err, nil)
}

func TestFSSContains(t *testing.T) {
	g_fss := createFloatSpanSet()
	g_fs := createFloatSpan()
	res, err := g_fss.Contains(g_fs)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_fss.Contains(2.2)
	assert.True(t, res)
	assert.Equal(t, err, nil)
	g_fss_2 := createFloatSpanSet2()
	res, err = g_fss.Contains(g_fss_2)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_fss.Contains("aaaa")
	assert.False(t, res)
	assert.NotEqual(t, err, nil)
}

func createFloatSpan2() *gomeos.FloatSpan {
	g_fs_in := "[1.1,3.3]"
	g_fs := gomeos.NewFloatSpan(g_fs_in)
	return g_fs
}
func TestFSSIsSame(t *testing.T) {
	g_fss := createFloatSpanSet()
	res, err := g_fss.IsSame(6.6)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	g_fs := createFloatSpan2()
	res, err = g_fss.IsSame(g_fs)
	assert.True(t, res)
	assert.Equal(t, err, nil)
	g_fss_2 := createFloatSpanSet2()
	res, err = g_fss.IsSame(g_fss_2)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_fss.IsSame("hahaha")
	assert.False(t, res)
	assert.NotEqual(t, err, nil)
}

func createFloatSpanSet3() *gomeos.FloatSpanSet {
	g_fss_in := "{[7.7,8.8]}"
	g_fss := gomeos.NewFloatSpanSet(g_fss_in)
	return g_fss
}

func TestFSSIsLeft(t *testing.T) {
	g_fss := createFloatSpanSet()
	res, err := g_fss.IsLeft(createFloatSpanSet3())
	assert.True(t, res)
	assert.Equal(t, err, nil)
	res, err = g_fss.IsLeft(2.2)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_fss.IsLeft("hahaha")
	assert.False(t, res)
	assert.NotEqual(t, err, nil)
}

func TestFSSIsOverOrLeft(t *testing.T) {
	g_fss := createFloatSpanSet()
	res, err := g_fss.IsOverOrLeft(createFloatSpanSet3())
	assert.True(t, res)
	assert.Equal(t, err, nil)
	res, err = g_fss.IsOverOrLeft(4.4)
	assert.True(t, res)
	assert.Equal(t, err, nil)
	res, err = g_fss.IsOverOrLeft("hahaha")
	assert.False(t, res)
	assert.NotEqual(t, err, nil)
}

func TestFSSIsRight(t *testing.T) {
	g_fss := createFloatSpanSet()
	res, err := g_fss.IsRight(createFloatSpanSet3())
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_fss.IsRight(2.2)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_fss.IsRight("hahaha")
	assert.False(t, res)
	assert.NotEqual(t, err, nil)
}

func TestFSSIsOverOrRight(t *testing.T) {
	g_fss := createFloatSpanSet()
	res, err := g_fss.IsOverOrRight(createFloatSpanSet3())
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_fss.IsOverOrRight(4.4)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_fss.IsOverOrRight("hahaha")
	assert.False(t, res)
	assert.NotEqual(t, err, nil)
}

func TestFSSDistance(t *testing.T) {
	g_fss := createFloatSpanSet()
	g_fs := createFloatSpan()
	res, err := g_fss.Distance(g_fs)
	assert.Equal(t, res, 1.1)
	assert.Equal(t, err, nil)
	res, err = g_fss.Distance(100.0)
	assert.Equal(t, res, 96.7)
	assert.Equal(t, err, nil)
	g_fss_2 := createFloatSpanSet2()
	res, err = g_fss.Distance(g_fss_2)
	assert.Equal(t, res, 0.0)
	assert.Equal(t, err, nil)
	res, err = g_fss.Distance("aaaa")
	assert.Equal(t, res, 0.0)
	assert.NotEqual(t, err, nil)
}

func TestFSSIntersection(t *testing.T) {
	g_fss := createFloatSpanSet()
	g_fs := createFloatSpan()
	res, err := g_fss.Intersection(g_fs)
	assert.Equal(t, res, nil)
	assert.Equal(t, err, nil)
	res, err = g_fss.Intersection(100.0)
	assert.Equal(t, res, nil)
	assert.Equal(t, err, nil)
	g_fss_2 := createFloatSpanSet2()
	res, err = g_fss.Intersection(g_fss_2)
	assert.Equal(t, res.FloatSpanSetOut(10), "{[1.1, 3.3]}")
	assert.Equal(t, err, nil)
	res, err = g_fss.Intersection("hahaha")
	assert.Equal(t, res, nil)
	assert.NotEqual(t, err, nil)
}

func TestFSSMinus(t *testing.T) {
	g_fss := createFloatSpanSet()
	g_fs := createFloatSpan()
	res, err := g_fss.Minus(g_fs)
	assert.Equal(t, res.FloatSpanSetOut(10), "{[1.1, 3.3]}")
	assert.Equal(t, err, nil)
	res, err = g_fss.Minus(2.2)
	assert.Equal(t, res.FloatSpanSetOut(10), "{[1.1, 2.2), (2.2, 3.3]}")
	assert.Equal(t, err, nil)
	g_fss_2 := createFloatSpanSet2()
	res, err = g_fss.Minus(g_fss_2)
	assert.Equal(t, res, nil)
	assert.Equal(t, err, nil)
	res, err = g_fss.Minus("hahaha")
	assert.Equal(t, res, nil)
	assert.NotEqual(t, err, nil)
}

func TestFSSUnion(t *testing.T) {
	g_fss := createFloatSpanSet()
	g_fs := createFloatSpan()
	res, err := g_fss.Union(g_fs)
	assert.Equal(t, res.FloatSpanSetOut(10), "{[1.1, 3.3], [4.4, 5.5]}")
	assert.Equal(t, err, nil)
	res, err = g_fss.Union(2.2)
	assert.Equal(t, res.FloatSpanSetOut(10), "{[1.1, 3.3]}")
	assert.Equal(t, err, nil)
	g_fss_2 := createFloatSpanSet2()
	res, err = g_fss.Union(g_fss_2)
	assert.Equal(t, res.FloatSpanSetOut(5), "{[1.1, 4.4), [9.9, 100]}")
	assert.Equal(t, err, nil)
	res, err = g_fss.Union("hahaha")
	assert.Equal(t, res, nil)
	assert.NotEqual(t, err, nil)
}
