package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestPrimeFactor(600851475143))
}

func largestPrimeFactor(number int64) int64 {
	var factor int64 = 1
	for i := int64(2); i*i <= number; i++ {
		if number == 1 {
			return factor
		} else if number%i != 0 {
			continue
		} else {
			factor = i
			for number%i == 0 {
				number /= i
			}
		}
	}
	return number
}
