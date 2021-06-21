package _35_分发糖果

func candy(ratings []int) int {
	l := len(ratings)
	s := make([]int, l)
	for i := 1; i < l; i++ {
		if ratings[i] > ratings[i-1] {
			s[i] = s[i-1] + 1
		}
	}

	for i := l - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			s[i] = max(s[i], s[i+1]+1)
		}
	}

	ans := 0
	for i := 0; i < l; i++ {
		ans += s[i]
	}

	return ans + l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
