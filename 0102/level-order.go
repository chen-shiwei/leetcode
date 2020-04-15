package _102

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

var result [][]int

func LevelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	apply(root, 0)
	return result
}

func apply(node *TreeNode, dep int) {
	if node == nil {
		return
	}

	if len(result) < dep+1 {
		result = append(result, []int{})
	}

	result[dep] = append(result[dep], node.Val)

	if node.Left != nil {
		apply(node.Left, dep+1)
	}
	if node.Right != nil {
		apply(node.Right, dep+1)
	}

	return

}

// 解法 2
func LevelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	q := &queue{}
	q.Push(root)

	var result [][]int
	for !q.Empty() {
		var nums []int
		l := q.Size()
		for i := 0; i < l; i++ {
			node := q.Pop()

			if node.Left != nil {
				q.Push(node.Left)
			}

			if node.Right != nil {
				q.Push(node.Right)
			}
			nums = append(nums, node.Val)
		}
		result = append(result, nums)

	}
	return result
}

type queue struct {
	data []*TreeNode
}

func (q *queue) Pop() *TreeNode {
	val := q.data[0]
	q.data = q.data[1:]
	return val
}

func (q *queue) Empty() bool {
	return len(q.data) == 0
}

func (q *queue) Push(v *TreeNode) {
	q.data = append(q.data, v)
}

func (q *queue) Size() int {
	return len(q.data)
}
