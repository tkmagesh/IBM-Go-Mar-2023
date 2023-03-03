package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("cannot divide by 0")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("something went wrong!", err)
			return
		}
		fmt.Println("Thank you!")
	}()
	divisor := 0
	q, r, err := divideClient(100, divisor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

//3rd party libray
func divide(x, y int) (quotient, remainder int) {
	fmt.Println("calculating quotient")
	if y == 0 {
		panic(DivideByZeroError)
	}
	quotient = x / y
	fmt.Println("calculating remainder")
	remainder = x % y
	return
}
