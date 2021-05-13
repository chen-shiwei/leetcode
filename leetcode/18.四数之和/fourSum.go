package _8_四数之和

import "sort"

func fourSum(nums []int, target int) [][]int {
	l := len(nums)

	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < l; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < l; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}

			left, right := j+1, l-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum > target {
					right--
				} else if sum < target {
					left++
				} else {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				}
			}
		}
	}

	return result

}
