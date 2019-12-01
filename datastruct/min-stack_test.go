package datastruct

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	minStack := NewMinStack()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	if minStack.GetMin() != -3 {
		t.Fatal(`minStack.GetMin() != -3`)
	}
	minStack.Pop()
	if minStack.Top() != 0 {
		t.Fatal(`minStack.Top() != 0`)
	}
	if minStack.GetMin() != -2 {
		t.Fatal(`minStack.GetMin() != -2`)
	}
	//minStack := MinStack{data: NewStack(), min: math.MinInt32}
	//
	//minStack.Push(1)
	//fmt.Println(minStack.data, minStack.min)
	//
	//minStack.Push(2)
	//fmt.Println(minStack.data, minStack.min)
	//
	//if minStack.Top() != 2 {
	//	t.Fatal(`minStack.Top() != 2`)
	//}
	//if minStack.GetMin() != 1 {
	//	t.Fatal(`minStack.GetMin() != 1`)
	//}
	//minStack.Pop()
	//
	//if minStack.GetMin() != 1 {
	//	t.Fatal(`minStack.GetMin() != 1`)
	//}
	//fmt.Println(minStack.data, minStack.min)
	//if minStack.Top() != 1 {
	//	t.Fatal(`minStack.Top() != 1`)
	//}
}

// -2 0 3
