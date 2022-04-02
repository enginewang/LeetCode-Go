package main

type MyStack struct {
	Queue []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.Queue = append(this.Queue, x)
}

// 1 2 3
// 1 2 3
func (this *MyStack) Pop() int {
	length := len(this.Queue)
	for length > 1 {
		last := this.Queue[0]
		this.Queue = this.Queue[1:]
		this.Queue = append(this.Queue, last)
		length--
	}
	top := this.Queue[0]
	this.Queue = this.Queue[1:]
	return top
}

func (this *MyStack) Top() int {
	length := len(this.Queue)
	last := 0
	for length > 0 {
		last = this.Queue[0]
		this.Queue = this.Queue[1:]
		this.Queue = append(this.Queue, last)
		length--
	}
	return last
}

func (this *MyStack) Empty() bool {
	if len(this.Queue) == 0 {
		return true
	} else {
		return false
	}
}

func main() {

}
