package main

import "fmt"

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	num := 1
	den := 1

	for i := 1; i < 10; i++ {
		for d := 1; d < i; d++ {
			for n := 1; n < d; n++ {
				a := n*10 + i
				b := i*10 + d
				if float32(n)/float32(d) == float32(a)/float32(b) {
					num *= a
					den *= b
				}
			}
		}
	}

	fmt.Println(den / gcd(num, den))
}
