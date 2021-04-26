package _8_四数之和

func fourSum(nums []int, target int) [][]int {
	l := len(nums)
	result := [][]int{}
	for i := 0; i < l; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < l; j++ {
			left, right := j+1, l-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum > 0 {
					right--
				} else if sum < 0 {
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

}
