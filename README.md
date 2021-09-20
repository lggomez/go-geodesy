# geodesy

go-geodesy is a package of geodesy-related utilities made in golang, including ellipsoid constants for development of geographic and geospatial implementations 

```
    import "github.com/lggomez/go-geodesy"
```

- [geodesy](#geodesy)
	- [Usage](#usage)
		- [Point definitions](#point-definitions)
			- [type Point](#type-point)
			- [func (Point) Antipode](#func-point-antipode)
			- [func (Point) Equals](#func-point-equals)
			- [func (Point) IsAntipode](#func-point-isantipode)
			- [func (Point) Lat](#func-point-lat)
			- [func (Point) LatRadians](#func-point-latradians)
			- [func (Point) Lon](#func-point-lon)
			- [func (Point) LonRadians](#func-point-lonradians)
		- [Calculating distances](#calculating-distances)
			- [func  Haversine](#func--haversine)
			- [func  VicentyInverse](#func--vicentyinverse)
		- [Ellipsoids](#ellipsoids)
			- [GRS-80](#grs-80)
			- [WGS-84](#wgs-84)

## Usage

### Point definitions

#### type Point

```go
type Point [2]float64
```

Point represents a latitude-longitude pair in decimal degrees

#### func (Point) Antipode

```go
func (p Point) Antipode() Point
```
Antipode returns a new point representing the geographical antipode of p

#### func (Point) Equals

```go
func (p Point) Equals(p2 Point) bool
```
Equals returns whether p is equal in latitude and longitude to p2

#### func (Point) IsAntipode

```go
func (p Point) IsAntipode(p2 Point) bool
```
IsAntipode returns whether p is the exact antipode of p2 or not

#### func (Point) Lat

```go
func (p Point) Lat() float64
```
Lat returns point p's latitude

#### func (Point) LatRadians

```go
func (p Point) LatRadians() float64
```
LatRadians returns point p's latitude in radians

#### func (Point) Lon

```go
func (p Point) Lon() float64
```
Lon returns point p's longitude

#### func (Point) LonRadians

```go
func (p Point) LonRadians() float64
```
LonRadians returns point p's longitude in radians

### Calculating distances
```
    import "github.com/lggomez/go-geodesy/distance"
```

#### func  Haversine

```go
func Haversine(p1, p2 geodesy.Point) float64
```
Haversine calculates the ellipsoidal distance in meters between 2 points using
the Haversine formula

#### func  VicentyInverse

```go
func VicentyInverse(p1, p2 geodesy.Point, accuracy float64, calculateAzimuth bool) (float64, float64, float64)
```

VicentyInverse calculates the ellipsoidal distance in meters and azimuth in degrees between 2 points using the inverse Vicenty formulae and the WGS-84 ellipsoid constants. 

The following notations are used:

    * a 	length of semi-major axis of the ellipsoid (radius at equator)
    * ƒ 	flattening of the ellipsoid
    * b = (1 − ƒ) a 	length of semi-minor axis of the ellipsoid (radius at the poles)
    * u1 = arctan( (1 − ƒ) tan lat1 ) 	reduced latitude for p1 (latitude on the auxiliary sphere);
    * u2 = arctan( (1 − ƒ) tan lat2 ) 	reduced latitude for p2 (latitude on the auxiliary sphere);
    * L1, L2 	longitude of the points;
    * L = L2 − L1 	difference in longitude of two points;
    * λ 	Difference in longitude of the points on the auxiliary sphere;
    * α1, α2 	forward azimuths at the points;
    * α 	forward azimuth of the geodesic at the equator, if it were extended that far;
    * s 	ellipsoidal distance between the two points;
    * σ 	angular separation between points;
    * σ1 	angular separation between the point and the equator;
    * σm 	angular separation between the midpoint of the line and the equator;

### Ellipsoids

```
     import "github.com/lggomez/go-geodesy/ellipsoids"
```

#### GRS-80
```go
const (
	// Geocentric gravitational constant GM, defined in (m^3)/(s^2)
	GRS80_GEOCENTRIC_GRAVITATIONAL_CONSTANT float64 = 3_986_005_000_000_000

	// Dynamical form factor J2; adimensional
	GRS80_DYNAMICAL_FORM_FACTOR float64 = 0.0108263

	// Dynamical form factor ω; defined in s^-1
	GRS80_ANGULAR_VELOCITY float64 = 0.0007292115
)
```
Defining physical constants

```go
const (
	// Semi minor axis b, defined in meters (m)
	GRS80_SEMI_MINOR_AXIS float64 = 6_356_752.314140

	// Aspect ratio (b/a); adimensional
	GRS80_ASPECT_RATIO float64 = 0.996647189318816362

	// Mean radius R1 = (2a+b)/3, defined in meters (m)
	GRS80_MEAN_RADIUS = 6_371_008.7714
	// Mean radius R2, defined in meters (m)
	GRS80_AUTHALIC_MEAN_RADIUS = 6_371_007.1810
	// Radius of a sphere of the same volume R3 = ((a^2)*b)^(1/3); defined in meters (m)
	GRS80_SPHERE_RADIUS = 6_371_000.7900
	// Polar radius of curvature = (a^2)/b; defined in meters (m)
	GRS80_POLAR_CURVATURE_RADIUS = 6_399_593.6259
	// Equatorial radius of curvature for a meridian = (b^2)/a; defined in meters (m)
	GRS80_MERIDIAN_CURVATURE_EQUATORIAL_RADIUS = 6_335_439.3271

	// Meridian quadrant (meridian quarter); defined in meters (m)
	// See https://en.wikipedia.org/wiki/Meridian_arc#Quarter_meridian
	GRS80_MERIDIAN_QUADRANT = 10_001_965.7293

	// Linear eccentricity c = sqrt((a^2)-(b^2)); defined in meters (m)
	GRS80_LINEAR_ECCENTRICITY = 521_854.0097
	// Eccentricity of elliptical section through poles e = sqrt((a^2)-(b^2))/a; adimensional
	GRS80_LINEAR_ECCENTRICITY_POLES = 0.0818191910435

	// Flattening f; adimensional
	GRS80_FLATTENING float64 = 0.003352810681183637418
	// Flattening inverse (1/f); adimensional
	GRS80_FLATTENING_INVERSE float64 = 1 / 0.003352810681183637418
)
```
Derived geometrical constants (all rounded)

```go
const (
	// Period of rotation (sidereal day) = 2π/ω; defined in seconds (s)
	GRS80_ROTATION_PERIOD float64 = 8_616.4100637
)
```
Derived physical constants (all rounded)

```go
const (
	// Semi major axis a, defined in meters (m)
	GRS80_SEMI_MAJOR_AXIS float64 = 6_378_137
)
```

Derived geometrical constants (all rounded)

#### WGS-84

```go
const (
	// WGS84_GEOCENTRIC_GRAVITATIONAL_CONSTANT Geocentric gravitational constant GM, defined in (m^3)/(s^2)
	WGS84_GEOCENTRIC_GRAVITATIONAL_CONSTANT float64 = 3_986_005_000_000_000

	// WGS84_DYNAMICAL_FORM_FACTOR Dynamical form factor J2; adimensional
	// See https://ahrs.readthedocs.io/en/latest/wgs84.html#ahrs.utils.wgs84.WGS.dynamical_form_factor
	WGS84_DYNAMICAL_FORM_FACTOR float64 = 0.0010826298213129219

	// WGS84_ANGULAR_VELOCITY Dynamical form factor ω; defined in s^-1
	WGS84_ANGULAR_VELOCITY float64 = 0.0007292115
)
```
Defining physical constants

```go
const (
	// WGS84_SEMI_MINOR_AXIS Semi minor axis b, defined in meters (m)
	WGS84_SEMI_MINOR_AXIS float64 = 6_356_752.31424518

	// WGS84_ASPECT_RATIO Aspect ratio (b/a); adimensional
	WGS84_ASPECT_RATIO float64 = 0.9966471893352525

	// WGS84_MEAN_RADIUS Mean radius R1 = (2a+b)/3, defined in meters (m)
	WGS84_MEAN_RADIUS = 6_371_008.771415059
	// WGS84_AUTHALIC_MEAN_RADIUS Mean radius R2, defined in meters (m)
	WGS84_AUTHALIC_MEAN_RADIUS = 6_371_007.1809182055
	// WGS84_SPHERE_RADIUS Radius of a sphere of the same volume R3 = ((a^2)*b)^(1/3); defined in meters (m)
	WGS84_SPHERE_RADIUS = 6_371_000.79000916
	// WGS84_POLAR_CURVATURE_RADIUS Polar radius of curvature = (a^2)/b; defined in meters (m)
	WGS84_POLAR_CURVATURE_RADIUS = 6_399_593.625758493
	// WGS84_MERIDIAN_CURVATURE_EQUATORIAL_RADIUS Equatorial radius of curvature for a meridian = (b^2)/a; defined in meters (m)
	WGS84_MERIDIAN_CURVATURE_EQUATORIAL_RADIUS = 6335439.327292821

	// WGS84_MERIDIAN_QUADRANT Meridian quadrant (meridian quarter); defined in meters (m)
	// See https://en.wikipedia.org/wiki/Meridian_arc#Quarter_meridian
	WGS84_MERIDIAN_QUADRANT = 10_001_965.729

	// WGS84_LINEAR_ECCENTRICITY Linear eccentricity c = sqrt((a^2)-(b^2)); defined in meters (m)
	WGS84_LINEAR_ECCENTRICITY = 521_854.0084234
	// WGS84_LINEAR_ECCENTRICITY_POLES Eccentricity of elliptical section through poles e = sqrt((a^2)-(b^2))/a; adimensional
	WGS84_LINEAR_ECCENTRICITY_POLES = 0.0818191918426205

	// WGS84_FLATTENING Flattening f; adimensional
	WGS84_FLATTENING float64 = 1 / 298.257223563
	// WGS84_FLATTENING_INVERSE Flattening inverse (1/f); adimensional
	WGS84_FLATTENING_INVERSE float64 = 298.257223563
)
```
Defining geometrical constants

```go
const (
	// WGS84_ROTATION_PERIOD Period of rotation (sidereal day) = 2π/ω; defined in seconds (s)
	WGS84_ROTATION_PERIOD float64 = 8_616.4100637
)
```
Derived physical constants (all rounded)

```go
const (
	// WGS84_SEMI_MAJOR_AXIS Semi major axis a, defined in meters (m)
	WGS84_SEMI_MAJOR_AXIS float64 = 6_378_137
)
```
Defining geometrical constants
