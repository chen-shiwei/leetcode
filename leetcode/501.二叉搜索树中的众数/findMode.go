package _01_二叉搜索树中的众数

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

func findMode(root *TreeNode) []int {

	var (
		preorder func(root *TreeNode)
		pre      *TreeNode
		count    int
		maxCount int
		result   []int
	)
	preorder = func(cur *TreeNode) {
		if cur == nil {
			return
		}

		preorder(cur.Left)

		if pre != nil {
			if cur.Val == pre.Val {
				count++
			} else {
				count = 1
			}
		} else {
			count = 1
		}
		if count == maxCount {
			result = append(result, cur.Val)
		} else if count > maxCount {
			maxCount = count
			result = result[:0]
			result = append(result, cur.Val)
		}
		pre = cur
		preorder(cur.Right)
	}

	preorder(root)

	fmt.Println(result)
	return result
}
