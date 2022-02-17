package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	var result []int
	// 这里只有26个字母，用数组代替map也可以
	window := make([]int, 26)
	need := make([]int, 26)
	for _, r := range p {
		need[r-'a']++
	}
	left := 0
	for right := 0; right < len(s); right++{
		if right==2{
			fmt.Println("1")
		}
		curChar := s[right]
		window[curChar-'a']++
		// 只要最新的多于need，必然要开始从左边减少，一直到这个元素的个数等于need为止
		if window[curChar-'a'] > need[curChar-'a']{
			for window[curChar-'a'] > need[curChar-'a']{
				window[s[left]-'a']--
				left++
			}
		}
		if right - left == len(p)-1{
			result = append(result, left)
		}
	}
	return result
}

func main() {
	//fmt.Println(sameString("abc", "bca"))
	//fmt.Println(sameString("abc", "bda"))
	fmt.Println(findAnagrams("cbgebabacd", "abc"))
}
