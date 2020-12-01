package _5_跳跃游戏

func canJump(nums []int) bool {
	l := len(nums)
	var step int
	// 可以跳最远的距离内的所有位置都能覆盖
	// 一步一步记录最远距离
	for i := 0; i < l; i++ {
		// 最远距离不能达到 所以false
		if step < i {
			return false
		}
		// can jump farthest
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
