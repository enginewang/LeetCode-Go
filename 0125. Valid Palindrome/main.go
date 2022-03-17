package main

import "fmt"

func isPalindrome(s string) bool {
	var str []rune
	fmt.Println([]rune(s))
	for _, item := range []rune(s) {
		if (item >= 97 && item <= 122) || (item >= 48 && item <= 57) {
			str = append(str, item)
		} else if item >= 65 && item <= 90 {
			str = append(str, item+32)
		}
	}
	fmt.Println(str)
	start := 0
	for ; start < len(str)-1-start; start++ {
		if str[start] != str[len(str)-1-start] {
			return false
		}
	}
	return true
}

func main() {
	//s := "A man, a plan, a canal: Panama"
	s := "ab_a"
	fmt.Println(isPalindrome(s))
}
