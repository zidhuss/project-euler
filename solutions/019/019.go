package main

import (
	"fmt"
)

func main() {
	dayOfYear := 365 % 7
	sundays := 0
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for year := 1901; year < 2001; year++ {
		for i := 0; i < 12; i++ {
			numberOfDays := months[i]
			if year%4 == 0 && i == 1 {
				numberOfDays = 29
			}
			day := dayOfYear % 7
			dayOfYear += numberOfDays
			if day == 6 {
				sundays++
			}
		}
	}
	fmt.Printf("%d\n", sundays)
}
