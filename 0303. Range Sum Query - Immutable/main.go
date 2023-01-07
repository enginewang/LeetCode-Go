package main

import "fmt"

type NumArray struct {
	//Nums []int
	Aux  []int
}

func Constructor(nums []int) NumArray {
	sum := 0
	na := new(NumArray)
	for _, v := range nums {
		sum += v
		na.Aux = append(na.Aux, sum)
	}
	return *na
}

func (this *NumArray) SumRange(left int, right int) int {
	if left != 0 {
		return this.Aux[right] - this.Aux[left-1]
	} else {
		return this.Aux[right]
	}
}

func main() {
	obj := Constructor([]int{1,2,3,4,5})
	fmt.Println(obj.SumRange(2, 4))
}
