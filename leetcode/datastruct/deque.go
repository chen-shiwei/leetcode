package datastruct

import "fmt"

type Deque struct {
	data []int
}

func NewDeque() *Deque {
	return &Deque{data: make([]int, 0)}
}

func (d *Deque) PushFront(x int) {
	d.data = append([]int{x}, d.data...)
}

func (d *Deque) PushBack(x int) {
	d.data = append(d.data, x)
}

func (d *Deque) PopFront() int {
	if d.Empty() {
		return -1
	}
	pop := d.data[0]
	d.data = d.data[1:]
	return pop
}

func (d *Deque) Empty() bool {
	return len(d.data) == 0
}

func (d *Deque) PopBack() int {
	if d.Empty() {
		return -1
	}
	pop := d.data[len(d.data)-1]
	d.data = d.data[:len(d.data)-1]
	return pop
}

func (d *Deque) Front() int {
	if d.Empty() {
		return -1
	}
	return d.data[0]
}

func (d *Deque) Back() int {
	if d.Empty() {
		return -1
	}
	return d.data[len(d.data)-1]
}

func (d *Deque) Println() {
	fmt.Println(d.data)
}
