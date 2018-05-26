package main

import (
	"fmt"
)

func main() {
	target := 200
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	ways := make([]int, target+1)
	ways[0] = 1

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= target; j++ {
			ways[j] += ways[j-coins[i]]
		}
	}

	fmt.Println(ways[target])
}
