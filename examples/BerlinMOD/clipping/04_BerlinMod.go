package main

//MEOS example: meos/examples/06_berlinmod_clip.c

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/MobilityDB/GoMEOS/gomeos"
)

// Constants
const (
	MaxLengthTrip   = 170001
	MaxLengthGeom   = 100001
	MaxLengthHeader = 1024
	MaxLengthName   = 101
	MaxLengthDate   = 12
	NoVehicles      = 5
	NoCommunes      = 19
)

// Struct Definitions
type CommuneRecord struct {
	id         int
	name       string
	population int
	geom       *gomeos.Geom // Assuming GSerialized is a custom struct for handling geometry data
}

type RegionRecord struct {
	name string
	geom *gomeos.Geom // Assuming GSerialized is a custom struct for handling geometry data
}

type TripRecord struct {
	tripid int
	vehid  int
	seq    int
	trip   *gomeos.TGeomPointSeq // Assuming Temporal is a custom struct for handling temporal data
}

// Arrays for computing results
var (
	communes       [NoCommunes]CommuneRecord
	distance       [NoVehicles + 1][NoCommunes + 3]float64
	tripBuffer     [MaxLengthTrip]byte
	geoBuffer      [MaxLengthGeom]byte
	headerBuffer   [MaxLengthHeader]byte
	dateBuffer     [MaxLengthDate]byte
	brusselsRegion RegionRecord
)

// Read communes from file
func readCommunes() int {
	file, err := os.Open("data/brussels_communes.csv")
	if err != nil {
		fmt.Println("Error opening input file 'brussels_communes.csv'")
		return 1
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields

	// Read the header line
	_, err = reader.Read()
	if err != nil {
		fmt.Println("Error reading header line:", err)
		return 1
	}

	noRecords := 0
	for {
		// Read each record (line)
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break // Exit loop if end of file is reached
			}
			fmt.Println("Error reading input file 'brussels_communes.csv'", err)
			return 1
		}

		// Check if the record has at least 4 fields
		if len(record) < 4 {
			fmt.Println("Commune record with missing values")
			return 1
		}

		// Parse the record fields
		id, _ := strconv.Atoi(record[0])
		population, _ := strconv.Atoi(record[2])
		geom := gomeos.NewGeom(record[3], -1)

		// Store the commune record
		communes[noRecords] = CommuneRecord{
			id:         id,
			name:       record[1],
			population: population,
			geom:       &geom,
		}
		noRecords++
	}

	fmt.Printf("%d commune records read\n", noRecords)
	return 0
}

func readBrusselsRegion() error {
	// Open the CSV file
	file, err := os.Open("data/brussels_region.csv")
	if err != nil {
		fmt.Println("Error opening input file 'brussels_region.csv'")
		return err
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields per record

	// Read the first line with the headers
	_, err = reader.Read()
	if err != nil {
		fmt.Println("Error reading headers from file")
		return err
	}

	// Read the next line for the Brussels region data
	record, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading region record")
		return err
	}

	// Check if we have exactly 2 fields
	if len(record) != 2 {
		fmt.Println("Region record with missing values")
		return fmt.Errorf("invalid record length")
	}

	// Assign the name and geometry
	brusselsRegion.name = record[0]
	geom := gomeos.NewGeom(record[1], -1)
	brusselsRegion.geom = &geom

	// Print success message
	fmt.Println("Brussels region record read")

	return nil
}

