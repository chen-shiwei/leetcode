package _215

func FindKthLargest(nums []int, k int) int {
	//Qp(nums)
	qp(nums)
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
}

func qp(nums []int) {
	l := len(nums)
	if l < 1 {
		return
	}

	flag := nums[0]
	flagIdx := 0
	last := l - 1
	for i := 1; i <= last; {
		if nums[i] < flag {
			nums[i], nums[flagIdx] = nums[flagIdx], nums[i]
			i++
			flagIdx++
		} else {
			nums[i], nums[last] = nums[last], nums[i]
			last--
		}
	}
	qp(nums[:flagIdx])
	qp(nums[flagIdx+1:])
}
