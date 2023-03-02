package main

import "fmt"

/*
	var name string = "Suresh"
*/

/*
	var name = "Suresh"
*/
/*
	name := "Suresh" // NOT SUPPORTED
*/

//unused varaible @package scope
var z int = 500

func main() {
	/*
		var name string
		name = "Magesh"
	*/
	/*
		var name string = "Magesh"
	*/

	//type inference
	/*
		var name = "Magesh"
	*/

	name := "Magesh"
	fmt.Printf("Hi %s, Have a nice day!\n", name)

	var n1 int = 100
	n1 = 200
	fmt.Println(n1)

	//Using multiple variables
	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/* combining declaration & initialization */
	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y int = 100, 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Sum of %d and %d is %d\n"
			result int    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	//type inference
	/*
		var x, y = 100, 200
		var str = "Sum of %d and %d is %d\n"
		var result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of %d and %d is %d\n"
			result    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	//using :=
	x, y, str := 100, 200, "Sum of %d and %d is %d\n"
	result := x + y
	fmt.Printf(str, x, y, result)
}
