package _056

// https://leetcode-cn.com/problems/merge-intervals/

// Merge 合并区间
func Merge(intervals [][]int) [][]int {

	quickSort(intervals)

	l := len(intervals)

	var mergedNums [][]int

	var i = 0
	for i < l {
		left := intervals[i][0]
		right := intervals[i][1]

		for i < l-1 && intervals[i+1][0] <= right {
			right = max(right, intervals[i+1][1])
			i++
		}
		mergedNums = append(mergedNums, []int{left, right})
		i++

	}

	return mergedNums

}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func bubbleSort(nums [][]int) {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		isChanged := false
		for j := 0; j < l-1-i; j++ {
			if nums[j][0] > nums[j+1][0] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isChanged = true
			}
		}
		if !isChanged {
			break
		}
	}
}

func quickSort(nums [][]int) {
	l := len(nums)
	if len(nums) < 1 {
		return
	}
	mid := nums[0][0]

	var head, tail = 0, l - 1
	for i := 1; i <= tail; {
		if nums[i][0] <= mid {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		} else {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		}
	}

	quickSort(nums[:head])

	quickSort(nums[head+1:])

}
