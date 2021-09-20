package geodesy

import "math"

// Point represents a latitude-longitude pair in decimal degrees
type Point [2]float64

// Lat returns point p's latitude
func (p Point) Lat() float64 {
	return p[0]
}

// LatRadians returns point p's latitude in radians
func (p Point) LatRadians() float64 {
	return (p[0]*math.Pi)/180
}

// Lon returns point p's longitude
func (p Point) Lon() float64 {
	return p[1]
}

// LonRadians returns point p's longitude in radians
func (p Point) LonRadians() float64 {
	return (p[1]*math.Pi)/180
}

// Antipode returns a new point representing the geographical antipode of p
func (p Point) Antipode() Point {
	return Point{-p[0], 180 - math.Abs(p[1])}
}

// IsAntipode returns whether p is the exact antipode of p2 or not
func (p Point) IsAntipode(p2 Point) bool {
	// Shorthand check to avoid Equals() calls between p and p2
	return ((p[0] == -p2[0]) && (p[1] == (180 - math.Abs(p2[1])))) ||
		(p2[0] == -p[0]) && (p2[1] == (180 - math.Abs(p[1])))
}

// Equals returns whether p is equal in latitude and longitude to p2
func (p Point) Equals(p2 Point) bool {
	return (p[0] == p2[0]) && (p[1] == p2[1])
}
