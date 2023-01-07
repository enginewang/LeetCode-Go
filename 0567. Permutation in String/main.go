package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	need, window := [26]int{}, [26]int{}
	for i := 0; i < len(s1); i++ { 
		need[s1[i]-'a']++
	}
	left := 0
	for i, c := range s2 {
		window[c-'a']++
		for left <= i && need[s2[left]-'a'] < window[s2[left]-'a'] {
			window[s2[left]-'a']--
			left++
		}
		if need == window {
			return true
		}
	}
	return false
}


func main() {
	fmt.Println(checkInclusion("ab", "eibpagaooo"))
}
