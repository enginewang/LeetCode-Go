package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	// 初始化棋盘
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	// 记录当前每一行的各个列号
	var path []int
	// 用dict的形式也记录，方便查询
	//pathDict := make(map[int]bool)
	var result [][]string
	backtrack(0, n, path, &board, &result)
	return result
}

func boardToString(board [][]string) []string {
	var result []string
	for _, row := range board {
		rowStr := strings.Join(row, "")
		result = append(result, rowStr)
	}
	return result
}

func backtrack(row int, n int, path []int, board *[][]string, result *[][]string) {
	// 遍历到了最后
	if len(path) == n {
		boardStr := boardToString(*board)
		*result = append(*result, boardStr)
	}
	for i := 0; i < n; i++ {
		if judgeValid(*board, row, i) {
			path = append(path, i)
			(*board)[row][i] = "Q"
			backtrack(row+1, n, path, board, result)
			path = path[:len(path)-1]
			(*board)[row][i] = "."
		}
	}
}

// 在第row行的第i列是否能放置棋子
func judgeValid(board [][]string, row int, col int) bool {
	// 同一个竖线已经有了
	for r := 0; r < row; r++ {
		if board[r][col] == "Q" {
			return false
		}
	}
	for r := row - 1; r >= 0; r-- {
		upLeftCol := col - row + r
		upRightCol := col + row - r
		if upLeftCol >= 0 && board[r][upLeftCol] == "Q" {
			return false
		}
		if upRightCol <= len(board)-1 && board[r][upRightCol] == "Q" {
			return false
		}
	}
	// 斜线判断，等会写
	return true
}

func main() {
	fmt.Println(solveNQueens(4))
}
