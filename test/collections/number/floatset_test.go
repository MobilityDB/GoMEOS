// collections/number/floatset_test.go
package number_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos"
	"github.com/alecthomas/assert/v2"
)

func createFloatSet() gomeos.FloatSet {
	return *gomeos.NewFloatSet("{1.23, 2.9, 3.6}")
}

func createFloatSet2() gomeos.FloatSet {
	return *gomeos.NewFloatSet("{2.9, 4.70, 5.97, 1.23, 3.6}")
}

func TestNewFloatSet(t *testing.T) {
	g_fs := createFloatSet()
	assert.Equal(t, g_fs.FloatSetOut(5), "{1.23, 2.9, 3.6}")
}

func TestFloatSetToIntSet(t *testing.T) {
	g_fs := createFloatSet()
	g_is := g_fs.ToIntSet()
	assert.Equal(t, g_is.IntSetOut(), "{1, 2, 3}")
}

func TestFloatSetToSpanSet(t *testing.T) {
	g_fs := createFloatSet()
	g_fss := g_fs.ToSpanSet()
	assert.Equal(t, g_fss.FloatSpanSetOut(5), "{[1.23, 1.23], [2.9, 2.9], [3.6, 3.6]}")
}

func TestFloatSetNumElements(t *testing.T) {
	g_fs := createFloatSet()
	assert.Equal(t, g_fs.NumElements(), 3)
}

func TestFloatSetElements(t *testing.T) {
	g_fs := createFloatSet()
	assert.Equal(t, g_fs.Elements(), []float64{1.23, 2.9, 3.6})
}

func TestFloatSetStartElement(t *testing.T) {
	g_fs := createFloatSet()
	assert.Equal(t, g_fs.StartElement(), 1.23)
}

func TestFloatSetEndElement(t *testing.T) {
	g_fs := createFloatSet()
	assert.Equal(t, g_fs.EndElement(), 3.6)
}

func TestFloatSetElementN(t *testing.T) {
	g_fs := createFloatSet()
	assert.Equal(t, g_fs.ElementN(0), 1.23)
	assert.Equal(t, g_fs.ElementN(1), 2.9)
	assert.Equal(t, g_fs.ElementN(2), 3.6)
}

func TestFloatSetShift(t *testing.T) {
	g_fs := createFloatSet()
	shift := g_fs.Shift(1.0)
	assert.Equal(t, shift.FloatSetOut(5), "{2.23, 3.9, 4.6}")
}

func TestFloatSetScale(t *testing.T) {
	g_fs := createFloatSet()
	scale := g_fs.Scale(2.0)
	assert.Equal(t, scale.FloatSetOut(5), "{1.23, 2.63928, 3.23}")
}

func TestFloatSetShiftScale(t *testing.T) {
	g_fs := createFloatSet()
	ss := g_fs.ShiftScale(1.0, 2.0)
	assert.Equal(t, ss.FloatSetOut(5), "{2.23, 3.63928, 4.23}")
}

func TestFloatSetContains(t *testing.T) {
	g_fs := createFloatSet()
	con, _ := g_fs.Contains(2.0)
	assert.False(t, con)
	g_fs2 := createFloatSet2()
	conSet, _ := g_fs2.Contains(g_fs)
	assert.False(t, conSet)
}

func TestFloatSetIsLeft(t *testing.T) {
	g_fs := createFloatSet()
	left, _ := g_fs.IsLeft(4.0)
	assert.True(t, left)
}

func TestFloatSetIsOverOrLeft(t *testing.T) {
	g_fs := createFloatSet()
	overleft, _ := g_fs.IsOverOrLeft(3.0)
	assert.False(t, overleft)
}

func TestFloatSetIsRight(t *testing.T) {
	g_fs := createFloatSet()
	right, _ := g_fs.IsRight(0.0)
	assert.True(t, right)
}

func TestFloatSetIsOverOrRight(t *testing.T) {
	g_fs := createFloatSet()
	overright, _ := g_fs.IsOverOrRight(3.0)
	assert.False(t, overright)
}

func TestFloatSetIntersection(t *testing.T) {
	g_is := createFloatSet()
	g_is2 := createFloatSet2()
	inter, _ := g_is.Intersection(&g_is2)
	assert.Equal(t, inter.FloatSetOut(5), "{1.23, 2.9, 3.6}")
}

func TestFloatSetMul(t *testing.T) {
	g_fs := createFloatSet()
	g_fs2 := createFloatSet2()
	mul, _ := g_fs.Mul(&g_fs2)
	assert.Equal(t, mul.FloatSetOut(5), "{1.23, 2.9, 3.6}")
}

func TestFloatSetMinus(t *testing.T) {
	g_is := createFloatSet()
	g_is2 := createFloatSet2()
	inter, _ := g_is2.Intersection(&g_is)
	assert.Equal(t, inter.FloatSetOut(5), "{1.23, 2.9, 3.6}")
}

func TestFloatSetSub(t *testing.T) {
	g_fs := createFloatSet()
	sub, _ := g_fs.Sub(2.0)
	assert.Equal(t, sub.FloatSetOut(5), "{1.23, 2.9, 3.6}")
}

func TestFloatSetDistance(t *testing.T) {
	g_fs := createFloatSet()
	dist, _ := g_fs.Distance(5.0)
	assert.Equal(t, dist, 1.4)
}

func TestFloatSetUnion(t *testing.T) {
	g_fs := createFloatSet()
	g_fs2 := createFloatSet2()
	union, _ := g_fs.Union(&g_fs2)
	assert.Equal(t, union.FloatSetOut(10), "{1.23, 2.9, 3.6, 4.7, 5.97}")
}

func TestFloatSetAdd(t *testing.T) {
	g_fs := createFloatSet()
	added, _ := g_fs.Add(4.0)
	assert.Equal(t, added.FloatSetOut(10), "{1.23, 2.9, 3.6, 4}")
}
