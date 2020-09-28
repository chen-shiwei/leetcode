package _146

import (
	"testing"
)

func TestConstructor(t *testing.T) {

	//cache = LRUCache(2)
	//
	//cache.put(1, 1);
	//cache.put(2, 2);
	//cache.get(1);       // 返回  1
	//cache.put(3, 3);    // 该操作会使得密钥 2 作废
	//cache.get(2);       // 返回 -1 (未找到)
	//cache.put(4, 4);    // 该操作会使得密钥 1 作废
	//cache.get(1);       // 返回 -1 (未找到)
	//cache.get(3);       // 返回  3
	//cache.get(4);       // 返回  4

	cache := Constructor(2)
	t.Run("cache.put(1,1)", func(t *testing.T) {
		cache.Put(1, 1)
	})
	t.Run("cache.put(2,2);", func(t *testing.T) {
		cache.Put(2, 2)
	})
	t.Run("cache.get(1);", func(t *testing.T) {
		if cache.Get(1) != 1 {
			t.FailNow()
		}
	})
	t.Run("cache.put(3,3)", func(t *testing.T) {
		cache.Put(3, 3)
	})

	t.Run("cache.get(2)", func(t *testing.T) {
		if v := cache.Get(2); v != -1 {
			t.Errorf("get %v", v)
		}
	})
	t.Run("cache.put(4,4)", func(t *testing.T) {
		cache.Put(4, 4)
	})
	t.Run("cache.get(1)", func(t *testing.T) {
		if cache.Get(1) != -1 {
			t.FailNow()
		}
	})
	t.Run("cache.get(3)", func(t *testing.T) {
		if cache.Get(3) != 3 {
			t.FailNow()
		}
	})

	t.Run("cache.get(4)", func(t *testing.T) {
		if cache.Get(4) != 4 {
			t.FailNow()
		}
	})

}
