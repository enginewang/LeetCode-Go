package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type NodeWithDepth struct {
	Node  *Node
	Depth int
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var queue []NodeWithDepth
	queue = append(queue, NodeWithDepth{Node: root, Depth: 0})
	for len(queue) != 0 {
		head := queue[0]
		if head.Node.Left != nil {
			queue = append(queue, NodeWithDepth{Node: head.Node.Left, Depth: head.Depth+1})
		}
		if head.Node.Right != nil {
			queue = append(queue, NodeWithDepth{Node: head.Node.Right, Depth: head.Depth+1})
		}
		if len(queue) > 1 && head.Depth == queue[1].Depth {
			head.Node.Next = queue[1].Node
		} else {
			head.Node.Next = nil
		}
		queue = queue[1:]
	}
	return root
}

func main() {

}
