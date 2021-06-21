package _746_使用最小花费爬楼梯

func minCostClimbingStairs(cost []int) int {
	var (
		l  = len(cost)
		dp = make([]int, l)
	)

	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < l; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}

	return min(dp[l-1], dp[l-2])
}

func min(a, b int) int {
	if a > b {
		return b

	}
	return a
}
