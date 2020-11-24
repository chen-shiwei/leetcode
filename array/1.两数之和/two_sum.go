package __两数之和

// twoSum 利用 dict o(1)
func twoSum1(nums []int, target int) []int {
	var dict = make(map[int]int, len(nums))
	var l = len(nums)
	for i := 0; i < l; i++ {
		dest := target - nums[i]
		index, ok := dict[dest]
		if ok {
			return []int{index, i}
		}
		dict[nums[i]] = i
	}
	return nil
}

// twoSum 两个for
func twoSum2(nums []int, target int) []int {
	var l = len(nums)
	for i := 0; i < l; i++ {
		dest := target - nums[i]
		for j := i + 1; j < l; j++ { // 因为j都和之前每个数匹配过
			if dest == nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}
