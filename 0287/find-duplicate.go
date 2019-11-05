package _287

func findDuplicate(nums []int) int {
	var l = len(nums)
	for i := 0; i < l; i++ {

		if nums[i] != i+1 {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
}

}
}
