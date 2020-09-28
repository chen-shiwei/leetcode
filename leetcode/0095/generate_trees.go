package _095

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 || n < 0 {
		return nil
	}

	return getTree(1, n)
}

func getTree(start, end int) []*TreeNode {
	var s []*TreeNode

	if start > end {
		s = append(s, nil)
		return s
	}

	if start == end {
		s = append(s, &TreeNode{
			Val:   start,
			Left:  nil,
			Right: nil,
		})
		return s
	}

	for i := start; i <= end; i++ {
		leftTree := getTree(start, i-1)
		rightTree := getTree(i+1, end)
		for _, lv := range leftTree {
			for _, jv := range rightTree {
				root := &TreeNode{Val: i}
				root.Left = lv
				root.Right = jv
				s = append(s, root)
			}
		}

	}

	return s
}
