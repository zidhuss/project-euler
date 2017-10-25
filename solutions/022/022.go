package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func nameScore(name string) (score int) {
	for _, letter := range name {
		score += int(letter - 64)
	}
	return
}

func main() {
	file, _ := os.Open("./resources/names.txt")
	defer file.Close()
	bs, _ := ioutil.ReadAll(file)
	names := strings.Split(string(bs), ",")
	sort.Strings(names)
	sum := 0
	for position, name := range names {
		name := strings.Replace(name, "\"", "", -1)
		sum += nameScore(name) * (position + 1)
	}
	fmt.Println(sum)
}
