package _098

import (
	"math"
)

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

func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var (
		isBST    = true
		checkBST func(node *TreeNode)
		flag     = math.MinInt32 - 1
	)
	checkBST = func(node *TreeNode) {
		if node == nil {
			return
		}

		checkBST(node.Left)

		if flag >= node.Val {
			isBST = false
			return
		}
		flag = node.Val

		checkBST(node.Right)

		return

	}

	checkBST(root)

	return isBST
}
