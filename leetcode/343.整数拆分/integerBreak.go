package _343_整数拆分_

func integerBreak(n int) int {
	if n <= 4 {
		return n
	}

	var start = 1

	for n > 4 {
		start *= 3
		n -= 3
	}

	return start * n
}
