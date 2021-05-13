package __无重复字符的最长子串

func lengthOfLongestSubstring(s string) int {
	var (
		l    = len(s)
		dict = make(map[uint8]int)
		left int
		max  int
	)
	if l == 0 {
		return 0
	}

	for i := 0; i < l; i++ {
		if idx, ok := dict[s[i]]; ok {
			if idx+1 > left {
				left = idx + 1
			}
		}
		dict[s[i]] = i
		if i-left+1 > max {
			max = i - left + 1
		}
	}

	return max
}
