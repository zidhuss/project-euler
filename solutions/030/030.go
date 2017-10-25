package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 355000; i > 2; i-- {
		j := i
		raisedSum := 0

		for j > 0 {
			raisedSum += pow((j % 10), 5)
			j /= 10
		}
		if i == raisedSum {
			sum += i
		}
	}
	fmt.Println(sum)
}

func pow(n, e int) int {
	s := 1
	for i := 0; i < e; i++ {
		s *= n
	}
	return s
}
