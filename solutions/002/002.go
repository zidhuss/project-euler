package main

import "fmt"

func main() {
	var fibo [36]int
	fibo[0], fibo[1] = 0, 1
	for i := 2; i < 34; i++ {
		fibo[i] = fibo[i-2] + fibo[i-1]
	}
	s := 0
	for _, x := range fibo {
		if x%2 == 0 {
			s += x
		}
	}
	fmt.Println(s)
}
