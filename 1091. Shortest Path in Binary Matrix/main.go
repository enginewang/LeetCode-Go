package main

import (
	"fmt"
)

type Grid struct {
	x int
	y int
}

type GridWithDepth struct {
	x     int
	y     int
	Depth int
}

// 广度优先，没什么好说的
func shortestPathBinaryMatrix(grid [][]int) int {
	if len(grid) == 0 || grid[0][0] != 0 || grid[len(grid)-1][len(grid)-1] != 0{
		return -1
	}
	director := []Grid{
		{1, 1},
		{0, 1},
		{1, 0},
		{-1, 1},
		{1, -1},
		{0, -1},
		{-1, 0},
		{-1, -1},
	}
	// 记录是否访问
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid))
	}
	// 初始化队列
	var queue []GridWithDepth
	queue = append(queue, GridWithDepth{x: 0, y: 0, Depth: 1})
	// 结束条件
	for len(queue) > 0 {
		head := queue[0]
		if head.x == len(grid)-1 && head.y == len(grid)-1{
			return head.Depth
		}
		// 等于true表示之前访问过了，那这次访问没有意义
		visited[head.x][head.y] = true
		for _, d := range director {
			newGrid := Grid{head.x + d.x, head.y + d.y}
			if newGrid.x == len(grid)-1 && newGrid.y == len(grid)-1{
				return head.Depth+1
			}
			if valid(newGrid.x, newGrid.y, len(grid)) && grid[newGrid.x][newGrid.y] == 0 && visited[newGrid.x][newGrid.y] != true {
				queue = append(queue, GridWithDepth{newGrid.x, newGrid.y, head.Depth + 1})
				// 这里直接把新加的置为1，下次不走了
				grid[newGrid.x][newGrid.y] = 1
			}
		}
		queue = queue[1:]
	}
	if len(queue) > 0 {
		return queue[0].Depth
	} else {
		return -1
	}
}

func valid(x int, y int, size int) bool {
	if x >= 0 && y >= 0 && x < size && y < size {
		return true
	} else {
		return false
	}
}

func main() {
	isConnected := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}
	fmt.Println(shortestPathBinaryMatrix(isConnected))
}
