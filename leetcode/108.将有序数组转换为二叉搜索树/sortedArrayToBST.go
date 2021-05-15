package _08_将有序数组转换为二叉搜索树

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

func sortedArrayToBST(nums []int) *TreeNode {

	var (
		mid int
		l   = len(nums)
	)
	if l == 0 {
		return nil
	}

	mid = l / 2

	root := &TreeNode{
		Val:   nums[mid],
		Left:  nil,
		Right: nil,
	}
	if l == 1 {
		return root
	}

	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}
