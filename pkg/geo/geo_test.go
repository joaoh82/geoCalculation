package geo

import (
	"testing"

	"github.com/joaoh82/housinganywhere/pkg/domain"
)

var originLocation = domain.Location{Lat: 51.925146, Long: 4.478617}
var disTests = []struct {
	latInt   float64 // input
	longInt  float64 // input
	expected float64 // expected result
}{
	{37.1768672, -3.608897, 1758.633},
	{52.36461880000000235, 4.93169289999999982, 57.844},
	{51.9245615, 4.492032399999999, 0.923},
	{52.1542113, 4.491864899999999, 25.495},
}

// TestDistance tests the Distance function that implemented the Haversin function
// Using as a base for the Earth Radius 6373km, which can vary a small amount depending
// on the source you use
func TestDistance(t *testing.T) {
	for _, tt := range disTests {
		actual := Distance(originLocation.Lat, originLocation.Long, tt.latInt, tt.longInt) / 1000
		if actual != tt.expected {
			t.Errorf("Dis(%f, %f): expected %f, actual %f", tt.latInt, tt.longInt, tt.expected, actual)
		}
	}
}
