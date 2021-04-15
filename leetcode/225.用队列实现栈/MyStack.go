package _25_用队列实现栈

type MyStack struct {
	master *Queue
	slave  *Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		master: &Queue{},
		slave:  &Queue{},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.slave.push(x)
	for !this.master.empty() {
		this.slave.push(this.master.pop())
	}
	this.master, this.slave = this.slave, this.master
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.master.pop()
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.master.top()
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.master.empty() && this.slave.empty()
}

type Queue struct {
	data []int
	l    int
}

func (q *Queue) empty() bool {
	return q.l == 0
}

func (q *Queue) push(n int) {
	q.data = append(q.data, n)
	q.l++
}

func (q *Queue) top() int {
	if q.empty() {
		return -1
	}
	return q.data[0]
}

func (q *Queue) pop() int {
	if q.empty() {
		return -1
	}
	n := q.data[0]
	q.data = q.data[1:]
	q.l--
	return n
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
