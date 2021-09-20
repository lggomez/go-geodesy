package distance_test

import (
	"testing"

	"github.com/lggomez/go-geodesy"
	"github.com/lggomez/go-geodesy/distance"
	"github.com/stretchr/testify/assert"
)

func TestHaversine(t *testing.T) {
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
				p1:               geodesy.Point{43.916325, -119.352141},
				p2:               geodesy.Point{-32.239202, 150.621015},
			},
			expectedDistance: 1.2424241373877214e+07,
		},
		{
			name: "OK/antipode",
			args: args{
				p1:               geodesy.Point{40.698470, -73.951442},
				p2:               geodesy.Point{-40.698470, 106.048558},
			},
			expectedDistance: 2.0015114352233686e+07,
		},
		{
			name: "OK/antipode2",
			args: args{
				p1:               geodesy.Point{40.698470, -73.951442},
				p2:               geodesy.Point{40.698470, -73.951442}.Antipode(),
			},
			expectedDistance: 2.0015114352233686e+07,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := distance.Haversine(tt.args.p1, tt.args.p2)
			assert.Equal(t, tt.expectedDistance, d)
		})
	}
}
