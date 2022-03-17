package main
// hard

import "fmt"

func superEggDrop(K int, N int) int {
	dp, step := make([]int, K+1), 0
	for ; dp[K] < N; step++ {
		for i := K; i > 0; i-- {
			dp[i] = (1 + dp[i] + dp[i-1])
		}
	}
	return step
}
func main() {
	fmt.Println(superEggDrop(3, 25))
}
