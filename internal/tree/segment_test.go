package tree_test

import (
	"testing"

	"github.com/edualb-challenge/treebabel/internal/tree"
)

func TestNewSegment(t *testing.T) {
	type input struct {
		vector []float64
	}
	type output struct {
		vector []float64
	}
	type testCase struct {
		name string
		in   input
		out  output
	}

	tests := []testCase{
		{
			name: "should return exponential of two array",
			in: input{
				vector: []float64{0, 0, 0, 0, 0, 0, 0, 0},
			},
			out: output{
				vector: []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			segTree := tree.NewSegment(tc.in.vector)
			if len(segTree.Seg) != len(tc.out.vector) {
				t.Errorf("unexpected length of segment tree, got: %d, wants %d", len(segTree.Seg), len(tc.out.vector))
				t.FailNow()
			}
			for i, v := range segTree.Seg {
				if v != tc.out.vector[i] {
					t.Errorf("unexpected value in segment tree, got: %f, wants %f", v, tc.out.vector[i])
					t.FailNow()
				}
			}
		})
	}
}

func TestSegmentSetMovingAverage(t *testing.T) {
	type input struct {
		vector []float64
		index  []int64
		value  []float64
	}
	type output struct {
		vector []float64
	}
	type testCase struct {
		name string
		in   input
		out  output
	}

	tests := []testCase{
		{
			name: "should return build tree with just left node with value",
			in: input{
				vector: []float64{0, 0, 0, 0, 0, 0, 0, 0},
				index:  []int64{4},
				value:  []float64{20},
			},
			out: output{
				vector: []float64{20, 20, 0, 0, 20, 0, 0, 0, 0, 0, 20, 0, 0, 0, 0, 0},
			},
		},
		{
			name: "should return build tree with left and right nodes with value",
			in: input{
				vector: []float64{0, 0, 0, 0, 0, 0, 0, 0},
				index:  []int64{4, 8},
				value:  []float64{20, 31},
			},
			out: output{
				vector: []float64{25.5, 20, 31, 0, 20, 0, 31, 0, 0, 0, 20, 0, 0, 0, 31, 0},
			},
		},
		{
			name: "should return build tree with left and right nodes with values when we set 3 times in the same index",
			in: input{
				vector: []float64{0, 0, 0, 0, 0, 0, 0, 0},
				index:  []int64{4, 8, 4, 4},
				value:  []float64{10, 31, 7, 3},
			},
			out: output{
				vector: []float64{25.5, 20, 31, 0, 20, 0, 31, 0, 0, 0, 20, 0, 0, 0, 31, 0},
			},
		},
		{
			name: "should return build tree with one value in right and more than one value in left",
			in: input{
				vector: []float64{0, 0, 0, 0, 0, 0, 0, 0},
				index:  []int64{4, 8, 3},
				value:  []float64{20, 31, 10},
			},
			out: output{
				vector: []float64{23, 15, 31, 0, 15, 0, 31, 0, 0, 10, 20, 0, 0, 0, 31, 0},
			},
		},
		{
			name: "should return build tree with more than one value in right and left",
			in: input{
				vector: []float64{0, 0, 0, 0, 0, 0, 0, 0},
				index:  []int64{4, 8, 3, 6},
				value:  []float64{20, 31, 10, 40},
			},
			out: output{
				vector: []float64{25.25, 15, 35.5, 0, 15, 40, 31, 0, 0, 10, 20, 0, 40, 0, 31, 0},
			},
		},
		{
			name: "should return build tree with more than one value in right and left when we set 2 times in the same index and one of them is 0",
			in: input{
				vector: []float64{0, 0, 0, 0, 0, 0, 0, 0},
				index:  []int64{4, 8, 3, 6, 4},
				value:  []float64{20, 31, 10, 40, 0},
			},
			out: output{
				vector: []float64{25.25, 15, 35.5, 0, 15, 40, 31, 0, 0, 10, 20, 0, 40, 0, 31, 0},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			segTree := tree.NewSegment(tc.in.vector)

			for i := range tc.in.index {
				segTree.Set(tc.in.index[i], tc.in.value[i])
			}

			for i, v := range segTree.Seg {
				if v != tc.out.vector[i] {
					t.Errorf("unexpected value in segment tree, got: %f (index %d), wants %f (index %d)", v, i, tc.out.vector[i], i)
				}
			}
		})
	}
}

