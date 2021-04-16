package _47_前K个高频元素

func topKFrequent(nums []int, k int) []int {
	var dict = make(map[int]int)

	for _, v := range nums {
		dict[v]++
	}

	var freqNums [][]int

	for key, value := range dict {
		freqNums = append(freqNums, []int{key, value})
	}

	QSort(freqNums)
	var items []int
	for _, item := range freqNums[:k] {
		items = append(items, item[0])
	}
	return items
}

func QSort(nums [][]int) {
	if len(nums) == 0 {
		return
	}
	selected := nums[0][1]
	var left, right = 0, len(nums) - 1

	for i := 1; i <= right; {
		if nums[i][1] >= selected {
			nums[left], nums[i] = nums[i], nums[left]
			left++
			i++
		} else {
			nums[right], nums[i] = nums[i], nums[right]
			right--
		}
	}

	QSort(nums[:left])
	QSort(nums[left+1:])

}
