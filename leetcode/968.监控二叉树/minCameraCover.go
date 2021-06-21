package _68_监控二叉树

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

func minCameraCover(root *TreeNode) int {
	// 0 无覆盖
	// 有覆盖，有摄像头
	// 有覆盖，无摄像头
	var fn func(root *TreeNode) int
	var result int
	fn = func(root *TreeNode) int {
		if root == nil {
			return 2
		}
		left := fn(root.Left)
		right := fn(root.Right)

		if left == 2 && right == 2 {
			return 0
		}

		if left == 0 || right == 0 {
			result++
			return 1
		}

		if left == 1 || right == 1 {
			return 2
		}

		return -1
	}

	if fn(root) == 0 {
		result++
	}

	return result
}
