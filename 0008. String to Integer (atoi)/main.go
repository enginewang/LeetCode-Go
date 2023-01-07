package main

import "fmt"

func myAtoi(s string) int {
	hasRead := false
	nextRead := false
	positive := 1
	number := 0
	for _, c := range []rune(s) {
		if nextRead && !(c >= '0' && c <= '9'){
			return 0
		}
		if !hasRead {
			if c == ' ' {
				continue
			} else if c == '-' || c == '+' {
				nextRead = true
				if c == '-' {
					positive = -1
				}
				continue
			}
		}
		if c >= '0' && c <= '9' {
			if nextRead {
				nextRead = false
			}
			hasRead = true
			number = number*10 + (int(c) - '0')
			if positive * number > 2147483647{
				return 2147483647
			}
			if positive * number < -2147483648{
				return -2147483648
			}
		} else {
			if nextRead {
				return 0
			}
			return positive * number
		}
	}
	return positive * number
}

func main() {
	fmt.Println(myAtoi("  --91283472332"))
}
