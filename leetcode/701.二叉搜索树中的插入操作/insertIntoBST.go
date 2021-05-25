package _01_二叉搜索树中的插入操作

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

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	var (
		fn func(node *TreeNode) *TreeNode
	)

	fn = func(node *TreeNode) *TreeNode {
		if node == nil {
			node = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			return node
		}

		if node.Val > val {
			node.Left = fn(node.Left)
		} else {
			node.Right = fn(node.Right)
		}

		return node
	}

	root = fn(root)

	return root

}
