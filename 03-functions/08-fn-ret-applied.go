package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)
	*/

	/*
		profileAdd := getProfileOperation(add)
		profileAdd(100, 200)

		profileSubtract := getProfileOperation(subtract)
		profileSubtract(100, 200)
	*/

	logAdd := getLogOperation(add)
	profileLogAdd := getProfileOperation(logAdd)
	profileLogAdd(100, 200)

	getProfileOperation(getLogOperation(subtract))(100, 200)
}

func getProfileOperation(operationFn func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		operationFn(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
}

func getLogOperation(operationFn func(int, int)) func(int, int) {
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
