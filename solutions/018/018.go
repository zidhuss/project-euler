package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func maximumPathSum(tree [][]int) int {

	for i := len(tree) - 2; i >= 0; i-- {
		row := tree[i]

		for j := 0; j < len(row); j++ {
			x := tree[i+1][j]
			y := tree[i+1][j+1]

			if x > y {
				row[j] += x
			} else {
				row[j] += y
			}
		}
	}

	return tree[0][0]
}

func main() {
	bs, _ := ioutil.ReadFile("./resources/tree.txt")
	lines := strings.Split(string(bs), "\n")

	tree := [][]int{}

	// Remove empty line at end
	lines = lines[0 : len(lines)-1]

	for _, line := range lines {
		row := []int{}
		numbers := strings.Split(line, " ")
		for _, number := range numbers {
			n, _ := strconv.Atoi(number)
			row = append(row, n)
		}
		tree = append(tree, row)
	}

	fmt.Println(maximumPathSum(tree))
}
