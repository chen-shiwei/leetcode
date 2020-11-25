package IsPail

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类 the head
 * @return bool布尔型
 */
func isPail(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}
	var (
		slow, fast  = head, head.Next
		pre, prePre *ListNode
	)
	// 12 21
	for fast != nil && fast.Next != nil {
		pre = slow

		slow = slow.Next
		fast = fast.Next.Next

		pre.Next = prePre
		prePre = pre

	}
	// 21 后
	p2 := slow.Next
	// 21 前翻转
	slow.Next = pre

	var p1 *ListNode
	if fast == nil {
		p1 = slow.Next
	} else {
		p1 = slow
	}

	for p1 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next

	}

	return true

}
