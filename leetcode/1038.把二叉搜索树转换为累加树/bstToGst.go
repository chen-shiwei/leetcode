package _038_把二叉搜索树转换为累加树

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

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var (
		fn  func(node *TreeNode)
		pre int
	)

	fn = func(node *TreeNode) {
		if node == nil {
			return
		}
		fn(node.Right)
		node.Val += pre
		pre = node.Val
		fn(node.Left)
	}

	fn(root)
	return root
}
