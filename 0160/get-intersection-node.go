package _160

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

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	dict := make(map[string]uint8, 0)

	var i int
	for headA != nil {
		i++
		dict[fmt.Sprintf("%p", headA)] = 0
		headA = headA.Next
	}

	for headB != nil {
		_, ok := dict[fmt.Sprintf("%p", headB)]
		if ok {
			return headB
		}
		headB = headB.Next
	}

	return nil
}
