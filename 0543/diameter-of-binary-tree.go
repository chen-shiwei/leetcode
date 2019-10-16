package _543

import "math"

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

func DiameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var max int
	maxDepth(root, &max)
	return max
}

func maxDepth(node *TreeNode, max *int) int {
	if node == nil {
		return 0
	}
	left := maxDepth(node.Left, max)
	right := maxDepth(node.Right, max)
	*max = int(math.Max(float64(left+right), float64(*max)))
	return int(math.Max(float64(left), float64(right))) + 1
}
