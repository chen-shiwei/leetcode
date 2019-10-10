package _101

import "fmt"

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

func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return diff(root.Left, root.Right)
}

func diff(left, right *TreeNode) bool {
	fmt.Println(left, right)
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return (left.Val == right.Val) && diff(left.Right, right.Left) && diff(left.Left, right.Right)
}
