package _136

func SingleNumber(nums []int) int {
	if len(nums) < 1 {
		return nums[0]
	}
	result := 0
	for _, v := range nums {
		result ^= v
	}
	return result
}
