package _9_螺旋矩阵II

func generateMatrix(n int) [][]int {

	result := make([][]int, n)

	for i := range result {
		result[i] = make([]int, n)
	}

	var start, end = 1, n * n
	var (
		left   = 0
		right  = n - 1
		top    = 0
		bottom = n - 1
	)
	// left -> right
	// top -> bottom
	// right -> left
	// bottom -> top

	for start <= end {
		for i := left; i <= right; i++ {
			result[top][i] = start
			start++
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			result[i][right] = start
			start++
		}
		right--
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			result[bottom][i] = start
			start++
		}
		bottom--
		if bottom < top {
			break
		}
		for i := bottom; i >= top; i-- {
			result[i][left] = start
			start++
		}
		left++
		if left > right {
			break
		}
	}

	return result
}
