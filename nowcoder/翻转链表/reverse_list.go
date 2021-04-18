package 翻转链表

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList(pHead *ListNode) *ListNode {
	// write code here

	var (
		cur     = pHead
		newList = &ListNode{}
	)

	for cur.Next != nil {
		tmp := cur.Next
		cur.Next = newList
		newList = cur
		cur = tmp
	}

	return cur
}

func ReverseList2(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil || pHead.Next == nil {
		return pHead
	}
	newList := ReverseList2(pHead.Next)
	t1 := pHead.Next
	t1.Next = pHead
	pHead.Next = nil
	return newList
}
