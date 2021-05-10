package _89_N叉树的前序遍历

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

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	var (
		stack  []*Node
		result []int
	)
	stack = append(stack, root)

	for len(stack) != 0 {
		l := len(stack)
		for i := 0; i < l; i++ {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			childLen := len(node.Children)
			for j := childLen - 1; j >= 0; j-- {
				stack = append(stack, node.Children[j])
			}
		}
	}

	return result
}
