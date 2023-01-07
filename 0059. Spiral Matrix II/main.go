package main

/*
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
*/

import "fmt"

func generateMatrix(n int) [][]int {
	// 先建立一个二维矩阵，开辟空间
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	// x，y是当前的位置（y是行，x是列），width，height是待完成的矩形的长宽
	// 初始如下：
	x, y, width, height := 0, 0, n, n
	num := 1
	// 结束条件就是，待完成的矩形长宽都为0
	for width > 0 && height > 0 {
		// 这里是最上面一行，从左往右
		if height > 0 {
			for i := x; i < x+width; i++ {
				matrix[y][i] = num
				num += 1
			}
			// 剩余矩形少了一行，height-1
			height -= 1
			// y=0的一行没了，所以y+1
			y += 1
		}
		// 最右边一行，从上往下
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
	generateMatrix(10)
}
