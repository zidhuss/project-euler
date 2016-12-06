package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func sliceProduct(x []int) (product int) {
	product = 1
	for _, i := range x {
		if i == 0 {
			return 0
		}
		product *= i
	}
	return
}

func adjacentProduct(numbers []int, size int) int {
	largest := sliceProduct(numbers[:size])
	for i := size + 1; i < len(numbers); i++ {
		current := sliceProduct(numbers[i-size : i])
		if current > largest {
			largest = current
		}
	}
	return largest
}

func main() {
	file, _ := os.Open("./resources/digits.txt")
	defer file.Close()
	bs, _ := ioutil.ReadAll(file)
	numberString := strings.Replace(string(bs), "\n", "", -1)
	var numbers []int
	for _, n := range numberString {
		digit, _ := strconv.Atoi(string(n))
		numbers = append(numbers, digit)
	}

	fmt.Println(adjacentProduct(numbers, 13))
}
