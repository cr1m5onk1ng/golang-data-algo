package main

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) insert(s string) {
	element := t.root
	for i := 0; i < len(s); i++ {
		char := s[i]
		index := char - 'a'
		child := element.children[index]
		if child == nil {
			element.children[index] = &TrieNode{}
		}
		element = element.children[index]
	}
	element.isEnd = true
}

func (t *Trie) search(s string) bool {
	element := t.root
	for i := 0; i < len(s); i++ {
		char := s[i]
		index := char - 'a'
		child := element.children[index]
		if child == nil {
			return false
		}
		element = element.children[index]
	}
	if !element.isEnd {
		return false
	}
	return true
}

func main() {

}
