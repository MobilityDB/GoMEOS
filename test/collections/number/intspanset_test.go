package number_test

import (
	"testing"

	gomeos "github.com/MobilityDB/GoMEOS"
	"github.com/alecthomas/assert/v2"
)

func createIntSpanSet() *gomeos.IntSpanSet {
	g_iss_in := "{[1,3]}"
	g_iss := gomeos.NewIntSpanSet(g_iss_in)
	return g_iss
}

func createIntSpanSet2() *gomeos.IntSpanSet {
	g_iss_in := "{[1,4), [9,100]}"
	g_iss := gomeos.NewIntSpanSet(g_iss_in)
	return g_iss
}

func TestISSInOutput(t *testing.T) {
	g_iss := createIntSpanSet()
	g_iss_out := g_iss.IntSpanSetOut()
	assert.Equal(t, g_iss_out, "{[1, 4)}")
}

func TestISSToSpan(t *testing.T) {
	g_iss := createIntSpanSet()
	g_is := g_iss.ToSpan()
	g_is_out := g_is.IntSpanOut()
	assert.Equal(t, g_is_out, "[1, 4)")
}

func TestISSConversion(t *testing.T) {
	g_iss := createIntSpanSet2()
	cfss := g_iss.ToFloatSpanSet()
	c_fss_out := cfss.FloatSpanSetOut(10)
	assert.Equal(t, c_fss_out, "{[1, 3], [9, 100]}")
}

func TestISSNumSpan(t *testing.T) {
	g_iss := createIntSpanSet2()
	assert.Equal(t, g_iss.NumSpans(), 2)
}

func TestISSWidth(t *testing.T) {
	g_iss := createIntSpanSet2()
	assert.Equal(t, g_iss.Width(true), 100)
	assert.Equal(t, g_iss.Width(false), 95)
}

func TestISSSpans(t *testing.T) {
	g_iss := createIntSpanSet2()
	s := g_iss.StartSpan()
	assert.Equal(t, s.IntSpanOut(), "[1, 4)")

	e := g_iss.EndSpan()
	assert.Equal(t, e.IntSpanOut(), "[9, 101)")

	n := g_iss.SpanN(1)
	assert.Equal(t, n.IntSpanOut(), "[9, 101)")

	sp := g_iss.Spans()[0]
	assert.Equal(t, sp.IntSpanOut(), "[1, 4)")
}

func TestISSShiftScale(t *testing.T) {
	g_iss := createIntSpanSet2()
	g_shift := g_iss.Shift(2)
	assert.Equal(t, g_shift.IntSpanSetOut(), "{[3, 6), [11, 103)}")
	g_scale := g_iss.Scale(80)
	assert.Equal(t, g_scale.IntSpanSetOut(), "{[1, 3), [7, 82)}")
	g_ss := g_iss.ShiftScale(10, 70)
	assert.Equal(t, g_ss.IntSpanSetOut(), "{[11, 13), [16, 82)}")
}

func createIntSpan() *gomeos.IntSpan {
	g_is_in := "[4,5]"
	g_is := gomeos.NewIntSpan(g_is_in)
	return g_is
}

func TestISSAdjacent(t *testing.T) {
	g_iss := createIntSpanSet()
	g_is := createIntSpan()
	res, err := g_iss.IsAdjacent(g_is)
	assert.True(t, res)
	assert.Equal(t, err, nil)
	res, err = g_iss.IsAdjacent(4)
	assert.True(t, res)
	assert.Equal(t, err, nil)
	g_iss_2 := createIntSpanSet2()
	res, err = g_iss.IsAdjacent(g_iss_2)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_iss.IsAdjacent("aaaa")
	assert.False(t, res)
	assert.NotEqual(t, err, nil)
}

func TestISSContains(t *testing.T) {
	g_iss := createIntSpanSet()
	g_is := createIntSpan()
	res, err := g_iss.Contains(g_is)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_iss.Contains(3)
	assert.True(t, res)
	assert.Equal(t, err, nil)
	g_iss_2 := createIntSpanSet2()
	res, err = g_iss.Contains(g_iss_2)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_iss.Contains("aaaa")
	assert.False(t, res)
	assert.NotEqual(t, err, nil)
}

func createIntSpan2() *gomeos.IntSpan {
	g_is_in := "[1,3]"
	g_is := gomeos.NewIntSpan(g_is_in)
	return g_is
}
func TestISSIsSame(t *testing.T) {
	g_iss := createIntSpanSet()
	res, err := g_iss.IsSame(6)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	g_is := createIntSpan2()
	res, err = g_iss.IsSame(g_is)
	assert.True(t, res)
	assert.Equal(t, err, nil)
	g_iss_2 := createIntSpanSet2()
	res, err = g_iss.IsSame(g_iss_2)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_iss.IsSame("hahaha")
	assert.False(t, res)
	assert.NotEqual(t, err, nil)
}

func createIntSpanSet3() *gomeos.IntSpanSet {
	g_iss_in := "{[7,8]}"
	g_iss := gomeos.NewIntSpanSet(g_iss_in)
	return g_iss
}

