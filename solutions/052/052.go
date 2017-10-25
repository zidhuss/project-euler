package main

import (
	"fmt"
)

func hasSameDigits(a, b int) bool {
	x := make([]int, 10)
	y := make([]int, 10)
	for a > 0 {
		x[a%10]++
		a /= 10
	}

	for b > 0 {
		y[b%10]++
		b /= 10
	}

	for i := 0; i < 10; i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func hasSameDigitsRange(n int) bool {
	for i := 2; i < 7; i++ {
		if hasSameDigits(n, n*i) != true {
			return false
		}
	}
	return true
}

func main() {
	i := 1
	for hasSameDigitsRange(i) != true {
		i++
	}
	fmt.Println(i)
}
