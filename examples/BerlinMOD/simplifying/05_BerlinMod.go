package main

//MEOS example: meos/examples/08_berlinmod_simplify.c

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	gomeos "github.com/MobilityDB/GoMEOS"
)

const (
	MaxLengthTrip   = 170001
	MaxLengthHeader = 1024
	MaxLengthDate   = 12
	MaxNoTrips      = 64
)

type TripRecord struct {
	TripID int
	VehID  int
	Day    time.Time
	Seq    int
	Trip   *gomeos.TGeomPointSeq
}

func main() {
	// Arrays to compute the results
	var trips [MaxNoTrips]TripRecord
	// var currInst [MaxNoTrips]int
	var tripsDP [MaxNoTrips]*gomeos.TGeomPointSeq
	var tripsSED [MaxNoTrips]*gomeos.TGeomPointSeq

	// Get start time
	startTime := time.Now()

	// Initialize MEOS
	gomeos.MeosInitialize("UTC")

	// Open the input CSV file
	file, err := os.Open("data/berlinmod_trips.csv")
	if err != nil {
		log.Fatalf("Error opening input file: %v\n", err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.FieldsPerRecord = -1 // To handle varying number of fields

	// Read the first line (headers)
	_, err = reader.Read()
	if err != nil {
		log.Fatalf("Error reading header: %v\n", err)
	}

	i := 0
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("Error reading record: %v\n", err)
		}

		// Extract values
		tripID, _ := strconv.Atoi(record[0])
		vehID, _ := strconv.Atoi(record[1])
		dateBuffer := record[2]
		seq, _ := strconv.Atoi(record[3])
		tripBuffer := record[4]

		// Transform the date string into a time.Time value
		day, err := time.Parse("2006-01-02", dateBuffer) // Example format: "YYYY-MM-DD"
		if err != nil {
			log.Fatalf("Error parsing date: %v\n", err)
		}

		// Transform the trip string into a Temporal value
		trip := gomeos.NewTGeomPointSeqFromWKB(tripBuffer)

		// Save the trip record
		trips[i] = TripRecord{
			TripID: tripID,
			VehID:  vehID,
			Day:    day,
			Seq:    seq,
			Trip:   trip,
		}

		i++
	}

	recordsIn := i

	fmt.Println("finish reading csv")

	for i := 0; i < recordsIn; i++ {
		// Simplify the trip using the DP and SED algorithms
		tripsDP[i] = gomeos.TemporalSimplifyDP(trips[i].Trip, &gomeos.TGeomPointSeq{}, 2, false)
		tripsSED[i] = gomeos.TemporalSimplifyDP(trips[i].Trip, &gomeos.TGeomPointSeq{}, 2, true)

		// Convert the day to a string
		dayStr := trips[i].Day.Format("2006-01-02")

		// Print the details of the trip
		fmt.Printf("Vehicle: %d, Date: %s, Seq: %d, No. of instants: %d, "+
			"No. of instants DP: %d, No. of instants SED: %d\n",
			trips[i].VehID, dayStr, trips[i].Seq,
			gomeos.TemporalNumInstants(trips[i].Trip),
			gomeos.TemporalNumInstants(tripsDP[i]),
			gomeos.TemporalNumInstants(tripsSED[i]))

	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("The program took %f seconds to execute\n", elapsedTime.Seconds())

	// Free memory (handled automatically in Go)
	// Finalize MEOS
	gomeos.MeosFinalize()
}
