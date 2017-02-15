package main

import (
	"fmt"
	"math/big"
)

func main() {

	s := make(map[string]int)

	z := new(big.Int)
	for a := int64(2); a <= 100; a++ {
		for b := int64(2); b <= 100; b++ {
			s[z.Exp(big.NewInt(a), big.NewInt(b), big.NewInt(0)).String()]++
		}
	}
	fmt.Println(len(s))
}
