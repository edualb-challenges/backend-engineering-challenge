package tree

func NewSegment(vector []float64, defaultValue float64) *Segment {
	tree := &Segment{}

	// preparing leaves. It needs to be a exponecial of two.
	n := 1
	for n < len(vector) {
		n = 2 * n
	}

	// building tree (with on more exponencial of two) and set all values as default value
	for i := 0; i < 2*n; i++ {
		tree.Seg = append(tree.Seg, defaultValue)
	}
	tree.DefaultValue = defaultValue
	tree.Size = int64(n)
	return tree
}
