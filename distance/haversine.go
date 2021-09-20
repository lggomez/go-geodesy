package distance

import (
	"math"

	"github.com/lggomez/go-geodesy"
	"github.com/lggomez/go-geodesy/ellipsoids"
)

// Haversine calculates the ellipsoidal distance in meters between 2 points using the Haversine formula
func Haversine(p1, p2 geodesy.Point) float64 {
	if p1.Equals(p2) {
		return 0
	}

	latRadians2 := p2.LatRadians()
	latRadians1 := p1.LatRadians()
	latDiff := math.Sin((latRadians2 - latRadians1) / 2)
	lonDiff := math.Sin((p2.LonRadians() - p1.LonRadians()) / 2)

	root := math.Sqrt(latDiff*latDiff + lonDiff*lonDiff*math.Cos(latRadians1)*math.Cos(latRadians2))

	d := 2 * ellipsoids.WGS84_MEAN_RADIUS * math.Asin(root)

	return d
}
