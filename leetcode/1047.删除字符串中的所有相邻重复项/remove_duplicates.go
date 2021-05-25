package _047_删除字符串中的所有相邻重复项

import "fmt"

func removeDuplicates(S string) string {
	var stack []byte
	for i := 0; i < len(S); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == S[i] {
			fmt.Println(stack[len(stack)-1], S[i], string(stack[:len(stack)-1]))
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, S[i])
	}
	fmt.Println(string(stack))
	return string(stack)
}
