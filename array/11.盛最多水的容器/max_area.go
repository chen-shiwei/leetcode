package _1_盛最多水的容器

import (
	"math"
)

func maxArea(height []int) int {
	var start, end = 0, len(height) - 1
	var area int
	for start <= end {
		var x = height[start]
		var y = end - start
		if height[start] > height[end] {
			x = height[end]
			end--
		} else {
			start++
		}
		area = int(math.Max(float64(y*x), float64(area)))
	}
	return area
}
