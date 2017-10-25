package main

import "fmt"

func sum(n, k int) int {
	return k * ((((n - 1) / k) * (((n - 1) / k) + 1)) / 2)
}

func main() {
	solution := sum(1000, 3) + sum(1000, 5) - sum(1000, 15)
	fmt.Println(solution)
}
