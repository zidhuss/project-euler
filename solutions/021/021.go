package main

import "fmt"

func sumDivisors(n int) int {
	sum := 0
	i := 2
	for ; i*i < n; i++ {
		if n%i == 0 {
			sum += i + (n / i)
		}
	}
	if i*i == n {
		sum += i
	}
	return sum + 1
}

func main() {
	amicableSum := 0
	for a := 2; a < 10000; a++ {
		b := sumDivisors(a)
		if (b > a) && (sumDivisors(b) == a) {
			amicableSum += a + b
		}
	}
	fmt.Println(amicableSum)
}
