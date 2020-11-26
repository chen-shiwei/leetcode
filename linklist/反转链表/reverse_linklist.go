package 反转链表

type Node struct {
	Val  int
	Next *Node
}

func reverseLinklist(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	newLinkList := reverseLinklist(head.Next)
	t := head.Next
	t.Next = head
	head.Next = nil

	return newLinkList
}
