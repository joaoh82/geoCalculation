package main

import (
	"fmt"
	"log"

	flags "github.com/jessevdk/go-flags"

	"github.com/joaoh82/housinganywhere/pkg/db"
	"github.com/joaoh82/housinganywhere/pkg/domain"
)

// Reaading the path using a commeand-line flag (the returned value is a pointer)
var opts struct {
	Path string `short:"p" description:"Path of the CSV file" required:"true"`
	DB   string `short:"d" description:"Type of db or data that is coming in" required:"true"`
}

// Constants representing the types of data types that could be informed in the tags
const (
	csv     = "csv"
	mongodb = "mongo"
)

var database db.Database

func main() {
	// Parsing the flags
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	// Checking to see which type of data was provided (Only csv actually implemented at the moment)
	switch opts.DB {
	case csv:
		// Setting the location of Housing Anywhere
		db.OriginLocation = domain.Location{Lat: 51.925146, Long: 4.478617}
		// Reading the CSV file
		database = db.NewCSVFile(opts.Path)
		handleResults()
	case mongodb:
		// Where the MongoDB setu would be
		log.Println("MongoDB setting up...")
		// Call handleResults() with the mongo database set up
	default:
		log.Println("Sorry, type of database not recognized.")
	}
}

// handleResults handles the results and presents the 5 closests
// and 5 furthest locations from housing anywhere
func handleResults() {
	list, err := database.ReadList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n//---------------//5 CLOSEST LOCATIONS//--------------//")
	fmt.Println("\n//ID: DISTANCE (in meters)")
	// Slicing the ordered list with the top 5 results
	for _, c := range list[:5] {
		fmt.Printf("%v: %.0fm\n", c.Id, c.Distance)
	}

	fmt.Println("\n//---------------//5 FURTHEST LOCATIONS//--------------//")
	fmt.Println("\n//ID: DISTANCE (in meters)")
	// Slicing the list with the bottom 5 results
	for _, c := range list[len(list)-5:] {
		fmt.Printf("%v: %.0fm\n", c.Id, c.Distance)
	}
}
