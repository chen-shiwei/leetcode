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
