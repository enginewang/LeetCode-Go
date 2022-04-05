package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func flatten(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	flatten(root.Left)
//	flatten(root.Right)
//	l := root.Left
//	r := root.Right
//	if l == nil {
//		return
//	}
//	root.Left = nil
//	root.Right = l
//	for l.Right != nil {
//		l = l.Right
//	}
//	l.Right = r
//}

func flatten(root *TreeNode) {
	helper(root)
}

// 返回最后第一个节点
func helper(root *TreeNode) *TreeNode {
	// 可以处理的情况
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	if root.Left == nil {
		return helper(root.Right)
	}
	if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		return helper(root.Right)
	}
	// 左右子树都flatten
	leftLast := helper(root.Left)
	rightLast := helper(root.Right)
	leftLast.Right = root.Right
	root.Right = root.Left
	root.Left = nil
	return rightLast
}

func main() {

}
