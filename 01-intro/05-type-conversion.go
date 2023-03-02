package main

import "fmt"

func main() {
	var i int32 = 100
	//use the typename as a function for type conversion
	var f float32 = float32(i)
	fmt.Println(f)
}
