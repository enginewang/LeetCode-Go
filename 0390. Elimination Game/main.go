package main

import "fmt"

func delTwo(start int, length int, step int, direct int) int{
	if length == 1 {
		return start
	} else if direct == 0 {
		return delTwo(start+step, int(length/2), step*2, 1-direct)
	} else {
		return delTwo(start + length%2*step, int(length/2), step*2, 1-direct)
	}
}

func lastRemaining(n int) int {
	return delTwo(1, n, 1, 0)
}


func main() {
	fmt.Println(lastRemaining(6))
}
