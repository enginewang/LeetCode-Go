package main

import "fmt"

// 要想达到O(1)的空间复杂度，方式是翻转后半个链表，然后就可以比较了

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	next := node.Next
	for next.Next != nil{
		node.Val = next.Val
		node = next
		next = next.Next
	}
	node.Val = next.Val
	node.Next = nil
}

func main() {
	node6 := ListNode{1, nil}
	node5 := ListNode{2, &node6}
	node4 := ListNode{3, &node5}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}
	deleteNode(&node3)
	fmt.Println(node1)
}
