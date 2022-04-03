package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var queue []*Node
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		var row []int
		for i := 0; i < length; i++ {
			cur := queue[0]
			row = append(row, cur.Val)
			queue = queue[1:]
			for _, node := range cur.Children {
				if node != nil {
					queue = append(queue, node)
				}
			}
		}
		result = append(result, row)
	}
	return result
}

func main() {

}
