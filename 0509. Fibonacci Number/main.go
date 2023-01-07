package main

import "fmt"

func fib(n int) int {
	memo := make(map[int]int)
	return dp(n, &memo)
}

func dp(n int, memo *map[int]int) int {
	if n == 0 || n == 1 {
		return n
	}
	if v, ok := (*memo)[n]; ok {
		return v
	}
	(*memo)[n] = fib(n-1) + fib(n-2)
	return (*memo)[n]
}

func main() {
	fmt.Println(fib(4))
}
