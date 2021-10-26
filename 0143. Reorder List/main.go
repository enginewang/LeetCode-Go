package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 第一步：找mid，分割为两个链表
// 第二步：将后一个链表reverse
// 第三步，合并两个链表

func findMidNode(head *ListNode) *ListNode {
	mid, tail := head, head
	for tail.Next != nil && tail.Next.Next != nil {
		mid = mid.Next
		tail = tail.Next.Next
	}
	return mid
}

func reverseLinkedList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur.Next != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	cur.Next = pre
	return cur
}

func mergeLinkedList(h1 *ListNode, h2 *ListNode) {
	// 直接修改h1
	cur1, cur2 := h1, h2
	for cur1 != nil && cur2 != nil{
		next1 := cur1.Next
		next2 := cur2.Next
		cur1.Next = cur2
		cur2.Next = next1
		cur1 = next1
		cur2 = next2
	}
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	mid := findMidNode(head)
	head2 := mid.Next
	mid.Next = nil
	head2 = reverseLinkedList(head2)
	mergeLinkedList(head, head2)
}


func main() {
	node5 := ListNode{5, nil}
	node4 := ListNode{4, &node5}
	node3 := ListNode{3, &node4}
	node2 := ListNode{2, &node3}
	node1 := ListNode{1, &node2}
	reorderList(&node1)
}
