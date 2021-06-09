package main

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) insert(k int) {
	if k < n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.insert(k)
		}

	} else if k > n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.insert(k)
		}
	}
}

func (n *Node) search(el int) bool {
	if n.Key == el {
		return true
	}
	if el > n.Key {
		r := n.Right
		if r != nil {
			return n.search(r.Key)
		}
	} else {
		l := n.Left
		if l != nil {
			return n.search(l.Key)
		}
	}
	return false
}

func main() {

}
