package nearestneighbors_test

import (
	"math/rand"
	"slices"
	"testing"

	nearestneighbors "github.com/natemcintosh/NearestNeighbors"
	"github.com/stretchr/testify/assert"
)

func FuzzExtrema(f *testing.F) {
	f.Add(int64(1), uint8(10))
	f.Fuzz(func(t *testing.T, seed int64, size uint8) {
		// Make sure it won't be empty
		if size == 0 {
			t.Skip()
		}

		// Use the seed to generate some random numbers
		r := rand.New(rand.NewSource(seed))

		// Make the slice
		s := make([]float64, size)
		// Fill it up with random values
		for i := 0; i < int(size); i++ {
			s[i] = r.Float64()
		}

		gotmn, gotmx := nearestneighbors.Extrema(s)
		wantmn := slices.Min(s)
		assert.Equal(t, wantmn, gotmn)

		wantmx := slices.Max(s)
		assert.Equal(t, wantmx, gotmx)
	})
}
