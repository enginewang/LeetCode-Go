package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
     if head == nil {
          return false
     }
     m := make(map[*ListNode] bool)
     m[head] = true
     cur := head.Next
     for cur != nil {
          if _, ok := m[cur]; ok {
               return true
          } else {
               m[cur] = true
          }
          cur = cur.Next
     }
     return false
}

func main() {
     node4 := ListNode{4, nil}
     node3 := ListNode{3, &node4}
     node2 := ListNode{2, &node3}
     node1 := ListNode{1, &node2}
     node4.Next = &node2
     fmt.Println(hasCycle(&node1))
}
