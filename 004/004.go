package main

import "fmt"

func isPalindrome(original int) bool {
	n, reverse := original, 0
	for n > 0 {
		reverse = reverse*10 + n%10
		n /= 10
	}
	return original == reverse
}

func main() {
	largest, last, first := 0, 999, 100

	for a := last; a >= first; a-- {
		for b := last; b >= first; b-- {
			calculated := a * b
			if calculated > largest && isPalindrome(calculated) {
				largest = calculated
			}
		}
	}
	fmt.Printf("%d\n", largest)
}
