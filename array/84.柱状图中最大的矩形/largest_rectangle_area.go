package _4_柱状图中最大的矩形

func largestRectangleArea(heights []int) int {
	l := len(heights)
	var area int
	for i := 0; i < l; i++ {
		var left = i
		var curHeight = heights[i]
		for left > 0 && heights[left-1] >= curHeight {
			left--
		}

		var right = i
		for right < l-1 && heights[right+1] >= curHeight {
			right++
		}
		width := right - left + 1
		curArea := width * curHeight
		if curArea > area {
			area = curArea
		}
	}
	return area
}
