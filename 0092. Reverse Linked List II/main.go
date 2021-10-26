package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	dummyNode := &ListNode{}
	pre := dummyNode
	dummyNode.Next = head
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i:=0; i<right-left;i++{
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}

func main() {
	node2 := ListNode{2, nil}
	node1 := ListNode{1, &node2}
	fmt.Println(reverseBetween(&node1, 1, 2))
}
