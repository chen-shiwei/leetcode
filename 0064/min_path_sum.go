package _064

func minPathSum(grid [][]int) int {

	xl := len(grid)
	yl := len(grid[0])

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}

	return grid[xl-1][yl-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
