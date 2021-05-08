package _59_重复的子字符串

func repeatedSubstringPattern(s string) bool {

	l := len(s)

	var next = make([]int, l)
	for i, j := 1, 0; i < l; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}

		if s[i] == s[j] {
			j++
		}

		next[i] = j
	}

	if next[l-1] != 0 && l%(l-next[l-1]) == 0 {
		return true
	}

	return false

}
