package db

import (
	"encoding/csv"
	"strings"
	"testing"

	"github.com/joaoh82/housinganywhere/pkg/domain"
)

func TestReadList(t *testing.T) {
	in := `"id","lat","lng"
    636618,50.853433,5.6841425
    22310,51.9214966,4.506075
    19328,52.14884880000000322,4.44838690000000003
    28257,52.0122851,4.3521148
    41124,51.917448,4.526022999999999
`

	var expedtedList = domain.Locations{
		domain.Location{Id: string(636618), Lat: 50.853433, Long: 5.6841425},
		domain.Location{Id: string(22310), Lat: 51.9214966, Long: 4.506075},
		domain.Location{Id: string(19328), Lat: 52.14884880000000322, Long: 4.44838690000000003},
		domain.Location{Id: string(28257), Lat: 52.0122851, Long: 4.3521148},
		domain.Location{Id: string(41124), Lat: 51.917448, Long: 4.526022999999999},
	}

	r := csv.NewReader(strings.NewReader(in))
	r.LazyQuotes = true
	csvfile := &CSVFile{r: r}
	resultList, err := csvfile.ReadList()
	if err != nil {
		t.Error("Error: ", err)
	}
	if len(resultList) != len(expedtedList) {
		t.Errorf("Locations length: expected %d, actual %d", len(expedtedList), len(resultList))
	}
}
