package main

import "fmt"

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	islandCount := 0
	n, m := len(grid1), len(grid1[0])
	var dfs func(grid [][]int, x int, y int, first bool, valid *bool)
	dfs = func(grid [][]int, x int, y int, first bool, valid *bool) {
		//if !*valid {
		//	return
		//}
		if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
			return
		}
		if grid1[x][y] == 0 {
			*valid = false
			//return
		}
		grid[x][y] = 0
		dfs(grid, x+1, y, false, valid)
		dfs(grid, x-1, y, false, valid)
		dfs(grid, x, y-1, false, valid)
		dfs(grid, x, y+1, false, valid)
	}
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if grid2[x][y] == 1 {
				valid := true
				dfs(grid2, x, y, true, &valid)
				if valid {
					islandCount++
				}
			}
		}
	}
	return islandCount
}

func main() {
	fmt.Println(countSubIslands([][]int{
		{1, 0, 1, 0, 1},
		{1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{1, 0, 1, 0, 1}},
		[][]int{
			{0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1},
			{0, 1, 0, 1, 0},
			{0, 1, 0, 1, 0},
			{1, 0, 0, 0, 1}},
	))
}
