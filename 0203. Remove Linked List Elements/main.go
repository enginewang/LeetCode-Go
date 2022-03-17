package main

// 要想达到O(1)的空间复杂度，方式是翻转后半个链表，然后就可以比较了

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Next.Val == val{
			cur.Next = cur.Next.Next
		}else {
			cur = cur.Next
		}
	}
	return head
}

func main() {
	node4 := ListNode{3, nil}
	node3 := ListNode{3, &node4}
	node2 := ListNode{3, &node3}
	node1 := ListNode{3, &node2}
	removeElements(&node1, 3)
}
