package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	To detect data races
		go run -race <program.go>
		go build -race <program.go>

*/
var count int64

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&count, 1)
}
