package main

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

很简答，遍历两个字符串，将字符作为key存在哈希表中。可以只用一个哈希表，一个字符串加一个减，如果最后哈希表的value全为0就说明是异位词。
*/

import "fmt"

func isAnagram(s string, t string) bool {
	d := make(map[rune]int)
	for _, i := range []rune(s) {
		d[i]++
	}
	for _, i := range []rune(t) {
		if d[i] <= 0 {
			return false
		}
		d[i]--
	}
	for _, value := range d {
		if value != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "vnagaramav"))
}
