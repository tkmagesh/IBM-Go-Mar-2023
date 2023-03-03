package calculator

import "fmt"

var op_count map[string]int

func init() {
	fmt.Println("calculator package initialized - [calc.go]")
	op_count = make(map[string]int)
}

func OpCount() map[string]int {
	return op_count
}
