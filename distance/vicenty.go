package distance

import (
	"math"

	"github.com/lggomez/go-geodesy"
	"github.com/lggomez/go-geodesy/ellipsoids"
)

const (
	defaultAccuracy float64 = 1e-12 // approximates to 0.06 mm
	maxIterations   int     = 50

	fullAngleRad = 2 * math.Pi
	radConversionFactor = 180/math.Pi
)

/*
	VicentyInverse calculates the ellipsoidal distance in meters and azimuth in degrees between 2 points using the
inverse Vicenty formulae and the WGS-84 ellipsoid constants. The following notations are used:
	a 	length of semi-major axis of the ellipsoid (radius at equator)
	ƒ 	flattening of the ellipsoid
	b = (1 − ƒ) a 	length of semi-minor axis of the ellipsoid (radius at the poles)
	u1 = arctan( (1 − ƒ) tan lat1 ) 	reduced latitude for p1 (latitude on the auxiliary sphere);
	u2 = arctan( (1 − ƒ) tan lat2 ) 	reduced latitude for p2 (latitude on the auxiliary sphere);
	L1, L2 	longitude of the points;
	L = L2 − L1 	difference in longitude of two points;
	λ 	Difference in longitude of the points on the auxiliary sphere;
	α1, α2 	forward azimuths at the points;
	α 	forward azimuth of the geodesic at the equator, if it were extended that far;
	s 	ellipsoidal distance between the two points;
	σ 	angular separation between points;
	σ1 	angular separation between the point and the equator;
	σm 	angular separation between the midpoint of the line and the equator;
*/
func VicentyInverse(p1, p2 geodesy.Point, accuracy float64, calculateAzimuth bool) (float64, float64, float64) {
	if p1.Equals(p2) {
		return 0, 0, 0
	}

	if p1.IsAntipode(p2) {
		// Antipodes are non-convergent
		return math.NaN(), math.NaN(), math.NaN()
	}

	ε := defaultAccuracy
	if accuracy > 0 {
		ε = accuracy
	}

	// Initial conditions setup
	a := ellipsoids.WGS84_SEMI_MAJOR_AXIS
	b := ellipsoids.WGS84_SEMI_MINOR_AXIS
	f := ellipsoids.WGS84_FLATTENING
	u1 := math.Atan((1 - f) * math.Tan(p1.LatRadians())) // Reduced latitude for p1
	u2 := math.Atan((1 - f) * math.Tan(p2.LatRadians())) // Reduced latitude for p2
	L := p2.LonRadians() - p1.LonRadians()               // Difference in longitude
	λ := L                                               // Difference in longitude of the points on the auxiliary sphere
	λ_prev := float64(0)
	f16Frac := f / 16
	sinu1, cosu1 := math.Sincos(u1)
	sinu2, cosu2 := math.Sincos(u2)

	// Loop variables
	cos2α := float64(0)
	sinσ, cosσ := float64(0), float64(0)
	sinλ, cosλ := float64(0), float64(0)
	cos2σₘ := float64(0)
	σ := float64(0)
	sinλ, cosλ = math.Sincos(λ)

	// Perform iterative evaluation of λ until it either converges to ε or reaches the maximum amount of iterations
	for i := 0; math.Abs(λ-λ_prev) > ε; i++ {
		// Test for divergence and nearly antipodal points
		if i > maxIterations {
			return math.NaN(), math.NaN(), math.NaN()
		}

		sinσ = math.Sqrt((cosu2*sinλ)*(cosu2*sinλ) +
			((cosu1*sinu2)-(sinu1*cosu2*cosλ))*((cosu1*sinu2)-(sinu1*cosu2*cosλ)))
		if sinσ == 0 {
			// Indeterminate sinα; It represents an end point coincident with,
			// or diametrically opposed to, the start point.
			return math.NaN(), math.NaN(), math.NaN()
		}

		cosσ = (sinu1 * sinu2) + (cosu1 * cosu2 * cosλ)
		σ = math.Atan2(sinσ, cosσ) // Angular separation between points

		sinα := (cosu1 * cosu2 * sinλ) / sinσ
		cos2α = 1 - (sinα * sinα)
		cos2σₘ = cosσ - ((2 * sinu1 * sinu2) / cos2α)

		C := float64(0)
		// Distances through the equator yield C = 0, so calculate if points do not fall on it
		if (p1.Lon() != 0) && (p2.Lon() != 0) {
			C = f16Frac * cos2α * (4 + f*(4-3*cos2α))
		}

		λ_prev = λ
		λ = L + (1-C)*f*sinα*(σ+C*sinσ*(cos2σₘ+C*cosσ*(-1+2*cos2σₘ*cos2σₘ)))
		sinλ, cosλ = math.Sincos(λ)
	}

	// Setup return variables
	α1 := math.NaN()
	α2 := math.NaN()
	d := float64(0)

	bSquared := b * b
	uSquared := cos2α * (((a * a) - bSquared) / bSquared)

	A := 1 + (uSquared/16384)*(4096+uSquared*(-768+uSquared*(320-175*uSquared)))
	B := (uSquared / 1024) * (256 + uSquared*(-128+uSquared*(74-47*uSquared)))
	Δσ := B * sinσ *
		(cos2σₘ + (B/4)*(cosσ*(-1+2*(cos2σₘ*cos2σₘ))-
			(B/6)*cos2σₘ*(-3+4*(sinσ*sinσ))*
				(-3+4*(cos2σₘ*cos2σₘ))))

	d = b * A * (σ - Δσ) // ellipsoidal distance in meters

	if calculateAzimuth {
		numα1 := cosu2*sinλ
		denomα1 := (cosu1*sinu2)-(sinu1*cosu2*cosλ)
		α1 = quadrantRadToDegree(math.Atan2(numα1, denomα1))

		numα2 := cosu1*sinλ
		denomα2 := (-sinu1*cosu2)+(cosu1*sinu2*cosλ)
		α2 = quadrantRadToDegree(math.Atan2(numα2, denomα2))
		α2 = math.Mod(α2 + 180, 360) // Normalize degree to north meridian as origin vector
	}

	return d, α1, α2
}

func quadrantRadToDegree(rad float64) float64 {
	if rad < 0 {
		return (rad + fullAngleRad) * radConversionFactor
	}

	return rad * radConversionFactor
}
