package _287

func FindDuplicate(nums []int) int {
	var l = len(nums)
	for i := 0; i < l; {
		if nums[i] != i+1 {
			if nums[i] == nums[nums[i]-1] {
				return nums[i]
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}
	return 0
}
