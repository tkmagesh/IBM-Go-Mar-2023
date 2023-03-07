package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("main started")
	wg.Add(1) // incrementing the counter by 1
	go f1()   //scheduling the execution of f1 through the scheduler
	f2()
	wg.Wait() // for the wg counter to become 0
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(1 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrementing the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
