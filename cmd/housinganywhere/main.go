package main

import (
	"log"

	flags "github.com/jessevdk/go-flags"

	"github.com/joaoh82/housinganywhere/pkg/db"
)

// Reaading the path using a commeand-line flag (the returned value is a pointer)
var opts struct {
	Path string `short:"p" long:"path" description:"Path of the CSV file" required:"true"`
}

var database db.Database

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	database = db.NewCSVFile(opts.Path)
	err = database.ReadList()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(Distance(51.925146, 4.478617, 37.1768672, -3.60897))
}
