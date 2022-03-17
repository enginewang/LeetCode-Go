package main

// 动态规划？
func maximalRectangle(matrix [][]int) int {
	// dp存的是
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			dp[i][j] = -1
		}
	}
	//for _, i := range dp {
	//	fmt.Println(i)
	//}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {

		}
	}
	return 0
}

func main() {
	maximalRectangle([][]int{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0},
	})
}
