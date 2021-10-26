package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	p1, p2 := head, head
	for p2.Next != nil && p2.Next.Next != nil{
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	if p2.Next != nil{
		p1 = p1.Next
	}
	return p1
}

func main() {
	node6 := ListNode{6, nil}
	node5 := ListNode{5, &node6}
	node4 := ListNode{4, &node5}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}
	fmt.Println(middleNode(&node1))
}
