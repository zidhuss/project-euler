package main

import (
	"fmt"
	"unicode/utf8"
)

func isPalindrome(original int) bool {
	n, reverse := original, 0
	for n > 0 {
		reverse = reverse*10 + n%10
		n /= 10
	}
	return original == reverse
}

func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func isPalindromeStr(s string) bool {
	return s == reverse(s)
}

func main() {
	sum := 0
	for i := 999999; i > 0; i -= 2 {
		if isPalindrome(i) && isPalindromeStr(fmt.Sprintf("%b", i)) {
			sum += i
		}
	}
	fmt.Printf("%d\n", sum)
}
