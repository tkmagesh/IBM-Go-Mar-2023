package main

import "fmt"

func main() {
	nextId := getIdGenerator()
	fmt.Println(nextId()) //=> 1
	fmt.Println(nextId()) //=> 2
	// counter = 10000
	fmt.Println(nextId()) //=> 3
	fmt.Println(nextId()) //=> 4
}

func getIdGenerator() func() int {
	var counter int
	return func() int {
		counter++
		return counter
	}
}
