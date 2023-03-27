package main

func combinationSum4(nums []int, target int) int {
	// 定义一个dp
	// dp[j]表示target为j时的组合的个数
	// 先初始化
	dp := make([]int, target+1)
	dp[0] = 1
	// 其实是排列，所以遍历顺序颠倒
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}

func main() {

}
