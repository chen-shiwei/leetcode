package _102

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

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var nodes [][]int

	apply(root, &nodes, 0)
	return nodes
}

func apply(node *TreeNode, nodes *[][]int, dep int) {
	if node == nil {
		return
	}

	if len(*nodes) < dep+1 {
		*nodes = append(*nodes, []int{})
	}

	(*nodes)[dep] = append((*nodes)[dep], node.Val)

	apply(node.Left, nodes, dep+1)
	apply(node.Right, nodes, dep+1)
	return
}
