package _224_基本计算器

import "fmt"

func calculate(s string) int {
	l := len(s)
	var (
		result int
		sign   = 1
		stack  = []int{1}
	)
	for i := 0; i < l; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = stack[len(stack)-1]
			i++
		case '-':
			sign = -stack[len(stack)-1]
			i++
		case '(':
			stack = append(stack, sign)
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		default:
			var n int
			for ; i < l && s[i] >= '0' && s[i] <= '9'; i++ {
				n = n*10 + int(s[i]-'0')
			}

			result += n * sign
			fmt.Println("res", result)
		}
	}
	return result
}
