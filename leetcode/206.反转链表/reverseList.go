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

func reverseList(head *ListNode) *ListNode {
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
