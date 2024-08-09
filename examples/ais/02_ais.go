package main

//MEOS example: meos/examples/02_ais_read.c
//MEOS example: meos/examples/03_ais_assemble.c
import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/MobilityDB/GoMEOS/gomeos"
)

// Define the AISRecord struct
type AISRecord struct {
	Timestamp time.Time
	MMSI      int64
	Latitude  float64
	Longitude float64
	SOG       float64
}

// Define the TripRecord struct
type TripRecord struct {
	MMSI int64
	// numinstants   int
	trip_instants []gomeos.TGeogPointInst
	SOG_instants  []gomeos.TFloatInst
	trip          gomeos.TGeogPointSeq
	SOG           gomeos.TFloatSeq
}

const (
	MAX_TRIPS         = 5
	NO_INSTANTS_BATCH = 1000
)

func main() {
	gomeos.MeosInitialize("UTC")
	// Open the CSV file
	file, err := os.Open("data/ais_instants.csv")
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
		gomeos.MeosFinalize()
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the header
	header, err := reader.Read()
	if err != nil {
		log.Fatalf("Error reading header: %v", err)
	}
	fmt.Println("Header:", header)

	// Initialize variables
	var trips [MAX_TRIPS]TripRecord
	numships := 0

	fmt.Printf("Reading the instants (one '*' marker every %d instants)\n", NO_INSTANTS_BATCH)

	// Process CSV records
	recordCount := 0
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// Parse the record
		timestamp, err := time.Parse("2006-01-02 15:04:05", record[0])
		if err != nil {
			log.Printf("Error parsing timestamp: %v", err)
			continue
		}

		mmsi, err := strconv.ParseInt(record[1], 10, 64)
		if err != nil {
			log.Printf("Error parsing MMSI: %v", err)
			continue
		}

		latitude, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Printf("Error parsing Latitude: %v", err)
			continue
		}

		longitude, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Printf("Error parsing Longitude: %v", err)
			continue
		}

		sog, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			log.Printf("Error parsing SOG: %v", err)
			continue
		}

		// Create an AISRecord instance
		rec := AISRecord{
			Timestamp: timestamp,
			MMSI:      mmsi,
			Latitude:  latitude,
			Longitude: longitude,
			SOG:       sog,
		}

		// Find the trip for the MMSI
		tripIndex := -1
		for i := 0; i < numships; i++ {
			if trips[i].MMSI == mmsi {
				tripIndex = i
				break
			}
		}

		// If no trip found, create a new one
		if tripIndex == -1 {
			if numships == MAX_TRIPS {
				log.Printf("The maximum number of ships in the input file is bigger than %d", MAX_TRIPS)
				break
			}
			tripIndex = numships
			trips[tripIndex].MMSI = mmsi
			numships++
		}

		// Add the record to the trip
		strPoint := fmt.Sprintf("SRID=4326;Point(%f %f)@%s+00", rec.Longitude, rec.Latitude, rec.Timestamp.Format("2006-01-02 15:04:05"))
		strFloat := fmt.Sprintf("%f@%s", rec.SOG, rec.Timestamp.Format("2006-01-02 15:04:05"))
		// fmt.Println(strPoint)
		// fmt.Println(strFloat)
		instPoint := gomeos.NewTGeogPointInst(strPoint)
		instFloat := gomeos.NewTFloatInst(strFloat)
		trips[tripIndex].trip_instants = append(trips[tripIndex].trip_instants, instPoint)
		trips[tripIndex].SOG_instants = append(trips[tripIndex].SOG_instants, instFloat)

		// Print marker for batches
		recordCount++
		if recordCount%NO_INSTANTS_BATCH == 0 {
			fmt.Print("*")
		}
	}

	fmt.Printf("\n%d records read.\n", recordCount)
	fmt.Printf("%d trips read.\n", numships)

	// Print the results
	for i := 0; i < numships; i++ {
		trip := trips[i]
		fmt.Printf("MMSI: %d, Number of input instants: %d\n", trip.MMSI, len(trip.SOG_instants))
		// Process further as needed, e.g., calculate distance traveled, time-weighted average SOG, etc.
	}

	// Write the output to a new CSV file
	outputFile, err := os.Create("data/ais_trips_instants.csv")
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Write the header
	writer.Write([]string{"mmsi", "trip_instants", "sog_instants"})

	// Write the trip data
	for i := 0; i < numships; i++ {
		trip := trips[i]
		for j, trip_instant := range trip.trip_instants {
			writer.Write([]string{
				strconv.FormatInt(trip.MMSI, 10),
				trip_instant.TPointOut(5),
				trip.SOG_instants[j].TPointOut(5),
			})
		}
	}
	for i := 0; i < numships; i++ {
		trips[i].trip = gomeos.TGeogPointSeqMake(trips[i].trip_instants, len(trips[i].trip_instants), true, true, "LINEAR", true)
		fmt.Printf("  Trip -> Number of instants: %d, Distance travelled %f\n", gomeos.TemporalNumInstants[gomeos.TGeogPointSeq](trips[i].trip), gomeos.TPointLength(trips[i].trip))
		trips[i].SOG = gomeos.TFloatSeqMake(trips[i].SOG_instants, len(trips[i].SOG_instants), true, true, "LINEAR", true)
		fmt.Printf("  Trip -> Number of instants: %d, Time-weighted average %f\n", gomeos.TemporalNumInstants[gomeos.TFloatSeq](trips[i].SOG), gomeos.TnumberTwavg(trips[i].SOG))
	}

	// Write the output to a new CSV file
	outputFile2, err2 := os.Create("data/ais_trips_seq.csv")
	if err2 != nil {
		log.Fatalf("Error creating output file: %v", err2)
	}
	defer outputFile2.Close()
	fmt.Println("create file!")
	writer2 := csv.NewWriter(outputFile2)
	defer writer2.Flush()

	// Write the header
	writer2.Write([]string{"mmsi", "trip", "sog"})
	fmt.Println("header finish!!")
	for i := 0; i < numships; i++ {
		writer2.Write([]string{
			strconv.FormatInt(trips[i].MMSI, 6),
			trips[i].trip.TPointOut(6),
			trips[i].SOG.TPointOut(6),
		})
		fmt.Println("one line written!")
	}
}
