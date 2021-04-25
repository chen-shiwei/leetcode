package _07_设计链表

type ListNode struct {
	Next *ListNode
	Val  int
}

type MyLinkedList struct {
	head *ListNode
	size int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	sentinel := &ListNode{Val: 0}
	return MyLinkedList{head: sentinel, size: 0}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.head

	var c int
	for cur.Next != nil && c <= index {
		c++
		cur = cur.Next
	}

	return cur.Val

}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &ListNode{
		Val: val,
	}
	newNode.Next = this.head.Next
	this.head.Next = newNode
	this.size++
	return
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	newListNode := &ListNode{Val: val}
	cur := this.head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newListNode
	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	newListNode := &ListNode{Val: val}
	cur := this.head
	var c int
	for cur.Next != nil && c < index {
		cur = cur.Next
		c++
	}
	newListNode.Next = cur.Next
	cur.Next = newListNode
	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	cur := this.head
	var c int
	for cur.Next != nil && c < index {
		cur = cur.Next
		c++
	}
	cur.Next = cur.Next.Next
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
