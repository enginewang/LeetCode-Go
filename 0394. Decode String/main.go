package main

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {
	var stack []byte
	for _, b := range []byte(s) {
		if b != ']' {
			stack = append(stack, b)
		} else {
			// 一对中括号里面的
			var inner []byte
			// 退栈，找出里面的内容
			for stack[len(stack)-1] != '[' {
				inner = append([]byte{stack[len(stack)-1]}, inner...)
				stack = stack[:len(stack)-1]
			}
			fmt.Println(inner)
			// 把'['退掉
			stack = stack[:len(stack)-1]
			// 找到前面的数字，拼起来
			var number []byte
			for len(stack) > 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
				number = append([]byte{stack[len(stack)-1]}, number...)
				stack = stack[:len(stack)-1]
			}
			num, _ := strconv.Atoi(string(number))
			for num > 0 {
				stack = append(stack, inner...)
				num--
			}
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(decodeString("2[abc]def3[a2[c]]"))
}
