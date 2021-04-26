package _202

func isHappy(n int) bool {
	var dict = make(map[int]struct{})
	for n != 1 {
		sum := getSum(n)
		if _, ok := dict[sum]; ok {
			return false
		} else {
			dict[sum] = struct{}{}
		}

		n = sum

	}
	return true

}

func getSum(n int) int {
	var sum int
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
