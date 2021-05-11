package _5_K个一组翻转链表

import "fmt"

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

	var left, right = pre, pre
	for left != nil {
		for i := 0; i < k; i++ {
			right = right.Next
			if right == nil {
				return hair.Next
			}
		}
		next := right.Next
		left = left.Next
		fmt.Println(left, right)
		newLeft, newRight := reverseLinkedList(left, right)
		// 修改反转的指针
		left.Next = newLeft
		newRight.Next = next
		// 移动下一位处理
		left = newRight
		right = newRight.Next
	}

	return hair.Next
}

func reverseLinkedList(head, tail *ListNode) (*ListNode, *ListNode) {
	var pre, cur = tail.Next, head
	for pre != tail {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return tail, head
}
