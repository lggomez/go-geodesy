package distance

import (
	"math/rand"
	"testing"
	"time"

	"github.com/lggomez/go-geodesy"
)

var nearPoints = make([][2]geodesy.Point, 10000, 10000)
var separatePoints = make([][2]geodesy.Point, 10000, 10000)
var result1, result2, result3 = float64(0), float64(0), float64(0)

func init() {
	rand.Seed(time.Now().UnixNano())

	initializeNearPoints()
	initializeSeparatePoints()
}

func initializeSeparatePoints() {
	// first world quadrant
	for i := 0; i < 10000; i++ {
		nearPoints[i] = [2]geodesy.Point{
			{-90 + rand.Float64()*(90 - -90), -180 + rand.Float64()*(180 - -180)},
			{-90 + rand.Float64()*(90 - -90), -180 + rand.Float64()*(180 - -180)},
		}
	}
}

func initializeNearPoints() {
	// first world quadrant
	for i := 0; i < 2500; i++ {
		nearPoints[i] = [2]geodesy.Point{
			{40 + rand.Float64()*(50-0), -150 + rand.Float64()*(-160 - -150)},
			{40 + rand.Float64()*(50-0), -150 + rand.Float64()*(-160 - -150)},
		}
	}

	// second world quadrant
	for i := 2501; i < 5000; i++ {
		nearPoints[i] = [2]geodesy.Point{
			{40 + rand.Float64()*(50-0), 150 + rand.Float64()*(160-150)},
			{40 + rand.Float64()*(50-0), 150 + rand.Float64()*(160-150)},
		}
	}

	// third world quadrant
	for i := 5001; i < 7500; i++ {
		nearPoints[i] = [2]geodesy.Point{
			{-50 + rand.Float64()*(-50 - -60), -150 + rand.Float64()*(-160 - -150)},
			{-50 + rand.Float64()*(-50 - -60), -150 + rand.Float64()*(-160 - -150)},
		}
	}

	// fourth world quadrant
	for i := 7501; i < 10000; i++ {
		nearPoints[i] = [2]geodesy.Point{
			{-50 + rand.Float64()*(-50 - -60), 150 + rand.Float64()*(160-150)},
			{-50 + rand.Float64()*(-50 - -60), 150 + rand.Float64()*(160-150)},
		}
	}
}

func Benchmark_DistanceNear(b *testing.B) {
	b.Run("Vincenty_noazimuth", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			idx := i % 10000
			result1, result2, result3 = VincentyInverse(nearPoints[idx][0], nearPoints[idx][1], -1, false)
		}
	})
	b.Run("Vincenty", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			idx := i % 10000
			result1, result2, result3 = VincentyInverse(nearPoints[idx][0], nearPoints[idx][1], -1, true)
		}
	})
	b.Run("Haversine", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			idx := i % 10000
			result1 = Haversine(nearPoints[idx][0], nearPoints[idx][1])
		}
	})
}

func Benchmark_DistanceFar(b *testing.B) {
	b.Run("Vincenty_noazimuth", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			idx := i % 10000
			result1, result2, result3 = VincentyInverse(separatePoints[idx][0], separatePoints[idx][1], -1, false)
		}
	})
	b.Run("Vincenty", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			idx := i % 10000
			result1, result2, result3 = VincentyInverse(separatePoints[idx][0], separatePoints[idx][1], -1, true)
		}
	})
	b.Run("Haversine", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			idx := i % 10000
			result1 = Haversine(separatePoints[idx][0], separatePoints[idx][1])
		}
	})
}
