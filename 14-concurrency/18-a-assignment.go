/*
Modify the below program to generate prime numbers using 3 instances of genPrimes functions and print the generated values
	genPrimes(3,100)
	genPrimes(101,200)
	genPrimes(201,300)
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
