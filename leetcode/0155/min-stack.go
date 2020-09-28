package _155

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() {
	s.data = s.data[0 : len(s.data)-1]
}

func (s *Stack) Top() int {
	if s.Empty() {
		return 0
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

type MinStack struct {
	data *Stack
	help *Stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	m := MinStack{
		data: NewStack(),
		help: NewStack(),
	}
	return m
}

func (m *MinStack) Push(x int) {
	m.data.Push(x)

	if x <= m.help.Top() || m.help.Empty() {
		m.help.Push(x)
	}
}

func (m *MinStack) Pop() {
	if !m.help.Empty() && m.help.Top() == m.Top() {
		m.help.Pop()
	}
	m.data.Pop()
}

func (m *MinStack) Top() int {
	return m.data.Top()
}

func (m *MinStack) GetMin() int {
	return m.help.Top()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
