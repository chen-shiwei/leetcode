package _13_找树左下角的值

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

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		maxDepth = -1
		val      int
	)

	var find func(node *TreeNode, depth int)
	find = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		find(node.Left, depth+1)
		if depth > maxDepth {
			fmt.Println(depth, root.Val)
			maxDepth = depth
			val = node.Val
		}
		find(node.Right, depth+1)

	}

	find(root, 0)

	return val
}
