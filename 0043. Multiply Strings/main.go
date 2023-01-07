package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	var result string
	// 进位
	next := make([]int, len(num1)+len(num2))
	for k := 0; k < len(num1)+len(num2); k++ {
		//fmt.Printf("开始倒数第%v位\n", k)
		cur := 0
		for i := 0; i <= k; i++ {
			if len([]rune(num1))-1-i >= 0 && len([]rune(num2))-1-(k-i) >= 0 {
				n1 := []rune(num1)[len([]rune(num1))-1-i] - '0'
				n2 := []rune(num2)[len([]rune(num2))-1-(k-i)] - '0'
				//fmt.Printf("%v %v\n", n1, n2)
				cur += int(n1) * int(n2)
			}
		}
		//fmt.Printf("加上上一位的进位%v\n", next[k])
		cur += next[k]
		result = string(rune(cur%10+'0')) + result
		// 进位存储
		m := 0
		//fmt.Printf("当前位计算结果为%v", cur)
		for cur > 0{
			cur = (cur - cur%10)/10
			if cur > 0{
				//fmt.Printf("进位%v %v\n", k+m+1, cur%10)
				next[k+m+1] += cur%10
			}
			//next[k+m+1] += cur%10
			m++
		}
		//fmt.Printf("结果为%v\n", result)
	}
	//去掉前面的0
	for i, c := range []rune(result){
		if c != '0'{
			result = result[i:]
			return result
		}
	}
	return "0"
}

func main() {
	fmt.Println(multiply("2", "4"))
	// 419254329864656431168468
}
