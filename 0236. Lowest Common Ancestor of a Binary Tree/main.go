package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 需要通过后序遍历
// p q 只是作为信息传递并不改变
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 如果遇到p或者q，就直接返回
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 如果左右都不是nil，也就是说明当前节点左边找到了一个，右边找到了一个，返回当前节点
	if left != nil && right != nil {
		return root
	}
	// 如果左边返回的不是nil，但是右边是nil，说明要从左边找
	if left != nil {
		return left
	}
	return right
}

func main() {
	Node0 := &TreeNode{Val: 0, Left: nil, Right: nil}
	Node1 := &TreeNode{Val: 1, Left: nil, Right: nil}
	Node2 := &TreeNode{Val: 2, Left: nil, Right: nil}
	Node3 := &TreeNode{Val: 3, Left: nil, Right: nil}
	Node4 := &TreeNode{Val: 4, Left: nil, Right: nil}
	Node5 := &TreeNode{Val: 5, Left: nil, Right: nil}
	Node6 := &TreeNode{Val: 6, Left: nil, Right: nil}
	Node7 := &TreeNode{Val: 7, Left: nil, Right: nil}
	Node8 := &TreeNode{Val: 8, Left: nil, Right: nil}
	Node3.Left = Node5
	Node3.Right = Node1
	Node5.Left = Node6
	Node5.Right = Node2
	Node2.Left = Node7
	Node2.Right = Node4
	Node1.Left = Node0
	Node1.Right = Node8
	lowestCommonAncestor(Node3, Node6, Node4)
}
