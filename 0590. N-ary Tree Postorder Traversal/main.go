package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil{
		return []int{}
	}
	var result []int
	for _, child := range root.Children{
		result = append(result, postorder(child)...)
	}
	result = append(result, root.Val)
	return result
}

func main() {

}
