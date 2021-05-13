package _54_最大二叉树

import (
	"math"
)

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

func constructMaximumBinaryTree(nums []int) *TreeNode {

	l := len(nums)
	if l == 0 {
		return nil
	}
	var (
		max = math.MinInt32
		idx int
	)
	for i := 0; i < l; i++ {
		if nums[i] > max {
			max = nums[i]
			idx = i
		}
	}

	rootVal := max
	root := &TreeNode{
		Val:   rootVal,
		Left:  nil,
		Right: nil,
	}
	if len(nums) == 1 {
		return root
	}
	root.Left = constructMaximumBinaryTree(nums[:idx])
	root.Right = constructMaximumBinaryTree(nums[idx+1:])

	return root

}
