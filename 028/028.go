package main

import (
	"fmt"
)

func main() {
	a := func(n int) int {
		return int(n*(n+2)/4) + int((n%4)/3) + 1
	}
	n := 1001
	sum := 0
	for i := 1; a(i) <= n*n; i++ {
		sum += a(i)
	}
	fmt.Println(sum)
}