func TestSegmentQueryMovingAverage(t *testing.T) {
	type input struct {
		vector     []float64
		index      []int64
		value      []float64
		indexLeft  int64
		indexRight int64
	}
	type output struct {
		value float64
	}
	type testCase struct {
		name string
		in   input
		out  output
	}

	// In order to facilitates the manutenability, we include the 'tree' comment
	// that represents the current tree the test is related with
	tests := []testCase{
		{
			name: "should return valid value when indexes are different",
			in: input{
				vector:     []float64{0, 0, 0, 0, 0, 0, 0, 0},
				index:      []int64{4},
				value:      []float64{20},
				indexLeft:  0,
				indexRight: 5,
			},
			out: output{
				// tree: []{20, 20, 0, 0, 20, 0, 0, 0, 0, 0, 20, 0, 0, 0, 0, 0},
				value: 20,
			},
		},
		{
			name: "should return 0 when indexes are equal",
			in: input{
				vector:     []float64{0, 0, 0, 0, 0, 0, 0, 0},
				index:      []int64{4},
				value:      []float64{20},
				indexLeft:  4,
				indexRight: 4,
			},
			out: output{
				// tree: []{20, 20, 0, 0, 20, 0, 0, 0, 0, 0, 20, 0, 0, 0, 0, 0},
				value: 0,
			},
		},
		{
			name: "should return valid value when indexes are next to (left side)",
			in: input{
				vector:     []float64{0, 0, 0, 0, 0, 0, 0, 0},
				index:      []int64{4},
				value:      []float64{20},
				indexLeft:  3,
				indexRight: 4,
			},
			out: output{
				// tree: []{20, 20, 0, 0, 20, 0, 0, 0, 0, 0, 20, 0, 0, 0, 0, 0},
				value: 20,
			},
		},
		{
			name: "should return valid value when indexes are next to (right side)",
			in: input{
				vector:     []float64{0, 0, 0, 0, 0, 0, 0, 0},
				index:      []int64{4},
				value:      []float64{20},
				indexLeft:  5,
				indexRight: 6,
			},
			out: output{
				// tree: []{20, 20, 0, 0, 20, 0, 0, 0, 0, 0, 20, 0, 0, 0, 0, 0},
				value: 0,
			},
		},
		{
			name: "should return root value with complex tree",
			in: input{
				vector:     []float64{0, 0, 0, 0, 0, 0, 0, 0},
				index:      []int64{4, 8, 3, 6},
				value:      []float64{20, 31, 10, 40},
				indexLeft:  0,
				indexRight: 8,
			},
			out: output{
				// tree: []{25.25, 15, 35.5, 0, 15, 40, 31, 0, 0, 10, 20, 0, 40, 0, 31, 0},
				value: 25.25,
			},
		},
		{
			name: "should return valid value when use indexes in the middle of the complex tree",
			in: input{
				vector:     []float64{0, 0, 0, 0, 0, 0, 0, 0},
				index:      []int64{4, 8, 2, 6},
				value:      []float64{20, 31, 10, 40},
				indexLeft:  3,
				indexRight: 6,
			},
			out: output{
				// tree: []{25.25, 15, 35.5, 0, 15, 40, 31, 0, 0, 10, 20, 0, 40, 0, 31, 0},
				value: 30,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			segTree := tree.NewSegment(tc.in.vector)

			for i := range tc.in.index {
				segTree.Set(tc.in.index[i], tc.in.value[i])
			}

			got := segTree.Query(tc.in.indexLeft, tc.in.indexRight)
			if got != tc.out.value {
				t.Errorf("unexpected value, got: %f, wants: %f", got, tc.out.value)
				t.FailNow()
			}
		})
	}
}
