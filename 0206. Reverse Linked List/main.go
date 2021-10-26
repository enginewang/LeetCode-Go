package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 翻转链表，非常经典的题目，面试高频，主要是几个指针要弄清楚，别被绕晕了。建议默下来
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	back := head
	head = head.Next
	// 这边记得要把最开始head的Next置为nil，不然就无限循环了。
	back.Next = nil
	for head.Next != nil {
		next := head.Next
		head.Next = back
		back = head
		head = next
	}
	head.Next = back
	return head
}


func main() {
	node3 := ListNode{3, nil}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}
	reverseList(&node1)
}
