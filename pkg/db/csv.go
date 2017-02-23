package db

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/joaoh82/housinganywhere/pkg/geo"

	"github.com/joaoh82/housinganywhere/pkg/domain"
)

var sortedList domain.Locations

// OriginLocation stores the location of Housing Anywhere, or
// any other place of origin to calculate the distance
var OriginLocation domain.Location

// CSVFile tyoe
type CSVFile struct {
	r *csv.Reader
}

// NewCSVFile creates a new csv.Reader with path passed
func NewCSVFile(path string) *CSVFile {
	// Setup reader
	csvIn, err := os.Open(path)
	if err != nil {
		// Checking if error is os.PathError and giving a more friendly message.
		// Just to show the conpect of type assertion here
		if _, ok := err.(*os.PathError); ok {
			log.Fatal("File not found. Please verify if the file is the provided path.")
		}
		log.Fatal(err)
	}
	r := csv.NewReader(csvIn)

	return &CSVFile{r: r}
}

// ReadList method is responsable for reading through the csv.Reader
// one record at a time using the Read() function from csv.Reader.Read()
func (f *CSVFile) ReadList() (domain.Locations, error) {
	for {
		rec, err := f.r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		// Checking if record is a string or a number, if it is a string
		// means we are on the first row and we move to the next one
		if rec[0] == "id" {
			rec, err = f.r.Read()
			if err != nil {
				return nil, err
			}
		}

		// Getting lat and long values
		id := rec[0]
		lat := rec[1]
		long := rec[2]

		// Since the record from the CSV is returned as a slice of string
		// we have to convert it to a float64
		latFloat, err := strconv.ParseFloat(lat, 64)
		if err != nil {
			return nil, err
		}
		longFloat, err := strconv.ParseFloat(long, 64)
		if err != nil {
			return nil, err
		}
		// calculate scores; THIS EXTERNAL METHOD CANNOT BE CHANGED
		distance := geo.Distance(OriginLocation.Lat, OriginLocation.Long, latFloat, longFloat)

		// Adds the location to the list using our struct type dnbLocation
		sortedList = append(sortedList, domain.Location{
			Id:       id,
			Lat:      latFloat,
			Long:     longFloat,
			Distance: distance,
		})
	}
	// Sorts the list using the standard sorting algorithm
	sort.Sort(sortedList)

	return sortedList, nil
}
