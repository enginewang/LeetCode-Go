package main

import "fmt"

func trailingZeroes(n int) int {
	result := 0
	for n/5 > 0 {
		result += n/5
		n /= 5
	}
	return result
}

func main() {
	fmt.Println(trailingZeroes(300))
}
