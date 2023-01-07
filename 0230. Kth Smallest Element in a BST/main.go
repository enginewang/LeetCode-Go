package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	result := 0
	traverse(root, &count, k, &result)
	return result
}

func traverse(root *TreeNode, count *int, k int, result *int) {
	if root == nil {
		return
	}
	traverse(root.Left, count, k, result)
	*count += 1
	if *count == k {
		*result = root.Val
		return
	}
	traverse(root.Right, count, k, result)
	return
}

func main() {

}
