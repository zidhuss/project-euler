package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("./resources/words.txt")
	defer f.Close()
	r := csv.NewReader(f)
	words, _ := r.Read()

	t := map[int]int{}

	for i := 0; i < 100; i++ {
		x := (i * (i + 1)) / 2
		t[x] = i
	}

	found := 0

	for _, word := range words {
		sum := 0
		for _, c := range word {
			sum += int(c - 64)
		}

		if _, ok := t[sum]; ok {
			found++
		}
	}

	fmt.Println(found)
}
