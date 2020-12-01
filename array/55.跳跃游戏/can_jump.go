package _5_跳跃游戏

func canJump(nums []int) bool {
	l := len(nums)
	var step int
	for i := 0; i < l; i++ {
		if step < i {
			return false
		}
		step = max(step, nums[i]+i)
	}
	return true
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
