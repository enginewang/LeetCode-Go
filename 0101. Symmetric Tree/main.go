package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return twoTreeIsSymmetric(root.Left, root.Right)
}

func twoTreeIsSymmetric(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if (a == nil && b != nil) || (a != nil && b == nil) || a.Val != b.Val {
		return false
	}
	if twoTreeIsSymmetric(a.Left, b.Right) && twoTreeIsSymmetric(a.Right, b.Left) {
		return true
	}
	return false
}

func main() {

}
