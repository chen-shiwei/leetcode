package _7_扰乱字符串

func isScramble(s1 string, s2 string) bool {
	// 记忆 减少重复计算
	var dist = make(map[string]bool)

	var dfs func(s1 string, s2 string) bool

	dfs = func(s1 string, s2 string) bool {
		// 保存当前s1,s2 已被计算
		var s = s1 + s2
		flag, ok := dist[s]
		if ok {
			return flag
		}

		l1 := len(s1)
		l2 := len(s2)
		if l1 != l2 {
			dist[s] = false
			return false
		}
		if s1 == s2 {
			dist[s] = true
			return true
		}

		var d [26]int

		for i := 0; i < len(s1); i++ {
			d[s1[i]-'a']++
			d[s2[i]-'a']--
		}

		for _, v := range d {
			if v != 0 {
				dist[s] = false
				return false
			}
		}

		for i := 1; i < l1; i++ {
			if dfs(s1[:i], s2[:i]) && dfs(s1[i:], s2[i:]) {
				dist[s] = true
				return true
			}

			if dfs(s1[:i], s2[l1-i:]) && dfs(s1[i:], s2[:l1-i]) {
				dist[s] = true
				return true
			}
		}

		dist[s] = false
		return false
	}
	return dfs(s1, s2)
}
