package _647

func countSubstrings(s string) int {
	l := len(s)

	var count int

	for i := 0; i < l*2-1; i++ {
		left, right := i/2, i/2+i%2 /*回文中心起始位置 除以2加上不能除尽的数 3/2+1*/
		for left >= 0 && right < l && s[left] == s[right] {
			left--
			right++
			count++
		}
	}

	return count
}
