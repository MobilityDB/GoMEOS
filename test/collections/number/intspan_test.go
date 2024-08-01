package number_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

// func createIntSpan() *number.IntSpan {
// 	g_is_in := "[4,5]"
// 	g_is := number.NewIntSpan(g_is_in)
// 	return g_is
// }

// func createIntSpan2() *number.IntSpan {
// 	g_is_in := "[1,3]"
// 	g_is := number.NewIntSpan(g_is_in)
// 	return g_is
// }

func TestIntSpanInOutput(t *testing.T) {
	g_is := createIntSpan()
	g_is_out := g_is.IntSpanOut()
	assert.Equal(t, g_is_out, "[4, 6)")
}

func TestIntSpanToSpanSet(t *testing.T) {
	g_is := createIntSpan()
	g_iss := g_is.ToSpanSet()
	g_iss_out := g_iss.IntSpanSetOut()
	assert.Equal(t, g_iss_out, "{[4, 6)}")
}

func TestIntSpanToFloatSpan(t *testing.T) {
	g_is := createIntSpan()
	g_fs := g_is.ToFloatSpan()
	g_fs_out := g_fs.FloatSpanOut(10)
	assert.Equal(t, g_fs_out, "[4, 5]")
}

func TestIntSpanAccessors(t *testing.T) {
	g_is := createIntSpan()
	assert.Equal(t, g_is.Lower(), 4)
	assert.Equal(t, g_is.Upper(), 6)
	assert.Equal(t, g_is.Width(), 2)
}

func TestIntSpanTransformations(t *testing.T) {
	g_is := createIntSpan()
	g_shift := g_is.Shift(2)
	assert.Equal(t, g_shift.IntSpanOut(), "[6, 8)")
	g_scale := g_is.Scale(3)
	assert.Equal(t, g_scale.IntSpanOut(), "[4, 8)")
	g_ss := g_is.ShiftScale(1, 2)
	assert.Equal(t, g_ss.IntSpanOut(), "[5, 8)")
}

func TestIntSpanIsAdjacent(t *testing.T) {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.IsAdjacent(g_is2)
	assert.True(t, res)
	assert.Equal(t, err, nil)
	res, err = g_is.IsAdjacent(5)
	assert.False(t, res)
	assert.Equal(t, err, nil)
}

func TestIntSpanContains(t *testing.T) {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.Contains(g_is2)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_is.Contains(200)
	assert.False(t, res)
	assert.Equal(t, err, nil)
}

func TestIntSpanIsSame(t *testing.T) {
	g_is := createIntSpan()
	res, err := g_is.IsSame(200)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	g_is2 := createIntSpan()
	res, err = g_is.IsSame(g_is2)
	assert.True(t, res)
	assert.Equal(t, err, nil)
}

func TestIntSpanIsLeft(t *testing.T) {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.IsLeft(g_is2)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_is.IsLeft(6)
	assert.True(t, res)
	assert.Equal(t, err, nil)
}

func TestIntSpanIsOverOrLeft(t *testing.T) {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.IsOverOrLeft(g_is2)
	assert.False(t, res)
	assert.Equal(t, err, nil)
	res, err = g_is.IsOverOrLeft(4)
	assert.False(t, res)
	assert.Equal(t, err, nil)
}

func TestIntSpanIsRight(t *testing.T) {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.IsRight(g_is2)
	assert.True(t, res)
	assert.Equal(t, err, nil)
	res, err = g_is.IsRight(3)
	assert.True(t, res)
	assert.Equal(t, err, nil)
}

func TestIntSpanIsOverOrRight(t *testing.T) {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.IsOverOrRight(g_is2)
	assert.True(t, res)
	assert.Equal(t, err, nil)
	res, err = g_is.IsOverOrRight(5)
	assert.False(t, res)
	assert.Equal(t, err, nil)
}

func TestIntSpanDistance(t *testing.T) {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.Distance(g_is2)
	assert.Equal(t, res, 1)
	assert.Equal(t, err, nil)
	res, err = g_is.Distance(10)
	assert.Equal(t, res, 5)
	assert.Equal(t, err, nil)
}

func TestIntSpanIntersection(t *testing.T) {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.Intersection(g_is2)
	assert.Equal(t, res, nil)
	assert.Equal(t, err, nil)
	res, err = g_is.Intersection(4)
	assert.Equal(t, res.IntSpanOut(), "[4, 5)")
	assert.Equal(t, err, nil)
}

func TestIntSpanMul(t *testing.T) {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.Mul(g_is2)
	assert.Equal(t, res, nil)
	assert.Equal(t, err, nil)
	res, err = g_is.Mul(4)
	assert.Equal(t, res.IntSpanOut(), "[4, 5)")
	assert.Equal(t, err, nil)
}

func TestIntSpanMinus(t *testing.T) {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.Minus(g_is2)
	assert.Equal(t, res.IntSpanSetOut(), "{[4, 6)}")
	assert.Equal(t, err, nil)
	res, err = g_is.Minus(4)
	assert.Equal(t, res.IntSpanSetOut(), "{[5, 6)}")
	assert.Equal(t, err, nil)
}

func TestIntSpanUnion(t *testing.T) {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.Union(g_is2)
	assert.Equal(t, res.IntSpanSetOut(), "{[1, 6)}")
	assert.Equal(t, err, nil)
	res, err = g_is.Union(6)
	assert.Equal(t, res.IntSpanSetOut(), "{[4, 7)}")
	assert.Equal(t, err, nil)
}
