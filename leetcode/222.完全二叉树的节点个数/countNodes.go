package _22_完全二叉树的节点个数

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

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		q []*TreeNode
		c int
	)

	q = append(q, root)

	for len(q) != 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			node := q[0]
			q = q[1:]
			c++
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return c
}
