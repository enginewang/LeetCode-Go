package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	var queue []*TreeNode
	queue = append(queue, root)
	var result int
	for len(queue) > 0 {
		length := len(queue)
		result = queue[0].Val
		for i := 0; i < length; i++ {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
	}
	return result
}

func main() {

}
