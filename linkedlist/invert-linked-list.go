package linkedlist

import "fmt"

func recursiveInvert(node *Node) *Node {
	if node.Next == nil {
		return node
	}

	fmt.Println("step 1")
	newHead := recursiveInvert(node.Next)

	fmt.Println("step 2")
	node.Next.Next = node
	node.Next = nil

	fmt.Println("step 3")
	return newHead
}

func eachInvert(head *Node) *Node {
	pre := head.Next
	cur := pre.Next
	pre.Next = nil

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	head.Next = pre
	return head
}
