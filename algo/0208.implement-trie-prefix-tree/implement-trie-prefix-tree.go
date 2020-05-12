package problem0208

// Trie
type Trie struct {
	val byte
	child [26]*Trie
	end int
}

// Constructor Initialize your data structure here
func Constructor() Trie  {
	return Trie{}
}

// Insert Inserts a word into the trie.
func (this *Trie) Insert(word string)  {
	node := this
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a'

		if node.child[idx] == nil {
			node.child[idx] = &Trie{val: word[i]}
		}

		node = node.child[idx]
	}
	node.end ++
}

// Search Returns if the word is in the trie.
func (this *Trie) Search(word string) bool  {
	node := this
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a'
		if node.child[idx] == nil {
			return false
		}
		node = node.child[idx]
	}

	if node.end > 0 {
		return true
	}

	return false
}

// StartsWith Returns if there is any word in the trie that starts with the given prefix.
func (this *Trie) StartsWith(prefix string) bool  {
	node := this

	size := len(prefix)

	for i := 0; i < size; i++ {
		idx := prefix[i] - 'a'
		if node.child[idx] == nil {
			return false
		}
		node = node.child[idx]
	}

	return true
}

