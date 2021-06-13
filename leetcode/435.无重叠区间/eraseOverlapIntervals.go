package _35_无重叠区间

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	l := len(intervals)
	result := 1
	end := intervals[0][1]
	for i := 1; i < l; i++ {
		if end <= intervals[i][0] {
			result++
			end = intervals[i][1]
		}
	}
	return l - result
}