func TestISSIsLeft(t *testing.T) {
	g_iss := createIntSpanSet()
	res, err := g_iss.IsLeft(createIntSpanSet3())
	assert.True(t, res)
	assert.Equal(t, err, nil)
	res, err = g_iss.IsLeft(2)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_iss.IsLeft("hahaha")
	assert.False(t, res)
	assert.NotEqual(t, err, nil)
}

func TestISSIsOverOrLeft(t *testing.T) {
	g_iss := createIntSpanSet()
	res, err := g_iss.IsOverOrLeft(createIntSpanSet3())
	assert.True(t, res)
	assert.Equal(t, err, nil)
	res, err = g_iss.IsOverOrLeft(4)
	assert.True(t, res)
	assert.Equal(t, err, nil)
	res, err = g_iss.IsOverOrLeft("hahaha")
	assert.False(t, res)
	assert.NotEqual(t, err, nil)
}

func TestISSIsRight(t *testing.T) {
	g_iss := createIntSpanSet()
	res, err := g_iss.IsRight(createIntSpanSet3())
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_iss.IsRight(2)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_iss.IsRight("hahaha")
	assert.False(t, res)
	assert.NotEqual(t, err, nil)
}

func TestISSIsOverOrRight(t *testing.T) {
	g_iss := createIntSpanSet()
	res, err := g_iss.IsOverOrRight(createIntSpanSet3())
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_iss.IsOverOrRight(4)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_iss.IsOverOrRight("hahaha")
	assert.False(t, res)
	assert.NotEqual(t, err, nil)
}

func TestISSDistance(t *testing.T) {
	g_iss := createIntSpanSet()
	g_is := createIntSpan()
	res, err := g_iss.Distance(g_is)
	assert.Equal(t, res, 1)
	assert.Equal(t, err, nil)
	res, err = g_iss.Distance(100)
	assert.Equal(t, res, 97)
	assert.Equal(t, err, nil)
	g_iss_2 := createIntSpanSet2()
	res, err = g_iss.Distance(g_iss_2)
	assert.Equal(t, res, 0)
	assert.Equal(t, err, nil)
	res, err = g_iss.Distance("aaaa")
	assert.Equal(t, res, 0)
	assert.NotEqual(t, err, nil)
}

func TestISSIntersection(t *testing.T) {
	g_iss := createIntSpanSet()
	g_is := createIntSpan()
	res, err := g_iss.Intersection(g_is)
	assert.Equal(t, res, nil)
	assert.Equal(t, err, nil)
	res, err = g_iss.Intersection(100)
	assert.Equal(t, res, nil)
	assert.Equal(t, err, nil)
	g_iss_2 := createIntSpanSet2()
	res, err = g_iss.Intersection(g_iss_2)
	assert.Equal(t, res.IntSpanSetOut(), "{[1, 4)}")
	assert.Equal(t, err, nil)
	res, err = g_iss.Intersection("aaaa")
	assert.Equal(t, res, nil)
	assert.NotEqual(t, err, nil)
}

func TestISSMul(t *testing.T) {
	g_iss := createIntSpanSet()
	g_is := createIntSpan()
	res, err := g_iss.Mul(g_is)
	assert.Equal(t, res, nil)
	assert.Equal(t, err, nil)
	res, err = g_iss.Mul(100)
	assert.Equal(t, res, nil)
	assert.Equal(t, err, nil)
	g_iss_2 := createIntSpanSet2()
	res, err = g_iss.Mul(g_iss_2)
	assert.Equal(t, res.IntSpanSetOut(), "{[1, 4)}")
	assert.Equal(t, err, nil)
	// res, err = g_iss.Mul("aaaa")
	// assert.Equal(t, res, nil)
	// assert.NotEqual(t, err, nil)
}

func TestISSMinus(t *testing.T) {
	g_iss := createIntSpanSet()
	g_is := createIntSpan()
	res, err := g_iss.Minus(g_is)
	assert.Equal(t, res.IntSpanSetOut(), "{[1, 4)}")
	assert.Equal(t, err, nil)
	res, err = g_iss.Minus(2)
	assert.Equal(t, res.IntSpanSetOut(), "{[1, 2), [3, 4)}")
	assert.Equal(t, err, nil)
	g_iss_2 := createIntSpanSet2()
	res, err = g_iss.Minus(g_iss_2)
	assert.Equal(t, res, nil)
	assert.Equal(t, err, nil)
	res, err = g_iss.Minus("aaaa")
	assert.Equal(t, res, nil)
	assert.NotEqual(t, err, nil)
}

func TestISSUnion(t *testing.T) {
	g_iss := createIntSpanSet()
	g_is := createIntSpan()
	res, err := g_iss.Union(g_is)
	assert.Equal(t, res.IntSpanSetOut(), "{[1, 6)}")
	assert.Equal(t, err, nil)
	res, err = g_iss.Union(2)
	assert.Equal(t, res.IntSpanSetOut(), "{[1, 4)}")
	assert.Equal(t, err, nil)
	g_iss_2 := createIntSpanSet2()
	res, err = g_iss.Union(g_iss_2)
	assert.Equal(t, res.IntSpanSetOut(), "{[1, 4), [9, 101)}")
	assert.Equal(t, err, nil)
	res, err = g_iss.Union("aaaa")
	assert.Equal(t, res, nil)
	assert.NotEqual(t, err, nil)
}
