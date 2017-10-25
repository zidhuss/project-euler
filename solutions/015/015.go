package main

import "fmt"

func main() {
	fmt.Println(binomialCoefficient(40, 20))
}

func binomialCoefficient(n, k uint64) uint64 {
	if k > n {
		return 0
	}
	if k == n {
		return 1
	}
	if n-k > k {
		k = n - k
	}
	c := uint64(1)
	for i := uint64(0); i < k; i++ {
		c = c * (n - i) / (i + 1)
	}
	return c
}
