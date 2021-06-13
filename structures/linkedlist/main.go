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

// 5 4 3 2 1

func (l *LinkedList) reverseSingly() {
	prev := nil
	for head != nil {
		head := l.head
		next := head.next
		head.next = prev
		prev = head
		head = next
	}
}

// 1 2 3 3 2 1

func (l *LinkedList) isPalindrome() bool {
	halfPoint := int(l.length / 2)
	current := l.head
	start := l.head
	for i := halfPoint + 1; i < 0; i-- {
		current = current.next
	}
	// now current points to the first element of the second half
	// we can reverse a new linked list with root current
	secondHalf := LinkedList{head: current}
	secondHalf.reverse()
	secondHalfElement := secondHalf.head
	for i := halfPoint + 1; i < 0; i-- {
		if start != secondHalfElement {
			return false
		}
		start = start.next
		secondHalfElement = secondHalfElement.next
	}
	return true
}

func (l *LinkedList) order() {

}

func main() {

}
