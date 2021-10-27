package main

import "fmt"

// 因为涉及的操作基本都是头尾的插入查找删除
// 用双链表很节约时间

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type MyCircularDeque struct {
	LeftDummyNode  *Node
	RightDummyNode *Node
	MaxSize        int
	Size           int
}

func Constructor(k int) MyCircularDeque {
	cq := MyCircularDeque{}
	cq.LeftDummyNode = &Node{Val: -1}
	cq.RightDummyNode = &Node{Val: -1}
	cq.LeftDummyNode.Right = cq.RightDummyNode
	cq.RightDummyNode.Left = cq.LeftDummyNode
	cq.Size = 0
	cq.MaxSize = k
	return cq
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.Size == this.MaxSize {
		return false
	}
	newNode := &Node{Val: value}
	newNode.Left = this.LeftDummyNode
	newNode.Right = this.LeftDummyNode.Right
	newNode.Right.Left = newNode
	this.LeftDummyNode.Right = newNode
	this.Size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.Size == this.MaxSize {
		return false
	}
	newNode := &Node{Val: value}
	newNode.Right = this.RightDummyNode
	newNode.Left = this.RightDummyNode.Left
	newNode.Left.Right = newNode
	this.RightDummyNode.Left = newNode
	this.Size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.Size == 0 {
		return false
	}
	this.LeftDummyNode.Right = this.LeftDummyNode.Right.Right
	this.LeftDummyNode.Right.Left = this.LeftDummyNode
	this.Size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.Size == 0 {
		return false
	}
	this.RightDummyNode.Left = this.RightDummyNode.Left.Left
	this.RightDummyNode.Left.Right = this.RightDummyNode
	this.Size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.Size == 0 {
		return -1
	} else {
		return this.LeftDummyNode.Right.Val
	}
}

func (this *MyCircularDeque) GetRear() int {
	if this.Size == 0 {
		return -1
	} else {
		return this.RightDummyNode.Left.Val
	}
}

func (this *MyCircularDeque) IsEmpty() bool {
	if this.Size == 0 {
		return true
	} else {
		return false
	}
}

func (this *MyCircularDeque) IsFull() bool {
	if this.Size == this.MaxSize {
		return true
	} else {
		return false
	}
}

func main() {
	obj := Constructor(3)
	param_1 := obj.InsertFront(1)
	param_2 := obj.InsertLast(2)
	param_3 := obj.DeleteFront()
	param_4 := obj.DeleteLast()
	param_5 := obj.GetFront()
	param_6 := obj.GetRear()
	param_7 := obj.IsEmpty()
	param_8 := obj.IsFull()
	fmt.Println(param_1, param_2, param_3, param_4, param_5, param_6, param_7, param_8)
}
