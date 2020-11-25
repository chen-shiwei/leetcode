package _370_上升下降字符串

import (
	"math"
)

func sortString(s string) string {
	var nums = make([]int, 26)
	for _, char := range s {
		nums[char-'a']++
	}
	l := len(s)
	var result = make([]rune, l)
	var i int
	for i < l {
		for a := 0; a < 26; a++ {
			if nums[a] != 0 {
				result[i] = rune(a + 'a')
				nums[a]--
				i++
			}
		}
		for b := 25; b >= 0; b-- {
			if nums[b] != 0 {
				result[i] = rune(b + 'a')
				nums[b]--
				i++
			}
		}
	}
	return string(result)
}

func removeIdxStr(s string, idx int) string {
	//if idx == len(s)-1 {
	//	return s[0:idx]
	//}
	return s[0:idx] + s[idx+1:]
}

func getMinChar(s string) (i int, minStr rune) {
	var minChar int32 = math.MaxInt32
	var idx int
	for _, c := range s {
		if c < minChar {
			minChar = c
			idx = i
		}
	}
	return idx, minChar
}

func getMaxChar(s string) (i int, maxStr rune) {
	var maxChar int32 = math.MinInt32
	var idx int
	for i, c := range s {
		if c > maxChar {
			maxChar = c
			idx = i
		}
	}
	return idx, maxChar
}
