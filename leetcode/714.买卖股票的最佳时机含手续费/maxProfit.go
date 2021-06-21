package _714_买卖股票的最佳时机含手续费

func maxProfit(prices []int, fee int) int {
	l := len(prices)
	var dp = make([][2]int, l)

	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][1]+prices[i]-fee, dp[i-1][0])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}

	return max(dp[l-1][1], dp[l-1][0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
