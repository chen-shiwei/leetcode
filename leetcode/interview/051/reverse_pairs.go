package _051

func reversePairs(nums []int) int {

	l := len(nums)
	if l < 2 {
		return 0
	}

	return split(nums, 0, l-1)

}

func split(nums []int, left, right int) int {
	if left >= right {
		return 0
	}
	mid := left + (right-left)/2

	c := split(nums, left, mid) + split(nums, mid+1, right)

	//temp := make([]int, right-left+1)
	var (
		temp []int
		i    = left
		j    = mid + 1
		//tempIndex  int
	)
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			//temp[tempIndex] = nums[i]
			temp = append(temp, nums[i])
			i++
			c += j - (mid + 1)
		} else {
			//temp[tempIndex] = nums[j]
			temp = append(temp, nums[j])
			j++
		}
	}

	for i <= mid {
		//temp[tempIndex] = nums[i]
		temp = append(temp, nums[i])
		i++
		c += right - (mid + 1) + 1
	}

	for j <= right {
		//temp[tempIndex] = nums[j]
		temp = append(temp, nums[j])

		//tempIndex++
		j++
	}

	copy(nums[left:right+1], temp)

	return c

	//return c + merge(nums, left, mid, right)
}

func merge(nums []int, left, mid, right int) int {

	temp := make([]int, right-left+1)

	var (
		leftIndex  = left
		rightIndex = mid + 1
		tempIndex  int
		c          int
	)

	for ; leftIndex <= mid && rightIndex <= right; tempIndex++ {
		if nums[leftIndex] <= nums[rightIndex] {
			temp[tempIndex] = nums[leftIndex]
			leftIndex++
			c += rightIndex - (mid + 1)
		} else {
			temp[tempIndex] = nums[rightIndex]
			rightIndex++
		}
	}

	for leftIndex <= mid {
		temp[tempIndex] = nums[leftIndex]
		tempIndex++
		leftIndex++
		c += right - (mid + 1) + 1
	}

	for rightIndex <= right {
		temp[tempIndex] = nums[rightIndex]
		tempIndex++
		rightIndex++
	}

	copy(nums[left:right+1], temp)

	for i := left; i <= right; i++ {
		nums[i] = temp[i-left]
	}

	return c

}

func dfs(nums []int, i, l int, c *int) {
	if i > l {
		return
	}

	if nums[i] > nums[i-1] {
		*c += 1
	}

	dfs(nums, i+1, l, c)
}
