package 剑指Offer

func maxSubArray(nums []int) int {
	l := len(nums)

	var result = nums[0]
	for i := 1; i < l; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		result = max(result, nums[i])
	}

	return result
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
