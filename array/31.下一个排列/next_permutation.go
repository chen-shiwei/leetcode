package _1_下一个排列

import "fmt"

func nextPermutation(nums []int) {

	l := len(nums)
	if l <= 1 {
		return
	}

	last := l - 1
	var i, j = l - 2, l - 1
	// 寻找 nums[i]<nums[j]
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	fmt.Println(nums[i], nums[j])

	// 存在 nums[i]<nums[j]
	if i >= 0 {
		for nums[i] >= nums[last] {
			last--
		}
		fmt.Println(nums[i], nums[last])
		// 交换
		nums[i], nums[last] = nums[last], nums[i]
	}

	// 反转
	a, b := j, l-1
	for a < b {
		nums[a], nums[b] = nums[b], nums[a]
		a = a + 1
		b = b - 1
	}

}
