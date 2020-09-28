package offer_10

func numWays(n int) int {
	if n <= 1 {
		return 1
	}

	var a, b = 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, (a+b)%1000000007
	}

	return b

}
