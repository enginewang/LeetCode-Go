package main

import (
	"fmt"
)

var memo map[int] bool

func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, i := range nums {
		sum += i
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	used := 0
	memo = make(map[int] bool)
	return backtrack(k, 0, nums, 0, used, target)
}


// k号桶思考是否要把nums[index]元素装进来。used表示这个元素是否已经被装到其他桶了
// 回溯会遍历所有的可能，如果凑不出的话，可能会出现重复遍历，只是桶换位子的情况。
// 需要在装满一个桶的时候存储used数组，下次再遇到的话就剪枝，用位图表示节约存储提高效率
func backtrack(k int, bucket int, nums []int, start int, used int, target int) bool {
	// 所有桶都装满了
	if k == 0 {
		return true
	}
	// 当前桶装满了，递归下一个桶的选择
	if bucket == target {
		res := backtrack(k-1, 0, nums, 0, used, target)
		memo[used] = res
		return res
	}
	if res, ok := memo[used]; ok {
		return res
	}
	// 从start开始往后找能装的元素
	for i := start; i < len(nums); i++ {
		// 跳过已经被装到其他桶的元素
		if (used >> i) & 1 == 1 {
			continue
		}
		// 跳过放进来会超过限制的元素
		if bucket+nums[i] > target {
			continue
		}
		// 装入当前元素
		used |= 1 << i
		bucket += nums[i]
		// 递归下一个
		if backtrack(k, bucket, nums, i+1, used, target) {
			return true
		}
		// 回溯
		used ^= 1 << i
		bucket -= nums[i]
	}
	return false
}

func main() {
	fmt.Println(canPartitionKSubsets([]int{10,1,10,9,6,1,9,5,9,10,7,8,5,2,10,8}, 11))
}
