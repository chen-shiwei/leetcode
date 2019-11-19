package sort

func QuickSort(nums []int) {
	l := len(nums)
	if len(nums) < 1 {
		return
	}
	mid := nums[0]

	var head, tail = 0, l - 1
	for i := 1; i <= tail; {
		if nums[i] <= mid {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		} else {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		}
	}

	QuickSort(nums[:head])

	QuickSort(nums[head+1:])

}
