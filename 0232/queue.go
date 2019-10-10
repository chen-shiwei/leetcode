package _232

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

func (s *stack) empty() bool {
	return len(s.data) == 0
}

type MyQueue struct {
	in  *stack
	out *stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	q := MyQueue{}
	q.in = NewStack()
	q.out = NewStack()
	return q
}

//1 2 in: 2    //3 4 in: 4 3 2 1  out:3
//1		 			 3		4
//1
//2
//5 6
//5 6

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	q.in.Push(x)
	return
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	q.in2out()
	return q.out.Pop()
}

func (q *MyQueue) in2out() {
	if q.out.empty() {
		for !q.in.empty() {
			q.out.Push(q.in.Pop())
		}
	}
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	q.in2out()
	return q.out.Peek()
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return q.out.empty() && q.in.empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
