package _8_旋转图像

import "fmt"

func rotate(matrix [][]int) {
	for _, v := range matrix {
		fmt.Printf("%+v \n", v)
	}
	fmt.Println("---")

	l := len(matrix)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for _, v := range matrix {
		fmt.Printf("%+v \n", v)
	}
	fmt.Println("---")

	for i := 0; i < l; i++ {
		for j := 0; j < l/2; j++ {
			matrix[i][j], matrix[i][l-j-1] = matrix[i][l-j-1], matrix[i][j]
		}
	}
	for _, v := range matrix {
		fmt.Printf("%+v \n", v)
	}
}
