package _208

import (
	"testing"
)

func TestConstructor(t *testing.T) {

	//Trie trie = new Trie();
	//
	//trie.insert("apple");
	//trie.search("apple");   // 返回 true
	//trie.search("app");     // 返回 false
	//trie.startsWith("app"); // 返回 true
	//trie.insert("app");
	//trie.search("app");     // 返回 true
	//trie := Constructor()
	//trie.Insert("apple")
	//if !trie.Search("apple") { // 返回 true
	//	t.Fatal(`//trie.search("apple");   // 返回 true`)
	//}
	//if trie.Search("app") {
	//	t.Fatal(`//trie.search("app");     // 返回 false`)
	//} // 返回 false
	//if !trie.StartsWith("app") {
	//	t.Fatal(`//trie.startsWith("app"); // 返回 true`)
	//} // 返回 true
	//trie.Insert("app")
	//if !trie.Search("app") {
	//	t.Fatal(`//trie.search("app");     // 返回 true`)
	//} // 返回 true

	//
	//["Trie","insert","search","search","search","startsWith","startsWith","startsWith"]
	//[[],["hello"],["hell"],["helloa"],["hello"],["hell"],["helloa"],["hello"]]
	//[null,null,false,false,true,true,false,true]
	trie := Constructor()
	trie.Insert("hello")
	if trie.Search("hell") {
		t.Fatal(`search hell false`)
	}

	if trie.Search("helloa") {
		t.Fatal(`search helloa false`)
	}

	if !trie.Search("hello") {
		t.Fatal(`search hello true`)
	}

	if !trie.StartsWith("hell") {
		t.Fatal(`startsWith hell true`)
	}

	if trie.StartsWith("helloa") {
		t.Fatal(`startsWith hell false`)
	}

	if !trie.StartsWith("hello") {
		t.Fatal(`startsWith hello true`)
	}
}

func TestA(t *testing.T) {
	m := map[int]string{1: "a", 2: "b"}
	for key := range m {
		t.Log(key)
	}
	s := []string{"a", "b"}
	for key := range s {
		t.Log(key)
	}
	for key := range "abc" {
		t.Log(key)
	}

}
