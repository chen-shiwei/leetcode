package _8_实现strStr

import "fmt"

func strStr(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)

	for i := 0; i < l1-l2+1; i++ {
		var s2 int
		for j := 0; j < l2; j++ {
			if haystack[i+j] != needle[j] {
				continue
			}
			s2++
		}

		if s2 == l2 {
			return i
		}
	}
	return -1
}

func strStrWithKMP(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)

	next := make([]int, l2)
	for i, j := 1, 0; i < l2; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}
	fmt.Println(next)
	for i, j := 0, 0; i < l1; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == l2 {
			return i - l2 + 1
		}
	}
	return -1
}
