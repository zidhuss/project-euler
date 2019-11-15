package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func minimumPathSum(matrix [][]int) int {
	max := len(matrix) - 1

	for i := max; i >= 0; i-- {
		row := matrix[i]

		for j := max; j >= 0; j-- {

			if i < max && j < max {
				x := row[j+1]
				y := matrix[i+1][j]

				if x < y {
					matrix[i][j] += x
				} else {
					matrix[i][j] += y
				}

			} else if i < max {
				matrix[i][j] += matrix[i+1][j]
			} else if j < max {
				matrix[i][j] += row[j+1]
			}
		}
	}

	return matrix[0][0]
}

func main() {
	bs, _ := ioutil.ReadFile("./resources/matrix.txt")
	lines := strings.Split(string(bs), "\n")

	// Remove empty line at end
	lines = lines[0 : len(lines)-1]

	size := 80

	matrix := make([][]int, size)

	for lineIdx, line := range lines {
		numbers := strings.Split(line, ",")
		row := make([]int, size)
		for numIdx, number := range numbers {
			n, _ := strconv.Atoi(number)
			row[numIdx] = n
		}
		matrix[lineIdx] = row
	}

	fmt.Println(minimumPathSum(matrix))
}
