package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val, Left: nil, Right: nil}
		return root
	}
	// 插入右子树
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
		return root
	} else {
		root.Left = insertIntoBST(root.Left, val)
		return root
	}
}

func main() {

}
