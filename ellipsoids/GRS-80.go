package ellipsoids
/*
	This file contains the constant definitions for the GRS-80 ellipsoid

	See https://en.wikipedia.org/wiki/Geodetic_Reference_System_1980 and
	http://geoweb.mit.edu/~tah/12.221_2005/grs80_corr.pdf ("Geodetic Reference System 1980")
	for more information
*/

// Defining geometrical constants
const (
	// Semi major axis a, defined in meters (m)
	GRS80_SEMI_MAJOR_AXIS float64 = 6_378_137
)

// Defining physical constants
const (
	// Geocentric gravitational constant GM, defined in (m^3)/(s^2)
	GRS80_GEOCENTRIC_GRAVITATIONAL_CONSTANT float64 = 3_986_005_000_000_000

	// Dynamical form factor J2; adimensional
	GRS80_DYNAMICAL_FORM_FACTOR float64 = 0.0108263

	// Dynamical form factor ω; defined in s^-1
	GRS80_ANGULAR_VELOCITY float64 = 0.0007292115
)

// Derived geometrical constants (all rounded)
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
	GRS80_FLATTENING_INVERSE float64 = 1/0.003352810681183637418
)

// Derived physical constants (all rounded)
const (
	// Period of rotation (sidereal day) = 2π/ω; defined in seconds (s)
	GRS80_ROTATION_PERIOD float64 = 8_616.4100637
)