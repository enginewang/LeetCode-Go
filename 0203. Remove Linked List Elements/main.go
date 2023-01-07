package main

// 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回新的头节点 。

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
	// 一直往后走，遇到了next等于val的就删掉next
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
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
