package _148

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {

	quickSort(head, nil)
	return head
}

func quickSort(head, end *ListNode) {
	if head == end || head.Next == end {
		return
	}
	mid := head.Val
	slow, fast := head, head.Next
	for fast != end {
		if fast.Val <= mid {
			slow = slow.Next
			slow.Val, fast.Val = fast.Val, slow.Val
		}
		fast = fast.Next
	}

	slow.Val, head.Val = head.Val, slow.Val
	quickSort(head, slow)
	quickSort(slow.Next, end)
}
