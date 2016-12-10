package main

import (
	"fmt"
)

func squareDigits(n int) (res int) {
	for n > 0 {
		d := n % 10
		n /= 10
		res += d * d
	}
	return
}

func main() {
	count := 0
	for i := 2; i < 10000000; i++ {
		x := i
		for x != 1 {
			if x == 89 {
				count++
				break
			}
			x = squareDigits(x)
		}
	}
	fmt.Println(count)
}
