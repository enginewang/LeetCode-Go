package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	var st []int
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			n1 := st[len(st)-1]
			n2 := st[len(st)-2]
			st = st[:len(st)-2]
			var result int
			switch token {
			case "+":
				result = n2 + n1
			case "-":
				result = n2 - n1
			case "*":
				result = n2 * n1
			case "/":
				result = n2 / n1
			}
			st = append(st, result)
		} else {
			num, _ := strconv.Atoi(token)
			st = append(st, num)
		}
		fmt.Println(st)
	}
	return st[0]
}

func main() {
	evalRPN([]string{"4", "13", "5", "/", "+"})
}
