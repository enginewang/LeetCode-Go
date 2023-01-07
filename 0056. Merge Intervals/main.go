package main

import (
	"fmt"
	"sort"
)

type myArr [][]int

func (b myArr) Len() int {
	return len(b)
}

func (b myArr) Less(i int, j int) bool {
	if b[i][0]-b[j][0] > 0 {
		return false
	} else {
		return true
	}
}

func (b myArr) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}

func merge(intervals [][]int) [][]int {
	var arr myArr = intervals
	var result [][]int
	sort.Sort(arr)
	// 区间部分最大值
	partialMax := arr[0][1]
	// 区间部分最小值
	partialMin := arr[0][0]
	for _, item := range arr {
		// 也就是不能合并，前一个写入结果
		if item[0] > partialMax {
			result = append(result, []int{partialMin, partialMax})
			partialMax = item[1]
			partialMin = item[0]
		} else if item[1] > partialMax {
			partialMax = item[1]
		}
	}
	result = append(result, []int{partialMin, partialMax})
	return result
}

func main() {
	fmt.Println(merge([][]int{{1, 10}, {5, 6}, {8, 9}, {11, 12}}))
}
