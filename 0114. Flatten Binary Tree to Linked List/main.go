package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	l := root.Left
	r := root.Right
	if l == nil {
		return
	}
	root.Left = nil
	root.Right = l
	for l.Right != nil {
		l = l.Right
	}
	l.Right = r
}

func main() {

}
