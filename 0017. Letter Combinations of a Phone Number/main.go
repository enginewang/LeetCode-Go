package main

import "fmt"

var numberMap = map[int][]string{
	2: {"a", "b", "c"},
	3: {"d", "e", "f"},
	4: {"g", "h", "i"},
	5: {"j", "k", "l"},
	6: {"m", "n", "o"},
	7: {"p", "q", "r", "s"},
	8: {"t", "u", "v"},
	9: {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	d := []rune(digits)
	var result []string
	if len(d) == 0 {
		return []string{}
	}
	if len(d) == 1 {
		return numberMap[int(d[0])-48]
	}
	for _, str := range letterCombinations(string(d[:len(d)-1])) {
		for _, cur := range numberMap[int(d[len(d)-1])-48] {
			result = append(result, str+cur)
		}
	}
	return result
}

func main() {
	fmt.Println(letterCombinations("23"))
}
