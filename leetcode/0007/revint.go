package _007

func Reverse(x int) int {
	var n int
	for x != 0 {
		n = n*10 + x%10
		x = x / 10
	}
	if n > 1<<31-1 || n < -(1<<31-1) {
		return 0
	}
	return n
}
