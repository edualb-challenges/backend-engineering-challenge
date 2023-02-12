package tree

import "fmt"

type Segment struct {
	Seg          []float64
	DefaultValue float64
	Size         int64
}

func (tree *Segment) Set(index int64, value float64) error {
	return tree.set(index, value, 0, 0, tree.Size)
}

func (tree *Segment) set(index int64, value float64, node, nodeLeft, nodeRight int64) error {
	if index < 0 || index > tree.Size {
		return fmt.Errorf("invalid index")
	}

	// reach the leaf
	if nodeRight-nodeLeft == 1 {
		tree.Seg[node] = value + tree.Seg[node]
		return nil
	}

	mid := tree.mid(nodeLeft, nodeRight)

	if index <= mid {
		// go to left
		tree.set(index, value, tree.leftNode(node), nodeLeft, mid)
	} else {
		// go to right
		tree.set(index, value, tree.rightNode(node), mid, nodeRight)
	}

	/* updating the node with the moving average */
	leftNode := tree.leftNode(node)
	rightNode := tree.rightNode(node)

	if tree.Seg[leftNode] == 0 {
		// just considering the right node even if it is 0
		tree.Seg[node] = tree.Seg[rightNode]
		return nil
	}
	if tree.Seg[rightNode] == 0 {
		// just considering the left node even if it is 0
		tree.Seg[node] = tree.Seg[leftNode]
		return nil
	}

	// just considering the medium between two nodes
	tree.Seg[node] = (tree.Seg[leftNode] + tree.Seg[rightNode]) / 2
	return nil
}

func (tree Segment) mid(nodeLeft, nodeRight int64) int64 {
	return nodeLeft + (nodeRight-nodeLeft)/2
}

func (tree Segment) leftNode(node int64) int64 {
	return 2*node + 1
}

func (tree Segment) rightNode(node int64) int64 {
	return 2*node + 2
}
