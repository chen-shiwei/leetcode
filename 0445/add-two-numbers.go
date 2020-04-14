package _445

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	temp1 := l1
	temp2 := l2

	stack1 := NewStack()
	stack2 := NewStack()

	var len1 int
	var len2 int
	for temp1 != nil {
		len1++
		stack1.Push(temp1.Val)
		temp1 = temp1.Next

	}

	for temp2 != nil {
		len2++
		stack2.Push(temp2.Val)
		temp2 = temp2.Next
	}

	length := len1
	if len2 > len1 {
		length = len2
	}

	var newHead = &ListNode{
		Val:  0,
		Next: nil,
	}

	var carry int

	for i := 0; i < length; i++ {
		var n int
		if !stack1.Empty() {
			n += stack1.Pop()
		}
		if !stack2.Empty() {
			n += stack2.Pop()
		}
		n = n + carry
		if n >= 10 {
			carry = 1
			n = n - 10
		} else {
			carry = 0
		}
		headInsert(newHead, n)
	}
	if carry > 0 {
		headInsert(newHead, carry)
	}

	return newHead.Next

}

func headInsert(head *ListNode, n int) {
	newNode := &ListNode{
		Val:  n,
		Next: nil,
	}
	newNode.Next = head.Next
	head.Next = newNode
}

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
