package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"

	flags "github.com/jessevdk/go-flags"
)

// Reaading the path using a commeand-line flag (the returned value is a pointer)
//var pathFlag = flag.String("path", "", "Path of CSV file")
var opts struct {
	Path string `short:"p" long:"path" description:"Path of the CSV file" required:"true"`
}

func init() {
	// Adding the option for a shart flag for path also
	//flag.StringVar(pathFlag, "p", "", "Path of CSV file")
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	// Setup reader
	csvIn, err := os.Open(opts.Path)
	if err != nil {
		// Checking if error is os.PathError and giving a more friendly message.
		// Just to show the conpect of type assertion here
		if _, ok := err.(*os.PathError); ok {
			log.Fatal("File not found. Please verify if the file is the provided path.")
		}
		log.Fatal(err)
	}
	r := csv.NewReader(csvIn)
	for {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		// Checking if record is a string or a number, if it is a string
		// means we are on the first row and we move to the next one
		if _, err := strconv.Atoi(rec[0]); err != nil {
			rec, err = r.Read()
		}

		// Getting lat and long values
		id := rec[0]
		lat := rec[1]
		long := rec[2]
		latFloat, err := strconv.ParseFloat(lat, 64)
		if err != nil {
			log.Fatalf("Record, error: %v, %v", lat, err)
		}
		longFloat, err := strconv.ParseFloat(long, 64)
		if err != nil {
			log.Fatalf("Record, error: %v, %v", long, err)
		}

		// calculate scores; THIS EXTERNAL METHOD CANNOT BE CHANGED
		distance := Distance(51.925146, 4.478617, latFloat, longFloat)
		fmt.Printf("%v: %f\n", id, distance)
	}

	//fmt.Println(Distance(51.925146, 4.478617, 37.1768672, -3.60897))
}

// Haversin(0) function
func haversin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// Distance function returns the distance (in meters) between two points of
// a given longitude and latitude relatively accurately (using a spherical
// approximation of the Earth) through the Haversin Distance Formula for
// great arc distance on a pshere with accuracy for small distances
// Point coordinates are supplied in degrees and converted into radius.
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// Converting to radians
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in Meters

	// Calculating
	h := haversin(la2-la1) + math.Cos(la1)*math.Cos(la2)*haversin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}
