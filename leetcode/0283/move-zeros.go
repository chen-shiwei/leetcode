package _283

func MoveZeroes(nums []int) {
	var length = len(nums)
	var i, j int
	for j < length {
		if nums[i] == 0 {
			if nums[j] == 0 {
				j++
			} else {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j++
			}
		} else {
			i++
			j++
		}
	}

}
