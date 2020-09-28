package _020

type stack struct {
	data []string
}

func NewStack() *stack {
	return new(stack)
}

func (s *stack) Pop() string {

	if len(s.data) < 1 {
		return ""
	}
	popVal := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return popVal
}

func (s *stack) Peek() string {
	if len(s.data) < 1 {
		return ""
	}
	popVal := s.data[len(s.data)-1]
	return popVal
}

func (s *stack) Push(x string) {
	s.data = append(s.data, x)
	return
}

func (s *stack) empty() bool {
	return len(s.data) == 0
}

func IsValid(s string) bool {
	if len(s) < 1 {
		return false
	}

	stack := NewStack()
	for _, word := range s {
		switch string(word) {
		case "{", "(", "[":
			stack.Push(string(word))
			break
		case "}":
			v := stack.Pop()
			if v != "{" {
				return false
			}
			break
		case ")":
			v := stack.Pop()
			if v != "(" {
				return false
			}
			break
		case "]":
			v := stack.Pop()
			if v != "[" {
				return false
			}
			break
		}
	}
	if stack.empty() {
		return true
	}
	return false
}
