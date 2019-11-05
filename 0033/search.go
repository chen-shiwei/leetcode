package _033

func Search(nums []int, target int) int {
	var l = len(nums)
	var head, tail = 0, l - 1

	var mid int
	for head <= tail {
		mid = head + (tail-head)/2
		if nums[mid] == target {
			return mid
		}

		if nums[head] <= nums[mid] {
			if target >= nums[head] && target < nums[mid] {
				tail = mid - 1
			} else {
				head = mid + 1
			}
		} else {
			if target <= nums[tail] && target > nums[mid] {
				head = mid + 1
			} else {
				tail = mid - 1
			}
		}
	}

	return -1
}
