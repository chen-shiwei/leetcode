package _198_打家劫舍

func rob(nums []int) int {
	l := len(nums)
	dp := make([]int, l+1)

	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= l; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[l]

}

//假设第 n 间没有被偷，那么此时 dp[n] = dp[n-1]dp[n]=dp[n−1] ，此时 dp[n+1] = dp[n] + num = dp[n-1] + numdp[n+1]=dp[n]+num=dp[n−1]+num ，即可以将 两种情况合并为一种情况 考虑；
//假设第 n 间被偷，那么此时 dp[n+1] = dp[n] + numdp[n+1]=dp[n]+num 不可取 ，因为偷了第 nn 间就不能偷第 n+1n+1 间。
func robDP(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		return max(nums[0], nums[1])
	}

	var cur, pre = max(nums[0], nums[1]), nums[0]
	for _, v := range nums[2:] {
		cur, pre = max(cur, pre+v), cur
	}
	return cur

	// 1 4 7 2 4 2
	//
	//cur pre  1 4
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
