/*
	Refactor the solution for assignment-01 using functions and error
	Ensure that each function has ONLY ONE responsibility (SRP)
*/
package main

import (
	"errors"
	"fmt"
)

var InvalidUserChoiceError error = errors.New("Invalid user choice")

func main() {
	for {
		userChoice, err := getUserChoice()
		if err == InvalidUserChoiceError {
			fmt.Println("Invalid user choice")
			continue
		}
		if err != nil {
			fmt.Println(err)
			continue
		}

		if userChoice == 5 {
			break
		}
		x, y, err := getOperands()
		if err != nil {
			fmt.Println(err)
			continue
		}
		processUserChoice(userChoice, x, y)
	}
}

func processUserChoice(userChoice, x, y int) {
	var result = 0
	switch userChoice {
	case 1:
		result = add(x, y)
	case 2:
		result = subtract(x, y)
	case 3:
		result = multiply(x, y)
	case 4:
		result = divide(x, y)
	}
	fmt.Println("Result :", result)
}
func getOperands() (x, y int, err error) {
	fmt.Println("Enter the value :")
	_, err = fmt.Scanln(&x, &y)
	return
}

func getUserChoice() (userChoice int, err error) {
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("Enter your choice:")
	n, err := fmt.Scanln(&userChoice)
	if n != 1 {
		return
	}
	if userChoice < 1 || userChoice > 5 {
		err = InvalidUserChoiceError
		return
	}
	return
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
