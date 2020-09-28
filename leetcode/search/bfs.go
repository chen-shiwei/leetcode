package search

import (
	"fmt"
)

func bfs(root *Node) {
	if root == nil {
		return
	}
	q := NewQueue()
	q.Push(root)

	for !q.Empty() {
		n := q.Pop()
		fmt.Println(n.data)

		if n.LeftNode != nil {
			q.Push(n.LeftNode)
		}

		if n.RightNode != nil {
			q.Push(n.RightNode)
		}
	}
	return
}

type Queue struct {
	data []*Node
}

func NewQueue() *Queue {
	return new(Queue)
}

func (s *Queue) Pop() *Node {

	if len(s.data) < 1 {
		return nil
	}
	popVal := s.data[0]
	s.data = s.data[1:]
	return popVal
}

func (s *Queue) Peek() *Node {
	if len(s.data) < 1 {
		return nil
	}
	popVal := s.data[0]
	return popVal
}

func (s *Queue) Push(x *Node) {
	s.data = append(s.data, x)
	return
}

func (s *Queue) Empty() bool {
	return len(s.data) == 0
}

func (s Queue) Len() int {
	return len(s.data)
}
