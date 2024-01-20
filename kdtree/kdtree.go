package kdtree

import (
	"errors"
	"math"

	nn "github.com/natemcintosh/NearestNeighbors"
)

type bounding_box struct {
	mins  [2]float64
	maxes [2]float64
}

type kdnode struct {
	left_child  *kdnode
	right_child *kdnode
	bbox        bounding_box
	inds        []int
}

type KDTree[T nn.Number] struct {
	rawx        []T
	rawy        []T
	parent_node *kdnode
}

const leaf_size = 10

func NewKDTree[T nn.Number](x, y []T) (KDTree[T], error) {
	if (len(x) == 0) || (len(y) == 0) {
		return KDTree[T]{nil, nil, nil},
			errors.New("please hand in slices with at least one element each")
	}

	xmin, xmax := nn.Extrema(x)
	ymin, ymax := nn.Extrema(y)

	bbox := bounding_box{
		mins:  [2]float64{float64(xmin), float64(ymin)},
		maxes: [2]float64{float64(xmax), float64(ymax)},
	}

	// The indices slice
	inds := make([]int, len(x))
	for i := 0; i < len(inds); i++ {
		inds[i] = i
	}

	// The parent node
	parent_node := kdnode{
		left_child:  &kdnode{},
		right_child: &kdnode{},
		bbox:        bbox,
		inds:        inds,
	}

	// Recursively partition the space into axis aligned rectangles
	// In each step, partition along the dimension with largest spread

	// Split on the median in partition dimension

	// Recurse until number of points in node is less than leaf node. Here hard coded to
	// be 10
}

func build_kdtree[T nn.Number](
	index int,
	x, y []T,
	bbox bounding_box,
	nodes []kdnode,
	inds []int,
	low, high int,
) {
	// Do not continue to partition if reached leaf size
	n_p := high - low + 1
	if n_p <= leaf_size {
		return
	}

	// Find the splitting point index
	mid_idx := find_split(low, leaf_size, n_p)

	// Find the dimension where the spread is maximal
	xspread := bbox.maxes[0] - bbox.mins[0]
	yspread := bbox.maxes[1] - bbox.mins[1]
	max_spread := 0.0
	split_dim := -1
	if xspread > yspread {
		max_spread = xspread
		split_dim = 0
	} else {
		max_spread = yspread
		split_dim = 1
	}

	// Get the split value
	// spit_val := math.
}

// find_split gets the middle index on which to split. Copied as closely as possible
// from the Julia version. The tree is split such that one of the sub trees has exactly
// 2^p points, and such that the left sub tree always has more points.
// This means that we can deterministically (with just some comparisons) find if we are
// at a leaf node and how many.
func find_split(low, leafsize, n_p int) int {
	// The number of leafs node left in the tree,
	// use `ceil` to count a partially filled node as 1.
	n_leafs := math.Ceil(float64(n_p) / float64(leaf_size))

	// Number of leftover nodes needed
	k := math.Floor(math.Log2(float64(n_leafs)))
	rest := n_leafs - math.Pow(2, k)

	// The conditionals here fulfill the desired splitting procedure but
	// can probably be written in a nicer way

	// Can fill less than two nodes -> leafsize to left node.
	var mid_idx int
	if n_p <= 2*leafsize {
		mid_idx = leafsize
	} else if rest > math.Pow(2, k-1) {
		// The last leaf node will be in the right sub tree -> fill the left
		// sub tree with
		// Last node over the "half line" in the row
		mid_idx = int(math.Pow(2, k)) * leafsize
	} else if rest == 0 {
		// Perfectly filling both sub trees -> half to left and right sub tree
		mid_idx = int(math.Pow(2, k-1)) * leafsize
	} else {
		// Else we fill the right sub tree -> send the rest to the left sub tree
		mid_idx = n_p - int(math.Pow(2, k-1))*leafsize
	}
	return mid_idx + low
}
