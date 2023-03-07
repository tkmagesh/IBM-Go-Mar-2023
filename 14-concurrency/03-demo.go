package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}
	fmt.Println("main started")
	wg.Add(1) // incrementing the counter by 1
	go f1(wg) //scheduling the execution of f1 through the scheduler
	f2()
	wg.Wait() // for the wg counter to become 0
	fmt.Println("main completed")
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() // decrementing the counter by 1
	fmt.Println("f1 started")
	time.Sleep(1 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
