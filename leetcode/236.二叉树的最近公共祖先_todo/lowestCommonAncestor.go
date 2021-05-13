package _36_二叉树的最近公共祖先_todo

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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return right
	}

	return left
}

func lowestCommonAncestor01(root, p, q *TreeNode) *TreeNode {

	var (
		result *TreeNode
		fn     func(node *TreeNode)
	)

	fn = func(node *TreeNode) {
		fmt.Println(1, node)
		if node == nil {
			return
		}
		fmt.Println(2)

		//if node.Left != nil || node.Right != nil {
		fmt.Println(node.Left, node.Right)
		if (node.Left == p && node.Right == q) || (node.Left == q && node.Right == p) {
			result = node
			return
		}
		//}
		fn(node.Left)
		fn(node.Right)
	}
	fn(root)
	return result

}
