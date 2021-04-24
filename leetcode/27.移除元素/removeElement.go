package _7_移除元素

func removeElement(nums []int, val int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	var i = 0
	for i < l {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			l--
		} else {
			i++
		}
	}

	return l
}

// 双指针
func removeElement1(nums []int, val int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	var quick, slow int

	// 用quick覆盖重复的slow(nums[quick] != val)
	for quick = 0; quick < l; quick++ {
		if nums[quick] != val {
			nums[slow] = nums[quick]
			slow += 1
		}
	}
	return slow
}
