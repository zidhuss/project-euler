package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if (n%2 == 0) || (n%3 == 0) {
		return false
	}
	i := 5
	for i*i <= n {
		if (n%i == 0) || (n%(i+2) == 0) {
			return false
		}
		i += 6
	}
	return true
}

func main() {
	primeCount := 0
	x := 3
	for {
		for i := 3; i > 0; i-- {
			n := x*x - i*(x-1)
			if isPrime(n) {
				primeCount++
			}
		}
		if 0.1 > (float64(primeCount) / float64(2*(x-1)+1)) {
			break
		}
		x += 2
	}
	fmt.Println(x)
}
