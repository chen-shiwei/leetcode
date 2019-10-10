package _164

//LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
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
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/lru-cache
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	val  int
	Next *ListNode
}

type LRUCache struct {
	current int
	count   int
	head    struct {
		Next *ListNode
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{}
	l.count = capacity
	l.current = 0
	l.head = struct{ Next *ListNode }{Next: &ListNode{
		val:  0,
		Next: nil,
	}}
	NewListNode(l.head.Next, capacity, 0)
	return l
}

func (this *LRUCache) Get(key int) int {
	node := this.head.Next
	if key > this.current {
		return -1
	}

	var val = -1
	for key > 0 && node != nil {
		val = node.val
		node = node.Next
		key--
	}

	return val
}

func (this *LRUCache) Put(key int, value int) {
	if this.current == this.count {
		cur := this.head.Next
		i := this.current
		for i > 1 {
			cur = cur.Next
			i--
		}
		cur.Next = nil
	}
	next := this.head.Next

	this.head.Next.val = value
	this.head.Next.Next = next
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func NewListNode(node *ListNode, c int, value int) {
	if c < 0 {
		return
	}
	node.val = value
	node.Next = &ListNode{}
	c--
	NewListNode(node.Next, c, value)
}
