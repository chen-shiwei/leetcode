package _169

func MajorityElement(nums []int) int {
	var l = len(nums)

	var dict = make(map[int]int, 0)
	var max, c int
	for i := 0; i < l; i++ {
		ic, ok := dict[nums[i]]
		if !ok {
			dict[nums[i]] = 1
			ic = 1
		} else {
			ic = ic + 1
			dict[nums[i]] = ic
		}

		if ic > c {
			max = nums[i]
			c = ic
		}
	}

	if c >= l/2 {
		return max
	}

	return 0
}
