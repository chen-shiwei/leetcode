package _337

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

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	money := root.Val

	if root.Left != nil {
		money += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		money += rob(root.Right.Left) + rob(root.Right.Right)
	}

	result := money
	if m := rob(root.Left) + rob(root.Right); m > result {
		return m
	}

	return result

}

func dfs(node *TreeNode, depth int, depths *[]int) {
	if node == nil {
		return
	}
	if len(*depths) <= depth {
		*depths = append(*depths, 0)
	}
	(*depths)[depth] += node.Val

	dfs(node.Left, depth+1, depths)
	dfs(node.Right, depth+1, depths)
}
