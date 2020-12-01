package _9_单词搜索

import "fmt"

func exist(board [][]byte, word string) bool {

	x := len(board)
	y := len(board[0])

	var boardShadow = make([][]bool, x)
	for i := range board {
		boardShadow[i] = make([]bool, y)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if check(board, boardShadow, i, j, word, 0) {
				return true
			}
		}
	}

	return false
}

func check(board [][]byte, boardShadow [][]bool, x, y int, word string, wordIdx int) bool {
	if boardShadow[x][y] {
		return false
	}

	if board[x][y] != word[wordIdx] {
		return false
	}

	fmt.Println(word, x, y)
	// 匹配到最后
	if wordIdx == len(word)-1 {
		return true
	}

	defer func() {
		// 记录被匹配过
		boardShadow[x][y] = true
	}()
	// top
	if x-1 >= 0 {
		return check(board, boardShadow, x-1, y, word, wordIdx+1)
	}

	// bottom
	if x+1 < len(board) {
		return check(board, boardShadow, x+1, y, word, wordIdx+1)
	}

	// left
	if y-1 > 0 {
		return check(board, boardShadow, x, y-1, word, wordIdx+1)
	}

	// right
	if y+1 < len(board[0]) {
		return check(board, boardShadow, x, y-1, word, wordIdx+1)
	}

	return false
}
