package main

import (
	"fmt"
	"log"
	"time"
)

type Operation func(int, int)

func main() {

	logAdd := getLogOperation(add)
	profileLogAdd := getProfileOperation(logAdd)
	profileLogAdd(100, 200)

	getProfileOperation(getLogOperation(subtract))(100, 200)
}

func getProfileOperation(operationFn Operation) Operation {
	return func(x, y int) {
		start := time.Now()
		operationFn(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
}

func getLogOperation(operationFn Operation) Operation {
	return func(x, y int) {
		log.Println("Operation started...")
		operationFn(x, y)
		log.Println("Operation completed!")
	}
}

//3rd party libraries
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
