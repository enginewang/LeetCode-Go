package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除倒数第N个节点，那么我们当前遍历的指针一定要指向 第N个节点的前一个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	dummyNode := &ListNode{Val: -1, Next: head}
	count := 0
	cur := head
	// 首先要遍历一遍算一下有多长，从而计算倒数第N个是正数第几个
	for cur != nil {
		cur = cur.Next
		count++
	}
	id := count - n
	cur = dummyNode
	for id > 0 {
		cur = cur.Next
		id--
	}
	cur.Next = cur.Next.Next
	return dummyNode.Next
}

func main() {
	node3 := ListNode{5, nil}
	removeNthFromEnd(&node3, 1)
}
