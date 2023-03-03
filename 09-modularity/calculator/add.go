package calculator

import "fmt"

func init() {
	fmt.Println("calculator package initialized - [add.go]")
}

func Add(x, y int) int {
	// op_count++
	op_count["add"]++
	return x + y
}
