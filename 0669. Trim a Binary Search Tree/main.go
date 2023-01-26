package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	// 否则root就是介于两者之间
	if root.Left != nil && root.Left.Val < low {
		root.Left = trimBST(root.Left.Right, low, high)
	}
	if root.Right != nil && root.Right.Val > high {
		root.Right = trimBST(root.Right.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

func main() {

}
