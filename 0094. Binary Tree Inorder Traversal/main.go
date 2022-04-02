package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	pTraversal(root, &result)
	return result
}

func pTraversal(cur *TreeNode, result *[]int) {
	if cur == nil {
		return
	}
	pTraversal(cur.Left, result)
	*result = append(*result, cur.Val)
	pTraversal(cur.Right, result)
}

//type MyTreeNode struct {
//	Node    *TreeNode
//	Visited bool
//}
//
//func inorderTraversal(root *TreeNode) []int {
//	var st []*MyTreeNode
//	var result []int
//	if root == nil {
//		return result
//	}
//	st = append(st, &MyTreeNode{Node: root, Visited: false})
//	for len(st) > 0 {
//		cur := st[len(st)-1]
//		if cur.Visited || (cur.Node.Left == nil && cur.Node.Right == nil) {
//			result = append(result, cur.Node.Val)
//			st = st[:len(st)-1]
//		} else {
//			// 先把自己取出来
//			st = st[:len(st)-1]
//			if cur.Node.Right != nil {
//				st = append(st, &MyTreeNode{Node: cur.Node.Right, Visited: false})
//			}
//			// 再把自己放回去
//			st = append(st, &MyTreeNode{Node: cur.Node, Visited: true})
//			if cur.Node.Left != nil {
//				st = append(st, &MyTreeNode{Node: cur.Node.Left, Visited: false})
//			}
//		}
//	}
//	return result
//}

func main() {

}
