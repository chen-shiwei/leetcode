package _200

func numIsLands(grid [][]byte) int {
	rows := len(grid)
	if rows < 1 {
		return 0
	}

	cols := len(grid[0])

	var count int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}

	return count
}

func dfs(grid [][]byte, i, j int) {
	if j >= len(grid[0]) || i >= len(grid) || i < 0 || j < 0 {
		return
	}

	if grid[i][j] == '2' || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '2'
	// top
	dfs(grid, i-1, j)
	// down
	dfs(grid, i+1, j)
	// left
	dfs(grid, i, j-1)
	// right
	dfs(grid, i, j+1)
}
