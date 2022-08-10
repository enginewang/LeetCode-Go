package main

import "fmt"

func maxProfit(prices []int) int {
	result := 0
	for i := 1; i < len(prices); i++ {
		result += max(0, prices[i]-prices[i-1])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
