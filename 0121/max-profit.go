package _121

import "math"

func MaxProfit(prices []int) int {
	var (
		maxProfit = 0
		min       = math.MaxInt32
	)
	l := len(prices)
	for i := 0; i < l; i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}

	}
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
