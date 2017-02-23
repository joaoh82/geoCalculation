package geo

import "math"

// Haversin(0) function
func haversin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// round helper function to round the distances mathematically
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

// Distance function returns the distance (in meters) between two points on
// earth of a given longitude and latitude relatively
// through the Haversin Distance Formula for great arc distance on a sphere with
// accuracy for small distances.
// Point coordinates are supplied in degrees and converted into radius.
func Distance(latOrigin, lonOrigin, latDest, lonDest float64) float64 {
	// Converting to radians
	var la1, lo1, la2, lo2, r float64
	la1 = latOrigin * math.Pi / 180
	lo1 = lonOrigin * math.Pi / 180
	la2 = latDest * math.Pi / 180
	lo2 = lonDest * math.Pi / 180

	r = 6373000 // Earth radius in Meters

	// Calculating
	h := haversin(la2-la1) + math.Cos(la1)*math.Cos(la2)*haversin(lo2-lo1)

	// Rounding and returning the result
	return round(2*r*math.Asin(math.Sqrt(h)), .5, 0)
}
