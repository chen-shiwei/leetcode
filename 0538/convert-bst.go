package _538

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

func ConvertBST(root *TreeNode) *TreeNode {
	var init = 0
	Sum(root, &init)
	return root
}

func Sum(root *TreeNode, preVal *int) {
	if root == nil {
		return
	}
	Sum(root.Right, preVal)
	root.Val = root.Val + *preVal
	*preVal = root.Val
	Sum(root.Left, preVal)
	return
}
