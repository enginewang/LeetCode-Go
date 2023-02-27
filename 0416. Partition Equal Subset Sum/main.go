package main

import "fmt"

//func canPartition(nums []int) bool {
//	sum := 0
//	for _, i := range nums {
//		sum += i
//	}
//	if sum%2 != 0 {
//		return false
//	}
//	target := sum / 2
//	// m的含义，如果后续直接出现了这个数字，就说明是true
//	m := make(map[int]bool)
//	m[target] = true
//	for _, num := range nums {
//		//fmt.Println(m)
//		if m[num] {
//			//fmt.Println(i, num)
//			return true
//		}
//		var l []int
//		for k, _ := range m {
//			if k-num > 0 {
//				l = append(l, k-num)
//			}
//		}
//		for _, j := range l {
//			m[j] = true
//		}
//	}
//	return false
//}

func canPartition(nums []int) bool {
	// 转换为01背包问题
	// 背包容量是sum/2
	// nums就是物品列表，体积和价值都是nums[i]
	sum, n := 0, len(nums)
	for _, i := range nums {
		sum += i
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	// dp[i][j]表示从nums选取前i个数，放入容量为j的背包的最大值
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, target+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = 0
	}
	for j := 0; j < target+1; j++ {
		if j >= nums[0] {
			dp[0][j] = nums[0]
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < target+1; j++ {
			// 不可能放得下
			if nums[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
			}
		}
	}
	if dp[len(dp)-1][len(dp[0])-1] == target {
		return true
	} else {
		return false
	}
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	fmt.Println(canPartition([]int{1, 1}))
}
