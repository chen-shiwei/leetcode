package _83_二叉搜索树节点最小距离

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var pre = -1
	var min = 1<<63 - 1
	var q []*TreeNode

	for root != nil || len(q) != 0 {
		for root != nil {
			q = append(q, root)
			root = root.Left
		}
		root = q[len(q)-1]
		q = q[:len(q)-1]
		if pre != -1 && root.Val-pre < min {
			min = root.Val - pre
		}
		pre = root.Val
		root = root.Right
	}

	return min
}

func minDiffInBST1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var pre = -1
	var min = 1<<63 - 1

	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		fmt.Println(root.Val)
		if pre != -1 && root.Val-pre < min {
			min = root.Val - pre
		}
		pre = root.Val
		dfs(root.Right)
	}

	dfs(root)
	return min
}
