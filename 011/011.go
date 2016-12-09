package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./resources/grid.txt")
	defer file.Close()
	bs, _ := ioutil.ReadAll(file)
	contents := string(bs[:len(bs)-1])
	contents = strings.Replace(contents, "\n", " ", -1)
	digits := strings.Split(contents, " ")

	grid := []int{0}
	for _, n := range digits {
		converted, _ := strconv.Atoi(n)
		grid = append(grid, converted)
	}

	largest := 0
	for i := 1; i < len(grid)-1; i++ {
		if i%20 < 18 { // East
			current := grid[i] * grid[i+1] * grid[i+2] * grid[i+3]
			if current > largest {
				largest = current
			}

			if i+60 < len(grid)-1 { // South East
				current := grid[i] * grid[i+21] * grid[i+42] * grid[i+63]
				if current > largest {
					largest = current
				}
			}

			if i-60 > 0 { // North East
				current := grid[i] * grid[i-19] * grid[i-38] * grid[i-57]
				if current > largest {
					largest = current
				}
			}
		}

		if i%20 > 3 { // West
			current := grid[i] * grid[i-1] * grid[i-2] * grid[i-3]
			if current > largest {
				largest = current
			}

			if i+60 < len(grid)-1 { // South West
				current := grid[i] * grid[i+19] * grid[i+38] * grid[i+57]
				if current > largest {
					largest = current
				}
			}

			if i-60 > 0 { // North West
				current := grid[i] * grid[i-21] * grid[i-42] * grid[i-63]
				if current > largest {
					largest = current
				}
			}
		}

		if i+60 < len(grid)-1 { // South
			current := grid[i] * grid[i+20] * grid[i+40] * grid[i+60]
			if current > largest {
				largest = current
			}
		}

		if i-60 > 0 { // North
			current := grid[i] * grid[i-20] * grid[i-40] * grid[i-60]
			if current > largest {
				largest = current
			}
		}
	}
	fmt.Println(largest)
}
