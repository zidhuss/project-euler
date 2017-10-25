package main

import (
	"fmt"
)

func main() {
	n := []string{"", "one", "two", "three", "four", "five", "six", "seven",
		"eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen",
		"fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens := []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy",
		"eighty", "ninety"}

	for _, ten := range tens {
		for i := 0; i < 10; i++ {
			n = append(n, ten+n[i])
		}
	}

	for i := 1; i < 10; i++ {
		n = append(n, n[i]+"hundred")
		for j := 1; j < 100; j++ {
			n = append(n, n[i]+"hundredand"+n[j])
		}
	}

	n = append(n, "onethousand")

	sum := 0
	for _, number := range n {
		sum += len(number)
	}

	fmt.Println(sum)
}
