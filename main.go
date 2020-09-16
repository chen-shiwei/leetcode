package main

import "fmt"

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

// stack1 123456
func Pop() int {
	if len(stack2) == 0 {
		if len(stack1) == 0 {
			return 0
		}
		for _, v := range stack1 {
			stack2 = append(stack2, v)
			stack1 = stack1[1:]
		}
		stack1 = nil
	}

	popNum := stack2[0]
	stack2 = stack2[1:]
	return popNum
}

func main() {
	Push(1)
	Push(2)

	Push(3)
	Push(4)

	fmt.Println(Pop())
	fmt.Println(Pop())
	fmt.Println(Pop())
	fmt.Println(Pop())
}
