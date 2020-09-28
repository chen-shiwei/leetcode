package _415

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)

	var (
		result = ""
		// 累加进制
		addBase = 0
	)

	for i, j := l1-1, l2-1; i >= 0 || j >= 0 || addBase != 0; /*还会向上进制*/ i, j = i-1, j-1 {
		var a, b int
		if i >= 0 {
			a = int(num1[i] - '0')
		}

		if j >= 0 {
			b = int(num2[j] - '0')
		}

		sum := a + b + addBase
		result = strconv.Itoa(sum%10) + result
		addBase = sum / 10
	}
	return result

}
