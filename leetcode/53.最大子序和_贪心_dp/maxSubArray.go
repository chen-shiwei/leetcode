package _3_最大子序和_贪心_dp

import "math"

func maxSubArray(nums []int) int {
	l := len(nums)
	var maxN = nums[0]
	for i := 1; i < l; i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		maxN = max(nums[i], maxN)
	}
	return maxN
}

func maxSubArrayWithDP(nums []int) int {
	l := len(nums)

	var dp = make([]int, l)
	dp[0] = nums[0]
	var ans int
	for i := 1; i < l; i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > ans {
			ans = nums[i]
		}
	}

	return ans
}

func maxSubArrayWithGreedy(nums []int) int {
	l := len(nums)

	var ans int
	var result = math.MinInt32
	for i := 0; i < l; i++ {
		ans += nums[i]
		if ans > result {
			result = ans
		}

		if ans < 0 {
			ans = 0
		}
	}

	return result
}

func maxSubArrayWithFor(nums []int) int {
	l := len(nums)
	var ans = math.MinInt32
	for i := 0; i < l; i++ {
		var tmp int
		for j := i; j < l; j++ {
			tmp += nums[j]
			if tmp > ans {
				ans = tmp
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
