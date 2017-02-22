package geo

import "math"

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
