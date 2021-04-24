package _44_比较含退格的字符串

// 从右开始退格--
func backspaceCompare(s string, t string) bool {
	var i = len(s) - 1
	var j = len(t) - 1
	var skipS, skipT int
	for i >= 0 || j >= 0 {
		for i >= 0 {
			// 遇到 #
			if s[i] == '#' {
				i--
				skipS++
				continue
			}
			// 上个字符是否是#
			if skipS > 0 {
				skipS--
				i--
				continue
			}
			break
		}

		for j >= 0 {
			// 遇到 #
			if t[j] == '#' {
				j--
				skipT++
				continue
			}
			// 上个字符是否是#
			if skipT > 0 {
				skipT--
				j--
				continue
			}
			break
		}

		if j >= 0 && i >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if j >= 0 || i >= 0 {
			return false
		}
		j--
		i--
	}

	return true

}
