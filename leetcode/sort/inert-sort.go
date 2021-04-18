package sort

import "sort"

// https://zh.wikipedia.org/wiki/%E6%8F%92%E5%85%A5%E6%8E%92%E5%BA%8F#Go

// InsertSort 插入排序
func InsertSort(nums []int) []int {

	l := len(nums)

	for i := 0; i < l; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			} else {
				break
			}
		}
	}

	return nums

}

func insertSort(nums sort.Interface, desc bool) {
	l := nums.Len()
	for i := 0; i < l; i++ {
		// 把当前 i 和 0:i-1 之间的数对比，比之小插入
		for j := i; j > 0; j-- {
			if nums.Less(j, j-1) {
				nums.Swap(j, j-1)
			} else {
				break
			}
		}
	}
}
