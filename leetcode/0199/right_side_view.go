package _0199

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

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var q []*TreeNode
	q = append(q, root)

	var nums []int
	for len(q) != 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			n := q[0]
			q = q[1:]

			if n.Left != nil {
				q = append(q, n.Left)
			}

			if n.Right != nil {
				q = append(q, n.Right)
			}

			if i == l-1 {
				nums = append(nums, n.Val)
			}
		}
	}

	return nums
}

type Queue struct {
	data []interface{}
}

func NewQueue() *Queue {
	return new(Queue)
}

func (s *Queue) Pop() interface{} {

	if len(s.data) < 1 {
		return 0
	}
	popVal := s.data[0]
	s.data = s.data[1:]
	return popVal
}

func (s *Queue) Peek() interface{} {
	if len(s.data) < 1 {
		return 0
	}
	popVal := s.data[0]
	return popVal
}

func (s *Queue) Push(x interface{}) {
	s.data = append(s.data, x)
	return
}

func (s *Queue) Empty() bool {
	return len(s.data) == 0
}

func (s Queue) Len() int {
	return len(s.data)
}

func rightSideViewDfs(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var nums []int
	dfs(root, 0, &nums)
	return nums
}

func dfs(root *TreeNode, depth int, nums *[]int) {
	if root == nil {
		return

	}
	if len(*nums) == depth {
		*nums = append(*nums, root.Val)
	}

	dfs(root.Right, depth+1, nums)
	dfs(root.Left, depth+1, nums)
}
