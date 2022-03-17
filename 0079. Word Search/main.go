package main

import "fmt"

var director = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			if dfs(x, y, 0, visited, board, word) == true {
				return true
			}
		}
	}
	return false
}

// count是目前是第几个字符
func dfs(x int, y int, count int, visited [][]bool, board [][]byte, word string) bool {
	if count == len(word) -1 {
		return true
	}
	// 出现这些问题，就跳过
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] != word[count] || visited[x][y] == true {
		return false
	}

	visited[x][y] = true
	for _, d := range director {
		xNew := x + d[0]
		yNew := y + d[1]
		if dfs(xNew, yNew, count+1, visited, board, word) == true {
			fmt.Println(xNew, yNew)
			return true
		}
	}
	visited[x][y] = false
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCESEEEFS"))
}
