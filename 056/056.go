package main

import (
	"fmt"
	"math/big"
)

func main() {
	largest := 0
	for a := 1; a < 101; a++ {
		for b := 1; b < 101; b++ {
			x := new(big.Int)
			x = x.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), new(big.Int))
			c := 0
			for _, i := range x.String() {
				c += int(i - '0')
			}
			if c > largest {
				largest = c
			}
		}
	}
	fmt.Println(largest)
}
