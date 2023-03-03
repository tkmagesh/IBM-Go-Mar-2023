package calculator

import "fmt"

func init() {
	fmt.Println("calculator package initialized - [subtract.go]")
}

func Subtract(x, y int) int {
	// op_count++
	op_count["subtract"]++
	return x - y
}
