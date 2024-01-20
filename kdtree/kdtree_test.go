package kdtree_test

import (
	"testing"

	"github.com/natemcintosh/NearestNeighbors/kdtree"
	"github.com/stretchr/testify/assert"
)

func TestKnnSimple(t *testing.T) {
	testCases := []struct {
		desc      string
		inx       []int
		iny       []int
		qx        int
		qy        int
		k         uint
		wantinds  []int
		wantdists []float64
	}{
		{
			desc:      "one point",
			inx:       []int{0},
			iny:       []int{0},
			qx:        0,
			qy:        0,
			k:         1,
			wantinds:  []int{0},
			wantdists: []float64{0},
		},
		{
			desc:      "two points",
			inx:       []int{-1, 1},
			iny:       []int{-1, 1},
			qx:        -1,
			qy:        -1,
			k:         1,
			wantinds:  []int{0},
			wantdists: []float64{0},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tree, _ := kdtree.NewKDTree(tC.inx, tC.iny)
			indices, distances := tree.Knn(tC.qx, tC.qy, tC.k)

			assert.Equal(t, tC.wantinds, indices)
			assert.InDeltaSlice(t, tC.wantdists, distances, 1e-6)
		})
	}
}

func TestKnnHarder(t *testing.T) {
	testCases := []struct {
		desc      string
		inx       []float64
		iny       []float64
		qx        float64
		qy        float64
		k         uint
		wantinds  []int
		wantdists []float64
	}{
		{
			desc:      "n=2,k=1",
			inx:       []float64{0.291848, 0.289646},
			iny:       []float64{0.618058, 0.219093},
			qx:        1.0,
			qy:        1.0,
			k:         1,
			wantinds:  []int{0},
			wantdists: []float64{0.804585960408029},
		},
		{
			desc:      "n=5,k=1",
			inx:       []float64{0.556478, 0.543232, 0.72787, 0.721939, 0.566468},
			iny:       []float64{0.561301, 0.135792, 0.420113, 0.0294034, 0.96774},
			qx:        0.5,
			qy:        0.5,
			k:         1,
			wantinds:  []int{0},
			wantdists: []float64{0.08335230724764078},
		},
		{
			desc:      "n=5,k=2",
			inx:       []float64{0.556478, 0.543232, 0.72787, 0.721939, 0.566468},
			iny:       []float64{0.561301, 0.135792, 0.420113, 0.0294034, 0.96774},
			qx:        0.5,
			qy:        0.5,
			k:         2,
			wantinds:  []int{0, 2},
			wantdists: []float64{0.08335230724764078, 0.2414674121235417},
		},
		{
			desc:      "n=10,k=5",
			inx:       []float64{2.76964, 1.72057, 4.50333, 1.41612, 3.34827, 1.71495, 3.10408, 4.34419, 2.66255, 0.826588},
			iny:       []float64{0.515162, 3.61385, 3.80529, 2.38902, 3.51642, 1.54887, 0.577434, 3.86879, 1.75372, 1.12639},
			qx:        1.0,
			qy:        1.0,
			k:         5,
			wantinds:  []int{9, 5, 3, 8, 0},
			wantdists: []float64{0.21458227529098237, 0.9013376561806172, 1.4500132915902555, 1.8254216624861987, 1.8348578802128732},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tree, _ := kdtree.NewKDTree(tC.inx, tC.iny)
			indices, distances := tree.Knn(tC.qx, tC.qy, tC.k)

			assert.Equal(t, tC.wantinds, indices)
			assert.InDeltaSlice(t, tC.wantdists, distances, 1e-5)
		})
	}
}

func TestInRange(t *testing.T) {
	testCases := []struct {
		desc     string
		inx      []float64
		iny      []float64
		qx       float64
		qy       float64
		r        float64
		wantinds []int
	}{
		{
			desc:     "n=10,r=3.14",
			inx:      []float64{2.76964, 1.72057, 4.50333, 1.41612, 3.34827, 1.71495, 3.10408, 4.34419, 2.66255, 0.826588},
			iny:      []float64{0.515162, 3.61385, 3.80529, 2.38902, 3.51642, 1.54887, 0.577434, 3.86879, 1.75372, 1.12639},
			qx:       5.0,
			qy:       5.0,
			r:        3.14,
			wantinds: []int{2, 4, 7},
		},
		{
			desc:     "n=10,r=5",
			inx:      []float64{2.76964, 1.72057, 4.50333, 1.41612, 3.34827, 1.71495, 3.10408, 4.34419, 2.66255, 0.826588},
			iny:      []float64{0.515162, 3.61385, 3.80529, 2.38902, 3.51642, 1.54887, 0.577434, 3.86879, 1.75372, 1.12639},
			qx:       1.0,
			qy:       1.0,
			r:        5.0,
			wantinds: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			desc:     "n=10,r=1",
			inx:      []float64{2.76964, 1.72057, 4.50333, 1.41612, 3.34827, 1.71495, 3.10408, 4.34419, 2.66255, 0.826588},
			iny:      []float64{0.515162, 3.61385, 3.80529, 2.38902, 3.51642, 1.54887, 0.577434, 3.86879, 1.75372, 1.12639},
			qx:       2.5,
			qy:       2.5,
			r:        1.0,
			wantinds: []int{8},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tree, _ := kdtree.NewKDTree(tC.inx, tC.iny)
			indices := tree.InRange(tC.qx, tC.qy, tC.r)

			assert.ElementsMatch(t, tC.wantinds, indices)
		})
	}
}
