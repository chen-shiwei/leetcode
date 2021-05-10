package _07_二叉树的层序遍历II

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var q []*TreeNode

	if root == nil {
		return nil
	}

	var result [][]int

	q = append(q, root)
	for len(q) != 0 {
		l := len(q)
		var nums []int
		for i := 0; i < l; i++ {
			node := q[0]
			q = q[1:]
			nums = append(nums, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		result = append(result, nums)
	}

	var left, right = 0, len(result) - 1
	if len(result) > 0 {
		for left < right {
			result[left], result[right] = result[right], result[left]
			left++
			right--
		}
	}

	return result
}
