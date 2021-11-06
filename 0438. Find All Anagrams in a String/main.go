package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	var result []int
	sDict := make([]int, 26)
	pDict := make([]int, 26)
	for _, r := range p {
		pDict[r-'a']++
	}
	left := 0
	for right := 0; right < len(s); right++{
		curChar := s[right]
		sDict[curChar-'a']++
		if sDict[curChar-'a'] > pDict[curChar-'a']{
			for sDict[curChar-'a'] > pDict[curChar-'a']{
				sDict[s[left]-'a']--
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
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
