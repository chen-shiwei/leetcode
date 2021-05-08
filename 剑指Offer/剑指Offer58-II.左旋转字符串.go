package 剑指Offer

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	reverse(b[0:n])
	reverse(b[n:])
	reverse(b)
	return string(b)
}

func reverse(b []byte) {
	l := len(b)
	if l < 2 {
		return
	}

	left, right := 0, l-1

	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}

	return
}
