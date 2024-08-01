// collections/number/floatspan_test.go
package number_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

// func createFloatSpan() *number.FloatSpan {
// 	return number.NewFloatSpan("[1.1, 3.3]")
// }

// func createFloatSpan2() *number.FloatSpan {
// 	return number.NewFloatSpan("[2.2, 4.4]")
// }

func TestNewFloatSpan(t *testing.T) {
	g_fs := createFloatSpan()
	assert.Equal(t, g_fs.FloatSpanOut(10), "[1.1, 3.3]")
}

func TestFloatSpanToIntSpan(t *testing.T) {
	g_fs := createFloatSpan()
	g_is := g_fs.ToIntSpan()
	assert.Equal(t, g_is.IntSpanOut(), "[1, 4)")
}

func TestFloatSpanLower(t *testing.T) {
	g_fs := createFloatSpan()
	assert.Equal(t, g_fs.Lower(), 1.1)
}

func TestFloatSpanUpper(t *testing.T) {
	g_fs := createFloatSpan()
	assert.Equal(t, g_fs.Upper(), 3.3)
}

func TestFloatSpanWidth(t *testing.T) {
	g_fs := createFloatSpan()
	assert.Equal(t, g_fs.Width(), 2.2)
}

func TestFloatSpanShiftScale(t *testing.T) {
	g_fs := createFloatSpan()
	assert.Equal(t, g_fs.ShiftScale(1.1, 2.2).FloatSpanOut(5), "[2.2, 4.4]")
}

func TestFloatSpanShift(t *testing.T) {
	g_fs := createFloatSpan()
	assert.Equal(t, g_fs.Shift(1.1).FloatSpanOut(5), "[2.2, 4.4]")
}

func TestFloatSpanScale(t *testing.T) {
	g_fs := createFloatSpan()
	assert.Equal(t, g_fs.Scale(2.2).FloatSpanOut(5), "[1.1, 3.3]")
}

func TestFloatSpanIsAdjacent(t *testing.T) {
	g_fs := createFloatSpan()
	is_adj, err := g_fs.IsAdjacent(3.4)
	assert.False(t, is_adj, true)
	assert.Equal(t, err, nil)
}

func TestFloatSpanContains(t *testing.T) {
	g_fs := createFloatSpan()
	contains, err := g_fs.Contains(2.2)
	assert.Equal(t, contains, true)
	assert.Equal(t, err, nil)
}

func TestFloatSpanIsSame(t *testing.T) {
	g_fs := createFloatSpan()
	is_same, err := g_fs.IsSame(2.2)
	assert.Equal(t, is_same, false)
	assert.Equal(t, err, nil)
}

func TestFloatSpanIsLeft(t *testing.T) {
	g_fs := createFloatSpan()
	is_left, err := g_fs.IsLeft(4.4)
	assert.Equal(t, is_left, true)
	assert.Equal(t, err, nil)
}

func TestFloatSpanIsOverOrLeft(t *testing.T) {
	g_fs := createFloatSpan()
	is_over_or_left, err := g_fs.IsOverOrLeft(3.3)
	assert.Equal(t, is_over_or_left, true)
	assert.Equal(t, err, nil)
}

func TestFloatSpanIsRight(t *testing.T) {
	g_fs := createFloatSpan()
	is_right, err := g_fs.IsRight(0.0)
	assert.Equal(t, is_right, true)
	assert.Equal(t, err, nil)
}

func TestFloatSpanIsOverOrRight(t *testing.T) {
	g_fs := createFloatSpan()
	is_over_or_right, err := g_fs.IsOverOrRight(2.2)
	assert.Equal(t, is_over_or_right, false)
	assert.Equal(t, err, nil)
}

func TestFloatSpanDistance(t *testing.T) {
	g_fs := createFloatSpan()
	g_fs_2 := createFloatSpan2()
	distance, err := g_fs.Distance(g_fs_2)
	assert.Equal(t, distance, 0)
	assert.Equal(t, err, nil)
	distance, err = g_fs.Distance(20)
	assert.Equal(t, distance, 4.625393849793197e+18)
	assert.Equal(t, err, nil)
}

func TestFloatSpanIntersection(t *testing.T) {
	g_fs := createFloatSpan()
	g_fs_2 := createFloatSpan2()
	res, err := g_fs.Intersection(g_fs_2)
	assert.Equal(t, res.FloatSpanOut(5), "[2.2, 3.3]")
	assert.Equal(t, err, nil)
}

func TestFloatSpanMinus(t *testing.T) {
	g_fs := createFloatSpan()
	res, err := g_fs.Minus(2.2)
	assert.Equal(t, res.FloatSpanSetOut(5), "{[4.4, 5.5]}")
	assert.Equal(t, err, nil)
}

func TestFloatSpanUnion(t *testing.T) {
	g_fs := createFloatSpan()
	res, err := g_fs.Union(2.2)
	assert.Equal(t, res.FloatSpanSetOut(5), "{[2.2, 2.2], [4.4, 5.5]}")
	assert.Equal(t, err, nil)
}
