package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bs, _ := ioutil.ReadFile("./resources/numbers.txt")
	lines := strings.Split(string(bs), "\n")
	sum := 0
	for _, line := range lines {
		n, _ := strconv.Atoi(line[:11])
		sum += n
	}
	fmt.Println(strconv.Itoa(sum)[:10])
}
