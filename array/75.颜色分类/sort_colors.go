package _5_颜色分类

func sortColors(nums []int) {

	var swapColor = func(nums []int, target int) (curIdx int) {
		l := len(nums)
		for i := 0; i < l-1; i++ {
			if nums[i] == target {
				nums[i], nums[curIdx] = nums[curIdx], nums[i]
				curIdx++
			}
		}
		return curIdx
	}

	curIdx := swapColor(nums, 0)
	swapColor(nums[curIdx:], 1)

	return
}
