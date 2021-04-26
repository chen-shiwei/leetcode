package _54_螺旋矩阵

func spiralOrder(matrix [][]int) []int {
	var slides = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var slidesIdx = 0
	rows := len(matrix)
	cols := len(matrix[0])
	var i, j int
	visited := make([][]bool, rows)

	for index := range visited {
		visited[index] = make([]bool, cols)
	}
	var result []int

	for start, end := 1, rows*cols; start <= end; start++ {
		result = append(result, matrix[i][j])
		visited[i][j] = true
		nextI, nextJ := slides[slidesIdx][0]+i, slides[slidesIdx][1]+j
		if nextI < 0 || nextJ < 0 || nextI >= rows || nextJ >= cols || visited[nextI][nextJ] {
			slidesIdx = (slidesIdx + 1) % 4
		}
		i, j = slides[slidesIdx][0]+i, slides[slidesIdx][1]+j
	}

	return result

}
func spiralOrder1(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])
	left, right, top, bottom := 0, cols-1, 0, rows-1
	end := rows * cols
	var result []int
	for start := 1; start <= end; {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
			start++
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
			start++
		}
		right--
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
			start++
		}
		bottom--
		if bottom < top {
			break
		}
		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
			start++
		}
		left++
		if left > right {
			break
		}
	}

	return result

}
