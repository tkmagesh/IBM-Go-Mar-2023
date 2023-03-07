package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}
	fmt.Println("main started")
	wg.Add(1)
	go printOdds(wg)

	wg.Add(1)
	go printEven(wg)

	wg.Wait()
	fmt.Println("main completed")
}

func printOdds(wg *sync.WaitGroup) {
	defer wg.Done() // decrementing the counter by 1
	for i := 1; i < 20; i += 2 {
		fmt.Println("Odd :", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printEven(wg *sync.WaitGroup) {
	defer wg.Done() // decrementing the counter by 1
	for i := 0; i < 20; i += 2 {
		fmt.Println("Even :", i)
		time.Sleep(500 * time.Millisecond)
	}
}
