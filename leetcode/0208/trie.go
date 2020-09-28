package _208

type TrieNode struct {
	Links [26]*TrieNode
	isEnd bool
}

func NewTrieNode(cap int) *TrieNode {
	t := TrieNode{}
	return &t
}

func (t *TrieNode) HasExist(s rune) bool {
	return t.Links[s-'a'] != nil
}

func (t *TrieNode) Put(s rune, node *TrieNode) {
	t.Links[s-'a'] = node
}

func (t *TrieNode) Get(s rune) *TrieNode {
	return t.Links[s-'a']
}

func (t *TrieNode) IsEnd() bool {
	return t.isEnd
}
func (t *TrieNode) SetEnd() {
	t.isEnd = true

}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	t := Trie{}
	t.root = NewTrieNode(26)
	return t
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	node := t.root
	for _, r := range word {
		if !node.HasExist(r) {
			node.Put(r, NewTrieNode(26))
		}
		node = node.Get(r)
	}
	node.SetEnd()
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, r := range word {
		if node != nil {
			node = node.Get(r)
		}
	}
	return node != nil && node.IsEnd()
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, r := range prefix {
		if node != nil {
			node = node.Get(r)
		}
	}
	return node != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
