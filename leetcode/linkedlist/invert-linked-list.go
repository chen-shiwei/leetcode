package linkedlist

func eachInvert(head *Node) *Node {
	pre := head.Next
	cur := head.Next.Next
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

func recursiveInvert(node *Node) *Node {
	if node.Next == nil {
		return node
	}

	newHead := recursiveInvert(node.Next)
	node.Next.Next = node
	node.Next = nil

	return newHead

}
