package main

import (
	"fmt"
)

func main() {
	n := 10001
	primes := []int{2}
	for x := 3; len(primes) < n+1; x += 2 {
		for _, p := range primes {
			if x%p == 0 { // Not prime
				break
			} else if p*p > x { // Is prime
				primes = append(primes, x)
				break
			}
		}
	}
	fmt.Printf("%d\n", primes[10000])
}
