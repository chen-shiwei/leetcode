package _232

import (
	"testing"
)

//queue.push(1);
//queue.push(2);
//queue.peek();  // 返回 1
//queue.pop();   // 返回 1
//queue.empty(); // 返回 false
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/implement-queue-using-stacks
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func TestQueue(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	t.Log(queue.Pop())

	queue.Push(5)
	queue.Push(6)
	t.Log(queue.Pop())

	//t.Log(queue.Peek())
	//t.Log(queue.Empty())
}

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
}
