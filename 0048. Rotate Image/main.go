package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)
	// 先转置
	for i, _ := range matrix {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 对于每行，左右颠倒
	for i, _ := range matrix {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func main() {
	rotate([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}})
}
