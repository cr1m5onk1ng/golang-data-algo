package main

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

func (b *bucket) Insert(k string) {
	newNode := &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
}

func (b *bucket) Search(k string) bool {
	node := b.head
	for node != nil {
		if node.key == k {
			return true
		}
		node = node.next
	}
	return false
}

func (b *bucket) Delete(key string) {
	previous := b.head
	if previous.key == key {
		b.head = previous.next
		return
	}
	for previous.next != nil {
		if previous.next.key == key {
			previous.next = previous.next.next
		}
		previous = previous.next
	}
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func (h *HashTable) Insert(key string) {
	hashed := hashFunction(key)
	bucket := h.array[hashed]
	bucket.Insert(key)
}

func (h *HashTable) Search(key string) bool {
	hashed := hashFunction(key)
	bucket := h.array[hashed]
	return bucket.Search(key)
}

func (h *HashTable) Delete(key string) {
	hashed := hashFunction(key)
	bucket := h.array[hashed]
	bucket.Delete(key)
}

func hashFunction(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func main() {

}
