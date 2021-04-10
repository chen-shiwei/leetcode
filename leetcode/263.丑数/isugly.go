package _63_丑数

func isUgly(n int) bool {
	if n == 0 {
		return false
	}

	var ugly = []int{2, 3, 5}
	for _, v := range ugly {
		for n%v == 0 {
			n = n / v
		}
	}

	return n == 1
}
