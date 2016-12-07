package main

import (
	"fmt"
	"math"
)

func main() {
	s := 1000
	r := int(math.Sqrt(float64(s)))
	for m := r; m > 1; m-- {
		for n := m; n > 2; n-- {
			a := m*m - n*n
			b := 2 * (m * n)
			c := m*m + n*n
			if a+b+c == s {
				fmt.Printf("%d\n", a*b*c)
				return
			}
		}
	}
}
