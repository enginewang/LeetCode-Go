package main

import "fmt"

// DFS，从边界的O开始，记录一下，然后最后没有在记录里的O都变成X
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	var dfs func(int, int)
	direct := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dfs = func(x int, y int) {
		//if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) || board[x][y] != 'O' {
		//	return
		//}
		if board[x][y] != 'O' {
			return
		}
		if visited[x][y] != true{
			visited[x][y] = true
		}
		for _, d := range direct{
			xNew := x + d[0]
			yNew := y + d[1]
			if xNew < 0 || yNew < 0 || xNew >= len(board) || yNew >= len(board[0]) || board[xNew][yNew] != 'O' || visited[xNew][yNew] == true{
				continue
			} else {
				dfs(xNew, yNew)
			}
		}
	}
	for j := 0; j < len(board[0]); j++{
		dfs(0, j)
		dfs(len(board)-1, j)
	}
	for i := 1; i < len(board)-1; i++{
		dfs(i, 0)
		dfs(i, len(board[0])-1)
	}
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			if board[x][y] == 'O' && visited[x][y] != true{
				board[x][y] = 'X'
			}
		}
	}
	fmt.Println(board)
}

func main() {
	grid := [][]byte{
		{'X', 'X', 'X', 'X', 'X'},
		{'X', 'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X', 'X'},
		{'X', 'O', 'X', 'X', 'X'},
	}
	//grid := [][]byte{
	//	{'O', 'O', 'O'},
	//	{'O', 'O', 'O'},
	//	{'O', 'O', 'O'},
	//}
	solve(grid)
}
