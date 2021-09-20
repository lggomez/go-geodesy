package distance_test

import (
	"math"
	"testing"

	"github.com/lggomez/go-geodesy"
	"github.com/lggomez/go-geodesy/distance"
	"github.com/stretchr/testify/assert"
)

func TestVicentyInverse(t *testing.T) {
	antipodeOrigin := geodesy.Point{40.698470, -73.951442}
	antipode := antipodeOrigin.Antipode()

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
			name: "OK/Geoscience_Australia_Testcase",
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
			name: "OK/antipode_aprox",
			args: args{
				p1:               antipodeOrigin,
				p2:               geodesy.Point{math.Ceil(antipode.Lat()), math.Ceil(antipode.Lon())},
				accuracy:         -1,
				calculateAzimuth: true,
			},
			expectedDistance: 1.990042228821669e+07,
			expectedAzimuth1: 323.2114038644566,
			expectedAzimuth2: 36.35011968130334,
		},
		{
			name: "FAIL/antipode",
			args: args{
				p1:               antipodeOrigin,
				p2:               antipode,
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
			d, az1, az2 := distance.VicentyInverse(tt.args.p1, tt.args.p2, tt.args.accuracy, tt.args.calculateAzimuth)
			if math.IsNaN(tt.expectedDistance) {
				assert.True(t, math.IsNaN(d))
			} else {
				assert.EqualValues(t, tt.expectedDistance, d)
			}
			if math.IsNaN(tt.expectedAzimuth1) {
				assert.True(t, math.IsNaN(az1))
			} else {
				assert.EqualValues(t, tt.expectedAzimuth1, az1)
			}
			if math.IsNaN(tt.expectedAzimuth2) {
				assert.True(t, math.IsNaN(az2))
			} else {
				assert.EqualValues(t, tt.expectedAzimuth2, az2)
			}
		})
	}
}
