package _09_长度最小的子数组

import "math"

// 暴力法 双for
func minSubArrayLenViolence(target int, nums []int) int {
	l := len(nums)

	var minLen = math.MaxInt32
	for i := 0; i < l; i++ {
		var sum int
		for j := i; j < l; j++ {
			sum += nums[i]
			if subLen := j - i + 1; sum >= target && subLen < minLen {
				minLen = subLen
				break
			}
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

func minSubArrayLenSlidingWindow(target int, nums []int) int {

	var left, right int
	l := len(nums)
	minLen := math.MaxInt32

	// 双指针滑动
	// right 移动控制是sum 增加
	// left 指针当sum >=target 时候更新 minlen 并向前移动
	var sum int
	for ; right < l; right++ {
		sum += nums[right]
		for sum >= target {
			subLen := right - left + 1
			if subLen < minLen {
				minLen = subLen
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}

	return minLen

}
