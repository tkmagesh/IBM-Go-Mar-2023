package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	/*
		var ch chan int
		ch = make(chan int)
	*/
	ch := make(chan int)
	wg.Add(1)
	go add(100, 200, wg, ch)
	wg.Wait()
	result := <-ch
	fmt.Println("result :", result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	result := x + y
	ch <- result
}
