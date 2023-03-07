package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	for {
		if data, isOpen := <-ch; isOpen {
			fmt.Println(data)
			continue
		}
		break
	}
	fmt.Println("main completed")
}

func fn(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		ch <- i * 10
	}
	fmt.Println("Job done")
	close(ch)
}
