package _52_用最少数量的箭引爆气球

import "sort"

func findMinArrowShots(points [][]int) int {

	l := len(points)
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	result := 1
	for i := 1; i < l; i++ {
		// 不沾边
		if points[i][0] > points[i-1][1] {
			result++
		} else {
			points[i][1] = min(points[i][1], points[i-1][1])
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
