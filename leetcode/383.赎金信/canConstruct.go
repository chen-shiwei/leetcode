package _83_赎金信

func canConstruct(ransomNote string, magazine string) bool {
	var dict = [26]int{}

	for i := 0; i < len(magazine); i++ {
		dict[magazine[i]-'a']++
	}

	// ransomNote有字符在 magazine没有
	for j := 0; j < len(ransomNote); j++ {
		dict[ransomNote[j]-'a']--
		if dict[ransomNote[j]-'a'] < 0 {
			return false
		}
	}

	return true
}
