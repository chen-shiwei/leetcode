package __无重复字符的最长子串

func lengthOfLongestSubstring(s string) int {
	var dict = make(map[string]int)
	var l = len(s)
	if l == 0 {
		return 0
	}

	var (
		start  int
		maxLen int
	)
	for i := 0; i < l; i++ {
		if workIdx, ok := dict[string(s[i])]; ok {
			start = max(workIdx+1, start)
		}

		dict[string(s[i])] = i
		maxLen = max(maxLen, i-start+1)
	}

	return maxLen

}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
