package _13_打家劫舍II

func rob(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		return max(nums[1], nums[0])
	}

	// 偷第一家和最后一家分为两个数组来计算
	return max(_rob(nums[1:]), _rob(nums[:l-1]))

}

func _rob(nums []int) int {
	var first, second = nums[0], max(nums[0], nums[1])
	for _, v := range nums[2:] {
		first, second = second, max(second, first+v)
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
