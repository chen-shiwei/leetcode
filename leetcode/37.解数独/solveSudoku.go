package _7_解数独

func solveSudoku(board [][]byte) {
	var (
		fn func() bool
	)

	fn = func() bool {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {
				if board[i][j] != '.' {
					continue
				}
				for k := '1'; k <= '9'; k++ {
					if isValid(i, j, byte(k), board) {
						board[i][j] = byte(k)
						if fn() {
							return true
						}
						board[i][j] = byte('.')
					}
				}
				return false
			}
		}
		return true
	}
	fn()
	return
}

func isValid(row, col int, v byte, board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		if board[row][i] == v {
			return false
		}
	}

	for i := 0; i < len(board[0]); i++ {
		if board[i][col] == v {
			return false
		}
	}

	var (
		// 倍数区间
		startRow = (row / 3) * 3
		startCol = (col / 3) * 3
	)
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == v {
				return false
			}
		}
	}
	return true

}
