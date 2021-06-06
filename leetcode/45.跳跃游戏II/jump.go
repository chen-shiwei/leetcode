package _5_跳跃游戏II

func jump(nums []int) int {
	l := len(nums)

	var next, cur, c int
	for i := 0; i < l; i++ {
		next = max(nums[i]+i, next)
		if i == cur {
			if i == l-1 {
				break
			}
			cur = next
			c++
			if next >= l-1 {
				break
			}
		}
	}
	return c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
