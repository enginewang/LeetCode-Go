package main

import (
	"fmt"
	"sort"
)

func advantageCount(nums1 []int, nums2 []int) []int {
	nums2Map, nums2SortedMap := make(map[int][]int), make(map[int]int)
	sortedNum1, sortedNum2 := make([]int, len(nums1)), make([]int, len(nums2))
	var newList []int
	copy(sortedNum1, nums1)
	copy(sortedNum2, nums2)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedNum1)))
	sort.Sort(sort.Reverse(sort.IntSlice(sortedNum2)))
	for i, v := range nums2 {
		if _, ok := nums2Map[v]; ok {
			nums2Map[v] = append(nums2Map[v], i)
		} else {
			nums2Map[v] = []int{i}
		}
	}
	for i, v := range sortedNum2 {
		nums2SortedMap[i] = nums2Map[v][0]
		nums2Map[v] = nums2Map[v][1:]
	}
	//fmt.Println(nums2)
	fmt.Println(nums2SortedMap)
	left, right := 0, len(sortedNum1)-1
	//fmt.Println(sortedNum1)
	//fmt.Println(sortedNum2)
	for i, _ := range sortedNum2 {
		if sortedNum2[i] < sortedNum1[left] {
			newList = append(newList, sortedNum1[left])
			left++
		} else {
			newList = append(newList, sortedNum1[right])
			right--
		}
	}
	fmt.Println(newList)
	for k, v := range nums2SortedMap {
		nums1[v] = newList[k]
	}
	fmt.Println(nums1)
	return nums1
}

func main() {
	advantageCount([]int{12, 24, 10, 8, 32}, []int{13, 25, 25, 32, 11})
}
