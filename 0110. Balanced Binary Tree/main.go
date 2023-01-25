package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 如果不平衡，返回-1，否则返回高度
func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := getHeight(node.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := getHeight(node.Right)
	if rightHeight == -1 {
		return -1
	}
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return -1
	} else {
		return max(leftHeight, rightHeight) + 1
	}
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return getHeight(root) != -1
}

func main() {

}
