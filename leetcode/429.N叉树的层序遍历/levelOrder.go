package _29_N叉树的层序遍历

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var (
		result [][]int
		q      []*Node
	)

	q = append(q, root)
	for len(q) != 0 {
		l := len(q)
		var nums []int
		for i := 0; i < l; i++ {
			node := q[0]
			q = q[1:]
			nums = append(nums, node.Val)
			childLen := len(node.Children)
			for j := 0; j < childLen; j++ {
				q = append(q, node.Children[j])
			}
		}
		result = append(result, nums)
	}
	return result
}
