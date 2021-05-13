package _151_翻转字符串里的单词

func reverseWords(s string) string {
	l := len(s)
	b := []byte(s)
	reverse(b, 0, l-1)
	left := 0
	for i := 0; i < l; i++ {
		if b[i] == ' ' {
			reverse(b, left, i-1)
			left = i + 1
		}
	}
	reverse(b, left, l-1)

	c := removeSpaces(b)
	return string(c)
}

func removeSpaces(b []byte) []byte {
	var slow, quick int

	l := len(b)
	for l > 0 && quick < l && b[quick] == ' ' {
		quick++
	}

	for ; quick < l; quick++ {
		if quick-1 > 0 && b[quick] == ' ' && b[quick-1] == ' ' {
			continue
		} else {
			b[slow] = b[quick]
			slow++
		}
	}
	if slow-1 > 0 && b[slow-1] == ' ' {
		return b[:slow-1]
	} else {
		return b[:slow]
	}
}

func reverse(b []byte, left, right int) {
	l := len(b)
	if l < 2 {
		return
	}

	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}

	return
}
