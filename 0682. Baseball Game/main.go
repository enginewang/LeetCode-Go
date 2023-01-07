package main

import "strconv"

func calPoints(ops []string) int {
	var score []int
	for _, op := range ops {
		if op == "+" {
			score = append(score, score[len(score)-2]+score[len(score)-1])
		} else if op == "D" {
			score = append(score, 2*score[len(score)-1])
		} else if op == "C" {
			score = score[:len(score)-1]
		} else {
			s, _ := strconv.Atoi(op)
			score = append(score, s)
		}
	}
	sum := 0
	for _, num := range score{
		sum += num
	}
	return sum
}

func main() {

}
