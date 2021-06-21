package _738_单调递增的数字

import (
	"strconv"
)

func monotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	l := len(s)

	flag := l
	for i := l - 1; i > 0; i-- {
		if s[i-1] > s[i] {
			s[i-1] -= 1
			flag = i
		}
	}
	for i := flag; i < l; i++ {
		s[i] = '9'
	}

	n, _ = strconv.Atoi(string(s))
	return n
}
