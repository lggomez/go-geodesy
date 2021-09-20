package geodesy

import "math"

type Point [2]float64

func (p Point) Lat() float64 {
	return p[0]
}

func (p Point) LatRadians() float64 {
	return (p[0]*math.Pi)/180
}

func (p Point) Lon() float64 {
	return p[1]
}

func (p Point) LonRadians() float64 {
	return (p[1]*math.Pi)/180
}

func (p Point) Antipode() Point {
	return Point{-p[0], 180 - math.Abs(p[1])}
}

func (p Point) IsAntipode(p2 Point) bool {
	// Shorthand check to avoid Equals() calls between p and p2
	return ((p[0] == -p2[0]) && (p[1] == (180 - math.Abs(p2[1])))) ||
		(p2[0] == -p[0]) && (p2[1] == (180 - math.Abs(p[1])))
}

func (p Point) Equals(p2 Point) bool {
	return (p[0] == p2[0]) && (p[1] == p2[1])
}
