package distance_test

import (
	"github.com/lggomez/go-geodesy"
	"github.com/lggomez/go-geodesy/distance"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHaversine(t *testing.T) {
	type args struct {
		p1               geodesy.Point
		p2               geodesy.Point
	}
	tests := []struct {
		name             string
		args             args
		expectedDistance float64
	}{
		{
			name:             "OK/equal_points",
			args:             args{
				p1:               geodesy.Point{-34.579340, -57.534954},
				p2:               geodesy.Point{-34.579340, -57.534954},
			},
			expectedDistance: 0,
		},
		{
			name:             "OK/sub_1k_km",
			args:             args{
				p1:               geodesy.Point{-37.550643, -56.51251},
				p2:               geodesy.Point{-34.555733, -58.520749},
			},
			expectedDistance: 378780.8521527565,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := distance.Haversine(tt.args.p1, tt.args.p2)
			assert.Equal(t, tt.expectedDistance, d)
		})
	}
}
