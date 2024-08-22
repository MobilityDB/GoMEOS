package main

//MEOS example: meos/examples/05_berlinmod_disassemble.c

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

type TripRecord struct {
	TripID int
	VehID  int
	Day    time.Time
	Seq    int
	Trip   *gomeos.TGeomPointSeq
}

func main() {
	// Arrays to compute the results
	var trips []TripRecord

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
		new_trip := TripRecord{
			TripID: tripID,
			VehID:  vehID,
			Day:    day,
			Seq:    seq,
			Trip:   trip,
		}
		trips = append(trips, new_trip)
		i++
	}

	recordsIn := i

	fmt.Println("finish reading csv")
	// Open the output CSV file
	outputFile, err := os.Create("data/berlinmod_instants.csv")
	if err != nil {
		log.Fatalf("Error creating output file: %v\n", err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Write the header line
	fmt.Println("start writing csv")
	writer.Write([]string{"tripid", "vehid", "day", "seqno", "geom", "t"})

	var nums int
	recordsOut := 0
	sum_nums := 0
	for i := 0; i < len(trips)-1; i++ {
		sum_nums += gomeos.TemporalNumInstants(trips[i].Trip)
	}
	fmt.Println(sum_nums)
	for i := 0; i < len(trips); i++ {
		nums = gomeos.TemporalNumInstants(trips[i].Trip)
		instants := make([]*gomeos.TGeomPointInst, nums)
		for i := 0; i < nums; i++ {
			instants[i] = gomeos.NewTGeomPointInstInner(nil)
		}
		insts := gomeos.TemporalInstants(trips[i].Trip, instants)
		for _, inst := range insts {
			dateStr := trips[i].Day.Format("2006-01-02")
			geomStr := gomeos.GeoAsEWKT(gomeos.TPointTrajectory(inst), 6)
			timeStr := inst.TimestampOut()
			writer.Write([]string{
				strconv.Itoa(trips[i].VehID),
				strconv.Itoa(trips[i].VehID),
				dateStr,
				strconv.Itoa(trips[i].Seq),
				geomStr,
				timeStr,
			})
			recordsOut++
		}
	}
	fmt.Printf("%d trip records read from file 'berlimod_trips.csv'.\n", recordsIn)
	fmt.Printf("%d observation records written in file 'berlimod_instants.csv'.\n", recordsOut)

	// Calculate the elapsed time
	elapsedTime := time.Since(startTime)
	fmt.Printf("The program took %f seconds to execute\n", elapsedTime.Seconds())

	// Free memory (handled automatically in Go)
	// Finalize MEOS
	gomeos.MeosFinalize()
}
