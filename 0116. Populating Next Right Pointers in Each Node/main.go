package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	connectTwoNodes(root.Left, root.Right)
	return root
}

func connectTwoNodes(node1 *Node, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	connectTwoNodes(node1.Left, node1.Right)
	connectTwoNodes(node1.Right, node2.Left)
	connectTwoNodes(node2.Left, node2.Right)
}

func main() {

}
