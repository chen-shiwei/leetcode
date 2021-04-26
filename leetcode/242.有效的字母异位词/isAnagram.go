package _42_有效的字母异位词

func isAnagram(s string, t string) bool {
	var dict [26]int

	for i := range s {
		dict[s[i]-'a']++
	}

	for i := range t {
		dict[t[i]-'a']--
	}

	for i := range dict {
		if dict[i] != 0 {
			return false
		}
	}

	return true
}
