package main

import "math"

type AVLNode struct {
	bf          int
	value       int
	height      int
	parent      *AVLNode
	left, right *AVLNode
}

func (n *AVLNode) leftRotate() {

}

func (n *AVLNode) rightRotate() {
	a := n
	b := n.left
	p := n.parent
	a.left = b.right
	if b.right != nil {
		b.right.parent = a
	}
	b.right = a
	a.parent = b
	b.parent = p
	if p != nil {
		if p.left == a {
			p.left = b
		} else {
			p.right = b
		}
	}
}

func (n *AVLNode) balance() {
	if n.bf == -2 {
		// left left unbalance case

		// left right case
		if n.left.bf <= 0 {
			// left left case
			n.rightRotate()

		} else {

			n.leftRotate()
			n.left.rightRotate()
		}

	} else if n.bf == +2 {

	}
}

func (n *AVLNode) update() {
	leftHeight := -1
	rightHeight := -1
	if n.left != nil {
		leftHeight = n.left.height
	}
	if n.right != nil {
		rightHeight = n.right.height
	}
	n.height = 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))
	n.bf = rightHeight - leftHeight
}

func contains(node *AVLNode, value int) bool {
	if node == nil {
		return false
	}
	if value < node.value {
		return contains(node.left, value)
	}
	if value > node.value {
		return contains(node.right, value)
	}
	return true
}

type AVLTree struct {
	root      *AVLNode
	nodeCount int
}

func (t *AVLTree) insert(value int) {

}

func main() {}
