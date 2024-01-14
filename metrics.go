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
