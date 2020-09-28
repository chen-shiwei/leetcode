package _021

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	head := new(ListNode)
	temp := head

	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				temp.Next = l1
				l1 = l1.Next
			} else {
				temp.Next = l2
				l2 = l2.Next
			}
		} else if l1 == nil && l2 != nil {
			temp.Next = l2
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			temp.Next = l1
			l1 = l1.Next
		}
		temp = temp.Next
	}

	return head.Next
}
