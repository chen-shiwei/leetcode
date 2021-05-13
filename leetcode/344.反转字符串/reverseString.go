package _44_反转字符串

func reverseString(s []byte) {
	l := len(s)
	if l < 2 {
		return
	}

	left, right := 0, l-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
