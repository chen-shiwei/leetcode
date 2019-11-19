package sort

import "sort"

// BubbleSort bubble implement
// https://zh.wikipedia.org/wiki/%E5%86%92%E6%B3%A1%E6%8E%92%E5%BA%8F
func BubbleSort(nums sort.Interface, desc bool) {
	l := nums.Len()
	for i := 0; i < l-1; /*除了最后一个不需要比较排序*/ i++ {
		isChanged := false
		for j := 0; j < l-1-i; j++ {
			if desc {
				if nums.Less(j, j+1) {
					nums.Swap(j, j+1)
					isChanged = true
				}
			} else {
				if !nums.Less(j, j+1) {
					nums.Swap(j, j+1)
					isChanged = true
				}
			}

		}
		if !isChanged {
			break
		}
	}
}
