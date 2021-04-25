package _019

import (
	"fmt"
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

// 时间复杂度：O(n) 空间复杂度：O(1)
func RemoveNthFromEnd1(head *ListNode, n int) *ListNode {
	var count int
	l := head
	for p := l; p != nil; p = p.Next {
		count++
	}

	if count < n {
		return head
	}

	if count == n {
		head = l.Next
		return head
	}

	var c int
	for p := l; p != nil; p = p.Next {
		c++
		if c == (count - n) {
			if p.Next == nil {
				p = nil
			} else {
				p.Next = p.Next.Next
			}
			break
		}
	}

	return head
}

func RemoveNthFromEnd2(head *ListNode, n int) *ListNode {

	// 防止删除头节点 做哨兵
	newListNode := new(ListNode)
	newListNode.Next = head

	var i, j = newListNode, newListNode

	var count int
	for i != nil {
		i = i.Next
		// 需要指向的是删除的上一个节点
		if count <= (n) {
			count++
		} else {
			j = j.Next
		}
	}

	j.Next = j.Next.Next

	return newListNode.Next
}

func printListNode(node *ListNode) {
	fmt.Println(node)
	if node != nil {
		printListNode(node.Next)
	}
}
