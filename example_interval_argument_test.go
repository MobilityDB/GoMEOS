package gomeos

// A handle the flat layer declares for a structure no model class wraps, an
// interval, is taken and answered as that handle, `*functions.Interval`, built
// by the flat layer's reader (`functions.IntervalIn`), as an `unsafe.Pointer`
// passes through. The window aggregates take one: each answer is the one
// MobilityDB's 042_temporal_waggfuncs test states for the same values, as
// #Example_temporalAggregateOverloads of example_aggregate_overloads_test.go
// states the aggregates'.

import (
	"fmt"

	"github.com/MobilityDB/GoMEOS/functions"
	"github.com/MobilityDB/GoMEOS/types"
)

func Example_intervalArgument() {
	i := func(s string) *types.TInt { v, _ := types.TIntFromText(s); return v }
	text := func(t *types.Temporal) string {
		if t == nil {
			return "<nil>"
		}
		s, _ := t.AsText(15)
		return s
	}
	window, err := functions.IntervalIn("5 minutes", -1)
	fmt.Println(err)

	wmin, err := types.TIntWMin([]*types.TInt{i("1@2001-01-01")}, window)
	expected := i("{[1@2001-01-01, 1@2001-01-01 00:05:00]}")
	fmt.Println(text(wmin) == text(&expected.Temporal), err)

	none, err := types.TIntWMin(nil, window)
	fmt.Println(none == nil, err)
	// Output:
	// <nil>
	// true <nil>
	// true <nil>
}
