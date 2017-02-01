package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 400000; i > 2; i-- {
		j := i
		factorialSum := 0

		for j > 0 {
			factorialSum += factorial(j % 10)
			j /= 10
		}
		if i == factorialSum {
			sum += i
		}
	}
	fmt.Println(sum)
}

func factorial(n int) int {
	s := 1
	for i := n; i > 0; i-- {
		s *= i
	}
	return s
}
