package __无重复字符的最长子串

func lengthOfLongestSubstring(s string) int {
	var (
		l      = len(s)
		dict   = make(map[uint8]int)
		start  int
		maxLen int
	)
	if l == 0 {
		return 0
	}

	// 滑动仓库 start->i
	for i := 0; i < l; i++ {
		if idx, ok := dict[s[i]]; ok {
			// id+1 左端出队
			start = max(idx+1, start)
		}
		dict[s[i]] = i
		maxLen = max(i-start+1, maxLen)
	}

	return maxLen
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
