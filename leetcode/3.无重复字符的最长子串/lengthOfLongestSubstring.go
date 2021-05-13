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

	// 滑动仓库 start->i
	for i := 0; i < l; i++ {
		if idx, ok := dict[s[i]]; ok {
			if idx+1 > left {
				// id+1 左端出队
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
