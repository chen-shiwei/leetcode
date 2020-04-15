package datastruct

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return new(Queue)
}

func (s *Queue) Pop() int {

	if len(s.data) < 1 {
		return 0
	}
	popVal := s.data[0]
	s.data = s.data[1:]
	return popVal
}

func (s *Queue) Peek() int {
	if len(s.data) < 1 {
		return 0
	}
	popVal := s.data[0]
	return popVal
}

func (s *Queue) Push(x int) {
	s.data = append(s.data, x)
	return
}

func (s *Queue) Empty() bool {
	return len(s.data) == 0
}

func (s Queue) Len() int {
	return len(s.data)
}
