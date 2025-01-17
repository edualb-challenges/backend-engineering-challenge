package tree

import "fmt"

type Segment struct {
	Seg  []float64
	Size int64
}

// Query get the moving average between index left and index right
func (tree Segment) Query(indexLeft, indexRight int64) float64 {
	// starting in the root of the tree
	return tree.query(indexLeft, indexRight, 0, 0, tree.Size)
}

// Set uses the 'value' to make a sum operation in the current value of the leaf ('index') and updates their parents until reach the root node
func (tree *Segment) Set(index int64, value float64) error {
	// starting in the root of the tree
	return tree.set(index, value, 0, 0, tree.Size)
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

func (tree *Segment) set(index int64, value float64, node, nodeLeft, nodeRight int64) error {
	if index < 0 || index > tree.Size {
		return fmt.Errorf("invalid index")
	}

	// are we in the leaf?
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

	/* updating the node with the moving average value */
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

func (tree Segment) query(indexLeft, indexRight, node, nodeLeft, nodeRight int64) float64 {
	// are we in a invalid index (out of range)?
	if indexLeft >= nodeRight {
		return 0
	}
	if indexRight <= nodeLeft {
		return 0
	}

	// are we in the range of the segment we are looking for?
	if indexLeft <= nodeLeft && indexRight >= nodeRight {
		return tree.Seg[node]
	}

	// are we in the leaf?
	if nodeRight-nodeLeft == 1 {
		return tree.Seg[node]
	}

	/* calculating and getting the moving average */
	mid := tree.mid(nodeLeft, nodeRight)
	leftNode := tree.leftNode(node)
	rightNode := tree.rightNode(node)

	leftValue := tree.query(indexLeft, indexRight, leftNode, nodeLeft, mid)
	rightValue := tree.query(indexLeft, indexRight, rightNode, mid, nodeRight)

	if leftValue == 0 {
		return rightValue
	}
	if rightValue == 0 {
		return leftValue
	}

	return (leftValue + rightValue) / 2
}
