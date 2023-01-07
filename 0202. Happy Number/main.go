package main

import "fmt"

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 {
		n = calSquare(n)
		// 如果已经存在了，就说明进入了无限循环
		if m[n] {
			return false
		}
		// 否则就放进去
		m[n] = true
	}
	return true
}

func calSquare(n int) int {
	result := 0
	for n > 0 {
		a := n % 10
		result += a * a
		n /= 10
	}
	return result
}

func main() {
	fmt.Println(isHappy(2))
}
