package _124

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var val = math.MinInt32
	maxSumDiff(root, &val)
	return val

}

func maxSumDiff(node *TreeNode, val *int) int {
	if node == nil {
		return 0
	}
	left := maxSumDiff(node.Left, val)
	right := maxSumDiff(node.Right, val)
	lmr := node.Val + max(0, left) + max(0, right) //不考虑更节点
	ret := node.Val + max(0, max(left, right))
	*val = max(*val, max(lmr, ret))
	return ret
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}
