package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func factorial(n int64) *big.Int {
	x := new(big.Int)
	x.MulRange(1, n)
	return x
}

func main() {
	var sum int64
	digits := factorial(100)
	for _, i := range digits.String() {
		digit, _ := strconv.Atoi(string(i))
		sum += int64(digit)
	}
	fmt.Printf("%d\n", sum)
}
