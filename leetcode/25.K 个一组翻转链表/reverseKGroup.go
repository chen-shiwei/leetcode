package _5_K_个一组翻转链表

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

func reverseKGroup(head *ListNode, k int) *ListNode {
	var n = head

	var stack []int

	var c = k
	for cur := head; cur != nil; cur = cur.Next {

		if c == 0 {
			var l = &ListNode{}
			for _, v := range stack {
				l.Val = v
				l.Next = &ListNode{}
			}
			stack = stack[:0]
			c = k
		} else {
			stack = append([]int{cur.Val}, stack...)
			c--
		}

	}

}
