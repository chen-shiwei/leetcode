package _146

type Node struct {
	Next, Prev *Node
	Key, Val   int
}

type DoubleList struct {
	Size       int
	Head, Tail *Node
}

func NewDoubleList() *DoubleList {
	d := &DoubleList{}
	d.Head = &Node{}
	d.Tail = &Node{}
	d.Head.Next = d.Tail
	d.Tail.Prev = d.Head
	return d
}

func (d *DoubleList) AddFirst(x *Node) {
	x.Next = d.Head.Next
	x.Prev = d.Head
	d.Head.Next.Prev = x
	d.Head.Next = x
	d.Size++
	return
}

func (d *DoubleList) Remove(x *Node) {
	x.Prev.Next = x.Next
	x.Next.Prev = x.Prev
	d.Size--
	return

}

func (d *DoubleList) RemoveLast() *Node {
	if d.Tail == d.Head {
		return nil
	}
	last := d.Tail.Prev
	d.Remove(last)
	return last
}
func (d *DoubleList) size() int {
	return d.Size
}

type LRUCache struct {
	Cap   int
	Cache *DoubleList
	Map   map[int]*Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{}
	lru.Cap = capacity
	lru.Map = make(map[int]*Node, 0)
	lru.Cache = NewDoubleList()
	return lru
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.Map[key]
	if !ok {
		return -1
	}
	this.Put(key, v.Val)
	return v.Val
}

func (this *LRUCache) Put(key int, value int) {
	x := &Node{}
	x.Key = key
	x.Val = value
	v, ok := this.Map[key]
	if ok {
		this.Cache.Remove(v)
		this.Cache.AddFirst(x)
		this.Map[key] = x
		return
	}
	if this.Cap == this.Cache.Size {
		last := this.Cache.RemoveLast()
		delete(this.Map, last.Key)
	}
	this.Cache.AddFirst(x)
	this.Map[key] = x
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
