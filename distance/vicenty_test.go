package distance_test

import (
	"math"
	"testing"

	"github.com/lggomez/go-geodesy"
	"github.com/lggomez/go-geodesy/distance"
	"github.com/stretchr/testify/assert"
)

func TestVicentyInverse(t *testing.T) {
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
			expectedDistance: 64985.585355322924,
			expectedAzimuth1: 330.834988,
			expectedAzimuth2: 152.017393,
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
