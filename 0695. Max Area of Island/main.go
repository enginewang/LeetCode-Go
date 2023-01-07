package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}
	curArea, maxArea := 0, 0
	var dfs func(grid [][]int, x int, y int, find bool)
	dfs = func(grid [][]int, x int, y int, find bool) {
		if find {
			curArea = 0
		}
		if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
			return
		}
		curArea += 1
		if curArea > maxArea {
			maxArea = curArea
		}
		grid[x][y] = 0
		for _, direction := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			newX := x + direction[0]
			newY := y + direction[1]
			dfs(grid, newX, newY, false)
		}
	}
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if grid[x][y] == 1 {
				dfs(grid, x, y, true)
			}
		}
	}
	return maxArea
}

func main() {
	fmt.Println(maxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}))
}
