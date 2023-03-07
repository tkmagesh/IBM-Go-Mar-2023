package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	fmt.Println("len(ch) :", len(ch))
	fmt.Println("Sending value - 1")
	ch <- 100
	fmt.Println("len(ch) :", len(ch))

	fmt.Println("Sending value - 2")
	ch <- 200
	fmt.Println("len(ch) :", len(ch))

	fmt.Println(<-ch)
	fmt.Println("len(ch) :", len(ch))

	fmt.Println(<-ch)
	fmt.Println("len(ch) :", len(ch))
}
