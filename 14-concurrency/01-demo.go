package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //scheduling the execution of f1 through the scheduler
	f2()
	// DO NOT use the following synchronization strategies
	time.Sleep(2 * time.Second) // blocking the execution of the function (main)
	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(1 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
