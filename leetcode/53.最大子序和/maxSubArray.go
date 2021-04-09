package _3_最大子序和

func maxSubArray(nums []int) int {
	l := len(nums)
	var maxN = nums[0]
	for i := 1; i < l; i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		maxN = max(nums[i], maxN)
	}
	return maxN
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
