package _389_找不同

func findTheDifference(s string, t string) byte {

	var cnt [26]int
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		cnt[t[i]-'a']--
		if cnt[t[i]-'a'] < 0 {
			return t[i]
		}
	}

	return 0
}
