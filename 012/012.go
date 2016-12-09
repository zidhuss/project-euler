package main

import (
	"fmt"
)

func divisorCount(n int) int {
	count := 2
	i := 2
	for i*i < n {
		if n%i == 0 {
			count += 2
		}
		i++
	}
	if i*i == n {
		count++
	}
	return count
}

func main() {
	triangleNumbers := []int{1}
	for i := 0; ; i++ {
		if divisorCount(triangleNumbers[i]) > 500 {
			break
		}
		triangleNumbers = append(triangleNumbers, triangleNumbers[i]+i+2)
	}
	fmt.Printf("%d\n", triangleNumbers[len(triangleNumbers)-1])
}
