package gomeos

// An aggregate several transitions serve under one name, one per type of its
// values (tsum is tint_tsum_transfn over tint and tfloat_tsum_transfn over
// tfloat), is a function of each value's class: the catalog names that class by
// the SQL type of the aggregate's values, the prefix of its leaf of the lattice.
// Each answer is the one MobilityDB's 040_temporal_aggfuncs test states for the
// same values, as #Example_rasterAggregates of example_aggregates_test.go states
// the raster aggregates'.

import (
	"fmt"

	"github.com/MobilityDB/GoMEOS/types"
)

func Example_temporalAggregateOverloads() {
	f := func(s string) *types.TFloat { v, _ := types.TFloatFromText(s); return v }
	i := func(s string) *types.TInt { v, _ := types.TIntFromText(s); return v }
	text := func(t *types.Temporal) string {
		if t == nil {
			return "<nil>"
		}
		s, _ := t.AsText(15)
		return s
	}

	sum, err := types.TFloatTSum([]*types.TFloat{
		f("[1@2001-01-01, 2@2001-01-03, 1@2001-01-05, 2@2001-01-07]"),
		f("[3@2001-01-02, 4@2001-01-06]"),
	})
	expected := f("{[1@2001-01-01, 1.5@2001-01-02), " +
		"[4.5@2001-01-02, 5.25@2001-01-03, 4.75@2001-01-05, 5.5@2001-01-06], " +
		"(1.5@2001-01-06, 2@2001-01-07]}")
	fmt.Println(text(sum) == text(&expected.Temporal), err)

	min, err := types.TIntTMinAgg([]*types.TInt{i("1@2001-01-01")})
	fmt.Println(text(min) == text(&i("{1@2001-01-01}").Temporal), err)

	none, err := types.TIntTMinAgg(nil)
	fmt.Println(none == nil, err)
	// Output:
	// true <nil>
	// true <nil>
	// true <nil>
}
