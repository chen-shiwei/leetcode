package main

func Find(nums []int, target int) int {
	l := len(nums)
	var start, end = 0, l - 1
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			var left = start
			for left > 0 && nums[left] == target {
				left--
			}
			return left
		}
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	// start
	for i := start; i < l; i++ {
		if nums[start] >= target {
			return i
		}
	}

	return -1

}
