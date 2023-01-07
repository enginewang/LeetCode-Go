package main


func spiralOrder(matrix [][]int) []int {
	var arr []int
	x, y, width, height := 0, 0, len(matrix[0]), len(matrix)
	for width > 0 && height > 0 {
		if height > 0 {
			for i := x; i < x+width; i++ {
				arr = append(arr, matrix[y][i])
			}
			height -= 1
			y += 1
		}
		if width > 0 {
			for i := y; i < y+height; i++ {
				arr = append(arr, matrix[i][x+width-1])
			}
			width -= 1
		}
		if height > 0 {
			for i := x + width - 1; i >= x; i-- {
				arr = append(arr, matrix[y+height-1][i])
			}
			height -= 1
		}
		if width > 0 {
			for i := y + height - 1; i >= y; i-- {
				arr = append(arr, matrix[i][x])
			}
			width -= 1
			x += 1
		}
	}
	return arr
}

func main() {
	spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}
