package main

import "fmt"

func main() {
	var f_name, l_name string
	fmt.Scanln(&f_name, &l_name)
	fmt.Printf("Hi %s %s, Have a good day!\n", f_name, l_name)

	var x, y int
	fmt.Scanln(&x, &y)
	fmt.Println(x + y)
}
