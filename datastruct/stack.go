package datastruct

type stack struct {
	data []interface{}
}

func NewStack() *stack {
	return new(stack)
}

func (s *stack) Pop() interface{} {

	if len(s.data) < 1 {
		return ""
	}
	popVal := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return popVal
}

func (s *stack) Peek() interface{} {
	if len(s.data) < 1 {
		return ""
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
