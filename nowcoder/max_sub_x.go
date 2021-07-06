package nowcoder

func getMaxSubX(nums []int) int {
	l := len(nums)

	var flag int
	var maxL int
	var left int
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			left = 0
			flag = 0
		} else if nums[i] > 0 {
			left++
		} else {
			left++
			flag++
		}

		if flag%2 == 0 {
			maxL = max(maxL, left)
		}
	}

	return maxL

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
