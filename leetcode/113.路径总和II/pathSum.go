package _13_路径总和II

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

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int

	var each func(node *TreeNode, targetSum int, nums []int)

	each = func(node *TreeNode, targetSum int, nums []int) {
		if node == nil {
			return
		}

		newSums := make([]int, len(nums), cap(nums))
		copy(newSums, nums)

		if IsLeafNode(node) && targetSum == node.Val {
			result = append(result, append(newSums, node.Val))
			return
		}

		each(node.Left, targetSum-node.Val, append(newSums, node.Val))
		each(node.Right, targetSum-node.Val, append(newSums, node.Val))
	}
	each(root, targetSum, []int{})

	return result
}

func IsLeafNode(node *TreeNode) bool {
	if node == nil {
		return false
	}
	return node.Left == nil && node.Right == nil
}
