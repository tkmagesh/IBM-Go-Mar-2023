/*
Write a goroutine that generates the prime numbers between the given start and end
The main function should print the generated numbers as and when they are generated
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go genPrimes(3, 100, ch)
	for primeNo := range ch {
		fmt.Println(primeNo)
	}
	fmt.Println("Done!")
}

func genPrimes(start, end int, ch chan int) {
	for no := start; no <= end; no++ {
		if isPrime(no) {
			ch <- no
			time.Sleep(500 * time.Millisecond)
		}
	}
	close(ch)
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
