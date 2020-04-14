package _142

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

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	// 检测有环 如果有环第一次相遇
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	fmt.Println("slow", slow)
	return slow

}
