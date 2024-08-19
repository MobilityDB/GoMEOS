package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	gomeos "github.com/MobilityDB/GoMEOS"
)

// Constants
const (
	MaxLengthHeader  = 1024
	MaxLengthSpanSet = 1024
	NumberGroups     = 10
)

func main() {

	// Start timer
	startTime := time.Now()

	gomeos.MeosInitialize("UTC")
	// Define state values for aggregating the spans
	var state [NumberGroups]*gomeos.FloatSpanSet
	for i := 0; i < NumberGroups; i++ {
		state[i] = &gomeos.FloatSpanSet{} // Initialize each element with a new SpanSet
	}
	// Open the CSV file
	file, err := os.Open("data/floatspanset.csv")
	if err != nil {
		fmt.Println("Error opening input file")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	// Read the header
	_, err = reader.Read()
	if err != nil {
		fmt.Println("Error reading header")
		return
	}

	fmt.Println("Processing records (one marker per span set)")

	noRecords := 0
	noNulls := 0

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// fmt.Println(len(record[0]))
		// fmt.Println(len(record[1]))
		if len(record) < 2 {
			fmt.Println("Record with missing values ignored")
			noNulls++
			continue
		}

		if len(record[1]) < 1 {
			fmt.Println("Record with missing values ignored")
			noNulls++
			continue
		}
		k, _ := strconv.Atoi(record[0])
		spansetBuffer := record[1]
		// fmt.Println(spansetBuffer)
		// Transform the string representing the span set into a SpanSet value
		ss := gomeos.NewFloatSpanSet(spansetBuffer)
		// fmt.Println("read fss success!")
		// Union the span set into the appropriate group
		groupIndex := k % NumberGroups
		state[groupIndex] = gomeos.SpansetUnionTransfn(state[groupIndex], ss, &gomeos.FloatSpanSet{})
		// fmt.Println("update ss success!")
		// Output the span set value read
		spansetOut := ss.FloatSpanSetOut(3)
		fmt.Printf("k: %d, spanset: %s\n", k, spansetOut)

		noRecords++
	}

	fmt.Printf("\n%d records read.\n%d incomplete records ignored.\n", noRecords, noNulls)

	// Compute the final result
	for i := 0; i < NumberGroups; i++ {
		final := gomeos.SpansetUnionFinalfn(state[i], &gomeos.FloatSpanSet{})

		// Print the accumulated span set
		fmt.Printf("----------\n")
		fmt.Printf("Group: %d\n", i+1)
		fmt.Printf("----------\n")
		spansetOut := final.FloatSpanSetOut(3)
		fmt.Printf("spanset: %s\n", spansetOut)
	}

	// Calculate the elapsed time
	elapsedTime := time.Since(startTime)
	fmt.Printf("The program took %f seconds to execute\n", elapsedTime.Seconds())

	// Finalize MEOS
	gomeos.MeosFinalize() // Replace with actual finalization
}
