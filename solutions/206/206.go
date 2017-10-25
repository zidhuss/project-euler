package main

import "fmt"

func isValid(n uint64) bool {
	if n < 1e18 || n%10 != 0 {
		return false
	}

	i := uint64(9)
	for n > 0 {
		n /= 100
		if n%10 != i {
			return false
		}
		i--
	}
	return true
}

func main() {
	for i := uint64(1e9); i < 1e19; i += 2 {
		if i%3 == 0 || i%7 == 0 {
			if isValid(i * i) {
				fmt.Println(i)
				break
			}
		}
	}
}
