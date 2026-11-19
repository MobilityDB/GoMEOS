package gomeos

// An aggregate of the catalog's `aggregates` is a function of the class of its
// values, named after the class as every function of a class is: the
// transition threaded over the values from a NULL state, then the final
// function. Each answer is the one MobilityDB's 500_raster test states for the
// same tiles, as #Example_rasterMapAlgebra of example_operators_test.go builds
// its rasters.

import (
	"fmt"

	"github.com/MobilityDB/GoMEOS/types"
)

// The example of the manual: two 2x2 tiles of one 32BF band, the second one
// pixel east and one south of the first, holding 1 and 3.
func aggregateTiles() []*types.Raster {
	tile := func(x, y, v float64) *types.Raster {
		empty, _ := types.RasterMakeEmpty(2, 2, x, y, 1, -1, 0, 0, 3857)
		r, _ := empty.AddBand(1, "32BF", v, -9999, true)
		return r
	}
	return []*types.Raster{tile(0, 2, 1), tile(1, 1, 3)}
}

func Example_rasterAggregates() {
	stats, err := types.RasterSummaryStatsAgg(aggregateTiles(), 1, true)
	count, _ := stats.Count()
	sum, _, _ := stats.Sum()
	mean, _, _ := stats.Mean()
	stddev, _, _ := stats.Stddev()
	minimum, _, _ := stats.Min()
	maximum, _, _ := stats.Max()
	fmt.Println(count, sum, mean, stddev, minimum, maximum, err)

	// The union takes the mean where the tiles overlap, {{1,1,NULL},{1,2,3},
	// {NULL,3,3}}: seven pixels of value, summing to 14.
	mode := "MEAN"
	union, err := types.RasterRasterUnion(aggregateTiles(), 1, &mode)
	ustats, _ := types.RasterSummaryStatsAgg([]*types.Raster{union}, 1, true)
	ucount, _ := ustats.Count()
	usum, _, _ := ustats.Sum()
	umin, _, _ := ustats.Min()
	umax, _, _ := ustats.Max()
	fmt.Println(ucount, usum, umin, umax, err)

	none, err := types.RasterSummaryStatsAgg(nil, 1, true)
	fmt.Println(none == nil, err)
	// Output:
	// 8 16 2 1 1 3 <nil>
	// 7 14 1 3 <nil>
	// true <nil>
}
