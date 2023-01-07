package main

func maxDepth(s string) int {
	maxD := 0
	depth := 0
	for _, c := range []rune(s){
		if c == '('{
			depth += 1
			if depth > maxD{
				maxD = depth
			}
		}
		if c == ')'{
			depth -= 1
		}
	}
	return maxD
}

func main() {

}
