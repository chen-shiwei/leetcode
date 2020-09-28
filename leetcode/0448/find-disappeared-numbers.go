package _448

// FindDisappearedNumbers
// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
func FindDisappearedNumbers(nums []int) []int {
	l := len(nums)

	var result []int
	for i := 0; i < l; i++ {
		index := (nums[i] - 1) % l
		nums[index] += l
	}

	for i := 0; i < l; i++ {
		if nums[i] <= l {
			result = append(result, i+1)
		}
	}

	return result
}
