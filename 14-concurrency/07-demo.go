package main

import (
	"fmt"
	"sync"
)

/*
	To detect data races
		go run -race <program.go>
		go build -race <program.go>

*/
type Counter struct {
	count int
	sync.Mutex
}

func (c *Counter) Add() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var counter Counter

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", counter.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Add()
}
