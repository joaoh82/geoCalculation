package main

import (
	"fmt"
	"log"

	flags "github.com/jessevdk/go-flags"

	"github.com/joaoh82/housinganywhere/pkg/db"
)

// Reaading the path using a commeand-line flag (the returned value is a pointer)
var opts struct {
	Path string `short:"p" description:"Path of the CSV file" required:"true"`
	DB   string `short:"d" description:"Type of db or data that is comming in" required:"true"`
}

const (
	csv     = "csv"
	mongodb = "mongo"
)

var database db.Database

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}
	switch opts.DB {
	case csv:
		database = db.NewCSVFile(opts.Path)
		list, err := database.ReadList()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\n//---------------//5 CLOSEST LOCATIONS//--------------//")
		for _, c := range list[:5] {
			fmt.Printf("%v: %f\n", c.Id, c.Distance)
		}
		fmt.Println("\n//---------------//5 FURTHEST LOCATIONS//--------------//")
		for _, c := range list[len(list)-5:] {
			fmt.Printf("%v: %f\n", c.Id, c.Distance)
		}

	case mongodb:
		log.Println("MongoDB setting up...")
	default:
		log.Println("Sorry, type of database not recognized.")
	}
}
