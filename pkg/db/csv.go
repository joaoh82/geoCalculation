package db

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/joaoh82/housinganywhere/pkg/geo"
)

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
	// csv := CSVFile{r: r}

	return &CSVFile{r: r}
}

func (f *CSVFile) ReadList() error {

	for {
		rec, err := f.r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			//log.Fatal(err)
			return err
		}

		// Checking if record is a string or a number, if it is a string
		// means we are on the first row and we move to the next one
		if _, err := strconv.Atoi(rec[0]); err != nil {
			rec, err = f.r.Read()
		}

		// Getting lat and long values
		id := rec[0]
		lat := rec[1]
		long := rec[2]
		latFloat, err := strconv.ParseFloat(lat, 64)
		if err != nil {
			//log.Fatalf("Record, error: %v, %v", lat, err)
			return err
		}
		longFloat, err := strconv.ParseFloat(long, 64)
		if err != nil {
			//log.Fatalf("Record, error: %v, %v", long, err)
			return err
		}

		// calculate scores; THIS EXTERNAL METHOD CANNOT BE CHANGED
		distance := geo.Distance(51.925146, 4.478617, latFloat, longFloat)
		fmt.Printf("%v: %f\n", id, distance)
	}
	return nil
}
