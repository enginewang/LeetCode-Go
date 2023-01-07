package main

/*
可以先遍历一个数组，生成一张哈希表map[int]bool
然后遍历第二个数组，如果map[k]=true就说明有相同元素，加到结果中，并设置map[k]=false。
*/
import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	hashMap := make(map[int]bool)
	var result []int
	for _, i := range nums1 {
		hashMap[i] = true
	}
	for _, j := range nums2 {
		if hashMap[j] {
			result = append(result, j)
			hashMap[j] = false
		}
	}
	return result
}

func main() {
	fmt.Println(intersection([]int{4, 5, 6, 6}, []int{4, 6, 6, 7}))
}
