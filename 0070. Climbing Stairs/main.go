package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	n1 := 1
	n2 := 2
	for i := 3; i <= n; i++ {
		temp := n2
		n2 = n2 + n1
		n1 = temp
	}
	return n2
}

func main() {
	fmt.Println(climbStairs(3))
}
