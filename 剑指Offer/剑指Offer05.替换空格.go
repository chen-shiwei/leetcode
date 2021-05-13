package 剑指Offer

func replaceSpace(s string) string {
	l := len(s)
	c := 0
	for i := 0; i < l; i++ {
		if s[i] == ' ' {
			c++
		}
	}

	l2 := l + c*2
	newS := make([]byte, l2)
	i, j := l-1, l2-1

	for i >= 0 {
		if s[i] == ' ' {
			newS[j] = '0'
			newS[j-1] = '2'
			newS[j-2] = '%'
			j = j - 3
			i--
		} else {
			newS[j] = s[i]
			j--
			i--
		}
	}
	return string(newS)
}
