package gomeos

// An operator the object model assigns a class is a method of it, named by its
// ooName. Where MEOS answers the operator over several types of second operand,
// the method takes it as `any` and calls the MEOS function of its type, as
// #Example_sequenceAtGeometry of example_types_io_test.go passes a value of a
// parent class through its embedded handle.

import (
	"bytes"
	"fmt"

	"github.com/MobilityDB/GoMEOS/types"
)

func Example_spanContainsByOperand() {
	span, _ := types.IntSpanFromText("[1, 10)")
	inner, _ := types.IntSpanFromText("[2, 5)")
	outer, _ := types.IntSpanFromText("[0, 20)")
	fmt.Println(span.Contains(5))
	fmt.Println(span.Contains(20))
	fmt.Println(span.Contains(&inner.Span))
	fmt.Println(span.Contains(&outer.Span))
	fmt.Println(span.Contains("5"))
	// Output:
	// true <nil>
	// false <nil>
	// true <nil>
	// false <nil>
	// false Contains of IntSpan takes no string
}

// As #Example_spanContainsByOperand, for an operand an ancestor's overload takes:
// TFloat's add takes a float, and TNumber's add a temporal number, which the
// method of TFloat calls through the TNumber it embeds. Each result is compared
// with the value written as text.
func Example_tfloatAddReachesItsAncestor() {
	temp, _ := types.TFloatFromText("[1@2024-01-01, 2@2024-01-02]")
	other, _ := types.TFloatFromText("[10@2024-01-01, 20@2024-01-02]")
	plusWant, _ := types.TFloatFromText("[2.5@2024-01-01, 3.5@2024-01-02]")
	sumWant, _ := types.TFloatFromText("[11@2024-01-01, 22@2024-01-02]")
	plus, err := temp.Add(1.5)
	fmt.Println(types.TFloatFromPointer(plus.Pointer()).String() == plusWant.String(), err)
	sum, err := temp.Add(&other.Temporal)
	fmt.Println(types.TFloatFromPointer(sum.Pointer()).String() == sumWant.String(), err)
	fmt.Println(types.TFloatFromPointer(sum.Pointer()).String() == plusWant.String())
	_, err = temp.Add(1)
	fmt.Println(err)
	// Output:
	// true <nil>
	// true <nil>
	// false
	// Add of TFloat takes no int
}

// As #Example_spanContainsByOperand, for a class whose functions the flat layer
// passes as a bare pointer: a raster's map algebra is its methods, a float or
// another raster added to it and a float it is subtracted from. Each result's
// pixels are compared with those of a band holding the expected value, a band of
// 64-bit floats, the pixel type MEOS answers map algebra in.
func Example_rasterMapAlgebra() {
	empty, _ := types.RasterMakeEmpty(2, 2, 0, 2, 1, -1, 0, 0, 3857)
	band := func(r *types.Raster, v float64) *types.Raster {
		b, _ := r.AddBand(1, "64BF", v, 0, false)
		return b
	}
	pixels := func(r *types.Raster) []byte {
		p, _ := r.BandPixels(1)
		return p
	}
	r := band(empty, 7)
	plus, err := r.Add(1.5)
	fmt.Println(bytes.Equal(pixels(plus), pixels(band(empty, 8.5))), err)
	twice, err := r.Add(r)
	fmt.Println(bytes.Equal(pixels(twice), pixels(band(empty, 14))), err)
	from, err := r.Rsub(10)
	fmt.Println(bytes.Equal(pixels(from), pixels(band(empty, 3))), err)
	fmt.Println(bytes.Equal(pixels(from), pixels(band(empty, -3))))
	_, err = r.Add("1.5")
	fmt.Println(err)
	// Output:
	// true <nil>
	// true <nil>
	// true <nil>
	// false
	// Add of Raster takes no string
}
