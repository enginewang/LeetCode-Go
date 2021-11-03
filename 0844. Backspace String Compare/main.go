package main

import "fmt"

func backspaceCompare(s string, t string) bool {
	return cleanString(s) == cleanString(t)
}

// remove #
func cleanString(s string) string {
	var result []rune
	for _, r := range s{
		if r == '#'{
			if len(result) > 0{
				result = result[:len(result)-1]
			}
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func main() {
	fmt.Println(backspaceCompare("y#fo##f", "y#f#o##f"))
}
