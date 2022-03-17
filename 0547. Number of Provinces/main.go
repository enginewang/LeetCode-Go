package main

import (
	"fmt"
)

func findCircleNum(isConnected [][]int) int {
	if len(isConnected) == 0 {
		return 0
	}
	// 记录是否被访问
	isVisited := make([]bool, len(isConnected))
	var dfs func(int)
	// 传入的是省份id
	dfs = func(i int) {
		isVisited[i] = true
		for j := 0; j < len(isConnected); j++{
			if isConnected[i][j] == 1 && !isVisited[j]{
				dfs(j)
			}
		}
	}
	num := 0
	for i := 0; i < len(isConnected); i++ {
		if isVisited[i] {
			continue
		}
		num++
		dfs(i)
	}
	return num
}

func main() {
	isConnected := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(findCircleNum(isConnected))
}
