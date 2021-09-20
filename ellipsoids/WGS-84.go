package ellipsoids
/*
	This file contains the constant definitions for the WGS-84 ellipsoid

	See https://es.wikipedia.org/wiki/WGS84, https://ahrs.readthedocs.io/en/latest/wgs84.html and
	https://espace.curtin.edu.au/bitstream/handle/20.500.11937/41042/146669_24797_compend.pdf?isAllowed=y&sequence=2 ("A Compendium Of Earth Constants Relevant To Australian Geodetic Science")
	for more information
*/

// Defining geometrical constants
const (
	// WGS84_SEMI_MAJOR_AXIS Semi major axis a, defined in meters (m)
	WGS84_SEMI_MAJOR_AXIS float64 = 6_378_137
)

// Defining physical constants
const (
	// WGS84_GEOCENTRIC_GRAVITATIONAL_CONSTANT Geocentric gravitational constant GM, defined in (m^3)/(s^2)
	WGS84_GEOCENTRIC_GRAVITATIONAL_CONSTANT float64 = 3_986_005_000_000_000

	// WGS84_DYNAMICAL_FORM_FACTOR Dynamical form factor J2; adimensional
	// See https://ahrs.readthedocs.io/en/latest/wgs84.html#ahrs.utils.wgs84.WGS.dynamical_form_factor
	WGS84_DYNAMICAL_FORM_FACTOR float64 = 0.0010826298213129219

	// WGS84_ANGULAR_VELOCITY Dynamical form factor ω; defined in s^-1
	WGS84_ANGULAR_VELOCITY float64 = 0.0007292115
)

// Derived geometrical constants (all rounded)
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
	WGS84_FLATTENING float64 = 1/298.257223563
	// WGS84_FLATTENING_INVERSE Flattening inverse (1/f); adimensional
	WGS84_FLATTENING_INVERSE float64 = 298.257223563
)

// Derived physical constants (all rounded)
const (
	// WGS84_ROTATION_PERIOD Period of rotation (sidereal day) = 2π/ω; defined in seconds (s)
	WGS84_ROTATION_PERIOD float64 = 8_616.4100637
)