package brute_force

import (
	"errors"
	"math"
	"slices"

	nn "github.com/natemcintosh/NearestNeighbors"
)

type BruteTree[T nn.Number] struct {
	x, y []T
}

func NewBruteTree[T nn.Number](x, y []T) (BruteTree[T], error) {
	if (len(x) == 0) || (len(y) == 0) {
		return BruteTree[T]{x, y},
			errors.New("please hand in slices with at least one element each")
	}

	return BruteTree[T]{x, y}, nil
}

func (tree BruteTree[T]) Knn(x_query_pt, y_query_pt T, k uint) (indices []int, distances []float64) {
	// If `k` is larger than the number of points in the tree, just return them all
	if k >= uint(len(tree.x)) {
		indices = make([]int, len(tree.x))
		distances = make([]float64, len(tree.x))
		for idx := range tree.x {
			indices[idx] = idx
			distances[idx] = nn.Hypot(x_query_pt-tree.x[idx], y_query_pt-tree.y[idx])
		}
		return indices, distances
	}

	// Make the index slice with the correct capacity, as well as a slice to track the
	// distance of each point to the query point
	indices = make([]int, k)
	distances = make([]float64, k)

	// Fill the indices and distances with the max values
	for idx := range indices {
		indices[idx] = math.MaxInt
		distances[idx] = math.MaxFloat64
	}

	// For each point in the tree, calculate its distance to the query point
	for idx := range tree.x {
		dist := nn.Hypot(x_query_pt-tree.x[idx], y_query_pt-tree.y[idx])
		n, _ := slices.BinarySearch(distances, dist)

		// If we're beyond the end of `k`, then do nothing
		if n >= int(k) {
			continue
		}

		// Insert the distance into the distances slice, and the x y position slices
		indices = slices.Insert(indices, n, idx)
		distances = slices.Insert(distances, n, dist)

		// Keep the slices at the correct length
		indices = indices[0:k]
		distances = distances[0:k]
	}

	return indices, distances
}
