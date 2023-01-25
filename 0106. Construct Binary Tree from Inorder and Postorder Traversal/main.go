package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	// 如果为空，直接返回nil
	if len(postorder) == 0 {
		return nil
	}
	// 后序遍历的最后一个肯定是当前的根节点，建立这个节点
	root := &TreeNode{postorder[len(postorder)-1], nil, nil}
	i := 0
	// 从中序遍历中找到根节点的位置，前后划分就是两段
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
