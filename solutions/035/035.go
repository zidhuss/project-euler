package main

import (
	"fmt"
	"math"
	"strconv"
)

func hasEven(str string) bool {
	for _, c := range str {
		if c == '0' || c == '2' || c == '5' || c == '4' || c == '6' || c == '8' {
			return true
		}
	}
	return false
}

func generatePrimeStringSet(limit int) map[string]bool {
	sievebound := (limit - 1) / 2
	crosslimit := int((math.Sqrt(float64(limit)) - 1) / 2)
	sieve := make([]bool, sievebound)

	for i := 1; i < crosslimit; i++ {
		if !sieve[i] {
			for j := 2 * i * (i + 1); j < sievebound; j += 2*i + 1 {
				sieve[j] = true
			}
		}
	}

	primeStrings := make(map[string]bool)

	for i, b := range sieve {
		if !b {
			str := strconv.Itoa(2*i + 1)
			if len(str) > 2 && !hasEven(str) {
				primeStrings[str] = true
			}
		}
	}
	return primeStrings
}

func main() {
	primeStrings := generatePrimeStringSet(1000000)

	// Circular primes less than 100
	count := 13

	for len(primeStrings) > 0 {

		var perms []string
		found := 0

		for k := range primeStrings {
			perms = permutations(k)
			break
		}

		for _, p := range perms {
			if primeStrings[p] {
				found++
				delete(primeStrings, p)
			}
		}

		if found == len(perms) {
			count += found
		}
	}

	fmt.Println(count)
}

func permutations(str string) []string {
	result := []string{str}
	for i := 1; i < len(str); i++ {
		perm := str[i:] + str[0:i]
		if perm[0] == '0' {
			continue
		}
		result = append(result, perm)
	}
	return result
}
