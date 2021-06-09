package main

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (l *LinkedList) prepend(value int) {
	node := &Node{value: value}
	newNext := l.head
	node.next = newNext
	newNext.prev = node
	l.head = node
}

func (l *LinkedList) push(value int) {

}

func swap(node1, node2 *Node) {
	node2.next = node1
	node1.next = node2
	node2.prev = node1.prev
	node1.prev = node2
}

func (l *LinkedList) reverseDoubly() {
	head, tail := l.head, l.tail
	for head != tail {
		swap(head, tail)
		head, tail = head.next, tail.prev
	}
}

//1 - 2 - 3 - 4 - 5 - nil

func (l *LinkedList) reverseSingly() {

}

func (l *LinkedList) order() {

}

func main() {

}
