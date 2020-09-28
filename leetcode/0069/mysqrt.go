package _069

func mySqrt(x int) int {
	var min, max = 0, x
	var result int

	for min <= max {
		mid := min + (max-min)/2
		// 8 -> 2.8 无限接近
		if mid*mid <= x {
			result = mid
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return result

}
