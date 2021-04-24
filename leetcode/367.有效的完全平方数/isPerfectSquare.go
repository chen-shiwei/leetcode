package _67_有效的完全平方数

func isPerfectSquare(num int) bool {

	var left, right = 0, num
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == num {
			return true
		}

		if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
