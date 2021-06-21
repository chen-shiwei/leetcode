package _376_摆动序列

func wiggleMaxLength(nums []int) int {
	var (
		cur    int
		pre    int
		result = 1
	)

	if len(nums) <= 1 {
		return len(nums)
	}

	for i := 1; i < len(nums); i++ {
		cur = nums[i] - nums[i-1]
		if (cur > 0 && pre <= 0) || (cur < 0 && pre >= 0) {
			result++
			pre = cur
		}
	}

	return result
}
