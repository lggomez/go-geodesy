package distance_test

import (
	"github.com/lggomez/go-geodesy"
	"github.com/lggomez/go-geodesy/distance"
	"github.com/stretchr/testify/assert"
	"testing"
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
		//{
		//	name:             "OK/equal_points",
		//	args:             args{
		//		p1:               geodesy.Point{-34.579340, -57.534954},
		//		p2:               geodesy.Point{-34.579340, -57.534954},
		//		accuracy:         -1,
		//		calculateAzimuth: true,
		//	},
		//	expectedDistance: 0,
		//	expectedAzimuth1: 0,
		//	expectedAzimuth2: 0,
		//},
		//{
		//	name:             "OK/equal_points_w_azimuth",
		//	args:             args{
		//		p1:               geodesy.Point{-34.579340, -57.534954},
		//		p2:               geodesy.Point{-34.579340, -57.534954},
		//		accuracy:         -1,
		//		calculateAzimuth: false,
		//	},
		//	expectedDistance: 0,
		//	expectedAzimuth1: 0,
		//	expectedAzimuth2: 0,
		//},
		{
			name:             "OK/Argentina_sub_500m",
			args:             args{
				p1:               geodesy.Point{-37.550643, -56.51251},
				p2:               geodesy.Point{-34.555733, -58.520749},
				accuracy:         -1,
				calculateAzimuth: false,
			},
			expectedDistance: 378358.863,
			expectedAzimuth1: 0,
			expectedAzimuth2: 0,
		},
		//{
		//	name:             "OK/Argentina_sub_500m_azimuth",
		//	args:             args{
		//		p1:               geodesy.Point{-37.550643, -56.51251},
		//		p2:               geodesy.Point{-34.555733, -58.520749},
		//		accuracy:         -1,
		//		calculateAzimuth: true,
		//	},
		//	expectedDistance: 378358.863,
		//	expectedAzimuth1: 330.834988,
		//	expectedAzimuth2: 152.017393,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, az1, az2 := distance.VicentyInverse(tt.args.p1, tt.args.p2, tt.args.accuracy, tt.args.calculateAzimuth)
			assert.Equal(t, tt.expectedDistance, d)
			assert.Equal(t, tt.expectedAzimuth1, az1)
			assert.Equal(t, tt.expectedAzimuth2, az2)
		})
	}
}
