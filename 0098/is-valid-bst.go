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
	flag = math.MinInt32 - 1
	check = true

	if root == nil {
		return true
	}
	checkBast(root)

	return check
}

var (
	flag  int  = math.MinInt32 - 1
	check bool = true
)

func checkBast(node *TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil {
		checkBast(node.Left)
	}
	if flag >= node.Val {
		check = false
		return
	}
	flag = node.Val
	if node.Right != nil {
		checkBast(node.Right)

	}
	return
}

//func checkBst(node, lower, upper *TreeNode) bool {
//	if node == nil {
//		return true
//	}
//
//	if lower != nil {
//		if node.Val <= lower.Val {
//			return false
//		}
//	}
//
//	if upper != nil {
//		if node.Val >= upper.Val {
//			return false
//		}
//	}
//
//	return checkBst(node.Left, lower, node) && checkBst(node.Right, node, upper)
//}

//	10
//5			 15
//		6			20
