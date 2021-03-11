package sort

import "sort"

// selectSort 选择排序
// 空间复杂度 O(1)
// 时间复杂度 O(n2)
// 原地排序、非稳定排序
func selectSort(nums sort.Interface, desc bool) {

	for i := 0; i < nums.Len()-1; /*最后一位不用排了*/ i++ {
		var min = i
		for j := i; j < nums.Len(); j++ {
			if desc {
				if !nums.Less(j, min) {
					min = j
				}
			} else {
				if nums.Less(j, min) {
					min = j
				}
			}

		}
		nums.Swap(min, i)
	}

	return
}
