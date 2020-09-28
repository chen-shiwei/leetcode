package offer_22

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {

	var (
		i     = 0
		node  = head
		jNode = head
	)

	for node.Next != nil {
		i++
		if i >= k {
			jNode = jNode.Next
		}
		node = node.Next
	}

	return jNode
}
