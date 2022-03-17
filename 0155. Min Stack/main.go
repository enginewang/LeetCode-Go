package main

import (
	"fmt"
)

// 一个关键的地方就在于，不需要求解所有元素最小的值，实际上只需要看当前插入的是否比min[]小
// 因为如果当前插入的更大，那么就不可能作为min范返回
type MinStack struct {
	s   []int
	min []int
}

func Constructor() MinStack {
	return MinStack{s: []int{}, min: []int{}}
}

func (this *MinStack) Push(val int) {
	this.s = append(this.s, val)
	if len(this.min) == 0 || val <= this.GetMin(){
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.s) < 1 {
		fmt.Println("Empty Stack")
	}
	if this.Top() == this.GetMin(){
		this.min = this.min[:len(this.min)-1]
	}
	this.s = this.s[:len(this.s)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
	fmt.Println("1")
}
