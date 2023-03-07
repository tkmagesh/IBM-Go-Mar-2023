/*
Modify the below program to generate prime numbers using 3 instances of genPrimes functions and print the generated values
	genPrimes(3,100)
	genPrimes(101,200)
	genPrimes(201,300)
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	primeData := []struct {
		start int
		end   int
	}{
		{start: 3, end: 100},
		{start: 101, end: 200},
		{start: 201, end: 300},
		{start: 301, end: 400},
		{start: 401, end: 500},
	}

	primeChannels := make([]<-chan int, 0)
	for _, pData := range primeData {
		ch := genPrimes(pData.start, pData.end)
		primeChannels = append(primeChannels, ch)
	}

	primeNosCh := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(len(primeChannels))
	for _, primeCh := range primeChannels {
		go func(ch <-chan int) {
			for data := range ch {
				primeNosCh <- data
			}
			wg.Done()
		}(primeCh)
	}

	doneCh := make(chan struct{})
	go func() {
		for primeNo := range primeNosCh {
			fmt.Println(primeNo)
		}
		close(doneCh)
	}()

	wg.Wait()
	close(primeNosCh)
	<-doneCh
	fmt.Println("Done!")
}

func genPrimes(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			if isPrime(no) {
				ch <- no
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(ch)
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
