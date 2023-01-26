package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	minDelta := math.MaxInt
	traverse(root, &minDelta, &[]int{})
	return minDelta
}

func traverse(root *TreeNode, min *int, result *[]int) {
	if root == nil {
		return
	}
	traverse(root.Left, min, result)
	if len(*result) > 0 && root.Val-(*result)[len(*result)-1] < *min {
		*min = root.Val - (*result)[len(*result)-1]
	}
	*result = append(*result, root.Val)
	traverse(root.Right, min, result)
}

func main() {

}
