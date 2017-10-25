package main

import (
	"fmt"
	"math/big"
)

func main() {
	fibo := []*big.Int{big.NewInt(0), big.NewInt(1)}

	for len(fibo[len(fibo)-1].String()) < 1000 {
		fibo = append(fibo, big.NewInt(0).Add(fibo[len(fibo)-2], fibo[len(fibo)-1]))
	}
	fmt.Println(len(fibo) - 1)
}
