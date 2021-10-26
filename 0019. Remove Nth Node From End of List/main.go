package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	dummyNode := &ListNode{Val: -1, Next: head}
	count := 0
	cur := head
	for cur != nil{
		cur = cur.Next
		count++
	}
	id := count - n
	cur = dummyNode
	for id > 0{
		cur = cur.Next
		id--
	}
	cur.Next = cur.Next.Next
	return dummyNode.Next
}

func main() {
	node3 := ListNode{5, nil}
	//node5 := ListNode{4, &node3}
	//node4 := ListNode{3, &node5}
	//node2 := ListNode{2, &node4}
	//node1 := ListNode{1, &node2}
	removeNthFromEnd(&node3, 1)
}