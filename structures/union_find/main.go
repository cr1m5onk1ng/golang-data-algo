package main

const N = 10

type UFElement struct {
	value  int
	parent *UFElement
	group  int
}

func (el *UFElement) findWithPathCompression() *UFElement {
	root := el
	for root.parent != nil {
		root = root.parent
	}

	element := el
	for element.parent != nil {
		oldParent := element.parent
		element.parent = root
		element = oldParent
	}
	return root
}

type UnionFind struct {
	elements [N]*UFElement
}

func (el *UFElement) find() *UFElement {
	root := el
	for root.parent != nil {
		root = root.parent
	}
	return root
}

func (uf *UnionFind) union(element1, element2 *UFElement) {
	root1, root2 := element1.findWithPathCompression(), element2.findWithPathCompression()
	if root1 == root2 {
		return
	}
	root2.parent = root1
}

func connected(element1, element2 *UFElement) bool {
	return element1.findWithPathCompression() == element2.findWithPathCompression()
}

func main() {

}
