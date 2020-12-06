package _81_最短无序连续子数组

import (
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	numsClone := make([]int, len(nums), cap(nums))
	copy(numsClone, nums)

	sort.Ints(nums)
	l := len(nums)
	var start, end = l, 0
	for i := 0; i < l; i++ {
		if nums[i] != numsClone[i] {
			if i < start {
				start = i
			}
			if i > end {
				end = i
			}
		}
	}

	if end-start > 0 {
		return end - start + 1
	}
	return 0

}
