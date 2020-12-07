package _52_乘积最大子数组

func maxProduct(nums []int) int {
	l := len(nums)
	var minNums, maxNums = make([]int, len(nums), cap(nums)), make([]int, len(nums), cap(nums))
	copy(minNums, nums)
	copy(maxNums, nums)

	for i := 1; i < l; i++ {
		maxNums[i] = max(maxNums[i-1]*nums[i], max(nums[i], minNums[i-1]*nums[i]))
		minNums[i] = min(minNums[i-1]*nums[i], min(nums[i], maxNums[i-1]*nums[i]))
	}

	var result = maxNums[0]
	for i := 1; i < l; i++ {
		result = max(result, maxNums[i])

	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
