package _5_搜索插入位置

func searchInsert(nums []int, target int) int {
	var left, right = 0, len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 取道最接近 target的左边的值
	return left

}
