package _04_左叶子之和

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

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := sumOfLeftLeaves(root.Left)
	right := sumOfLeftLeaves(root.Right)
	// 是否是有效的左叶子节点
	var mid int
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		mid = root.Left.Val
	}

	return mid + left + right
}
