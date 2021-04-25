package _03_移除链表元素

import (
	"fmt"
	"time"
)

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

// sentinel
func removeElementsSentinel(head *ListNode, val int) *ListNode {
	sentinel := &ListNode{Next: head, Val: 0}
	cur := sentinel
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}

	return sentinel.Next
}

func removeElementsRecursion(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	fmt.Println("before", time.Now().UnixNano(), head.Val)
	// 返回的后面的节点，不是得话就不返回
	head.Next = removeElementsRecursion(head.Next, val)
	fmt.Println("after", time.Now().UnixNano(), head.Val)
	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}
