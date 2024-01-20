package nearestneighbors

import (
	"math"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Hypot[T Number](x, y T) float64 {
	return math.Hypot(float64(x), float64(y))
}

// Extrema calculates the min and max of a slice. Returns 2 items, first is
// min, second is max. Should in theory be a little faster than calling `slices.Min()`,
// and then `slices.Max()`.
func Extrema[S ~[]T, T Number](x S) (T, T) {
	if len(x) < 1 {
		panic("slices.Min: empty list")
	}
	mn := x[0]
	mx := x[0]
	for i := 1; i < len(x); i++ {
		mn = min(mn, x[i])
		mx = max(mx, x[i])
	}
	return mn, mx
}
