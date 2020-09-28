package FindKthToTail

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * @param pListHead ListNode类
 * @param k int整型
 * @return ListNode类
 */
func FindKthToTail(pListHead *ListNode, k int) *ListNode {
	// write code here

	var quick, slow *ListNode = pListHead, pListHead

	var i int
	for quick != nil {
		quick = quick.Next
		if i >= k {
			slow = slow.Next
			continue
		}
		i++
	}

	if i < k {
		return nil
	}

	return slow
}
