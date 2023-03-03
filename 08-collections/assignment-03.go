/* Write a function that returns all the prime numbers beween the given start and end */

package main

import "fmt"

func main() {
	primes := getPrimes(3, 100)
	fmt.Println(primes)
}

func getPrimes(start, end int) []int {
	primes := make([]int, 0)
	for i := start; i <= end; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
