package datastruct

type stack struct {
	data []int
}

func NewStack() *stack {
	return new(stack)
}

func (s *stack) Pop() int {

	if len(s.data) < 1 {
		return 0
	}
	popVal := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return popVal
}

func (s *stack) Peek() int {
	if len(s.data) < 1 {
		return 0
	}
	popVal := s.data[len(s.data)-1]
	return popVal
}

func (s *stack) Push(x int) {
	s.data = append(s.data, x)
	return
}

func (s *stack) Empty() bool {
	return len(s.data) == 0
}
func (s stack) Len() int {
	return len(s.data)
}
