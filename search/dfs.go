package search

import (
	"fmt"
)

type Node struct {
	data      int
	LeftNode  *Node
	MidNode   *Node
	RightNode *Node
}

func dfs(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.data)

	dfs(root.LeftNode)
	dfs(root.RightNode)
	return
}

func dfsWithStack(root *Node) {
	stack := NewStack()
	stack.Push(root)
	for !stack.Empty() {
		node := stack.Pop()
		fmt.Println(node.data)
		if node.RightNode != nil {
			stack.Push(node.RightNode)
		}
		if node.LeftNode != nil {
			stack.Push(node.LeftNode)
		}
	}
}

type stack struct {
	data []*Node
}

func NewStack() *stack {
	return new(stack)
}

func (s *stack) Pop() *Node {

	if len(s.data) < 1 {
		return nil
	}
	popVal := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return popVal
}

func (s *stack) Peek() *Node {
	if len(s.data) < 1 {
		return nil
	}
	popVal := s.data[len(s.data)-1]
	return popVal
}

func (s *stack) Push(x *Node) {
	s.data = append(s.data, x)
	return
}

func (s *stack) Empty() bool {
	return len(s.data) == 0
}
func (s stack) Len() int {
	return len(s.data)
}
