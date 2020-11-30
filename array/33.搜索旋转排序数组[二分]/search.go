package _3_搜索旋转排序数组_二分_

func search(nums []int, target int) int {

	l := len(nums)
	var start, end = 0, l - 1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}

		// 前段数组有序
		if nums[start] <= nums[mid] {
			// target 在前段数组
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			// target 在后段数组
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}

	}

	return -1

}
