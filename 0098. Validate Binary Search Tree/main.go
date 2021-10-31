package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return valid(root, nil, nil)
}

func valid(root *TreeNode, min *int, max *int) bool {
	if root == nil {
		return true
	}
	if (max != nil && root.Val >= *max) || (min != nil && root.Val <= *min) {
		return false
	}
	return valid(root.Left, min, &root.Val) && valid(root.Right, &root.Val, max)
}

func main() {

}
