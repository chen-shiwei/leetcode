package _4_在排序数组中查找元素的第一个和最后一个位置

func searchRange(nums []int, target int) []int {

	var start, end = 0, len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			var a, b int
			for i := mid; i >= start; i-- {
				if nums[i] == target {
					a = i
				} else {
					break
				}
			}

			for i := mid; i <= end; i++ {
				if nums[i] == target {
					b = i
				} else {
					break
				}
			}
			return []int{a, b}
		}

		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return []int{-1, -1}
}
