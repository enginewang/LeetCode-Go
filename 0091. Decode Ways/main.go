package main

import (
	"fmt"
)

func numDecodings(s string) int {
	if len(s) == 0{
		return 0
	}
	if len(s) == 1{
		if s[0] != '0' {
			return 1
		} else {
			return 0
		}
	}
	if len(s) == 2{
		if s[1] == '0'{
			if s[0] == '0'{
				return 0
			} else {
				return 1
			}
		} else {
			if s[0] == '0'{
				return 0
			} else if (s[1] > '6' && s[0] == '2') || s[0] > '2'{
				return 1
			} else {
				return 2
			}
		}

	}
	// 如果最后一个是0，不能单独组成
	if s[len(s)-1] == '0'{
		if s[len(s)-2] > '2'{
			return 0
		}else {
			// 实际上等于前两个
			return numDecodings(s[:len(s)-2])
		}
	} else{
		// 否则的话，单独就可以做一个
		if (s[len(s)-1] > '6' && s[len(s)-2] == '2') || s[len(s)-2] > '2' || s[len(s)-2] == '0' {
			return numDecodings(s[:len(s)-1])
		} else {
			return numDecodings(s[:len(s)-1]) + numDecodings(s[:len(s)-2])
		}
	}
}

func main() {
	fmt.Println(numDecodings("10011"))
}
