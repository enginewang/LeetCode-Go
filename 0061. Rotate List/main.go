package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	length := 1
	cur := head
	for cur.Next != nil {
		length += 1
		cur = cur.Next
	}
	tail := cur
	cur = head
	if k%length == 0 {
		return head
	}
	for i := 0; i < length-k%length-1; i++ {
		cur = cur.Next
	}
	newHead := cur.Next
	cur.Next = nil
	tail.Next = head
	return newHead
}
func main() {

}
