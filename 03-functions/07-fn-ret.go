package main

import "fmt"

func main() {
	fn := getFn()
	fn()
}

func getFn() func() {
	return func() {
		fmt.Println("anon fn invoked")
	}
}
