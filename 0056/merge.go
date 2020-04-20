package _056

// https://leetcode-cn.com/problems/merge-intervals/

// Merge 合并区间
func Merge(intervals [][]int) [][]int {
	l := len(intervals)
	if l < 2 {
		return intervals
	}
	quickSort(intervals)

	var result = make([][]int, 0)

	for i := 0; i < l; {
		left := intervals[i][0]
		right := intervals[i][1]
		for i < l-1 && intervals[i+1][0] <= right {
			if intervals[i+1][1] > right {
				right = intervals[i+1][1]
			}
			i++
		}

		result = append(result, []int{left, right})
		i++
	}
	return result

	//quickSort(intervals)
	//
	//l := len(intervals)
	//
	//var mergedNums [][]int
	//
	//var i = 0
	//for i < l {
	//	left := intervals[i][0]
	//	right := intervals[i][1]
	//
	//	for i < l-1 && intervals[i+1][0] <= right {
	//		right = max(right, intervals[i+1][1])
	//		i++
	//	}
	//	mergedNums = append(mergedNums, []int{left, right})
	//	i++
	//
	//}
	//
	//return mergedNums
}

func quickSort(intervals [][]int) {
	l := len(intervals)
	if l < 2 {
		return
	}
	mid := intervals[0]
	head, tail := 0, l-1

	for i := 1; i <= tail; {
		if intervals[i][0] < mid[0] {
			intervals[i], intervals[head] = intervals[head], intervals[i]
			i++
			head++
		} else {
			intervals[i], intervals[tail] = intervals[tail], intervals[i]
			tail--
		}
	}

	quickSort(intervals[:head])
	quickSort(intervals[head+1:])

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
