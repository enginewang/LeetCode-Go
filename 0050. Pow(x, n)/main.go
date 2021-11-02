package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0{
		return 1.0/myPow(x, -n)
	}
	if n % 2 == 0{
		return myPow(x*x, n/2)
	} else {
		return myPow(x*x, n/2)*x
	}
}

func main() {
	fmt.Println(myPow(2.1, 3))
}
