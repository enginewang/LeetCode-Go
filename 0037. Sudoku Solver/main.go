package main

func solveSudoku(board [][]byte) {
	var backtrack func([][]byte) bool
	backtrack = func(board [][]byte) bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] == '.' {
					for num := '1'; num <= '9'; num++ {
						if isAvailable(board, i, j, byte(num)) {
							board[i][j] = byte(num)
							if backtrack(board) {
								return true
							}
							board[i][j] = '.'
						}
					}
					return false
				} else {
					continue
				}
			}
		}
		return true
	}
	backtrack(board)
}

func isAvailable(board [][]byte, row int, col int, number byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == number {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if board[j][col] == number {
			return false
		}
	}
	for r := row / 3 * 3; r < row/3*3+3; r++ {
		for c := col / 3 * 3; c < col/3*3+3; c++ {
			if board[r][c] == number {
				return false
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
}
