package main

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (n *Node) insert(k int) {
	if k < n.key {
		if n.right == nil {
			n.right = &Node{key: k}
		} else {
			n.right.insert(k)
		}

	} else if k > n.key {
		if n.left == nil {
			n.left = &Node{key: k}
		} else {
			n.left.insert(k)
		}
	}
}

func (n *Node) search(el int) bool {
	if n.key == el {
		return true
	}
	if el > n.key {
		r := n.right
		if r != nil {
			return n.search(r.key)
		}
	} else {
		l := n.left
		if l != nil {
			return n.search(l.key)
		}
	}
	return false
}

func (n *Node) inOrder() []int {
	for n != nil {

	}
}

func main() {

}
