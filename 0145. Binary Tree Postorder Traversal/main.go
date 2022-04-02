package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func postorderTraversal(root *TreeNode) []int {
//	var result []int
//	pTraversal(root, &result)
//	return result
//}
//
//func pTraversal(cur *TreeNode, result *[]int) {
//	if cur == nil {
//		return
//	}
//	pTraversal(cur.Left, result)
//	pTraversal(cur.Right, result)
//	*result = append(*result, cur.Val)
//}

type MyTreeNode struct {
	Node    *TreeNode
	Visited bool
}

func postorderTraversal(root *TreeNode) []int {
	var st []*MyTreeNode
	var result []int
	if root == nil {
		return result
	}
	st = append(st, &MyTreeNode{Node: root, Visited: false})
	for len(st) > 0 {
		cur := st[len(st)-1]
		if cur.Visited || (cur.Node.Left == nil && cur.Node.Right == nil) {
			result = append(result, cur.Node.Val)
			st = st[:len(st)-1]
		} else {
			if cur.Node.Right != nil {
				st = append(st, &MyTreeNode{Node: cur.Node.Right, Visited: false})
			}
			if cur.Node.Left != nil {
				st = append(st, &MyTreeNode{Node: cur.Node.Left, Visited: false})
			}
			cur.Visited = true
		}
	}
	return result
}

func main() {

}
