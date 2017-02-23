package geo

import "math"

// Haversin(0) function
func haversin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// round helper function to round the distances mathmatically
func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
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

	r = 6373000 // Earth radius in Meters

	// Calculating
	h := haversin(la2-la1) + math.Cos(la1)*math.Cos(la2)*haversin(lo2-lo1)

	return round(2*r*math.Asin(math.Sqrt(h)), .5, 0)
}
