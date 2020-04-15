package _542

func updateMatrix(matrix [][]int) [][]int {
	for i, nums := range matrix {
		for j, num := range nums {
			if num == 1 {
				matrix[i][j] = breadthFirstSearch(matrix, i, j)
			}
		}

	}
	return matrix
}

func breadthFirstSearch(matrix [][]int, i, j int) int {

}
