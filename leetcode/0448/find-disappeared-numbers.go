package _448

// FindDisappearedNumbers
// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
func FindDisappearedNumbers(nums []int) []int {
	l := len(nums)

	var result []int
	// 把出现的数对应的索引加n
	for i := 0; i < l; i++ {
		index := (nums[i] - 1) % l
		nums[index] += l
	}

	// 找出没加n的
	for i := 0; i < l; i++ {
		if nums[i] <= l {
			result = append(result, i+1)
		}
	}

	return result
}
