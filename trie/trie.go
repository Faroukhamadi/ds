package trie

const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{},
	}
}

func (t *Trie) Insert(word string) {
	n := len(word)
	currentNode := t.root
	for i := 0; i < n; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(word string) bool {
	lastNode := t.searchWord(word)
	if lastNode == nil {
		return false
	}
	return lastNode.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.searchWord(prefix) != nil
}

func (t *Trie) searchWord(word string) *Node {
	n := len(word)
	currentNode := t.root
	for i := 0; i < n; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return nil
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode
}
