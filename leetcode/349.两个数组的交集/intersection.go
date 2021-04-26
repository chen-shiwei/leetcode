package _49_两个数组的交集

func intersection(nums1 []int, nums2 []int) []int {
	var dict = make(map[int]int)

	for _, v := range nums1 {
		dict[v] = 1
	}

	var result []int
	for _, v := range nums2 {
		if _, ok := dict[v]; ok {
			dict[v] = 2
		}
	}

	for i, v := range dict {
		if v == 2 {
			result = append(result, i)
		}
	}

	return result
}
