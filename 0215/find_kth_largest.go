package _215

func FindKthLargest(nums []int, k int) int {
	Qp(nums)
	return nums[(len(nums) - k)]
}

func Qp(nums []int) {
	l := len(nums)
	if l < 1 {
		return
	}

	mid := nums[0]
	first := 0
	last := l - 1
	for i := 1; i <= last; {
		if nums[i] < mid {
			nums[first], nums[i] = nums[i], nums[first]
			i++
			first++
		} else {
			nums[last], nums[i] = nums[i], nums[last]
			last--
		}
	}

	Qp(nums[:first])
	Qp(nums[first+1:])
}
