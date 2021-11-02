package main

import "fmt"

func maxProfit(prices []int) int {
	prevMin := 0
	max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] - prices[prevMin] > max {
			max = prices[i] - prices[prevMin]
		}
		if prices[i] < prices[prevMin]{
			prevMin = i
		}
	}
	return max
}

func main() {
	fmt.Println(maxProfit([]int{7}))
}
