package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var noOfGoroutines int
	wg := &sync.WaitGroup{}
	rand.Seed(7)
	flag.IntVar(&noOfGoroutines, "count", 0, "Number of goroutines to spawn")
	flag.Parse()
	fmt.Printf("%d goroutines are going to be scheduled.. Hit ENTER to start\n", noOfGoroutines)
	fmt.Scanln()
	for i := 1; i <= noOfGoroutines; i++ {
		wg.Add(1)
		go fn(i, wg)
	}
	wg.Wait()
	fmt.Println("Hit ENTER to shutdown...")
	fmt.Scanln()

}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
