package main

import (
	"fmt"
	"math"
)

func main() {
	limit := 1000
	sievebound := (limit - 1) / 2
	crosslimit := int((math.Sqrt(float64(limit)) - 1) / 2)
	sieve := make([]bool, sievebound)

	for i := 1; i < crosslimit; i++ {
		if !sieve[i] {
			for j := 2 * i * (i + 1); j < sievebound; j += 2*1 + 1 {
				sieve[j] = true
			}
		}
	}

	longest := 1
	var n, l int

	for i := 1; i < sievebound; i++ {
		if !sieve[i] {
			n = 2*i + 1
			l = cycleLength(n)
			if l > longest {
				longest = l + 1
			}

		}
	}

	fmt.Println(longest)
}

func cycleLength(b int) int {
	if b%2 == 0 || b%5 == 0 {
		return 0
	}

	a := 1
	t := 0

	for {
		a = a * 10 % b
		t++
		if a == 1 {
			return t
		}
	}
}
