package _5_三数之和

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
		// 因为已排序  如果 nums[i] > 0 那么nums[left] nums[j] 不会为0
		if nums[i] > 0 {
			return result
		}
		// 错误去重方法，将会漏掉-1,-1,2 这种情况
		/*
		   if nums[i] == nums[i + 1] {
		       continue
		   }
		*/
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, l-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 跳过值 一样的数
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 跳过值 一样的数
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			}
		}

	}

	return result
}
