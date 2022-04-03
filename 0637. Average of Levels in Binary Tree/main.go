package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	var result []float64
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		sum, count := 0, 0
		for i := 0; i < length; i++ {
			cur := queue[0]
			sum += cur.Val
			count += 1
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		fmt.Println(sum, count)
		result = append(result, float64(sum)/float64(count))
	}
	return result
}

func main() {

}
