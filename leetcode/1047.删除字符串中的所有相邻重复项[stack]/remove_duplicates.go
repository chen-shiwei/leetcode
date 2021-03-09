package _047_删除字符串中的所有相邻重复项_stack_

func removeDuplicates(S string) string {
	var stack []byte
	for i := 0; i < len(S); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, S[i])
	}
	return string(stack)
}
