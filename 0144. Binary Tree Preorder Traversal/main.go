package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{root.Val}
	left := preorderTraversal(root.Left)
	result = append(result, left...)
	right := preorderTraversal(root.Right)
	result = append(result, right...)
	return result
}

func main() {

}
