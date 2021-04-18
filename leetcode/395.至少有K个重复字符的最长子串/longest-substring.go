package _395_至少有K个重复字符的最长子串

import "strings"

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}

	var cnt [26]int

	// 数组字典统计字符出现次数
	for _, c := range s {
		cnt[c-'a']++
	}

	var sep byte
	// 如果字符出现但次数小于k，表示该包含该字符的子串都不满足k个重复子串，去掉该字符，查找该字符之外，也就是该字符的两边
	for i, count := range cnt {
		if count > 0 && count < k {
			sep = byte(i) + 'a'
			break
		}
	}
	// 如果没有出现小于k个的字符，表示都是满足k个重复子串
	if sep == 0 {
		return len(s)
	}

	var result int
	for _, substr := range strings.Split(s, string(sep)) {
		result = max(result, longestSubstring(substr, k))
	}

	return result
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
