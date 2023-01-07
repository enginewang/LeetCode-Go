package main

// 1 2 3
//

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		// 将inStack全部转移到outStack
		for len(this.inStack) > 0 {
			last := this.inStack[len(this.inStack)-1]
			this.inStack = this.inStack[:len(this.inStack)-1]
			this.outStack = append(this.outStack, last)
		}

	}
	out := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return out
}

func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		// 将inStack全部转移到outStack
		inLength := len(this.inStack)
		for inLength > 0 {
			last := this.inStack[len(this.inStack)-1]
			this.inStack = this.inStack[:len(this.inStack)-1]
			this.outStack = append(this.outStack, last)
			inLength--
		}

	}
	return this.outStack[len(this.outStack)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.inStack) == 0 && len(this.outStack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {

}
