package main

import (
	"fmt"
	"sort"
)

func longestWord(words []string) string {
	sort.Strings(words)
	m := make(map[string]bool)
	m[""] = true
	longest := ""
	for _, w := range words {
		if _, ok := m[w[:len(w)-1]]; ok {
			m[w] = true
			if len(w) > len(longest) {
				longest = w
			}
		}
	}
	fmt.Println(longest)
	return longest
}

func main() {
	longestWord([]string{"w", "wo", "wor", "worl", "world", "appl", "a"})
}
