package _1_N皇后_回溯

import "fmt"

func solveNQueens(n int) [][]string {
	var (
		paths [][]string
		path  = make([][]byte, n)
		fn    func(col, row int)
	)

	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = '.'
	}

	for i := range path {
		c := make([]byte, n)
		copy(c, b)
		path[i] = c
	}

	fn = func(col, row int) {
		if row == n {
			var tmp []string
			for _, v := range path {
				tmp = append(tmp, string(v))
			}
			paths = append(paths, tmp)
			return
		}

		for i := 0; i < n; i++ {
			if IsValid(path, row, i, n) {
				path[row][i] = 'Q'
				fn(i, row+1)
				path[row][i] = '.'
			}
		}
	}
	fn(0, 0)

	return paths
}

func IsValid(chessboard [][]byte, row, col int, n int) bool {
	// 列
	for i := 0; i < n; i++ {
		if chessboard[i][col] == 'Q' {
			return false
		}
	}
	// 左斜线
	i := row - 1
	j := col - 1
	for i >= 0 && j >= 0 {
		fmt.Println(i, j)
		if chessboard[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}

	a := row - 1
	b := col + 1

	for a >= 0 && b < n {
		if chessboard[a][b] == 'Q' {
			return false
		}
		a--
		b++
	}

	return true
}
