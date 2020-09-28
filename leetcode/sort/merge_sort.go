package sort

func MergeSort(nums []int) {

	l := len(nums)
	if l < 2 {
		return
	}

	split(nums, 0, l-1)

}

func split(nums []int, left, right int) {

	if left >= right {
		return
	}
	mid := left + (right-left)/2
	split(nums, left, mid)
	split(nums, mid+1, right)
	merge(nums, left, mid, right)

}

func merge(nums []int, left, mid, right int) {

	tempNums := make([]int, right-left+1)
	leftIndex, rightIndex := left, mid+1
	var tempNumsIndex = 0

	for ; leftIndex <= mid && rightIndex <= right; tempNumsIndex++ {
		if nums[leftIndex] <= nums[rightIndex] {
			tempNums[tempNumsIndex] = nums[leftIndex]
			leftIndex++
		} else {
			tempNums[tempNumsIndex] = nums[rightIndex]
			rightIndex++
		}
	}

	var (
		start = leftIndex
		end   = mid
	)

	if rightIndex < right {
		start = rightIndex
		end = right
	}

	for start <= end {
		tempNums[tempNumsIndex] = nums[start]
		start++
		tempNumsIndex++
	}

	//for leftIndex <= mid {
	//	tempNums[tempNumsIndex] = nums[leftIndex]
	//	tempNumsIndex++
	//	leftIndex++
	//}
	//
	//for rightIndex <= right {
	//	tempNums[tempNumsIndex] = nums[rightIndex]
	//	tempNumsIndex++
	//	rightIndex++
	//}

	copy(nums[left:right+1], tempNums)
}
