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

func ConvertBST1(root *TreeNode) *TreeNode {
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
