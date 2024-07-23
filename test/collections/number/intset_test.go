package number_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos/collections/number"
	"github.com/alecthomas/assert/v2"
)

func createIntSet() *number.IntSet {
	return number.NewIntSet("{1, 2, 3}")
}

func createIntSet2() *number.IntSet {
	return number.NewIntSet("{3, 4, 5}")
}

func TestNewIntSet(t *testing.T) {
	g_is := createIntSet()
	assert.Equal(t, g_is.IntSetOut(), "{1, 2, 3}")
}

func TestIntSetToFloatSet(t *testing.T) {
	g_is := createIntSet()
	g_fs := g_is.ToFloatSet()
	assert.Equal(t, g_fs.FloatSetOut(5), "{1, 2, 3}")
}

func TestIntSetToSpanSet(t *testing.T) {
	g_is := createIntSet()
	g_iss := g_is.ToSpanSet()
	assert.Equal(t, g_iss.IntSpanSetOut(), "{[1, 4)}")
}

func TestIntSetNumElements(t *testing.T) {
	g_is := createIntSet()
	assert.Equal(t, g_is.NumElements(), 3)
}

func TestIntSetElements(t *testing.T) {
	g_is := createIntSet()
	assert.Equal(t, g_is.Elements(), []int{1, 2, 3})
}

func TestIntSetStartElement(t *testing.T) {
	g_is := createIntSet()
	assert.Equal(t, g_is.StartElement(), 1)
}

func TestIntSetEndElement(t *testing.T) {
	g_is := createIntSet()
	assert.Equal(t, g_is.EndElement(), 3)
}

func TestIntSetElementN(t *testing.T) {
	g_is := createIntSet()
	assert.Equal(t, g_is.ElementN(0), 1)
	assert.Equal(t, g_is.ElementN(1), 2)
	assert.Equal(t, g_is.ElementN(2), 3)
}

func TestIntSetShift(t *testing.T) {
	g_is := createIntSet()
	g_is_shift := g_is.Shift(1)
	assert.Equal(t, g_is_shift.IntSetOut(), "{2, 3, 4}")
}

func TestIntSetScale(t *testing.T) {
	g_is := createIntSet()
	g_is_scale := g_is.Scale(2)
	assert.Equal(t, g_is_scale.IntSetOut(), "{1, 2, 4}")
}

func TestIntSetShiftScale(t *testing.T) {
	g_is := createIntSet()
	g_is_ss := g_is.ShiftScale(1, 2)
	assert.Equal(t, g_is_ss.IntSetOut(), "{2, 3, 5}")
}

func TestIntSetContains(t *testing.T) {
	g_is := createIntSet()
	con, _ := g_is.Contains(2)
	assert.True(t, con)
	g_is2 := createIntSet2()
	conSet, _ := g_is2.Contains(g_is)
	assert.False(t, conSet)
}

func TestIntSetIsLeft(t *testing.T) {
	g_is := createIntSet()
	left, _ := g_is.IsLeft(4)
	assert.True(t, left)
}

func TestIntSetIsOverOrLeft(t *testing.T) {
	g_is := createIntSet()
	overleft, _ := g_is.IsOverOrLeft(3)
	assert.True(t, overleft)
}

func TestIntSetIsRight(t *testing.T) {
	g_is := createIntSet()
	right, _ := g_is.IsRight(0)
	assert.True(t, right)
}

func TestIntSetIsOverOrRight(t *testing.T) {
	g_is := createIntSet()
	overright, _ := g_is.IsOverOrRight(3)
	assert.False(t, overright)
}

func TestIntSetIntersection(t *testing.T) {
	g_is := createIntSet()
	g_is2 := createIntSet2()
	inter, _ := g_is.Intersection(g_is2)
	assert.Equal(t, inter.IntSetOut(), "{3}")
}

func TestIntSetMinus(t *testing.T) {
	g_is := createIntSet()
	g_is2 := createIntSet2()
	min, _ := g_is.Minus(g_is2)
	assert.Equal(t, min.IntSetOut(), "{1, 2}")
}

func TestIntSetDistance(t *testing.T) {
	g_is := createIntSet()
	dist, _ := g_is.Distance(5)
	assert.Equal(t, dist, 2)
}

func TestIntSetSub(t *testing.T) {
	g_is := createIntSet()
	sub, _ := g_is.Sub(2)
	assert.Equal(t, sub.IntSetOut(), "{1, 3}")
}

func TestIntSetMul(t *testing.T) {
	g_is := createIntSet()
	g_is2 := createIntSet2()
	mul, _ := g_is.Mul(g_is2)
	assert.Equal(t, mul.IntSetOut(), "{3}")
}

// func TestIntSetUnion(t *testing.T) {
// 	g_is := createIntSet()
// 	g_is2 := createIntSet2()
// 	union, _ := g_is.Union(g_is2)
// 	assert.Equal(t, union.IntSetOut(), "{1, 2, 3, 4, 5}")
// }

// func TestIntSetAdd(t *testing.T) {
// 	g_is := createIntSet()
// 	added, _ := g_is.Add(4)
// 	assert.Equal(t, added.IntSetOut(), "{1, 2, 3, 4}")
// }