// Function to print the distance matrix
func matrixPrint(distance [NoVehicles + 1][NoCommunes + 3]float64, allCommunes bool) {
	var buf strings.Builder

	// Print table header
	buf.WriteString("\n                --")
	for j := 1; j < NoCommunes+2; j++ {
		if allCommunes || distance[NoVehicles][j] != 0 {
			buf.WriteString("---------")
		}
	}
	buf.WriteString("\n                | Commmunes\n    --------------")
	for j := 1; j < NoCommunes+2; j++ {
		if allCommunes || distance[NoVehicles][j] != 0 {
			buf.WriteString("---------")
		}
	}
	buf.WriteString("\nVeh | Distance | ")
	for j := 1; j < NoCommunes+1; j++ {
		if allCommunes || distance[NoVehicles][j] != 0 {
			buf.WriteString(fmt.Sprintf("   %2d   ", j))
		}
	}
	buf.WriteString("|  Inside | Outside\n")
	for j := 0; j < NoCommunes+3; j++ {
		if allCommunes || distance[NoVehicles][j] != 0 {
			buf.WriteString("---------")
		}
	}
	buf.WriteString("\n")

	// Print for each vehicle
	for i := 0; i < NoVehicles; i++ {
		buf.WriteString(fmt.Sprintf(" %2d | %8.3f |", i+1, distance[i][0]))
		for j := 1; j <= NoCommunes; j++ {
			if allCommunes || distance[NoVehicles][j] != 0 {
				buf.WriteString(fmt.Sprintf(" %7.3f", distance[i][j]))
			}
		}
		for j := NoCommunes + 1; j < NoCommunes+3; j++ {
			buf.WriteString(fmt.Sprintf(" | %7.3f", distance[i][j]))
		}
		buf.WriteString("\n")
	}

	// Print the total row
	for j := 0; j < NoCommunes+3; j++ {
		if allCommunes || distance[NoVehicles][j] != 0 {
			buf.WriteString("---------")
		}
	}
	buf.WriteString(fmt.Sprintf("\n    | %8.3f |", distance[NoVehicles][0]))
	for j := 1; j <= NoCommunes; j++ {
		if allCommunes || distance[NoVehicles][j] != 0 {
			buf.WriteString(fmt.Sprintf(" %7.3f", distance[NoVehicles][j]))
		}
	}
	for j := NoCommunes + 1; j < NoCommunes+3; j++ {
		buf.WriteString(fmt.Sprintf(" | %7.3f", distance[NoVehicles][j]))
	}
	buf.WriteString("\n")
	for j := 0; j < NoCommunes+3; j++ {
		if allCommunes || distance[NoVehicles][j] != 0 {
			buf.WriteString("---------")
		}
	}
	buf.WriteString("\n\n")
	fmt.Print(buf.String())
}

func main() {
	// Initialize MEOS
	gomeos.MeosInitialize("UTC")

	// Read communes and Brussels region data
	if readCommunes() != 0 || readBrusselsRegion() != nil {
		log.Fatal("Failed to read input files")
	}

	// Open the CSV file
	file, err := os.Open("data/berlinmod_trips.csv")
	if err != nil {
		fmt.Println("Error opening input file 'berlinmod_trips.csv'")
		return
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields

	// Read the header line
	_, err = reader.Read()
	if err != nil {
		fmt.Println("Error reading header line:", err)
		return
	}

	fmt.Println("Processing trip records (one '*' marker per trip)")

	noRecords := 0
	for {
		// Read each record (line)
		record, err := reader.Read()
		if err != nil {
			break // Exit loop if no more records or error
		}

		// Check if the record has at least 5 fields
		if len(record) < 5 {
			fmt.Println("Trip record with missing values")
			return
		}

		// Parse the record fields
		tripid, _ := strconv.Atoi(record[0])
		vehid, _ := strconv.Atoi(record[1])
		seq, _ := strconv.Atoi(record[3])
		tripWKB := record[4]
		// fmt.Println("tripWKB", tripWKB)

		// fmt.Println("parse the record successfully!")
		// Create a new TripRecord
		tripRec := TripRecord{
			tripid: tripid,
			vehid:  vehid,
			seq:    seq,
			trip:   gomeos.NewTGeomPointSeqFromWKB(tripWKB),
		}

		// fmt.Println("new trip record created!")
		noRecords++
		fmt.Print("*")

		// Compute the total distance
		d := gomeos.TPointLength(tripRec.trip) / 1000
		// fmt.Println("distance calculated!", d)
		distance[tripRec.vehid-1][0] += d
		distance[NoVehicles][0] += d

		// Loop for each commune
		for i := 0; i < NoCommunes; i++ {
			// fmt.Println("loop commune ", i)
			// fmt.Println("trip:", tripRec.trip.TPointOut(5))
			// fmt.Println("commune:", communes[i].geom.GeomOut())
			atgeom := gomeos.TpointAtGeomTime(tripRec.trip, &gomeos.TGeomPointSeq{}, communes[i].geom)
			// fmt.Println("atgeomtime success!")
			if atgeom.Inner() != nil {
				d = gomeos.TPointLength(atgeom) / 1000
				distance[tripRec.vehid-1][i+1] += d
				distance[tripRec.vehid-1][NoCommunes+1] += d
				distance[NoVehicles][i+1] += d
				distance[NoVehicles][NoCommunes+1] += d
			}
		}

		// Compute the distance outside Brussels Region
		minusgeom := gomeos.TpointMinusGeomTime(tripRec.trip, &gomeos.TGeomPointSeq{}, brusselsRegion.geom)
		if minusgeom.Inner() != nil {
			// Calculate the distance of the trip outside Brussels Region
			d := gomeos.TPointLength(minusgeom) / 1000
			// Add to the corresponding row and column totals
			distance[tripRec.vehid-1][NoCommunes+2] += d
			distance[NoVehicles][NoCommunes+2] += d
		}
	}

	fmt.Printf("\n%d trip records read.\n\n", noRecords)
	matrixPrint(distance, true)
}
