package _41_反转字符串II

func reverseStr(s string, k int) string {

	b := []byte(s)

	l := len(s)

	for i := 0; i < l; {
		var right int
		if i+k >= l {
			right = i + (l - i)
		} else {
			right = i + k
		}
		reverse(b, i, right-1)
		i = right + k
	}

	return string(b)
}

func reverse(s []byte, left, right int) {
	l := len(s)
	if l < 2 {
		return
	}

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	return
}
