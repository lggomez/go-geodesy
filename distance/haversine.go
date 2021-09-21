package distance

import (
	"math"

	"github.com/lggomez/go-geodesy"
	"github.com/lggomez/go-geodesy/ellipsoids"
)

// Haversine calculates the ellipsoidal distance in meters between 2 points
// using the Haversine formula and the WGS-84 ellipsoid constants
func Haversine(p1, p2 geodesy.Point) float64 {
	if p1.Equals(p2) {
		return 0
	}

	φ1 := p1.LatRadians()
	φ2 := p2.LatRadians()

	λ1 := p1.LonRadians()
	λ2 := p2.LonRadians()

	latHalfVersine := math.Sin((φ2 - φ1) / 2)
	lonHalfVersine := math.Sin((λ2 - λ1) / 2)

	h := math.Sqrt((latHalfVersine * latHalfVersine) +
		(lonHalfVersine * lonHalfVersine) *
			math.Cos(φ1)*math.Cos(φ2))

	if h > 1 {
		// d is only real for 0<=h<=1
		return math.NaN()
	}

	// Main inverse haversine formula
	return 2 * ellipsoids.WGS84_MEAN_RADIUS * math.Asin(h)
}
