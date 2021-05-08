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
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}

		}
		tmp := tail.Next
		head, tail = reverseLinkedList(head, tail)
		pre.Next = head
		tail.Next = tmp
		pre = tail
		head = tail.Next
	}

	return hair.Next
}

func reverseLinkedList(head, tail *ListNode) (*ListNode, *ListNode) {
	var pre, cur = tail.Next, head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return tail, head
}
