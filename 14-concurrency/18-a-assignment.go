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
	ch := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go genPrimes(3, 100, ch, wg)

	wg.Add(1)
	go genPrimes(101, 200, ch, wg)

	wg.Add(1)
	go genPrimes(201, 300, ch, wg)

	//using a waitgroup
	/*
		wg2 := &sync.WaitGroup{}
		wg2.Add(1)
		go func() {
			for primeNo := range ch {
				fmt.Println(primeNo)
			}
			fmt.Println("All prime numbers received!")
			wg2.Done()
		}()

		wg.Wait()
		close(ch)
		wg2.Wait()
		fmt.Println("Done!")
	*/

	//using a channel
	done := make(chan struct{})
	go func() {
		for primeNo := range ch {
			fmt.Println(primeNo)
		}
		fmt.Println("All prime numbers received!")
		// done <- struct{}{}
		close(done)
	}()

	wg.Wait()
	close(ch)
	<-done
	fmt.Println("Done!")
}

func genPrimes(start, end int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for no := start; no <= end; no++ {
		if isPrime(no) {
			ch <- no
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
