package _015

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	l := len(nums)
	if l < 3 {
		return nil
	}
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < l; i++ {
		// 因为已排序  如果 nums[i] > 0 那么nums[l] nums[j] 不会为0
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, l-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				// 跳过值 一样的数
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				// 跳过值 一样的数
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			}
		}

	}

	return result
}
