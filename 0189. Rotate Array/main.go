package main

import "fmt"


// 1. 用另一个数组append，比较笨
func rotate(nums []int, k int)  {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}
// 2. 数组翻转，数组翻转比较容易，直接从头尾往中间走，两个分布交换即可
// 翻转nums，然后翻转nums[:k]，然后饭庄nums[k:]即可。

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}
