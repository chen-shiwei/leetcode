package _26_删除有序数组中的重复项

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}

	var i, j = 0, 1

	for j < l {
		if nums[i] == nums[j] {
			nums = append(nums[:i], nums[i+1:]...)
			l--
		} else {
			i++
			j++
		}
	}

	return len(nums)
}
func removeDuplicates1(nums []int) int {
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
