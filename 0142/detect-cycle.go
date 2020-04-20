package _142

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
	// slow 走一步，fast 走两步
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next
		// 第一次相遇 环
		if slow == fast {
			break
		}
	}

	// 假设head 到 回环起始点距离 为x
	// 回环起始点到相遇点为 为 y，slow 走的距离为 x+y；
	// 设 y相遇点到环起始点的距离为z， 因为fast要和slow相遇它走的距离是 2(x+y) = x+y+y+z => 2x+2y = 2y+x+z => 2x=x+z => x=z，
	// 假设head 到 回环起始点距离 = y相遇点到环起始点的距离
	// 让 slow 从 head 开始走 速度为1，fast 已速度为1开始走，同样的距离，它们将在环起始点相遇

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
