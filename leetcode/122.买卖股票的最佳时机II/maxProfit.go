package _22_买卖股票的最佳时机II

func maxProfit(prices []int) int {
	var ans int

	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			ans += prices[i] - prices[i-1]
		}
	}

	return ans
}

func maxProfitWithDP(prices []int) int {
	var l = len(prices)

	var dp = make([][2]int, len(prices))
	// 现金
	dp[0][0] = 0
	// 股票
	dp[0][1] -= prices[0]

	for i := 1; i < l; i++ {
		// 昨天的股票以今天价格卖了换成现金
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 昨天的现金买了今天的股票
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return max(dp[l-1][0], dp[l-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
