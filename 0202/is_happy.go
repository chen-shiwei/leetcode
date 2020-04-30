package _202

func isHappy(n int) bool {
	var slow, first = n, squareSum(n)
	for slow != first {
		slow = squareSum(slow)
		first = squareSum(squareSum(first))
	}
	return slow == 1
}

func squareSum(n int) int {
	var sum = 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}
