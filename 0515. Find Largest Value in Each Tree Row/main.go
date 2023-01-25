package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var queue []*TreeNode
	var result []int
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		rowMax := -math.MaxInt64
		for i := 0; i < length; i++ {
			if queue[0].Val > rowMax {
				rowMax = queue[0].Val
			}
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		result = append(result, rowMax)
	}
	return result
}

func main() {

}
