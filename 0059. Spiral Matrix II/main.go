package main

import "fmt"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix{
		matrix[i] = make([]int, n)
	}
	x, y, width, height := 0, 0, n, n
	num := 1
	for width > 0 && height > 0 {
		if height > 0 {
			for i := x; i < x+width; i++ {
				matrix[y][i] = num
				num += 1
			}
			height -= 1
			y += 1
		}
		if width > 0 {
			for i := y; i < y+height; i++ {
				matrix[i][x+width-1] = num
				num += 1
			}
			width -= 1
		}
		if height > 0 {
			for i := x + width - 1; i >= x; i-- {
				matrix[y+height-1][i] = num
				num += 1
			}
			height -= 1
		}
		if width > 0 {
			for i := y + height - 1; i >= y; i-- {
				matrix[i][x] = num
				num += 1
			}
			width -= 1
			x += 1
		}
	}
	fmt.Println(matrix)
	return matrix
}

func main() {
	generateMatrix(3)
}
