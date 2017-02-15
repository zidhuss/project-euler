package main

import (
	"fmt"
	"math/big"
)

func main() {
	z := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil)
	sum := 0
	for _, x := range z.String() {
		sum += int(x - 48)
	}
	fmt.Println(sum)
}
