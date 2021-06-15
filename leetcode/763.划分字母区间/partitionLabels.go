package _763_划分字母区间

func partitionLabels(s string) []int {
	l := len(s)
	var record = make([]int, 26)
	for i := 0; i < l; i++ {
		record[s[i]-'a'] = i
	}

	var result []int
	var left, right int
	for i := 0; i < l; i++ {
		right = max(right, record[s[i]-'a'])
		if i == right {
			result = append(result, right-left+1)
			left = i + 1
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
