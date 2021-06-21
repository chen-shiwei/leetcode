package _5_跳跃游戏

func canJump(nums []int) bool {

	l := len(nums)
	var ans int
	if l <= 1 {
		return true
	}
	for i := 0; i <= ans; i++ {
		ans = max(ans, nums[i]+i)
		if ans >= l-1 {
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
