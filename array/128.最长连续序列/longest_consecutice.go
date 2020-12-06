package _28_最长连续序列

func longestConsecutive(nums []int) int {

	l := len(nums)
	var long int
	var dict = make(map[int]bool)
	for _, v := range nums {
		dict[v] = true
	}

	for i := 0; i < l; i++ {
		var curLong = 0
		var n = nums[i]
		if !dict[nums[i]-1] {
			curLong = 1
			for dict[n+1] {
				n++
				curLong++
			}
			if curLong > long {
				long = curLong
			}
		}
	}

	return long
}
