package _4_在排序数组中查找元素的第一个和最后一个位置

func searchRange(nums []int, target int) []int {
	var left, right = 0, len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			var min = mid
			for i := mid; i >= left; i-- {
				if nums[i] == target {
					min = i
				} else {
					break
				}
			}
			var max = mid
			for i := mid; i <= right; i++ {
				if nums[i] == target {
					max = i
				} else {
					break
				}
			}
			return []int{min, max}
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return []int{-1, -1}

}
