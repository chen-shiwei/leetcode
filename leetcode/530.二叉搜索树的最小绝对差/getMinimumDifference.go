package _30_二叉搜索树的最小绝对差

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

func getMinimumDifference(root *TreeNode) int {
	var (
		result = math.MaxInt32
		pre    *TreeNode
	)

	var minDifference func(node *TreeNode)
	minDifference = func(node *TreeNode) {
		if node == nil {
			return
		}
		minDifference(node.Left)
		if pre != nil {
			result = min(result, node.Val-pre.Val)
		}
		pre = node
		minDifference(node.Right)
	}

	minDifference(root)

	return result
}

func getMinimumDifferenceWithStack(root *TreeNode) int {
	var (
		result = math.MaxInt32
		stack  []*TreeNode
	)

	if root == nil {
		return 0
	}

	var (
		pre *TreeNode
		cur *TreeNode = root
	)

	for len(stack) != 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != nil {
				result = min(result, cur.Val-pre.Val)
			}
			pre = cur
			cur = cur.Right
		}

	}

	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
