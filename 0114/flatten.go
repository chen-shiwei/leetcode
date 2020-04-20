package _114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {

	if root == nil {
		return
	}

	stack := NewStack()
	stack.Push(root)
	var pre *TreeNode
	for !stack.Empty() {
		node := stack.Pop().(*TreeNode)
		if pre != nil {
			pre.Right = node
			pre.Left = nil
		}
		// 栈 先进后出 先打印左节点，所以右节点先进去
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}

		pre = node
	}

	return
}

type stack struct {
	data []interface{}
}

func NewStack() *stack {
	return new(stack)
}

func (s *stack) Pop() interface{} {

	if len(s.data) < 1 {
		return 0
	}
	popVal := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return popVal
}

func (s *stack) Peek() interface{} {
	if len(s.data) < 1 {
		return nil
	}
	popVal := s.data[len(s.data)-1]
	return popVal
}

func (s *stack) Push(x interface{}) {
	s.data = append(s.data, x)
	return
}

func (s *stack) Empty() bool {
	return len(s.data) == 0
}
func (s stack) Len() int {
	return len(s.data)
}
