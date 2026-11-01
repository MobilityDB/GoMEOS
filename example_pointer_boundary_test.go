package gomeos

// The generated package and this one each declare their own cgo view of the
// MEOS structs, so a `*C.Temporal` from one is a DIFFERENT Go type in the other
// and Inner() cannot cross between them. The opaque pointer is what both sides
// agree on, which is what these examples exercise: a handle built here reaches
// a generated wrapper, the pointer survives the crossing unchanged, and an
// absent handle stays absent.

import (
	"fmt"
	"unsafe"

	"github.com/MobilityDB/GoMEOS/functions"
)

func ExampleTemporalFromPointer() {
	tb := NewTBoolSeq("{FALSE@2022-10-01, FALSE@2022-10-02, TRUE@2022-10-03}")

	// Hand this package's handle to the generated package through the pointer,
	// and let a generated wrapper read it.
	crossed := functions.TemporalFromPointer(unsafe.Pointer(tb.Inner()))
	instants, err := functions.TemporalNumInstants(crossed)
	subtype, err2 := functions.TemporalSubtype(crossed)
	fmt.Println(instants, err, subtype, err2)
	// Output:
	// 3 <nil> Sequence <nil>
}

func Example_temporalPointer() {
	tb := NewTBoolSeq("{FALSE@2022-10-01, TRUE@2022-10-03}")
	crossed := functions.TemporalFromPointer(unsafe.Pointer(tb.Inner()))

	// The pointer survives the round trip unchanged, so the boundary carries
	// the handle rather than copying what it points at.
	fmt.Println(crossed.Pointer() == unsafe.Pointer(tb.Inner()))
	// Output:
	// true
}

func ExampleTemporalFromPointer_absent() {
	// A MEOS entry answering NULL crosses as a nil handle rather than a
	// dereference, and a nil handle answers a nil pointer.
	absent := functions.TemporalFromPointer(nil)
	fmt.Println(absent == nil, absent.Pointer() == nil)
	// Output:
	// true true
}
