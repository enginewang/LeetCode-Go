package main

import "fmt"

func isAnagram(s string, t string) bool {
	d := make(map[rune] int)
	for _, i := range []rune(s){
		d[i]++
	}
	for _, i := range []rune(t){
		if d[i] <= 0{
			return false
		}
		d[i]--
	}
	for _, value := range d{
		if value != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "vnagaramav"))
}
