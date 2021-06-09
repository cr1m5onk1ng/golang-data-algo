package main

type MaxHeap struct {
	array []int
}

func leftChild(index int) int {
	return index*2 + 1
}

func rightChild(index int) int {
	return index*2 + 2
}

func parent(index int) int {
	return (index - 1) / 2
}

func (h *MaxHeap) extract() int {
	largest := h.array[0]
	h.swap(0, len(h.array)-1)
	h.array = h.array[:len((h.array))-1]
	h.maxHeapifyDown(0)
	return largest
}

func (h *MaxHeap) insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) remove(index int) {

}

func (h *MaxHeap) maxHeapifyDown(index int) {
	left, right := leftChild(0), rightChild(0)
	lastIndex := len(h.array) - 1
	childToCompare := 0
	for left <= lastIndex {
		if left == lastIndex {
			childToCompare = left
		} else if h.array[left] > h.array[right] {
			childToCompare = left
		} else {
			childToCompare = right
		}
		if h.array[childToCompare] < h.array[index] {
			h.swap(childToCompare, index)
			index = childToCompare
			left, right = leftChild(index), rightChild(index)
		} else {
			return
		}
	}
}

func (h *MaxHeap) swap(index1, index2 int) {
	temp := h.array[index1]
	h.array[index1] = h.array[index2]
	h.array[index2] = temp
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	parentIndex := parent(index)
	for h.array[parentIndex] < h.array[index] {
		h.swap(parentIndex, index)
		index = parentIndex
	}
}

func main() {

}
