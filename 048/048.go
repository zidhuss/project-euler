package main

import (
	"fmt"
	"math/big"
)

func main() {
	sum := new(big.Int)
	for i := 1; i <= 1000; i++ {
		n := big.NewInt(int64(i))
		sum.Add(sum, n.Exp(n, n, new(big.Int)))
	}
	numberString := sum.String()
	fmt.Println(numberString[len(numberString)-10:])
}
