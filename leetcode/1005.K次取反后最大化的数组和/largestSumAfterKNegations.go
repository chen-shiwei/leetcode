package _005_K次取反后最大化的数组和

import (
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	l := len(nums)
	sort.Slice(nums, func(i, j int) bool {
		a := nums[i]
		if nums[i] < 0 {
			a = nums[i] * -1
		}
		b := nums[j]
		if nums[j] < 0 {
			b = nums[j] * -1
		}
		return a > b
	})
	var ans int
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			k--
			nums[i] = -nums[i]
		}
	}

	for k > 0 {
		nums[l-1] = -nums[l-1]
		k--
	}
	for i := 0; i < len(nums); i++ {
		ans += nums[i]
	}

	return ans
}
