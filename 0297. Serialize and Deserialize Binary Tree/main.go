package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}


func dfs(valsPtr *[]string) *TreeNode {
	val := (*valsPtr)[0]
	*valsPtr = (*valsPtr)[1:]
	if val == "null" {
		return nil
	}
	valInt, _ := strconv.Atoi(val)
	node := &TreeNode{Val: valInt}
	if len(*valsPtr) > 0 {
		node.Left = dfs(valsPtr)
	}
	if len(*valsPtr) > 0 {
		node.Right = dfs(valsPtr)
	}
	return node
}

func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	return dfs(&vals)
}

func main() {
	this := Constructor()
	tree := this.deserialize("4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2")
	str := this.serialize(tree)
	fmt.Println(str)
}
