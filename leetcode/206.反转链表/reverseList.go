package _06_反转链表

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

func reverseListWithStack(head *ListNode) *ListNode {
	var s = []int{}

	for cur := head; cur != nil; cur = cur.Next {
		s = append([]int{cur.Val}, s...)
	}
	fmt.Println(s)
	var h = &ListNode{}
	n := h
	for _, v := range s {
		n.Val = v
		n.Next = &ListNode{}
		n = n.Next
	}

	for cur := h; cur != nil; cur = cur.Next {
		fmt.Println(cur.Val)
	}

	return nil
}

func reverseListWithEach(head *ListNode) *ListNode {
	var (
		pre *ListNode
		cur = head
	)

	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}

// 递归
func reverseListWithRecursion(head *ListNode) *ListNode {

	var reverse func(pre, cur *ListNode) *ListNode

	reverse = func(pre, cur *ListNode) *ListNode {
		if cur == nil {
			return pre
		}
		tmp := cur.Next
		cur.Next = pre
		// 相比 each 缺少
		// pre = cur
		// cur = tmp
		return reverse(cur, tmp)
	}

	return reverse(nil, head)

}
