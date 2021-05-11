package _57_二叉树的所有路径

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

var result []string

func binaryTreePaths(root *TreeNode) []string {
	dfs(root, "")
	return result
}

func dfs(node *TreeNode, paths string) {
	if node == nil {
		return
	}
	paths += fmt.Sprintf("%d", node.Val)
	if node.Right == nil && node.Left == nil {
		result = append(result, paths)
	}
	paths += "->"
	dfs(node.Left, paths)
	dfs(node.Right, paths)
}
