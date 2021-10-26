package main

import "fmt"

// 二维的求和数组

type NumMatrix struct {
	Matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	var numMatrix NumMatrix
	numMatrix.Matrix = make([][]int, len(matrix)+1)
	for i := range numMatrix.Matrix {
		numMatrix.Matrix[i] = make([]int, len(matrix[0])+1)
	}
	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[i-1]); j++ {
			numMatrix.Matrix[i][j] = numMatrix.Matrix[i-1][j] + numMatrix.Matrix[i][j-1] - numMatrix.Matrix[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	for i := 0; i <= len(matrix); i++ {
		fmt.Println(numMatrix.Matrix[i])
	}
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	return numMatrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.Matrix[row2+1][col2+1] - this.Matrix[row1][col2+1] - this.Matrix[row2+1][col1] + this.Matrix[row1][col1]
}

func main() {
	//obj := Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
	//fmt.Println(obj.SumRegion(2, 1, 4, 3))
	_ = Constructor([][]int{{-4, -5}})
}
