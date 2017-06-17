package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bs, _ := ioutil.ReadFile("resources/cipher.txt")
	cipher := strings.Split(strings.Replace(string(bs), "\n", "", -1), ",")

	size := len(cipher)

	sum := 0
	key := [3]int{'g', 'o', 'd'}

	for i := 0; i < size; i++ {
		x, _ := strconv.Atoi(cipher[i])
		sum += x ^ (key[i%3])
	}

	fmt.Println(sum)
}
