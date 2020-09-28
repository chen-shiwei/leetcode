package _008

import "strconv"

func MyAtoi(str string) int {
	if len(str) < 1 {
		return 0
	}
	strconv.Atoi()
	var symbol = 1
	str0 := str[0]
	if str0 == '-' {
		symbol *= -1
		str = str[1:]
	}
	if str0 == '+'
}
