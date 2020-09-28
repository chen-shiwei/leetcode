package datastruct

type TrieNode struct {
	Links map[string]*TrieNode
	isEnd bool
}

func NewTrieNode() *TrieNode {
	t := TrieNode{
		Links: make(map[string]*TrieNode),
	}
	return &t
}

func (t *TrieNode) HasExist(s string) bool {
	_, ok := t.Links[s]
	return ok
}

func (t *TrieNode) Put(s string, node *TrieNode) {
	t.Links[s] = node
}

func (t *TrieNode) Get(s string) *TrieNode {
	return t.Links[s]
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
	t.root = NewTrieNode()
	return t
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	node := t.root
	for _, r := range word {
		if !node.HasExist(string(r)) {
			node.Put(string(r), NewTrieNode())
		}
		node = node.Get(string(r))
	}
	node.SetEnd()
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, r := range word {
		if node != nil {
			node = node.Get(string(r))
		}
	}
	return node != nil && node.IsEnd()
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, r := range prefix {
		if node != nil {
			node = node.Get(string(r))
		}
	}
	return node != nil
}
