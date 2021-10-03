package distance_test

import (
	"math"
	"testing"

	"github.com/lggomez/go-geodesy"
	"github.com/lggomez/go-geodesy/distance"
	"github.com/stretchr/testify/assert"
)

func TestHaversine(t *testing.T) {
	antipodeOriginNW := geodesy.Point{42.358312, -95.310466}
	antipodeNW := antipodeOriginNW.Antipode()

	antipodeOriginNE := geodesy.Point{62.379312, 99.612962}
	antipodeNE := antipodeOriginNE.Antipode()

	antipodeOriginSW := geodesy.Point{-54.839747, 66.500319}
	antipodeSW := antipodeOriginSW.Antipode()

	antipodeOriginSE := geodesy.Point{-46.272337, 169.398118}
	antipodeSE := antipodeOriginSE.Antipode()

	type args struct {
		p1 geodesy.Point
		p2 geodesy.Point
	}
	tests := []struct {
		name             string
		args             args
		expectedDistance float64
	}{
		{
			name: "OK/equal_points",
			args: args{
				p1: geodesy.Point{-34.579340, -57.534954},
				p2: geodesy.Point{-34.579340, -57.534954},
			},
			expectedDistance: 0,
		},
		{
			name: "OK/sub_1k_km",
			args: args{
				p1: geodesy.Point{-37.550643, -56.51251},
				p2: geodesy.Point{-37.5507, -56.5126},
			},
			expectedDistance: 10.15491528106308,
		},
		{
			name: "OK/Geoscience_Australia_Testcase",
			args: args{
				p1: geodesy.Point{-37.57037203, 144.25295244},
				p2: geodesy.Point{-37.39101561, 143.55353839},
			},
			expectedDistance: 64_858.3251025962,
		},
		{
			name: "OK/NW-SE_over_10000km",
			args: args{
				p1: geodesy.Point{43.916325, -119.352141},
				p2: geodesy.Point{-32.239202, 150.621015},
			},
			expectedDistance: 1.2424241373877214e+07,
		},
		{
			name: "OK/equator",
			args: args{
				p1: geodesy.Point{0, -71.313379},
				p2: geodesy.Point{0, -73.15691},
			},
			expectedDistance: 204991.57653826568,
		},
		{
			name: "OK/antipodeNW_aprox",
			args: args{
				p1: antipodeOriginNW,
				p2: geodesy.Point{antipodeNW.Lat() + 0.5, antipodeNW.Lon() + 0.5},
			},
			expectedDistance: 1.9945887707923543e+07,
		},
		{
			name: "OK/antipodeNE_aprox",
			args: args{
				p1: antipodeOriginNE,
				p2: geodesy.Point{antipodeNE.Lat() + 0.1, antipodeNE.Lon() + 0.1},
			},
			expectedDistance: 1.39540965209409e+07,
		},
		{
			name: "OK/antipodeSE_aprox",
			args: args{
				p1: antipodeOriginSE,
				p2: geodesy.Point{antipodeSE.Lat() + 0.1, antipodeSE.Lon() + 0.1},
			},
			expectedDistance: 1.8384089351069085e+07,
		},
		{
			name: "OK/antipodeSW_aprox",
			args: args{
				p1: antipodeOriginSW,
				p2: geodesy.Point{antipodeSW.Lat() + 0.1, antipodeSW.Lon() + 0.1},
			},
			expectedDistance: 1.2938700886812994e+07,
		},
		{
			name: "OK/antipode",
			args: args{
				p1: geodesy.Point{40.698470, -73.951442},
				p2: geodesy.Point{-40.698470, 106.048558},
			},
			expectedDistance: 2.0015114352233686e+07,
		},
		{
			name: "OK/antipode2",
			args: args{
				p1: geodesy.Point{40.698470, -73.951442},
				p2: geodesy.Point{40.698470, -73.951442}.Antipode(),
			},
			expectedDistance: 2.0015114352233686e+07,
		},
		{
			name: "FAIL/invalid_p1",
			args: args{
				p1:               geodesy.Point{geodesy.LatUpperBound+1, -57.534954},
				p2:               geodesy.Point{-34.579340, -57.534954},
			},
			expectedDistance: math.NaN(),
		},
		{
			name: "FAIL/invalid_p2",
			args: args{
				p1:               geodesy.Point{-34.579340, -57.534954},
				p2:               geodesy.Point{geodesy.LatUpperBound+1, -57.534954},
			},
			expectedDistance: math.NaN(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := distance.Haversine(tt.args.p1, tt.args.p2)
			if math.IsNaN(tt.expectedDistance) {
				assert.True(t, math.IsNaN(d), "got %f", d)
			} else {
				assert.EqualValues(t, tt.expectedDistance, d)
			}
		})
	}
}
