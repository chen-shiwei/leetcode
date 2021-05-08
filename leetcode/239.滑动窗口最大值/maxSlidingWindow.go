package _39_滑动窗口最大值

import (
	"fmt"
	"github.com/chen-shiwei/leetcode/leetcode/datastruct"
)

type MyQueue struct {
	dq *datastruct.Deque
}

func NewMyQueue() *MyQueue {
	return &MyQueue{dq: datastruct.NewDeque()}
}

func (m *MyQueue) Push(x int) {
	for !m.dq.Empty() && x > m.dq.Back() {
		m.dq.PopBack()
	}
	m.dq.PushBack(x)
}

func (m *MyQueue) Pop(x int) {
	if !m.dq.Empty() && x == m.dq.Front() {
		fmt.Println("pop", m.dq.PopFront())
	}
}

func (m *MyQueue) Front() int {
	return m.dq.Front()
}

func maxSlidingWindow(nums []int, k int) []int {
	l := len(nums)
	//var q []int

	qu := NewMyQueue()

	//push := func(x int) {
	//	if len(q) != 0 && x > q[len(q)-1] {
	//		q = q[:len(q)-1]
	//	}
	//	q = append(q, x)
	//}
	//
	//pop := func(x int) {
	//	if len(q) != 0 && q[0] == x {
	//		q = q[1:]
	//	}
	//}

	var result []int
	for i := 0; i < k; i++ {
		qu.Push(nums[i])
	}

	qu.dq.Println()
	//result = append(result, q[0])
	result = append(result, qu.Front())
	for i := k; i < l; i++ {
		qu.Pop(nums[i-k])
		qu.Push(nums[i])
		qu.dq.Println()
		result = append(result, qu.Front())
		//result = append(result, q[0])
	}
	return result
}

func maxSlidingWindow1(nums []int, k int) []int {
	l := len(nums)
	var q []int

	push := func(x int) {
		for len(q) != 0 && x > q[len(q)-1] {
			q = q[:len(q)-1]
		}
		q = append(q, x)
	}

	pop := func(x int) {
		if len(q) != 0 && q[0] == x {
			fmt.Println("pop", q[0])
			q = q[1:]
		}
	}

	var result []int
	for i := 0; i < k; i++ {
		push(nums[i])
	}

	result = append(result, q[0])
	for i := k; i < l; i++ {
		pop(nums[i-k])
		push(nums[i])
		result = append(result, q[0])
	}
	return result
}
