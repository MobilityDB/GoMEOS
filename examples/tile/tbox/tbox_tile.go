package main

import (
	"fmt"

	"github.com/MobilityDB/GoMEOS/gomeos"
)

func main() {
	// Buffer for output data
	const maxLineLength = 1024
	outputBuffer := make([]byte, maxLineLength)

	// Parameters to choose between temporal integer or float box, and enable/disable splits
	intSpan := false
	valueSplit := true
	timeSplit := false

	// Initialize MEOS
	gomeos.MeosInitialize("UTC")

	// Initialize values for tiling
	box := func() *TBox {
		if intSpan {
			return tboxIn("TBOXINT XT([1,10],[2020-03-01, 2020-03-10])")
		}
		return tboxIn("TBOXFLOAT XT([1,10],[2020-03-01, 2020-03-10])")
	}()
	interval := pgIntervalIn("5 days")
	tOrigin := pgTimestamptzIn("2020-03-01")

	var boxes []TBox
	var spans []Span
	var count int

	// Perform tiling
	if valueSplit {
		if intSpan {
			boxes, count = tintboxTileList(box, 5, timeSplit, 1, tOrigin)
		} else {
			boxes, count = tfloatboxTileList(box, 5, timeSplit, 1, tOrigin)
		}
	} else {
		spans, count = tstzspanBucketList(nil, interval, tOrigin)
	}

	// Print the input value to split
	boxStr := tboxOut(box, 3)
	fmt.Println("------------------")
	fmt.Println("| Value to split |")
	fmt.Println("------------------\n")
	fmt.Println(boxStr)

	// Output the resulting tiles
	fmt.Println("--------")
	fmt.Println("| Tiles |")
	fmt.Println("--------\n")

	for i := 0; i < count; i++ {
		var tileStr string
		if valueSplit {
			tileStr = tboxOut(&boxes[i], 3)
		} else {
			tileStr = floatspanOut(&spans[i], 3)
		}
		line := fmt.Sprintf("%d: %s\n", i+1, tileStr)
		copy(outputBuffer, line)
		fmt.Print(string(outputBuffer))
	}

	// Print information about the result
	fmt.Printf("\nNumber of tiles: %d\n", count)

	// Clean up allocated objects (handled by garbage collector in Go)
	meosFinalize()
}
