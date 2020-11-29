package search

func BinarySearch(nums []int, val int) (i int) {
	low, height := 0, len(nums)-1

	for low <= height {
		m := (low + height) / 2
		if nums[m] == val {
			return m
		}
		if nums[m] > val {
			height = m - 1
		} else if nums[m] < val {
			low = m + 1
		}
	}

	return 0

}
