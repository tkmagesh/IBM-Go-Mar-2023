/* Functions as values to variables */
package main

import "fmt"

func main() {

	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi there!")
	}
	sayHi()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	var getGreetMsg func(string) string
	getGreetMsg = func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a nice day!", userName)
	}
	greetMsg := getGreetMsg("Suresh")
	fmt.Println(greetMsg)

	var operation func(int, int) int
	operation = func(x, y int) int {
		return x + y
	}
	addResult := operation(100, 200)
	fmt.Println("add Result :", addResult)

	operation = func(x, y int) int {
		return x * y
	}
	multiplyResult := operation(100, 200)
	fmt.Println("multiply Result :", multiplyResult)

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}
	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

}
