package _337

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

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := NewQueue()
	q.Push(root)
	var (
		top  = root.Val
		down = 0
	)

	for !q.Empty() {
		node := q.Pop().(*TreeNode)
		fmt.Println(node)
		if node.Left != nil {
			q.Push(node.Left)
		}
		if node.Right != nil {
			q.Push(node.Right)
		}
	}
	return 0

}

type Queue struct {
	data []interface{}
}

func NewQueue() *Queue {
	return new(Queue)
}

func (s *Queue) Pop() interface{} {

	if len(s.data) < 1 {
		return nil
	}
	popVal := s.data[0]
	s.data = s.data[1:]
	return popVal
}

func (s *Queue) Peek() interface{} {
	if len(s.data) < 1 {
		return nil
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
