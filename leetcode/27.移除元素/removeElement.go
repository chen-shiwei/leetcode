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
