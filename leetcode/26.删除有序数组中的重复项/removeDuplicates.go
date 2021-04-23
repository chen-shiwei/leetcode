package _6_删除有序数组中的重复项

func removeDuplicates(nums []int) int {
	l := len(nums)

	// 第一个数和之前不重复,不用覆盖
	var slow = 1
	for quick := 1; quick < l; quick++ {
		if nums[quick] != nums[quick-1] {
			nums[slow] = nums[quick]
			slow += 1
		}
	}
	return slow
}
