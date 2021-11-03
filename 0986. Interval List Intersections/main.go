package main

import "fmt"

// 双指针，不过头尾是列表
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	start1, start2 := 0, 0
	var result [][]int
	// 一条到最后，结束
	for start1 != len(firstList) && start2 != len(secondList) {
		intersect, id := intersectList(firstList[start1], secondList[start2])
		if intersect != nil {
			result = append(result, intersect)
		}

		if id == 0 {
			start1++
		} else {
			start2++
		}
	}
	return result
}

// 辅助函数，返回两个列表的交集，以及哪个列表的结尾在后面，第二个参数为0表示第一个尾巴在前，移动第一个，为1表示移动第二个
func intersectList(firstList []int, secondList []int) ([]int, int) {
	// 没有重叠情况
	if firstList[1] < secondList[0] {
		return nil, 0
	}
	if firstList[0] > secondList[1] {
		return nil, 1
	}
	// 有重叠情况
	if firstList[0] < secondList[0] {
		if firstList[1] < secondList[1] {
			return []int{secondList[0], firstList[1]}, 0
		} else {
			return secondList, 1
		}
	} else {
		if firstList[1] > secondList[1] {
			return []int{firstList[0], secondList[1]}, 1
		} else {
			return firstList, 0
		}
	}
}

func main() {
	fmt.Println(intervalIntersection([][]int{
	}, [][]int{
		{1, 5},
		{8, 12},
		{15, 24},
		{25, 26},
	},
	))
}
