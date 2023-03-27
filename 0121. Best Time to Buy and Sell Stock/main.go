package main

import "fmt"

//func maxProfit(prices []int) int {
//	prevMin := 0
//	max := 0
//	for i := 1; i < len(prices); i++ {
//		if prices[i] - prices[prevMin] > max {
//			max = prices[i] - prices[prevMin]
//		}
//		if prices[i] < prices[prevMin]{
//			prevMin = i
//		}
//	}
//	return max
//}

func maxProfit(prices []int) int {
	// 持有股票的当前现金
	holdStock := make([]int, len(prices))
	// 未持有股票的当前现金
	noHoldStock := make([]int, len(prices))
	holdStock[0], noHoldStock[0] = -prices[0], 0
	for i := 1; i < len(prices); i++ {
		// 第i天持有股票，两种情况
		// 1. 第i-1天也持有股票
		// 2. 第i-1天未持有股票，第i天购买。那么之前一定都没有买
		holdStock[i] = max(holdStock[i-1], -prices[i])
		// 第i天未持有股票，两种情况
		// 1. 第i-1天也未持有股票
		// 2. 第i-1天持有股票，第i天售出
		noHoldStock[i] = max(noHoldStock[i-1], holdStock[i-1]+prices[i])
	}
	return noHoldStock[len(prices)-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
