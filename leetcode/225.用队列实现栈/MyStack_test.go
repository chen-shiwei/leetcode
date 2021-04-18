package _25_用队列实现栈

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	myStack := Constructor()
	myStack.Push(1)
	myStack.Push(2)
	fmt.Println(myStack.master)
	fmt.Println(myStack.slave)
	if myStack.Top() != 2 {
		t.Error("top no 返回 2")
	} // 返回 2
	fmt.Println(myStack.master)
	fmt.Println(myStack.slave)
	if myStack.Pop() != 2 {
		t.Error("pop no 返回 2")
	} // 返回 2
	if myStack.Empty() {
		t.Error("empty no 返回 false")
	} // 返回 False

}
