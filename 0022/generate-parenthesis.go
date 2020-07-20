package _022

func GenerateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	var result []string
	dfs("", n, n, result)
	return result
}

func dfs(curStr string, left, right int, result []string) {
	if left == 0 && right == 0 {

	}
}
