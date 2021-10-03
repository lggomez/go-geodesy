package distance_test

import (
	"math"
	"testing"

	"github.com/lggomez/go-geodesy"
	"github.com/lggomez/go-geodesy/distance"
	"github.com/stretchr/testify/assert"
)

func TestVincentyInverse(t *testing.T) {
	antipodeOriginNW := geodesy.Point{42.358312, -95.310466}
	antipodeNW := antipodeOriginNW.Antipode()

	antipodeOriginNE := geodesy.Point{62.379312, 99.612962}
	antipodeNE := antipodeOriginNE.Antipode()

	antipodeOriginSW := geodesy.Point{-54.839747, 66.500319}
	antipodeSW := antipodeOriginSW.Antipode()

	antipodeOriginSE := geodesy.Point{-46.272337, 169.398118}
	antipodeSE := antipodeOriginSE.Antipode()

	// Antipodes will fail to converge on the exact points
	// and within a short radius, so induce a non-convergent
	// point pair
	antipodeNonconvOrigin := geodesy.Point{42.35831235, -95.31046631}
	antipodeNonconv := geodesy.Point{antipodeNonconvOrigin.Antipode().Lat() + 0.00000001, antipodeNonconvOrigin.Antipode().Lon() + 0.00000001}

	type args struct {
		p1               geodesy.Point
		p2               geodesy.Point
		accuracy         float64
		calculateAzimuth bool
	}
	tests := []struct {
		name             string
		args             args
		expectedDistance float64
		expectedAzimuth1 float64
		expectedAzimuth2 float64
	}{
		{
			name: "OK/equal_points",
			args: args{
				p1:               geodesy.Point{-34.579340, -57.534954},
				p2:               geodesy.Point{-34.579340, -57.534954},
				accuracy:         -1,
				calculateAzimuth: true,
			},
			expectedDistance: 0,
			expectedAzimuth1: 0,
			expectedAzimuth2: 0,
		},
		{
			name: "OK/equal_points_w_azimuth",
			args: args{
				p1:               geodesy.Point{-34.579340, -57.534954},
				p2:               geodesy.Point{-34.579340, -57.534954},
				accuracy:         -1,
				calculateAzimuth: false,
			},
			expectedDistance: 0,
			expectedAzimuth1: 0,
			expectedAzimuth2: 0,
		},
		{
			name: "OK/SW_sub_1k_km",
			args: args{
				p1:               geodesy.Point{-37.550643, -56.51251},
				p2:               geodesy.Point{-37.5507, -56.5126},
				accuracy:         -1,
				calculateAzimuth: false,
			},
			expectedDistance: 10.162235455135733,
			expectedAzimuth1: math.NaN(),
			expectedAzimuth2: math.NaN(),
		},
		{
			name: "OK/SW_sub_1k_km_less_accurate",
			args: args{
				p1:               geodesy.Point{-37.550643, -56.51251},
				p2:               geodesy.Point{-37.5507, -56.5126},
				accuracy:         1e-6,
				calculateAzimuth: false,
			},
			expectedDistance: 10.149099956337418,
			expectedAzimuth1: math.NaN(),
			expectedAzimuth2: math.NaN(),
		},
		{
			name: "OK/SW/SW_sub_500km",
			args: args{
				p1:               geodesy.Point{-37.550643, -56.51251},
				p2:               geodesy.Point{-34.555733, -58.520749},
				accuracy:         -1,
				calculateAzimuth: false,
			},
			expectedDistance: 378358.8626108233,
			expectedAzimuth1: math.NaN(),
			expectedAzimuth2: math.NaN(),
		},
		{
			name: "OK/SE_Geoscience_Australia_Testcase",
			args: args{
				p1:               geodesy.Point{-37.57037203, 144.25295244},
				p2:               geodesy.Point{-37.39101561, 143.55353839},
				accuracy:         -1,
				calculateAzimuth: true,
			},
			expectedDistance: 64_985.585355322924,
			expectedAzimuth1: 287.62433093048827,
			expectedAzimuth2: 108.04992409663924,
		},
		{
			name: "OK/NW_over_1000km",
			args: args{
				p1:               geodesy.Point{43.916325, -119.352141},
				p2:               geodesy.Point{27.049648, -84.467283},
				accuracy:         -1,
				calculateAzimuth: true,
			},
			expectedDistance: 3.6377487944727764e+06,
			expectedAzimuth1: 109.26559384340659,
			expectedAzimuth2: 310.1608440829011,
		},
		{
			name: "OK/NW-SE_over_10000km",
			args: args{
				p1:               geodesy.Point{43.916325, -119.352141},
				p2:               geodesy.Point{-32.239202, 150.621015},
				accuracy:         -1,
				calculateAzimuth: true,
			},
			expectedDistance: 1.2410562861916382e+07,
			expectedAzimuth1: 245.76137499804267,
			expectedAzimuth2: 50.994718337013865,
		},
		{
			name: "OK/equator",
			args: args{
				p1:               geodesy.Point{0, -71.313379},
				p2:               geodesy.Point{0, -73.15691},
				accuracy:         -1,
				calculateAzimuth: true,
			},
			expectedDistance: 205220.9321815281,
			expectedAzimuth1: 270,
			expectedAzimuth2: 90,
		},
		{
			name: "OK/antipodeNW_aprox",
			args: args{
				p1:               antipodeOriginNW, // this coordinate seems especially hard to converge
				p2:               geodesy.Point{antipodeNW.Lat() + 0.5, antipodeNW.Lon() + 0.5},
				accuracy:         -1,
				calculateAzimuth: true,
			},
			expectedDistance: 1.9939429136002023e+07,
			expectedAzimuth1: 335.01113953197097,
			expectedAzimuth2: 24.780010599686307,
		},
		{
			name: "OK/antipodeNE_aprox",
			args: args{
				p1:               antipodeOriginNE,
				p2:               geodesy.Point{antipodeNE.Lat() + 0.1, antipodeNE.Lon() + 0.1},
				accuracy:         -1,
				calculateAzimuth: true,
			},
			expectedDistance: 1.392052684154077e+07,
			expectedAzimuth1: 190.81695568835386,
			expectedAzimuth2: 10.780632958360911,
		},
		{
			name: "OK/antipodeSE_aprox",
			args: args{
				p1:               antipodeOriginSE,
				p2:               geodesy.Point{antipodeSE.Lat() + 0.1, antipodeSE.Lon() + 0.1},
				accuracy:         -1,
				calculateAzimuth: true,
			},
			expectedDistance: 1.8383995064596053e+07,
			expectedAzimuth1: 277.98586894108723,
			expectedAzimuth2: 97.20292896924542,
		},
		{
			name: "OK/antipodeSW_aprox",
			args: args{
				p1:               antipodeOriginSW,
				p2:               geodesy.Point{antipodeSW.Lat() + 0.1, antipodeSW.Lon() + 0.1},
				accuracy:         -1,
				calculateAzimuth: true,
			},
			expectedDistance: 1.2906415827849764e+07,
			expectedAzimuth1: 28.09660240174846,
			expectedAzimuth2: 208.17248801650823,
		},
		{
			name: "FAIL/divergent",
			args: args{
				p1:               antipodeNonconvOrigin,
				p2:               antipodeNonconv,
				accuracy:         -1,
				calculateAzimuth: true,
			},
			expectedDistance: math.NaN(),
			expectedAzimuth1: math.NaN(),
			expectedAzimuth2: math.NaN(),
		},
		{
			name: "FAIL/antipode",
			args: args{
				p1:               antipodeOriginNW,
				p2:               antipodeNW,
				accuracy:         -1,
				calculateAzimuth: true,
			},
			expectedDistance: math.NaN(),
			expectedAzimuth1: math.NaN(),
			expectedAzimuth2: math.NaN(),
		},
		{
			name: "FAIL/invalid_p1",
			args: args{
				p1:               geodesy.Point{geodesy.LatUpperBound+1, -57.534954},
				p2:               geodesy.Point{-34.579340, -57.534954},
				accuracy:         -1,
				calculateAzimuth: true,
			},
			expectedDistance: math.NaN(),
			expectedAzimuth1: math.NaN(),
			expectedAzimuth2: math.NaN(),
		},
		{
			name: "FAIL/invalid_p2",
			args: args{
				p1:               geodesy.Point{-34.579340, -57.534954},
				p2:               geodesy.Point{geodesy.LatUpperBound+1, -57.534954},
				accuracy:         -1,
				calculateAzimuth: true,
			},
			expectedDistance: math.NaN(),
			expectedAzimuth1: math.NaN(),
			expectedAzimuth2: math.NaN(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, az1, az2 := distance.VincentyInverse(tt.args.p1, tt.args.p2, tt.args.accuracy, tt.args.calculateAzimuth)
			if math.IsNaN(tt.expectedDistance) {
				assert.True(t, math.IsNaN(d), "got %f", d)
			} else {
				assert.EqualValues(t, tt.expectedDistance, d)
			}
			if math.IsNaN(tt.expectedAzimuth1) {
				assert.True(t, math.IsNaN(az1), "got %f", az1)
			} else {
				assert.EqualValues(t, tt.expectedAzimuth1, az1)
			}
			if math.IsNaN(tt.expectedAzimuth2) {
				assert.True(t, math.IsNaN(az2), "got %f", az2)
			} else {
				assert.EqualValues(t, tt.expectedAzimuth2, az2)
			}
		})
	}
}
