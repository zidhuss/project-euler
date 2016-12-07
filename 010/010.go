package main

import (
	"fmt"
	"math"
)

func main() {
	limit := 2000000
	sievebound := (limit - 1) / 2
	crosslimit := int((math.Sqrt(float64(limit)) - 1) / 2)
	sieve := make([]bool, sievebound)

	for i := 1; i < crosslimit; i++ {
		if !sieve[i] {
			for j := 2 * i * (i + 1); j < sievebound; j += 2*i + 1 {
				sieve[j] = true
			}
		}
	}

	sum := 2
	for i := 1; i < sievebound; i++ {
		if !sieve[i] {
			sum += 2*i + 1
		}
	}
	fmt.Println(sum)
}
