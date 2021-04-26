package _54_四数相加II

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	dict := make(map[int]int)

	for _, a := range nums1 {
		for _, b := range nums2 {
			dict[a+b]++
		}
	}

	count := 0
	for _, c := range nums3 {
		for _, d := range nums4 {

			val := 0 - (c + d)
			n, ok := dict[val]
			if ok {
				count += n
			}
		}
	}
	return count
}
