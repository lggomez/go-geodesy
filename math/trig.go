package math

import "math"

// Sin2 calculates the Taylor expansion of sin²x up to the 5th term,
// allowing for similar precision for x in the approximate domain of (-2;2)
func Sin2(x float64) float64 {
	return math.Pow(x, 2) -
		math.Pow(x, 4)/3 +
		(2*math.Pow(x, 6))/45 -
		math.Pow(x, 8)/315 +
		(2*math.Pow(x, 10))/14175
}

// Cos2 calculates the Taylor expansion of cos²x up to the 5th term,
// allowing for similar precision for x in the approximate domain of (-2;2)
func Cos2(x float64) float64 {
	return 1 - math.Pow(x, 2) +
		math.Pow(x, 4)/3 -
		(2*math.Pow(x, 6))/45 +
		math.Pow(x, 8)/315 -
		(2*math.Pow(x, 10))/14175
}
