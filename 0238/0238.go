package _238

func ProductExceptSelf(nums []int) []int {

	l := len(nums)
	var L, R, result = make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
	L[0], R[l-1] = 1, 1

	for i := 1; i < l; i++ {
		L[i] = L[i-1] * nums[i-1]
	}

	for i := l - 2; i >= 0; i-- {
		R[i] = R[i+1] * nums[i+1]
	}

	for i, _ := range result {
		result[i] = L[i] * R[i]
	}

	return result
}
