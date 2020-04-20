package _4

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	xMax := len(matrix[0])
	yMax := len(matrix)
	// 起点
	x, y := xMax-1, 0
	var exist bool
	for x >= 0 && y < yMax {
		fmt.Println(y, x)
		if matrix[y][x] == target {
			exist = true
			break
		} else if matrix[y][x] > target {
			x--
		} else if matrix[y][x] < target {
			y++
		}
	}
	if exist {
		return true
	}
	return false

}
