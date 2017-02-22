package domain

// Location struct
type Location struct {
	Id       string
	Lat      float64
	Long     float64
	Distance float64
}

//Locations is the collection that will implement the sort.Interface
type Locations []Location

// Implementinng the sort.Interface on our Locations type
func (slice Locations) Len() int {
	return len(slice)
}

func (slice Locations) Less(i, j int) bool {
	return slice[i].Distance < slice[j].Distance
}

func (slice Locations) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
