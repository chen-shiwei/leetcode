package _208_todo

type TrieNode struct {
	Value string
	Links []TrieNode
	Exist bool
}

func NewTrieNode(cap int) TrieNode {
	t := TrieNode{
		Links: make([]TrieNode, 0, cap),
		Exist: false,
	}
	return t
}

func (t *TrieNode) Put() {
	t.Exist = true
}

func (t *TrieNode) HasExist() bool {
	return t.Exist
}

type Trie struct {
}

/** Initialize your data structure here. */
func Constructor() Trie {

}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {

}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {

}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
