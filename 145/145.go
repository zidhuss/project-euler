package main

import "fmt"

func allOdd(n int) bool {
	for n > 0 {
		x := n % 10
		if x%2 == 0 {
			return false
		}
		n /= 10
	}
	return true
}

func reverse(n int) int {
	res := 0
	for n > 0 {
		x := n % 10
		res *= 10
		res += x
		n /= 10
	}
	return res
}

func main() {
	count := 0
	for i := 11; i < 1e8; i++ {
		if i%10 == 0 {
			continue
		}
		if allOdd(i + reverse(i)) {
			count++
		}
	}
	fmt.Println(count)
}
