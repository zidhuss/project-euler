package main

import (
	"fmt"
)

func main() {
	n := 1000000
	sequences := make([]int, n)
	longest := 1

	for i := 2; i < n; i++ {
		found := false
		count := 0
		c := i
		for c != 1 {
			// Have we already reahed this number before?
			if c < n && sequences[c] > 0 {
				sequences[i] = count + sequences[c]
				found = true
				break
			}

			if c%2 == 0 {
				c = c / 2
			} else {
				c = 3*c + 1
			}
			count++
		}

		if !found {
			sequences[i] = count + 1
		}

		if sequences[i] > sequences[longest] {
			longest = i
		}
	}
	fmt.Println(longest)
}
