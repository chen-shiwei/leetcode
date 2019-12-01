package datastruct

import "math"

// MinStack 时间复杂度o(1) 空间复杂度o(1)
type MinStack struct {
	data *stack
	min  int
}

func NewMinStack() MinStack {
	return MinStack{data: NewStack(), min: math.MinInt32}
}

// 2 0 2
// 1 -1 1
// 3 2 1
// 0 -1 0

func (m *MinStack) Push(v int) {
	if m.data.Empty() {
		m.min = v
	}
	m.data.Push(v - m.min)
	if v < m.min {
		m.min = v

	}
}

func (m *MinStack) Top() int {
	if m.data.Peek().(int) < 0 {
		return m.min + m.data.Peek().(int)
	} else {
		return m.min + m.data.Peek().(int)
	}
}

func (m *MinStack) GetMin() int {
	return m.min
}

func (m *MinStack) Pop() {
	top := m.data.Peek().(int)
	if top < 0 {
		m.min = m.min - top
	}
	m.data.Pop()

}
