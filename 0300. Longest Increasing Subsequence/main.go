package main

func lengthOfLIS(nums []int) int {
	// memo[n]表示包含n个元素的最小递增子序列的长度
	dp := make([]int, len(nums))
	max := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		res := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				this := dp[j] + 1
				if this > res {
					res = this
				}
				dp[i] = res
				if res > max {
					max = res
				}
			}
		}
	}
	return max
}

func main() {
	lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})
}
