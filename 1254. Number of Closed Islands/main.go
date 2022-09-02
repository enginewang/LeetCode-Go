package main

func closedIsland(grid [][]int) int {
	closedIslandCount := 0
	n, m := len(grid), len(grid[0])
	var dfs func(grid [][]int, x int, y int)
	dfs = func(grid [][]int, x int, y int) {
		if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 1 {
			return
		}
		grid[x][y] = 1
		dfs(grid, x+1, y)
		dfs(grid, x-1, y)
		dfs(grid, x, y+1)
		dfs(grid, x, y-1)
	}
	for y := 0; y < m; y++ {
		dfs(grid, 0, y)
		dfs(grid, n-1, y)
	}
	for x := 1; x < n-1; x++ {
		dfs(grid, x, 0)
		dfs(grid, x, m-1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				dfs(grid, i, j)
				closedIslandCount += 1
			}
		}
	}
	return closedIslandCount
}

func main() {

}
