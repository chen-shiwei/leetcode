package 面试题_08_03

func findMagicIndex(nums []int) int {
	//var i, j = 0, len(nums)
	//if j == 0 {
	//	return 0
	//}

	l := len(nums)
	for i := 0; i < l; i++ {

		if i == nums[i] {
			return i
		}

		if nums[i] > i+1 {
			i = nums[i] - 1
		}
	}
	return 0
}
