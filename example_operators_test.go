package gomeos

// An operator the object model assigns a class is a method of it, named by its
// ooName. Where MEOS answers the operator over several types of second operand,
// the method takes it as `any` and calls the MEOS function of its type, as
// #Example_sequenceAtGeometry of example_types_io_test.go passes a value of a
// parent class through its embedded handle.

import (
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
