package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	next := head.Next
	head.Next = swapPairs(head.Next.Next)
	next.Next = head
	return next
}


func main() {
	node4 := ListNode{4, nil}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}
	swapPairs(&node1)
}
