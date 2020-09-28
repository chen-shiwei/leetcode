package 面试题_08_11

func waysToChange(n int) int {
	if n < 2 {
		return n
	}
	switch {
	case n >= 25:
		dfs(n, 1, 5, 10, 25)
		fallthrough
	case n >= 10:
		dfs(n, 1, 5, 10)
		fallthrough
	case n >= 5:
		dfs(n, 1, 5)
		fallthrough
	case n >= 1:
		dfs(n, 1)
	}

}

func dfs(n int, subN ...int) {

}
