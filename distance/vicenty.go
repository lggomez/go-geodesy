package distance

import (
	"github.com/lggomez/go-geodesy"
	"github.com/lggomez/go-geodesy/ellipsoids"
	"math"
)

const (
	defaultAccuracy float64 = 1e-12 // approximates to 0.06 mm
	maxIterations   int     = 100
)

func VicentyInverse(p1, p2 geodesy.Point, accuracy float64, calculateAzimuth bool) (float64, float64, float64) {
	if p1.Equals(p2) {
		return 0, 0, 0
	}

	ε := defaultAccuracy
	if accuracy > 0 {
		ε = accuracy
	}

	// Initial conditions setup
	a := ellipsoids.WGS84_SEMI_MAJOR_AXIS
	b := ellipsoids.WGS84_SEMI_MINOR_AXIS
	f := ellipsoids.WGS84_FLATTENING
	u1 := math.Atan((1 - f) * math.Tanh(p1.LatRadians()))
	u2 := math.Atan((1 - f) * math.Tanh(p2.LatRadians()))
	L := p2.LonRadians() - p1.LonRadians()
	λ := L
	λ_prev := float64(0)
	f16Frac := f / 16
	sin_u2, cos_u2 := math.Sincos(u2)
	sin_u1, cos_u1 := math.Sincos(u1)

	// Loop variables
	cos2_α := float64(0)
	sin_σ, cos_σ := float64(0), float64(0)
	sin_λ, cos_λ := float64(0), float64(0)
	cos_2σm := float64(0)
	σ := float64(0)

	// Perform iterative evaluation of λ until it either converges to ε or reaches the maximum amount of iterations
	for i := 0; math.Abs(λ-λ_prev) > ε; i++ {
		// Test for divergence and nearly antipodal points
		if i > maxIterations || (math.Abs(λ) > math.Pi) {
			return math.NaN(), math.NaN(), math.NaN()
		}

		sin_λ, cos_λ = math.Sincos(λ)
		sin_σ = math.Sqrt((cos_u2*sin_λ)*(cos_u2*sin_λ) +
			((cos_u1*sin_u2)-(sin_u1*cos_u2*cos_λ))*((cos_u1*sin_u2)-(sin_u1*cos_u2*cos_λ)))
		if sin_σ == 0 {
			return 0, 0, 0
		}

		cos_σ = (sin_u1 * sin_u2) + (cos_u1 * cos_u2 * cos_λ)
		σ = math.Atan2(sin_σ, cos_σ)

		sin_α := (cos_u1 * cos_u2 * sin_λ) / sin_σ
		cos2_α = 1 - (sin_α * sin_α)
		cos_2σm = cos_σ - ((2 * sin_u1 * sin_u2) / cos2_α)

		C := f16Frac * cos2_α * (4 + f*(4-3*cos2_α))

		λ_prev = λ
		λ = L + (1-C)*f*sin_α*(
			σ+C*sin_σ*(
				cos_2σm+C*cos_σ*(
					-1+2*cos_2σm*cos_2σm)))
	}

	// Setup return variables
	α1 := float64(0)
	α2 := float64(0)
	d := float64(0)

	bSquared := b * b
	uSquared := cos2_α * (((a * a) - bSquared) / bSquared)

	A := 1 + (uSquared/16384)*(4096+uSquared*(-768+uSquared*(320-175*uSquared)))
	B := (uSquared / 1024) * (256 + uSquared*(-128+uSquared*(74-47*uSquared)))
	Δσ := B * sin_σ * (
		cos_2σm + (B/4)*(
			cos_σ*(-1+2*(cos_2σm*cos_2σm)) -
				(B/6)*cos_2σm*(-3+4*(sin_σ*sin_σ))*(-3+4*(cos_2σm*cos_2σm))))

	d = b * A * (σ - Δσ)

	if calculateAzimuth {
		α1 = math.Atan2(cos_u2*sin_λ, (cos_u1*sin_u2)-(sin_u1*cos_u2*cos_λ))
		α2 = math.Atan2(cos_u2*sin_λ, -(sin_u1*cos_u2)+(cos_u2*sin_u2*cos_λ))
	}

	return d, α1, α2
}
