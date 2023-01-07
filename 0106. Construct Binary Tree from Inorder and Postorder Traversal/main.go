package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{postorder[len(postorder)-1], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			break
		}
	}
	root.Left = buildTree(inorder[:i], postorder[:len(inorder[:i])])
	root.Right = buildTree(inorder[i+1:], postorder[len(inorder[:i]):len(postorder)-1])
	return root
}

func main() {

}
