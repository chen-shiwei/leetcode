package _04_二分查找

func search(nums []int, target int) int {
	// 左闭右闭区间[left,right]
	// right 也是有效范围
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
	return -1
}

func search1(nums []int, target int) int {
	// 左闭右开区间[left,right)
	// right 也是有效范围
	var left, right = 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			// 因为 right 一定得是无效的，如果有效 left< right永远走不到
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}
