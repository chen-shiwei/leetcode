package _77_有序数组的平方

func sortedSquares(nums []int) []int {
	var l = len(nums)
	// 双指针
	var left, right = 0, l - 1

	var result = make([]int, l, l)

	// 因为 -4,-1,0,3,10
	// 比较最左右，最大的倒序插入结果数据
	// 0n时间复杂度
	for i := l - 1; i >= 0; i-- {
		if a, b := nums[left]*nums[left], nums[right]*nums[right]; a >= b {
			result[i] = a
			left++
		} else {
			result[i] = b
			right--
		}
	}

	return result
}
