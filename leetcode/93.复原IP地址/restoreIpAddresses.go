package _3_复原IP地址

func restoreIpAddresses(s string) []string {
	var (
		paths []string
		fn    func(s string, startIdx int, n int)
	)

	fn = func(s string, startIdx int, n int) {
		if n == 3 {
			if isValidIP(s, startIdx, len(s)-1) {
				paths = append(paths, s)
			}
			return
		}

		for i := startIdx; i < len(s); i++ {
			if !isValidIP(s, startIdx, i) {
				continue
			}
			s = s[:i+1] + "." + s[i+1:]
			n++
			fn(s, i+2, n)
			n--
			s = s[:i+1] + s[i+2:]
		}
	}
	fn(s, 0, 0)
	return paths
}

// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
func isValidIP(s string, start, end int) bool {
	if start > end {
		return false
	}

	// 有前导 0
	if s[start] == '0' && start != end {
		return false
	}

	var num = 0
	for i := start; i <= end; i++ {
		if s[i] > '9' || s[i] < '0' {
			return false
		}
		// 累加
		num = num*10 + int(s[i]-'0')
		if num > 255 {
			return false
		}
	}

	return true
}
