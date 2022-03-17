package main

import "fmt"

// 二维变一维的二分搜索
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix)==0 || len(matrix[0])==0{
		return false
	}
	i, j := 0, len(matrix[0])-1
	for j >= 0 && i < len(matrix){
		if matrix[i][j] == target{
			return true
		}
		if matrix[i][j] > target{
			j--
		} else {
			i++
		}
	}
	return false
}

func main() {
	var m [][]int
	for i := 0; i < 3; i++{
		row := make([]int, 4)
		m = append(m, row)
	}
	for i := 0; i < 3; i++{
		for j := 0; j < 4; j++ {
			m[i][j] = 4*i+j
		}
	}
	fmt.Println(searchMatrix(m, 1))
}
