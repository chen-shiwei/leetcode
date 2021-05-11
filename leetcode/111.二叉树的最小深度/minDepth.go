package _11_二叉树的最小深度

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

func minDepth(root *TreeNode) int {
	return getDepth(root)
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := getDepth(root.Left)
	right := getDepth(root.Right)

	if root.Left == nil && root.Right != nil {
		return right + 1
	}
	if root.Left != nil && root.Right == nil {
		return left + 1
	}

	if root.Right == nil && root.Left == nil {
		return 1
	}

	return min(left, right) + 1

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
