package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// K个一组翻转链表，高频面试题，建议默写

func hasKNode(head *ListNode, k int) bool {
	cur := head
	for ; k > 0; k-- {
		if cur != nil {
			cur = cur.Next
		} else {
			return false
		}
	}
	return true
}

//func reverseKGroup(head *ListNode, k int) *ListNode {
//    if k == 1{
//         return head
//    }
//    firstTime := true
//    var result *ListNode
//    var tempTail *ListNode
//    for hasKNode(head, k){
//         p := 0
//         tempHead := head
//         back := head
//         head = head.Next
//         for head.Next != nil {
//              p += 1
//              if p == k-1 {break}
//              next := head.Next
//              head.Next = back
//              back = head
//              head = next
//         }
//         tempHead.Next = head.Next
//         if firstTime{
//              result = head
//              firstTime = false
//         } else {
//              tempTail.Next = head
//         }
//         tempTail = tempHead
//         head.Next = back
//         head = tempHead.Next
//    }
//    return result
//}

// 递归
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 || !hasKNode(head, k) {
		return head
	}
	left, thisHead := head, head
	cur := head.Next
	next := cur.Next
	thisHead.Next = nil
	for i := 0; i < k-1; i++ {
		cur.Next = left
		left = cur
		cur = next
		if next != nil {
			next = next.Next
		}
	}
	thisHead.Next = reverseKGroup(cur, k)
	return left
}

func main() {
	node2 := ListNode{2, nil}
	node1 := ListNode{1, &node2}
	result := reverseKGroup(&node1, 2)
	fmt.Println(result)
	//fmt.Println(hasKNode(&node1, 6))
}
