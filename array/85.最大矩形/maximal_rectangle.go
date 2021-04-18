package _5_最大矩形

func maximalRectangle(matrix [][]byte) int {

	x := len(matrix)

	y := len(matrix[0])
	if x == 0 {
		return 0
	}

	var area int
	var heights = make([]int, y)
	for row := 0; row < x; row++ {
		for col := 0; col < y; col++ {
			if matrix[row][col] == '1' {
				heights[col] += 1
			} else {
				heights[col] = 0
			}
		}

		curArea := largestRectangleArea(heights)
		if curArea > area {
			area = curArea
		}
	}

	return area

}

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
