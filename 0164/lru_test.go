package _164

import (
	"fmt"
	"testing"
)

func TestLRUCache_Put(t *testing.T) {
	lru := Constructor(2)

	t.Run("//cache.put(1, 1);", func(t *testing.T) {
		lru.Put(1, 1)
	})
	t.Run("//cache.put(2, 2);", func(t *testing.T) {
		lru.Put(2, 2)
	})
	t.Run("//cache.get(1);       // 返回  1", func(t *testing.T) {
		if lru.Get(1) != 1 {
			t.Fail()
		}
	})
	t.Run("//cache.put(3, 3);    // 该操作会使得密钥 2 作废", func(t *testing.T) {
		lru.Put(3, 3)
	})
	t.Run("//cache.get(2);       // 返回 -1 (未找到)", func(t *testing.T) {
		if lru.Get(2) != -1 {
			t.Fail()
		}
	})
	t.Run("//cache.put(4, 4);    // 该操作会使得密钥 1 作废", func(t *testing.T) {
		lru.Put(4, 4)
	})
	t.Run("//cache.get(1);       // 返回 -1 (未找到)", func(t *testing.T) {
		if lru.Get(1) != -1 {
			t.Fail()
		}
	})
	t.Run("//cache.get(3);       // 返回  3", func(t *testing.T) {
		if lru.Get(3) != 3 {
			t.Fail()
		}
	})
	t.Run("//cache.get(4);       // 返回  4", func(t *testing.T) {
		if lru.Get(4) != 4 {
			t.Fail()
		}
	})

}

func TestConstructor(t *testing.T) {
	var words = []string{"a", "b"}

	fmt.Println(words[0] > words[1])
}
