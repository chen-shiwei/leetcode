package _50_逆波兰表达式求值

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	l := len(tokens)
	var stack []int
	for i := 0; i < l; i++ {
		switch tokens[i] {
		case "+":
			if len(stack) >= 2 {
				num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, num1+num2)
			}
		case "-":
			if len(stack) >= 2 {
				num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, num1-num2)
			}
		case "*":
			if len(stack) >= 2 {
				num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, num1*num2)
			}
		case "/":
			if len(stack) >= 2 {
				num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, num1/num2)
			}
		default:
			n, _ := strconv.Atoi(tokens[i])
			stack = append(stack, n)
		}
	}

	if len(stack) > 0 {
		result := stack[0]
		stack = stack[1:]
		return result
	}

	return -1
}
