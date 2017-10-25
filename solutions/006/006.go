package main

import (
	"fmt"
)

func main() {
	var sumSquare, squareSum = 0, 0
	for i := 1; i <= 100; i++ {
		squareSum += i * i
		sumSquare += i
	}
	fmt.Println(sumSquare*sumSquare - squareSum)
}
