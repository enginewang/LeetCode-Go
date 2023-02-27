package main

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	memo := make([]int, n+1)
	memo[0] = 1
	memo[1] = 1
	memo[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			memo[i] = max(memo[i], max((i-j)*j, memo[i-j]*j))
		}
	}
	return memo[n]
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {

}
